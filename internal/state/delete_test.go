package state

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/pkg/testutils"
	"github.com/stretchr/testify/require"
)

func TestDeletePendingUpload(t *testing.T) {
	initial := backend.NewTestKV()
	ctx := context.Background()
	err := PutPendingUpload(ctx, initial, "foo", &PendingUpload{
		Filename: "foo.tar.gz",
		Function: &Function{
			Meta: Meta{Name: "foo"},
		},
	})
	require.NoError(t, err)
	err = PutPendingUpload(ctx, initial, "bar", &PendingUpload{
		Filename: "bar.tar.gz",
		Function: &Function{
			Meta: Meta{Name: "foo"},
		},
	})
	require.NoError(t, err)

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
			ctx := context.Background()
			kv := initial.Copy()
			err := DeletePendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			testutils.AssertGolden(
				t,
				testutils.SnapshotJSONMap(kv.Data),
				fmt.Sprintf("testdata/TestDeletePendingUpload-%s.yaml", test.TestName),
			)
		})
	}
}
