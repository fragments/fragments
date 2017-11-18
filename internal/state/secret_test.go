package state

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/pkg/testutils"
	"github.com/stretchr/testify/assert"
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

			testutils.AssertGolden(
				t,
				testutils.SnapshotStringMap(kv.Data),
				fmt.Sprintf("testdata/TestPutUserCredentials-%s.yaml", test.TestName),
			)
		})
	}
}

func TestUserAWSCredentials(t *testing.T) {
	ctx := context.Background()
	kv := backend.NewTestKV()
	err := PutUserCredentials(ctx, kv, "foo", "AWSACCESSKEY", "AWSSECRETKEY")
	require.NoError(t, err)

	tests := []struct {
		TestName        string
		Snapshot        string
		Name            string
		AccessKeyID     string
		SecretAccessKey string
		Error           bool
	}{
		{
			TestName: "NotFound",
			Name:     "bar",
			Error:    true,
		},
		{
			TestName:        "Found",
			Name:            "foo",
			AccessKeyID:     "AWSACCESSKEY",
			SecretAccessKey: "AWSSECRETKEY",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			actual, err := UserAWSCredentials(ctx, kv, test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			creds, err := actual.Get()
			require.NoError(t, err)
			assert.Equal(t, test.AccessKeyID, creds.AccessKeyID)
			assert.Equal(t, test.SecretAccessKey, creds.SecretAccessKey)
		})
	}
}
