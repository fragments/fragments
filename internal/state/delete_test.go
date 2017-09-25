package state

import (
	"context"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeletePendingUpload(t *testing.T) {
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
				err := PutPendingUpload(ctx, kv, test.Token, test.Existing)
				require.NoError(t, err)
			}
			err := DeletePendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			actual, err := GetPendingUpload(ctx, kv, test.Token)
			require.NoError(t, err)
			assert.Nil(t, actual)
		})
	}
}
