package state

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/pkg/snapshot"
	"github.com/stretchr/testify/require"
)

func TestPutUserCredentials(t *testing.T) {
	tests := []struct {
		TestName string
		Name     string
		Username string
		Password string
		Error    bool
	}{
		{
			TestName: "NoName",
			Error:    true,
		},
		{
			TestName: "Ok",
			Name:     "test",
			Username: "foo",
			Password: "bar",
		},
	}

	ctx := context.Background()

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			kv := backend.NewTestKV()
			err := PutUserCredentials(ctx, kv, test.Name, test.Username, test.Password)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			snapshot.AssertString(t, kv.TestString(), fmt.Sprintf("testdata/TestPutUserCredentials-%s.txt", test.TestName), *update)
		})
	}
}