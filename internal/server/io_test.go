package server

import (
	"context"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_getFunction(t *testing.T) {
	existing := &Function{
		Meta: Meta{
			Name: "foo",
		},
	}

	ctx := context.Background()
	kv := backend.NewMemoryKV()
	err := putResource(ctx, kv, ResourceTypeFunction, existing)
	require.NoError(t, err)

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
			Name:     existing.Name(),
			Expected: existing,
			Error:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := getFunction(ctx, kv, test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func Test_getPendingUpload(t *testing.T) {
	function := &Function{
		Meta: Meta{
			Name: "foo",
		},
		Runtime: "go",
	}
	existing := &PendingUpload{
		Filename:         "foo.tar.gz",
		PreviousFilename: "bar.tar.gz",
		Function:         function,
	}

	ctx := context.Background()
	kv := backend.NewMemoryKV()
	err := putPendingUpload(ctx, kv, "token", existing)
	require.NoError(t, err)

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
			Expected: existing,
			Error:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := getPendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func Test_getEnvironment(t *testing.T) {
	existing := &Environment{
		Meta: Meta{
			Name: "foo",
		},
		Infrastructure: InfrastructureTypeAWS,
	}

	ctx := context.Background()
	kv := backend.NewMemoryKV()
	err := putResource(ctx, kv, ResourceTypeEnvironment, existing)
	require.NoError(t, err)

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
			Name:     existing.Name(),
			Expected: existing,
			Error:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := getEnvironment(ctx, kv, test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func Test_putResource(t *testing.T) {
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
			err := putResource(ctx, kv, test.ResourceType, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			switch test.ResourceType {
			case ResourceTypeFunction:
				actual, err := getFunction(ctx, kv, test.Input.Name())
				require.NoError(t, err)
				assert.Equal(t, test.Input, actual)
			case ResourceTypeEnvironment:
				actual, err := getEnvironment(ctx, kv, test.Input.Name())
				require.NoError(t, err)
				assert.Equal(t, test.Input, actual)
			}
		})
	}
}

func Test_putPendingUpload(t *testing.T) {
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
			err := putPendingUpload(ctx, kv, test.Token, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			actual, err := getPendingUpload(ctx, kv, test.Token)
			require.NoError(t, err)
			assert.Equal(t, test.Input, actual)
		})
	}
}

func Test_deletePendingUpload(t *testing.T) {
	ctx := context.Background()
	kv := backend.NewMemoryKV()

	pendingUpload := &PendingUpload{
		Filename: "foo.tar.gz",
		Function: &Function{
			Meta: Meta{
				Name: "foo",
			},
		},
	}

	tests := []struct {
		TestName string
		Existing *PendingUpload
		Token    string
		Error    bool
	}{
		{
			TestName: "No token",
			Error:    true,
		},
		{
			TestName: "Not found",
			Token:    "foo",
			Error:    true,
		},
		{
			TestName: "Deleted",
			Existing: pendingUpload,
			Token:    "foo",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			if test.Existing != nil {
				err := putPendingUpload(ctx, kv, test.Token, test.Existing)
				require.NoError(t, err)
			}
			err := deletePendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			actual, err := getPendingUpload(ctx, kv, test.Token)
			require.NoError(t, err)
			assert.Nil(t, actual)
		})
	}
}
