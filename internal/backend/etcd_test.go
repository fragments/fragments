package backend

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testETCDEndpoint = fmt.Sprintf("127.0.0.1:%s", os.Getenv("ETCD_TEST_LISTEN_PORT"))

// testClient creates an ETCD client that can be used for preparing test data
func testClient(t *testing.T) (*clientv3.Client, func()) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{testETCDEndpoint},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		t.Fatal(errors.Wrap(err, "could not create test ETCD client"))
	}
	return cli, func() {
		if err := cli.Close(); err != nil {
			t.Fatal(errors.Wrap(err, "could not close test ETCD client"))
		}
	}
}

func testClean(t *testing.T, client *clientv3.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if _, err := client.Delete(ctx, "", clientv3.WithPrefix()); err != nil {
		t.Fatal(errors.Wrap(err, "could not clean test ETCD db"))
	}
}

// testGet gets a value for verificataion purposes
func testGet(t *testing.T, client *clientv3.Client, key string) string {
	ctx := context.Background()
	res, err := client.Get(ctx, key)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "could not get test value: %s", key))
	}
	if res.Count < 1 {
		t.Fatalf("test get did not return any values: %s", key)
	}
	return string(res.Kvs[0].Value)
}

// testPut prepares the test ETCD database by inserting a test value
func testPut(t *testing.T, client *clientv3.Client, key, value string) {
	ctx := context.Background()
	if _, err := client.Put(ctx, key, value); err != nil {
		t.Fatal(errors.Wrapf(err, "could not put test value: %s", key))
	}
}

func TestNewETCDClient(t *testing.T) {
	tests := []struct {
		TestName  string
		Endpoints []string
		Error     bool
	}{
		{
			TestName:  "No endpoints",
			Endpoints: nil,
			Error:     true,
		},
		{
			TestName:  "No listener",
			Endpoints: []string{"127.0.0.1:1"},
			Error:     true,
		},
		{
			TestName:  "OK",
			Endpoints: []string{testETCDEndpoint},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			client, err := NewETCDClient(test.Endpoints, 250*time.Millisecond)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			defer client.Close()
		})
	}
}

func TestPut(t *testing.T) {
	tc, close := testClient(t)
	defer close()

	testClean(t, tc)
	testPut(t, tc, "bar", "bar")

	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	tests := []struct {
		TestName string
		Context  context.Context
		Key      string
		Value    string
		Error    bool
	}{
		{
			TestName: "Empty key",
			Context:  ctx,
			Key:      "",
			Value:    "empty",
			Error:    true,
		},
		{
			TestName: "Context canceled",
			Context:  ctxCanceled,
			Key:      "canceled",
			Value:    "context",
			Error:    true,
		},
		{
			TestName: "OK",
			Context:  ctx,
			Key:      "foo",
			Value:    "foo",
		},
		{
			TestName: "Overwrite",
			Context:  ctx,
			Key:      "bar",
			Value:    "foobar",
		},
	}

	client, err := NewETCDClient([]string{testETCDEndpoint}, 1*time.Second)
	require.NoError(t, err)
	defer client.Close()

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {

			err = client.Put(test.Context, test.Key, test.Value)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			actual := testGet(t, tc, test.Key)
			assert.Equal(t, test.Value, actual)
		})
	}
}

func TestGet(t *testing.T) {
	tc, close := testClient(t)
	defer close()

	testClean(t, tc)
	testPut(t, tc, "foo", "bar")

	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	tests := []struct {
		TestName string
		Context  context.Context
		Key      string
		Expected string
		Error    bool
		NotFound bool
	}{
		{
			TestName: "Context canceled",
			Context:  ctxCanceled,
			Key:      "foo",
			Error:    true,
		},
		{
			TestName: "Context canceled",
			Context:  ctxCanceled,
			Key:      "foo",
			Error:    true,
		},
		{
			TestName: "Not found",
			Context:  ctx,
			Key:      "baz",
			NotFound: true,
		},
		{
			TestName: "Found",
			Context:  ctx,
			Key:      "foo",
			Expected: "bar",
		},
	}

	client, err := NewETCDClient([]string{testETCDEndpoint}, 1*time.Second)
	require.NoError(t, err)
	defer client.Close()

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := client.Get(test.Context, test.Key)
			if test.Error || test.NotFound {
				require.Error(t, err)

				switch errors.Cause(err).(type) {
				case *ErrNotFound:
					assert.True(t, test.NotFound)
					assert.Contains(t, err.Error(), test.Key)
				default:
					assert.False(t, test.NotFound)
				}

				return
			}
			require.NoError(t, err)

			assert.EqualValues(t, test.Expected, actual)
		})
	}
}
