package state

import (
	"context"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFunction(t *testing.T) {
	existing := &Function{
		Meta: Meta{
			Name: "foo",
		},
	}

	ctx := context.Background()
	kv := backend.NewMemoryKV()
	err := PutModel(ctx, kv, ModelTypeFunction, existing)
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
			actual, err := GetFunction(ctx, kv, test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestGetPendingUpload(t *testing.T) {
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
	err := PutPendingUpload(ctx, kv, "token", existing)
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
			actual, err := GetPendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.Expected, actual)
		})
	}
}

func TestGetEnvironment(t *testing.T) {
	existing := &Environment{
		Meta: Meta{
			Name: "foo",
		},
		Infrastructure: InfrastructureTypeAWS,
	}

	ctx := context.Background()
	kv := backend.NewMemoryKV()
	err := PutModel(ctx, kv, ModelTypeEnvironment, existing)
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
	existing := &Deployment{
		Meta: Meta{
			Name: "foo",
		},
	}

	ctx := context.Background()
	kv := backend.NewMemoryKV()
	err := PutModel(ctx, kv, ModelTypeDeployment, existing)
	require.NoError(t, err)

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
			Name:     existing.Name(),
			Expected: existing,
			Error:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
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
