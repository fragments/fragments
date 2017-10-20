// +build integration

package backend

import (
	"context"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

// etcdValidator is used for integration tests to validate the actual state in
// vault after the ETCD client has interacted with it.
type etcdValidator struct {
	*clientv3.Client
}

func newETCDValidator(t *testing.T, endpoints ...string) *etcdValidator {
	t.Helper()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	require.NoError(t, err)
	if _, err := cli.Delete(context.Background(), "", clientv3.WithPrefix()); err != nil {
		t.Fatal(errors.Wrap(err, "could not clean test ETCD db"))
	}
	return &etcdValidator{cli}
}

func (e *etcdValidator) put(t *testing.T, key, value string) {
	t.Helper()
	_, err := e.Put(context.Background(), key, string(value))
	require.NoError(t, err)
}

func (e *etcdValidator) get(t *testing.T, key string) (string, bool) {
	t.Helper()
	res, err := e.Get(context.Background(), key)
	require.NoError(t, err)
	if res.Count < 1 {
		return "", false
	}
	return string(res.Kvs[0].Value), true
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
			defer func() {
				_ = client.Close()
			}()
		})
	}
}
