package state

import (
	"context"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutResource(t *testing.T) {
	ctx := context.Background()
	kv := backend.NewMemoryKV()

	tests := []struct {
		TestName     string
		Input        Resource
		ResourceType ResourceType
		Error        bool
	}{
		{
			TestName: "No resource",
			Error:    true,
		},
		{
			TestName: "No name",
			Input:    &Function{},
			Error:    true,
		},
		{
			TestName: "Function",
			Input: &Function{
				Meta: Meta{
					Name: "foo",
				},
			},
			ResourceType: ResourceTypeFunction,
		},
		{
			TestName: "Environment",
			Input: &Environment{
				Meta: Meta{
					Name: "foo",
				},
				Infrastructure: InfrastructureTypeAWS,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := PutResource(ctx, kv, test.ResourceType, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			switch test.ResourceType {
			case ResourceTypeFunction:
				actual, err := GetFunction(ctx, kv, test.Input.Name())
				require.NoError(t, err)
				assert.Equal(t, test.Input, actual)
			case ResourceTypeEnvironment:
				actual, err := GetEnvironment(ctx, kv, test.Input.Name())
				require.NoError(t, err)
				assert.Equal(t, test.Input, actual)
			}
		})
	}
}

func TestPutPendingUpload(t *testing.T) {
	ctx := context.Background()
	kv := backend.NewMemoryKV()

	function := &Function{
		Meta: Meta{
			Name: "foo",
		},
	}

	tests := []struct {
		TestName string
		Input    *PendingUpload
		Token    string
		Error    bool
	}{
		{
			TestName: "No pending upload",
			Error:    true,
		},
		{
			TestName: "No filename",
			Token:    "token",
			Input: &PendingUpload{
				Filename: "",
				Function: function,
			},
			Error: true,
		},
		{
			TestName: "No function",
			Token:    "token",
			Input: &PendingUpload{
				Filename: "foo.tar.gz",
				Function: nil,
			},
			Error: true,
		},
		{
			TestName: "No token",
			Input: &PendingUpload{
				Filename: "foo.tar.gz",
				Function: function,
			},
			Error: true,
		},
		{
			TestName: "Ok",
			Token:    "token",
			Input: &PendingUpload{
				Filename: "foo.tar.gz",
				Function: function,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := PutPendingUpload(ctx, kv, test.Token, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			actual, err := GetPendingUpload(ctx, kv, test.Token)
			require.NoError(t, err)
			assert.Equal(t, test.Input, actual)
		})
	}
}
