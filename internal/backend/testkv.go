package backend

import (
	"context"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// TestKV keeps data in memory. It should only be used for unit tests.
type TestKV struct {
	Data  map[string]string
	mu    sync.Mutex
	locks map[string]sync.Locker
}

// NewTestKV creates a new in key-value backend for tests.
// Snapshots are loaded to set the initial state.
func NewTestKV() *TestKV {
	kv := &TestKV{
		Data:  make(map[string]string),
		locks: make(map[string]sync.Locker),
	}

	return kv
}

// Put adds or overwrites a key to the in memory store.
func (t *TestKV) Put(ctx context.Context, key, value string) error {
	if key == "" {
		return errors.New("key is empty")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	t.mu.Lock()
	t.Data[key] = value
	t.mu.Unlock()
	return nil
}

// Get returns a value from the in memory store.
func (t *TestKV) Get(ctx context.Context, key string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	t.mu.Lock()
	v, ok := t.Data[key]
	t.mu.Unlock()
	if !ok {
		return "", &NotFoundError{key}
	}
	return v, nil
}

// Delete removes a key from the in memory store.
func (t *TestKV) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	t.mu.Lock()
	_, ok := t.Data[key]
	t.mu.Unlock()
	if !ok {
		return &NotFoundError{key}
	}
	t.mu.Lock()
	delete(t.Data, key)
	t.mu.Unlock()
	return nil
}

// List lists keys in the test store.
func (t *TestKV) List(ctx context.Context, root string) (map[string]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}

	out := make(map[string]string)
	t.mu.Lock()
	for k, v := range t.Data {
		if strings.HasPrefix(k, root) {
			key := strings.TrimPrefix(k, root)
			out[key] = v
		}
	}
	t.mu.Unlock()

	return out, nil
}

// Lock locks a key on the test kv. The key is locked for concurrent access
// until unlocked by calling the returned function.
func (t *TestKV) Lock(ctx context.Context, key string) (func(), error) {
	if err := ctx.Err(); err != nil {
		return func() {}, err
	}

	t.mu.Lock()
	locker, exists := t.locks[key]
	if !exists {
		locker = new(sync.Mutex)
		t.locks[key] = locker
	}
	t.mu.Unlock()

	locker.Lock()

	return locker.Unlock, nil
}

// Copy returns a copy of the TestKV, including its data. This is meant for
// unit tests, where a single instance of the TestKV is created, and copies of
// it are mutated.
func (t *TestKV) Copy() *TestKV {
	t.mu.Lock()
	defer t.mu.Unlock()
	newData := make(map[string]string)
	for k, v := range t.Data {
		newData[k] = v
	}
	return &TestKV{
		Data: newData,
	}
}
