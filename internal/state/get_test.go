package state

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFunction(t *testing.T) {
	snapshotFile := "testdata/TestGetFunction.yaml"
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutModel(ctx, kv, ModelTypeFunction, &Function{
			Meta: Meta{
				Name: "foo",
			},
			Runtime:  "go",
			Checksum: "abc",
		})
		require.NoError(t, err)
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		TestName string
		Name     string
		Expected *Function
		Error    bool
	}{
		{
			TestName: "NoName",
			Name:     "",
			Expected: nil,
			Error:    true,
		},
		{
			TestName: "NotFound",
			Name:     "bar",
			Expected: nil,
			Error:    false,
		},
		{
			TestName: "Found",
			Name:     "foo",
			Expected: &Function{
				Meta: Meta{
					Name: "foo",
				},
				Runtime:  "go",
				Checksum: "abc",
			},
			Error: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := backend.NewTestKV()
			kv.LoadSnapshot(snapshotFile)
			actual, err := GetFunction(ctx, kv, test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.EqualValues(t, test.Expected, actual)
		})
	}
}

func TestGetPendingUpload(t *testing.T) {
	snapshotFile := "testdata/TestGetPendingUpload.yaml"
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutPendingUpload(ctx, kv, "token", &PendingUpload{
			Filename: "foo.tar.gz",
			Function: &Function{
				Meta: Meta{
					Name: "foo",
				},
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
		Expected *PendingUpload
		Error    bool
	}{
		{
			TestName: "NoName",
			Token:    "",
			Expected: nil,
			Error:    true,
		},
		{
			TestName: "NotFound",
			Token:    "baz",
			Expected: nil,
			Error:    false,
		},
		{
			TestName: "Found",
			Token:    "token",
			Expected: &PendingUpload{
				Filename: "foo.tar.gz",
				Function: &Function{
					Meta: Meta{
						Name: "foo",
					},
				},
			},
			Error: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := backend.NewTestKV()
			kv.LoadSnapshot(snapshotFile)
			actual, err := GetPendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.EqualValues(t, test.Expected, actual)
		})
	}
}

func TestGetEnvironment(t *testing.T) {
	snapshotFile := "testdata/TestGetEnvironment.yaml"
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutModel(ctx, kv, ModelTypeEnvironment, &Environment{
			Meta: Meta{
				Name: "foo",
			},
			Infrastructure: InfrastructureTypeAWS,
		})
		require.NoError(t, err)
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		TestName string
		Name     string
		Expected *Environment
		Error    bool
	}{
		{
			TestName: "NoName",
			Name:     "",
			Expected: nil,
			Error:    true,
		},
		{
			TestName: "NotFound",
			Name:     "bar",
			Expected: nil,
			Error:    false,
		},
		{
			TestName: "Found",
			Name:     "foo",
			Expected: &Environment{
				Meta: Meta{
					Name: "foo",
				},
				Infrastructure: InfrastructureTypeAWS,
			},
			Error: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := backend.NewTestKV()
			kv.LoadSnapshot(snapshotFile)
			actual, err := GetEnvironment(ctx, kv, test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestGetDeployment(t *testing.T) {
	snapshotFile := "testdata/TestGetDeployment.yaml"
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutModel(ctx, kv, ModelTypeDeployment, &Deployment{
			Meta: Meta{
				Name: "foo",
			},
			EnvironmentLabels: map[string]string{
				"foo": "foo",
			},
			FunctionLabels: map[string]string{
				"bar": "bar",
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
		Name     string
		Expected *Deployment
		Error    bool
	}{
		{
			TestName: "NoName",
			Name:     "",
			Expected: nil,
			Error:    true,
		},
		{
			TestName: "NotFound",
			Name:     "bar",
			Expected: nil,
			Error:    false,
		},
		{
			TestName: "Found",
			Name:     "foo",
			Expected: &Deployment{
				Meta: Meta{
					Name: "foo",
				},
				EnvironmentLabels: map[string]string{
					"foo": "foo",
				},
				FunctionLabels: map[string]string{
					"bar": "bar",
				},
			},
			Error: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := backend.NewTestKV()
			kv.LoadSnapshot(snapshotFile)
			actual, err := GetDeployment(ctx, kv, test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, test.Expected, actual)
		})
	}
}
