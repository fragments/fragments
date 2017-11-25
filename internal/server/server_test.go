package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	fsmocks "github.com/fragments/fragments/internal/filestore/mocks"
	"github.com/fragments/fragments/internal/state"
	"github.com/fragments/fragments/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutFunction(t *testing.T) {
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := state.PutModel(ctx, initial, state.ModelTypeFunction, &state.Function{
		Meta: state.Meta{
			Name: "existing",
			Labels: map[string]string{
				"code":   "initial",
				"config": "initial",
			},
		},
		AWS:            &state.FunctionAWS{Timeout: 3, Memory: 256},
		Checksum:       "ABC",
		Runtime:        "nodejs",
		SourceFilename: "existing.tar.gz",
	})
	require.NoError(t, err)

	tests := []struct {
		TestName string
		Function *state.Function
		Token    string
		Response *UploadRequest
		Error    bool
	}{
		{
			TestName: "NoInput",
			Function: nil,
			Error:    true,
		},
		{
			TestName: "NoName",
			Function: &state.Function{},
			Error:    true,
		},
		{
			TestName: "CreateNew",
			Function: &state.Function{
				Meta: state.Meta{
					Name: "new",
					Labels: map[string]string{
						"code":   "new",
						"config": "new",
					},
				},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:  "nodejs",
				Checksum: "new",
			},
			Token: "newtoken",
			Response: &UploadRequest{
				Token: "newtoken",
				URL:   "https://newtoken",
			},
		},
		{
			TestName: "UpdateCode",
			Function: &state.Function{
				Meta: state.Meta{
					Name: "existing",
					Labels: map[string]string{
						"code":   "updated",
						"config": "initial",
					},
				},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:  "nodejs",
				Checksum: "UPDATED",
			},
			Token: "codetoken",
			Response: &UploadRequest{
				Token: "codetoken",
				URL:   "https://codetoken",
			},
		},
		{
			TestName: "UpdateConfig",
			Function: &state.Function{
				Meta: state.Meta{
					Name: "existing",
					Labels: map[string]string{
						"config": "updated",
						"code":   "initial",
					},
				},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 512},
				Runtime:  "nodejs",
				Checksum: "ABC",
			},
			Response: nil,
		},
		{
			TestName: "UpdateCodeAndConfig",
			Function: &state.Function{
				Meta: state.Meta{
					Name: "existing",
					Labels: map[string]string{
						"code":   "updated",
						"config": "updated",
					},
				},
				AWS:      &state.FunctionAWS{Timeout: 10, Memory: 1024},
				Runtime:  "nodejs",
				Checksum: "ABC123",
			},
			Token: "token",
			Response: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "NoChange",
			Function: &state.Function{
				Meta:     state.Meta{Name: "existing"},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:  "nodejs",
				Checksum: "ABC",
			},
			Response: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()

			mockSourceStore := &fsmocks.SourceTarget{}
			mockSourceStore.
				On("NewUploadURL", test.Token).
				Return(fmt.Sprintf("https://%s", test.Token), nil)

			kv := initial.Copy()

			s := &Server{
				StateStore:    kv,
				SourceStore:   mockSourceStore,
				GenerateToken: func() string { return test.Token },
			}

			res, err := s.PutFunction(ctx, test.Function)
			if test.Error {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.Response, res)

			testutils.AssertGolden(
				t,
				testutils.SnapshotJSONMap(kv.Data),
				fmt.Sprintf("testdata/TestPutFunction-%s.yaml", test.TestName),
			)
		})
	}
}

func TestConfirmUpload(t *testing.T) {
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := state.PutModel(ctx, initial, state.ModelTypeFunction, &state.Function{
		Meta:           state.Meta{Name: "existing"},
		AWS:            &state.FunctionAWS{Timeout: 3, Memory: 256},
		Runtime:        "go",
		SourceFilename: "previous.tar.gz",
		Checksum:       "foo",
	})
	require.NoError(t, err)
	err = state.PutPendingUpload(ctx, initial, "new", &state.PendingUpload{
		Filename: "new.tar.gz",
		Function: &state.Function{
			Meta:     state.Meta{Name: "new"},
			AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
			Runtime:  "go",
			Checksum: "new",
		},
	})
	require.NoError(t, err)
	err = state.PutPendingUpload(ctx, initial, "update-config", &state.PendingUpload{
		Filename: "foo.tar.gz",
		Function: &state.Function{
			Meta:     state.Meta{Name: "existing"},
			AWS:      &state.FunctionAWS{Timeout: 5, Memory: 1024},
			Runtime:  "nodejs",
			Checksum: "foo",
		},
	})
	require.NoError(t, err)
	err = state.PutPendingUpload(ctx, initial, "update-code", &state.PendingUpload{
		Filename: "bar.tar.gz",
		Function: &state.Function{
			Meta:     state.Meta{Name: "existing"},
			AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
			Runtime:  "go",
			Checksum: "updated",
		},
	})
	require.NoError(t, err)

	tests := []struct {
		TestName string
		Token    string
		Error    bool
	}{
		{
			TestName: "NoToken",
			Token:    "",
			Error:    true,
		},
		{
			TestName: "NoPendingUpload",
			Token:    "baz",
			Error:    true,
		},
		{
			TestName: "New",
			Token:    "new",
		},
		{
			TestName: "UpdatedConfig",
			Token:    "update-config",
		},
		{
			TestName: "UpdateCode",
			Token:    "update-code",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()

			mockSourceStore := &fsmocks.SourceTarget{}
			mockSourceStore.
				On("Persist", ctx, test.Token).
				Return(nil)

			kv := initial.Copy()

			s := &Server{
				StateStore:    kv,
				SourceStore:   mockSourceStore,
				GenerateToken: func() string { return test.Token },
			}

			err := s.ConfirmUpload(ctx, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			mockSourceStore.AssertExpectations(t)

			testutils.AssertGolden(
				t,
				testutils.SnapshotJSONMap(kv.Data),
				fmt.Sprintf("testdata/TestConfirmUpload-%s.yaml", test.TestName),
			)
		})
	}
}

func TestCreateEnvironment(t *testing.T) {
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := state.PutModel(ctx, initial, state.ModelTypeEnvironment, &state.Environment{
		Meta:           state.Meta{Name: "existing"},
		Infrastructure: state.InfrastructureTypeAWS,
	})
	require.NoError(t, err)

	tests := []struct {
		TestName string
		Input    *EnvironmentInput
		Error    bool
	}{
		{
			TestName: "NoInput",
			Input:    nil,
			Error:    true,
		},
		{
			TestName: "NoName",
			Input:    &EnvironmentInput{},
			Error:    true,
		},
		{
			TestName: "Existing",
			Input: &EnvironmentInput{
				Name: "existing",
			},
			Error: true,
		},
		{
			TestName: "New",
			Input: &EnvironmentInput{
				Name: "new",
				Labels: map[string]string{
					"new": "true",
				},
				Infrastructure: state.InfrastructureTypeAWS,
				Username:       "user",
				Password:       "pass",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()

			kv := initial.Copy()
			secretsKV := backend.NewTestKV()

			s := &Server{
				StateStore:    kv,
				SecretStore:   secretsKV,
				SourceStore:   nil,
				GenerateToken: nil,
			}

			err := s.CreateEnvironment(ctx, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			testutils.AssertGolden(
				t,
				testutils.SnapshotJSONMap(kv.Data),
				fmt.Sprintf("testdata/TestCreateEnvironment-%s-state.yaml", test.TestName),
			)
			testutils.AssertGolden(
				t,
				testutils.SnapshotStringMap(secretsKV.Data),
				fmt.Sprintf("testdata/TestCreateEnvironment-%s-secrets.yaml", test.TestName),
			)
		})
	}
}

func TestPutDeployment(t *testing.T) {
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := state.PutModel(ctx, initial, state.ModelTypeDeployment, &state.Deployment{
		Meta:              state.Meta{Name: "existing"},
		EnvironmentLabels: map[string]string{},
		FunctionLabels:    map[string]string{},
	})
	require.NoError(t, err)

	tests := []struct {
		TestName string
		Input    *state.Deployment
		Error    bool
	}{
		{
			TestName: "NoInput",
			Input:    nil,
			Error:    true,
		},
		{
			TestName: "NoName",
			Input:    &state.Deployment{},
			Error:    true,
		},
		{
			TestName: "New",
			Input: &state.Deployment{
				Meta:              state.Meta{Name: "new"},
				EnvironmentLabels: map[string]string{"foo": "foo"},
				FunctionLabels:    map[string]string{"bar": "bar"},
			},
		},
		{
			TestName: "Update",
			Input: &state.Deployment{
				Meta: state.Meta{
					Name: "existing",
					Labels: map[string]string{
						"foo": "bar",
					},
				},
				EnvironmentLabels: map[string]string{"foo": "foo"},
				FunctionLabels:    map[string]string{"bar": "bar"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()

			kv := initial.Copy()
			s := &Server{
				StateStore: kv,
			}

			err := s.PutDeployment(ctx, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			testutils.AssertGolden(
				t,
				testutils.SnapshotJSONMap(kv.Data),
				fmt.Sprintf("testdata/TestPutDeployment-%s-state.yaml", test.TestName),
			)
		})
	}
}
