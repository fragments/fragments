package state

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/pkg/snapshot"
	"github.com/stretchr/testify/require"
)

func TestDeletePendingUpload(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutPendingUpload(ctx, kv, "foo", &PendingUpload{
			Filename: "foo.tar.gz",
			Function: &Function{
				Meta: Meta{Name: "foo"},
			},
		})
		require.NoError(t, err)
		err = PutPendingUpload(ctx, kv, "bar", &PendingUpload{
			Filename: "bar.tar.gz",
			Function: &Function{
				Meta: Meta{Name: "foo"},
			},
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "testdata/TestDeletePendingUpload.json")
	}

	tests := []struct {
		TestName string
		Token    string
		Error    bool
	}{
		{
			TestName: "NoToken",
			Error:    true,
		},
		{
			TestName: "NotFound",
			Token:    "invalid",
			Error:    true,
		},
		{
			TestName: "Deleted",
			Token:    "foo",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			kv := backend.NewTestKV("testdata/TestDeletePendingUpload.json")
			ctx := context.Background()
			err := DeletePendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			snapshot.AssertString(t, kv.TestString(), fmt.Sprintf("testdata/TestDeletePendingUpload-%s.txt", test.TestName), *update)
		})
	}
}
