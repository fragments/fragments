package server

import (
	"context"
	"flag"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	fsmocks "github.com/fragments/fragments/internal/filestore/mocks"
	"github.com/fragments/fragments/internal/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("test.update", false, "update test snapshots")

func TestPutFunction(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := state.PutModel(ctx, kv, state.ModelTypeFunction, &state.Function{
			Meta:     state.Meta{Name: "foo"},
			AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
			Checksum: "foo",
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "TestPutFunction.json")
	}

	foo := &state.Function{
		Meta:     state.Meta{Name: "foo"},
		AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
		Checksum: "foo",
	}
	bar := &state.Function{
		Meta:     state.Meta{Name: "bar"},
		AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
		Checksum: "bar",
	}
	fooCode := &state.Function{
		Meta:     state.Meta{Name: "foo"},
		AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
		Checksum: "foobar",
	}
	fooConfig := &state.Function{
		Meta:     state.Meta{Name: "foo"},
		AWS:      &state.FunctionAWS{Timeout: 3, Memory: 512},
		Checksum: "foo",
	}

	tests := []struct {
		TestName string
		Function *state.Function
		Token    string
		Response *UploadRequest
		Error    bool
	}{
		{
			TestName: "No input",
			Function: nil,
			Error:    true,
		},
		{
			TestName: "No name",
			Function: &state.Function{},
			Error:    true,
		},
		{
			TestName: "No existing",
			Function: bar,
			Token:    "token",
			Response: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "Update code",
			Function: fooCode,
			Token:    "token",
			Response: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "Update config",
			Function: fooConfig,
			Response: nil,
		},
		{
			TestName: "No change",
			Function: foo,
			Response: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := backend.NewTestKV("TestPutFunction.json")

			mockSourceStore := &fsmocks.SourceTarget{}
			mockSourceStore.
				On("NewUploadURL", test.Token).
				Return(fmt.Sprintf("https://%s", test.Token), nil)

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

			kv.AssertSnapshot(t, fmt.Sprintf("TestPutFunction-%s.json", test.TestName), *update)
		})
	}
}

func TestConfirmUpload(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := state.PutModel(ctx, kv, state.ModelTypeFunction, &state.Function{
			Meta:     state.Meta{Name: "foo"},
			AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
			Checksum: "foo",
		})
		require.NoError(t, err)
		err = state.PutPendingUpload(ctx, kv, "foo-config", &state.PendingUpload{
			Filename: "foo.tar.gz",
			Function: &state.Function{
				Meta:     state.Meta{Name: "foo"},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 1024},
				Checksum: "foo",
			},
		})
		require.NoError(t, err)
		err = state.PutPendingUpload(ctx, kv, "foo-code", &state.PendingUpload{
			Filename: "bar.tar.gz",
			Function: &state.Function{
				Meta:     state.Meta{Name: "bar"},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Checksum: "updated",
			},
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "TestConfirmUpload.json")
	}

	tests := []struct {
		TestName string
		Token    string
		Error    bool
	}{
		{
			TestName: "No token",
			Token:    "",
			Error:    true,
		},
		{
			TestName: "No pending upload",
			Token:    "baz",
			Error:    true,
		},
		{
			TestName: "Updated config",
			Token:    "foo-config",
		},
		{
			TestName: "Update code",
			Token:    "foo-code",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := backend.NewTestKV("TestConfirmUpload.json")

			mockSourceStore := &fsmocks.SourceTarget{}
			mockSourceStore.
				On("Persist", ctx, test.Token).
				Return(nil)

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

			kv.AssertSnapshot(t, fmt.Sprintf("TestConfirmUpload-%s.json", test.TestName), *update)
		})
	}
}

func TestCreateEnvironment(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := state.PutModel(ctx, kv, state.ModelTypeEnvironment, &state.Environment{
			Meta:           state.Meta{Name: "foo"},
			Infrastructure: state.InfrastructureTypeAWS,
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "TestCreateEnvironment.json")
	}

	tests := []struct {
		TestName string
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
			Input: &EnvironmentInput{
				Name: "foo",
			},
			Error: true,
		},
		{
			TestName: "New",
			Input: &EnvironmentInput{
				Name: "bar",
				Labels: map[string]string{
					"foo": "bar",
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
			kv := backend.NewTestKV("TestCreateEnvironment.json")

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

			kv.AssertSnapshot(t, fmt.Sprintf("TestCreateEnvironment-%s-state", test.TestName), *update)
			secretsKV.AssertSnapshot(t, fmt.Sprintf("TestCreateEnvironment-%s-secrets", test.TestName), *update)
		})
	}
}

func TestPutDeployment(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := state.PutModel(ctx, kv, state.ModelTypeDeployment, &state.Deployment{
			Meta:              state.Meta{Name: "foo"},
			EnvironmentLabels: map[string]string{},
			FunctionLabels:    map[string]string{},
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "TestPutDeployment.json")
	}

	tests := []struct {
		TestName string
		Input    *state.Deployment
		Error    bool
	}{
		{
			TestName: "No input",
			Input:    nil,
			Error:    true,
		},
		{
			TestName: "No name",
			Input:    &state.Deployment{},
			Error:    true,
		},
		{
			TestName: "New",
			Input: &state.Deployment{
				Meta:              state.Meta{Name: "bar"},
				EnvironmentLabels: map[string]string{"foo": "foo"},
				FunctionLabels:    map[string]string{"bar": "bar"},
			},
		},
		{
			TestName: "Update",
			Input: &state.Deployment{
				Meta: state.Meta{
					Name: "foo",
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
			kv := backend.NewTestKV("TestPutDeployment.json")

			s := &Server{
				StateStore: kv,
			}

			err := s.PutDeployment(ctx, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			kv.AssertSnapshot(t, fmt.Sprintf("TestPutDeployment-%s", test.TestName), *update)
		})
	}
}
