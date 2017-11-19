package server

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	fsmocks "github.com/fragments/fragments/internal/filestore/mocks"
	"github.com/fragments/fragments/internal/state"
	"github.com/fragments/fragments/pkg/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("test.update", false, "update test snapshots")

func TestPutFunction(t *testing.T) {
	snapshotFile := "testdata/TestPutFunction.yaml"
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := state.PutModel(ctx, kv, state.ModelTypeFunction, &state.Function{
			Meta: state.Meta{
				Name: "foo",
				Labels: map[string]string{
					"snapshot": "yes",
				},
			},
			AWS:            &state.FunctionAWS{Timeout: 3, Memory: 256},
			Checksum:       "foo",
			Runtime:        "nodejs",
			SourceFilename: "source.tar.gz",
		})
		require.NoError(t, err)
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		TestName string
		Function *state.Function
		Snapshot string
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
				Meta:     state.Meta{Name: "foo"},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:  "nodejs",
				Checksum: "foo",
			},
			Token: "token",
			Response: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "UpdateCode",
			Snapshot: snapshotFile,
			Function: &state.Function{
				Meta: state.Meta{
					Name: "foo",
					Labels: map[string]string{
						"code": "updated",
					},
				},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:  "nodejs",
				Checksum: "foobar",
			},
			Token: "token",
			Response: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "UpdateConfig",
			Snapshot: snapshotFile,
			Function: &state.Function{
				Meta: state.Meta{
					Name: "foo",
					Labels: map[string]string{
						"config": "updated",
					},
				},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 512},
				Runtime:  "nodejs",
				Checksum: "foo",
			},
			Response: nil,
		},
		{
			TestName: "UpdateCodeAndConfig",
			Snapshot: snapshotFile,
			Function: &state.Function{
				Meta: state.Meta{
					Name: "foo",
					Labels: map[string]string{
						"code":   "updated",
						"config": "updated",
					},
				},
				AWS:      &state.FunctionAWS{Timeout: 10, Memory: 1024},
				Runtime:  "nodejs",
				Checksum: "foobar",
			},
			Token: "token",
			Response: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "NoChange",
			Snapshot: snapshotFile,
			Function: &state.Function{
				Meta:     state.Meta{Name: "foo"},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:  "nodejs",
				Checksum: "foo",
			},
			Response: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := backend.NewTestKV()
			if test.Snapshot != "" {
				kv.LoadSnapshot(snapshotFile)
			}

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

			snapshot.AssertString(t, kv.Snapshot(), fmt.Sprintf("testdata/TestPutFunction-%s.yaml", test.TestName), *update)
		})
	}
}

func TestConfirmUpload(t *testing.T) {
	snapshotFile := "testdata/TestConfirmUpload.yaml"
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := state.PutModel(ctx, kv, state.ModelTypeFunction, &state.Function{
			Meta:           state.Meta{Name: "foo"},
			AWS:            &state.FunctionAWS{Timeout: 3, Memory: 256},
			Runtime:        "go",
			SourceFilename: "previous.tar.gz",
			Checksum:       "foo",
		})
		require.NoError(t, err)
		err = state.PutPendingUpload(ctx, kv, "new", &state.PendingUpload{
			Filename: "new.tar.gz",
			Function: &state.Function{
				Meta:     state.Meta{Name: "new"},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:  "go",
				Checksum: "new",
			},
		})
		require.NoError(t, err)
		err = state.PutPendingUpload(ctx, kv, "foo-config", &state.PendingUpload{
			Filename: "foo.tar.gz",
			Function: &state.Function{
				Meta:     state.Meta{Name: "foo"},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 1024},
				Runtime:  "nodejs",
				Checksum: "foo",
			},
		})
		require.NoError(t, err)
		err = state.PutPendingUpload(ctx, kv, "foo-code", &state.PendingUpload{
			Filename: "bar.tar.gz",
			Function: &state.Function{
				Meta:     state.Meta{Name: "bar"},
				AWS:      &state.FunctionAWS{Timeout: 3, Memory: 256},
				Runtime:  "go",
				Checksum: "updated",
			},
		})
		require.NoError(t, err)
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

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
			Token:    "foo-config",
		},
		{
			TestName: "UpdateCode",
			Token:    "foo-code",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := backend.NewTestKV()
			kv.LoadSnapshot(snapshotFile)

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

			snapshot.AssertString(t, kv.Snapshot(), fmt.Sprintf("testdata/TestConfirmUpload-%s.yaml", test.TestName), *update)
		})
	}
}

func TestCreateEnvironment(t *testing.T) {
	snapshotFile := "testdata/TestCreateEnvironment.yaml"
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := state.PutModel(ctx, kv, state.ModelTypeEnvironment, &state.Environment{
			Meta:           state.Meta{Name: "foo"},
			Infrastructure: state.InfrastructureTypeAWS,
		})
		require.NoError(t, err)
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

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
			kv := backend.NewTestKV()
			kv.LoadSnapshot(snapshotFile)

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

			snapshot.AssertString(t, kv.Snapshot(), fmt.Sprintf("testdata/TestCreateEnvironment-%s-state.yaml", test.TestName), *update)
			snapshot.AssertString(t, secretsKV.Snapshot(), fmt.Sprintf("testdata/TestCreateEnvironment-%s-secrets.yaml", test.TestName), *update)
		})
	}
}

func TestPutDeployment(t *testing.T) {
	snapshotFile := "testdata/TestPutDeployment.yaml"
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := state.PutModel(ctx, kv, state.ModelTypeDeployment, &state.Deployment{
			Meta:              state.Meta{Name: "foo"},
			EnvironmentLabels: map[string]string{},
			FunctionLabels:    map[string]string{},
		})
		require.NoError(t, err)
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

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
			kv := backend.NewTestKV()
			kv.LoadSnapshot(snapshotFile)

			s := &Server{
				StateStore: kv,
			}

			err := s.PutDeployment(ctx, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			snapshot.AssertString(t, kv.Snapshot(), fmt.Sprintf("testdata/TestPutDeployment-%s-state.yaml", test.TestName), *update)
		})
	}
}
