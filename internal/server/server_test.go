package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	fsmocks "github.com/fragments/fragments/internal/filestore/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutFunction(t *testing.T) {
	foo := &Function{
		Meta:     Meta{Name: "foo"},
		AWS:      &FunctionAWS{Memory: 256},
		Checksum: "foo",
	}
	fooCode := &Function{
		Meta:     Meta{Name: "foo"},
		AWS:      &FunctionAWS{Memory: 256},
		Checksum: "foobar",
	}
	fooConfig := &Function{
		Meta:     Meta{Name: "foo"},
		AWS:      &FunctionAWS{Memory: 512},
		Checksum: "foo",
	}

	tests := []struct {
		TestName string
		Input    *Function
		Existing *Function
		Token    string
		Result   *Function
		Expected *UploadRequest
		Error    bool
	}{
		{
			TestName: "No input",
			Input:    nil,
			Error:    true,
		},
		{
			TestName: "No name",
			Input:    &Function{},
			Error:    true,
		},
		{
			TestName: "No existing",
			Input:    foo,
			Token:    "token",
			Expected: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "Update code",
			Input:    fooCode,
			Existing: foo,
			Token:    "token",
			Expected: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "Update config",
			Input:    fooConfig,
			Existing: foo,
			Result:   fooConfig,
		},
		{
			TestName: "No change",
			Input:    foo,
			Existing: foo,
			Expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			mockKV := backend.NewMemoryKV()
			if test.Existing != nil {
				_ = putResource(ctx, mockKV, ResourceTypeFunction, test.Existing)
			}

			mockSourceStore := &fsmocks.SourceTarget{}
			mockSourceStore.
				On("NewUploadURL", test.Token).
				Return(fmt.Sprintf("https://%s", test.Token), nil)

			s := &Server{
				StateStore:    mockKV,
				SourceStore:   mockSourceStore,
				GenerateToken: func() string { return test.Token },
			}

			actual, err := s.PutFunction(ctx, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.Expected, actual)

			if test.Result != nil {
				current, _ := getFunction(ctx, mockKV, test.Input.Meta.Name)
				assert.Equal(t, current, test.Input)
			}
		})
	}
}

func TestConfirmUpload(t *testing.T) {
	foo := &Function{
		Meta:     Meta{Name: "foo"},
		AWS:      &FunctionAWS{Memory: 256},
		Checksum: "foo",
	}
	uploadNew := &PendingUpload{
		Filename: "foo.tar.gz",
		Function: foo,
	}
	uploadUpdate := &PendingUpload{
		Filename:         "bar.tar.gz",
		PreviousFilename: "foo.tar.gz",
		Function:         foo,
	}

	tests := []struct {
		TestName string
		Token    string
		Pending  *PendingUpload
		Error    bool
	}{
		{
			TestName: "No token",
			Token:    "",
			Error:    true,
		},
		{
			TestName: "No pending upload",
			Token:    "token",
			Error:    true,
		},
		{
			TestName: "New upload",
			Token:    "token",
			Pending:  uploadNew,
		},
		{
			TestName: "Update code",
			Token:    "token",
			Pending:  uploadUpdate,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			mockKV := backend.NewMemoryKV()
			if test.Pending != nil {
				_ = putPendingUpload(ctx, mockKV, test.Token, test.Pending)
			}

			mockSourceStore := &fsmocks.SourceTarget{}
			mockSourceStore.
				On("Persist", ctx, test.Token).
				Return(nil)

			s := &Server{
				StateStore:    mockKV,
				SourceStore:   mockSourceStore,
				GenerateToken: func() string { return test.Token },
			}

			err := s.ConfirmUpload(ctx, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}

			mockSourceStore.AssertExpectations(t)

			require.NoError(t, err)

			u, _ := getPendingUpload(ctx, mockKV, test.Token)
			assert.Nil(t, u)

			expected := test.Pending.Function
			expected.SourceFilename = test.Pending.Filename

			actual, _ := getFunction(ctx, mockKV, expected.Meta.Name)
			assert.EqualValues(t, expected, actual)
		})
	}
}

func TestCreateEnvironment(t *testing.T) {
	foo := &Environment{
		Meta:           Meta{Name: "foo"},
		Infrastructure: InfrastructureTypeAWS,
	}

	tests := []struct {
		TestName string
		Existing *Environment
		Input    *EnvironmentInput
		Error    bool
	}{
		{
			TestName: "No input",
			Input:    nil,
			Error:    true,
		},
		{
			TestName: "No name",
			Input:    &EnvironmentInput{},
			Error:    true,
		},
		{
			TestName: "Existing",
			Existing: foo,
			Input: &EnvironmentInput{
				Name: "foo",
			},
			Error: true,
		},
		{
			TestName: "New",
			Existing: foo,
			Input: &EnvironmentInput{
				Name: "bar",
				Labels: map[string]string{
					"foo": "bar",
				},
				Infrastructure: InfrastructureTypeAWS,
				Username:       "user",
				Password:       "pass",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			mockKV := backend.NewMemoryKV()
			if test.Existing != nil {
				_ = putResource(ctx, mockKV, ResourceTypeEnvironment, test.Existing)
			}

			mockSecrets := backend.NewMemoryKV()

			s := &Server{
				StateStore:    mockKV,
				SecretStore:   mockSecrets,
				SourceStore:   nil,
				GenerateToken: nil,
			}

			err := s.CreateEnvironment(ctx, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			expected := &Environment{
				Meta: Meta{
					Name:   test.Input.Name,
					Labels: test.Input.Labels,
				},
				Infrastructure: test.Input.Infrastructure,
			}

			actual, _ := getEnvironment(ctx, mockKV, test.Input.Name)
			assert.EqualValues(t, expected, actual)

			user, err := mockSecrets.Get(ctx, fmt.Sprintf("user/%s/%s", test.Input.Name, keySecretUser))
			require.NoError(t, err)
			assert.Equal(t, test.Input.Username, user)
			pass, err := mockSecrets.Get(ctx, fmt.Sprintf("user/%s/%s", test.Input.Name, keySecretPass))
			require.NoError(t, err)
			assert.Equal(t, test.Input.Password, pass)
		})
	}
}
