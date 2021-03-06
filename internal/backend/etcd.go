package backend

import (
	"context"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
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

// List lists keys in ETCD that have root as a prefix.
func (e *ETCD) List(ctx context.Context, root string) (map[string]string, error) {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}
	res, err := e.client.Get(ctx, root, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	out := make(map[string]string)
	for _, kv := range res.Kvs {
		key := string(kv.Key)
		key = strings.TrimPrefix(key, root)
		out[key] = string(kv.Value)
	}

	return out, nil
}

// Lock creates a new distributes lock on a key. Any future locks on the same
// key block until the lock is released. The lock is released if the context is
// cancelled.
//
// When no longer needed, the lock must be unlocked by calling the returned
// unlock function.
//
// Note: if the context is cancelled due to a timeout, this will still wait to
// acquire the lock. The lock will timeout automatically by a value set on the
// ETCD server (default 60 seconds).
func (e *ETCD) Lock(ctx context.Context, key string) (func(), error) {
	ses, err := concurrency.NewSession(e.client, concurrency.WithContext(ctx))
	if err != nil {
		return nil, errors.Wrap(err, "could not get etcd session for lock")
	}
	lock := concurrency.NewLocker(ses, key)
	lock.Lock()
	return lock.Unlock, nil
}

// Close closes the connection to ETCD.
func (e *ETCD) Close() error {
	return e.client.Close()
}
