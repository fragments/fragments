package backend

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
)

// ETCD is a wrapper around the ETCD client that implements the backend.KV
// interface.
type ETCD struct {
	client *clientv3.Client
}

// NewETCDClient returns a new ETCD v3 client. Returns an error in case the
// connection to all endpoints fails.
func NewETCDClient(endpoints []string, dialTimeout time.Duration) (*ETCD, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to etcd")
	}
	return &ETCD{
		client: cli,
	}, nil
}

// Put stores a value in ETCD. The key is created if it doesn't exist, if it
// exists it is overwritten. Any watchers on the key will be notified.
func (e *ETCD) Put(ctx context.Context, key, value string) error {
	if _, err := e.client.Put(ctx, key, value); err != nil {
		return errors.Wrap(err, "could not put key")
	}
	return nil
}

// Get retrieves values from ETCD. Returns NotFoundError in case the key does not
// exist.
func (e *ETCD) Get(ctx context.Context, key string) (string, error) {
	res, err := e.client.Get(ctx, key, clientv3.WithLimit(1))
	if err != nil {
		return "", errors.Wrapf(err, "could not get key: %s", key)
	}
	if res.Count < 1 {
		return "", &NotFoundError{key}
	}

	return string(res.Kvs[0].Value), nil
}

// Delete deletes a key from ETCD. Returns NotFoundError in case the key does not
// exist.
func (e *ETCD) Delete(ctx context.Context, key string) error {
	res, err := e.client.Delete(ctx, key)
	if err != nil {
		return errors.Wrapf(err, "could not delete key: %s", key)
	}
	if res.Deleted < 1 {
		return &NotFoundError{key}
	}
	return nil
}

// Close closes the connection to ETCD.
func (e *ETCD) Close() error {
	return e.client.Close()
}
