package state

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/require"
)

func TestDeletePendingUpload(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutPendingUpload(ctx, kv, "token", &PendingUpload{
			Filename: "filename.tar.gz",
			Function: &Function{
				Meta: Meta{Name: "foo"},
			},
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "TestDeletePendingUpload.json")
	}

	tests := []struct {
		TestName string
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
			Token:    "token",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			kv := backend.NewTestKV("TestDeletePendingUpload.json")
			ctx := context.Background()
			err := DeletePendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			kv.AssertSnapshot(t, fmt.Sprintf("TestDeletePendingUpload-%s.json", test.TestName), *update)
		})
	}
}
