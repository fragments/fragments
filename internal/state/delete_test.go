package state

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/pkg/snapshot"
	"github.com/stretchr/testify/require"
)

func TestDeletePendingUpload(t *testing.T) {
	snapshotFile := "testdata/TestDeletePendingUpload.yaml"
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
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
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
			kv := backend.NewTestKV()
			kv.LoadSnapshot(snapshotFile)
			ctx := context.Background()
			err := DeletePendingUpload(ctx, kv, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			snapshot.AssertString(t, kv.Snapshot(), fmt.Sprintf("testdata/TestDeletePendingUpload-%s.yaml", test.TestName), *update)
		})
	}
}
