package state

import (
	"context"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFunction(t *testing.T) {
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := PutModel(ctx, initial, ModelTypeFunction, &Function{
		Meta: Meta{
			Name: "existing",
		},
		Runtime:  "go",
		Checksum: "abc",
	})
	require.NoError(t, err)

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
			Name:     "nonexisting",
			Expected: nil,
			Error:    false,
		},
		{
			TestName: "Found",
			Name:     "existing",
			Expected: &Function{
				Meta: Meta{
					Name: "existing",
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
			kv := initial.Copy()
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
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := PutPendingUpload(ctx, initial, "existing", &PendingUpload{
		Filename: "existing.tar.gz",
		Function: &Function{
			Meta: Meta{
				Name: "existingfunc",
			},
		},
	})
	require.NoError(t, err)

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
			Token:    "nonexisting",
			Expected: nil,
			Error:    false,
		},
		{
			TestName: "Found",
			Token:    "existing",
			Expected: &PendingUpload{
				Filename: "existing.tar.gz",
				Function: &Function{
					Meta: Meta{
						Name: "existingfunc",
					},
				},
			},
			Error: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := initial.Copy()
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
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := PutModel(ctx, initial, ModelTypeEnvironment, &Environment{
		Meta: Meta{
			Name: "existing",
		},
		Infrastructure: InfrastructureTypeAWS,
	})
	require.NoError(t, err)

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
			Name:     "nonexisting",
			Expected: nil,
			Error:    false,
		},
		{
			TestName: "Found",
			Name:     "existing",
			Expected: &Environment{
				Meta: Meta{
					Name: "existing",
				},
				Infrastructure: InfrastructureTypeAWS,
			},
			Error: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			kv := initial.Copy()
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
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := PutModel(ctx, initial, ModelTypeDeployment, &Deployment{
		Meta: Meta{
			Name: "existing",
		},
		EnvironmentLabels: map[string]string{
			"foo": "foo",
		},
		FunctionLabels: map[string]string{
			"bar": "bar",
		},
	})
	require.NoError(t, err)

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
			Name:     "nonexisting",
			Expected: nil,
			Error:    false,
		},
		{
			TestName: "Found",
			Name:     "existing",
			Expected: &Deployment{
				Meta: Meta{
					Name: "existing",
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
			kv := initial.Copy()
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
