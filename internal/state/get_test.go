package state

import (
	"context"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFunction(t *testing.T) {
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
		kv.SaveSnapshot(t, "TestGetFunction.json")
	}

	tests := []struct {
		TestName string
		Name     string
		Expected *Function
		Error    bool
	}{
		{
			TestName: "No name",
			Name:     "",
			Expected: nil,
			Error:    true,
		},
		{
			TestName: "Not found",
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
			kv := backend.NewTestKV("TestGetFunction.json")
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
		kv.SaveSnapshot(t, "TestGetPendingUpload.json")
	}

	tests := []struct {
		TestName string
		Token    string
		Expected *PendingUpload
		Error    bool
	}{
		{
			TestName: "No name",
			Token:    "",
			Expected: nil,
			Error:    true,
		},
		{
			TestName: "Not found",
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
			kv := backend.NewTestKV("TestGetPendingUpload.json")
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
		kv.SaveSnapshot(t, "TestGetEnvironment.json")
	}

	tests := []struct {
		TestName string
		Name     string
		Expected *Environment
		Error    bool
	}{
		{
			TestName: "No name",
			Name:     "",
			Expected: nil,
			Error:    true,
		},
		{
			TestName: "Not found",
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
			kv := backend.NewTestKV("TestGetEnvironment.json")
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
		kv.SaveSnapshot(t, "TestGetDeployment.json")
	}

	tests := []struct {
		TestName string
		Name     string
		Expected *Deployment
		Error    bool
	}{
		{
			TestName: "No name",
			Name:     "",
			Expected: nil,
			Error:    true,
		},
		{
			TestName: "Not found",
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
			kv := backend.NewTestKV("TestGetDeployment.json")
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
