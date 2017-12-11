package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	fsmocks "github.com/fragments/fragments/internal/filestore/mocks"
	"github.com/fragments/fragments/internal/model"
	"github.com/fragments/fragments/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutFunction(t *testing.T) {
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := putFunction(ctx, initial, &model.Function{
		Name: "existing",
		Labels: map[string]string{
			"code":   "initial",
			"config": "initial",
		},
		AWS:            &model.FunctionAWS{Timeout: 3, Memory: 256},
		Checksum:       "ABC",
		Runtime:        "nodejs",
		SourceFilename: "existing.tar.gz",
	})
	require.NoError(t, err)

	tests := []struct {
		TestName string
		Function *model.Function
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
			Function: &model.Function{},
			Error:    true,
		},
		{
			TestName: "CreateNew",
			Function: &model.Function{
				Name: "new",
				Labels: map[string]string{
					"code":   "new",
					"config": "new",
				},
				AWS:      &model.FunctionAWS{Timeout: 3, Memory: 256},
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
			Function: &model.Function{
				Name: "existing",
				Labels: map[string]string{
					"code":   "updated",
					"config": "initial",
				},
				AWS:      &model.FunctionAWS{Timeout: 3, Memory: 256},
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
			Function: &model.Function{
				Name: "existing",
				Labels: map[string]string{
					"config": "updated",
					"code":   "initial",
				},
				AWS:      &model.FunctionAWS{Timeout: 3, Memory: 512},
				Runtime:  "nodejs",
				Checksum: "ABC",
			},
			Response: nil,
		},
		{
			TestName: "UpdateCodeAndConfig",
			Function: &model.Function{
				Name: "existing",
				Labels: map[string]string{
					"code":   "updated",
					"config": "updated",
				},
				AWS:      &model.FunctionAWS{Timeout: 10, Memory: 1024},
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
			Function: &model.Function{
				Name: "existing",
				Labels: map[string]string{
					"code":   "initial",
					"config": "initial",
				},
				AWS:            &model.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:        "nodejs",
				Checksum:       "ABC",
				SourceFilename: "existing.tar.gz",
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

			s := New(kv, nil, mockSourceStore)
			s.GenerateToken = func() string {
				return test.Token
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
	err := putFunction(ctx, initial, &model.Function{
		Name:           "existing",
		AWS:            &model.FunctionAWS{Timeout: 3, Memory: 256},
		Runtime:        "go",
		SourceFilename: "previous.tar.gz",
		Checksum:       "foo",
	})
	require.NoError(t, err)
	err = putPendingUpload(ctx, initial, &model.PendingUpload{
		Token:    "new",
		Filename: "new.tar.gz",
		Function: &model.Function{
			Name:     "new",
			AWS:      &model.FunctionAWS{Timeout: 3, Memory: 256},
			Runtime:  "go",
			Checksum: "new",
		},
	})
	require.NoError(t, err)
	err = putPendingUpload(ctx, initial, &model.PendingUpload{
		Token:    "update-config",
		Filename: "foo.tar.gz",
		Function: &model.Function{
			Name:     "existing",
			AWS:      &model.FunctionAWS{Timeout: 5, Memory: 1024},
			Runtime:  "nodejs",
			Checksum: "foo",
		},
	})
	require.NoError(t, err)
	err = putPendingUpload(ctx, initial, &model.PendingUpload{
		Token:    "update-code",
		Filename: "bar.tar.gz",
		Function: &model.Function{
			Name:     "existing",
			AWS:      &model.FunctionAWS{Timeout: 3, Memory: 256},
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

			s := New(kv, nil, mockSourceStore)
			s.GenerateToken = func() string {
				return test.Token
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
	err := putEnvironment(ctx, initial, &model.Environment{
		Name:           "existing",
		Infrastructure: model.InfrastructureTypeAWS,
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
				Infrastructure: model.InfrastructureTypeAWS,
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

			s := New(kv, secretsKV, nil)

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
	err := putDeployment(ctx, initial, &model.Deployment{
		Name:              "existing",
		EnvironmentLabels: map[string]string{},
		FunctionLabels:    map[string]string{},
	})
	require.NoError(t, err)

	tests := []struct {
		TestName string
		Input    *model.Deployment
		Error    bool
	}{
		{
			TestName: "NoInput",
			Input:    nil,
			Error:    true,
		},
		{
			TestName: "NoName",
			Input:    &model.Deployment{},
			Error:    true,
		},
		{
			TestName: "New",
			Input: &model.Deployment{
				Name:              "new",
				EnvironmentLabels: map[string]string{"foo": "foo"},
				FunctionLabels:    map[string]string{"bar": "bar"},
			},
		},
		{
			TestName: "Update",
			Input: &model.Deployment{
				Name:              "existing",
				EnvironmentLabels: map[string]string{"foo": "foo"},
				FunctionLabels:    map[string]string{"bar": "bar"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()

			kv := initial.Copy()
			s := New(kv, nil, nil)

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
