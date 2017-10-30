package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"sync"
	"testing"

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
// The snapshots are loaded in the order they are passed in. In case snapshots
// contain duplicate keys the keys from the latter snapshots overwrite the
// earlier ones.
// Panics if a snapshot fails to load.
func NewTestKV(snapshots ...string) *TestKV {
	kv := &TestKV{
		Data:  make(map[string]string),
		locks: make(map[string]sync.Locker),
	}

	for _, snapshot := range snapshots {
		data, err := ioutil.ReadFile(snapshot)
		if err != nil {
			log.Panic(err)
		}
		temp := make(map[string]string)
		if err := json.Unmarshal(data, &temp); err != nil {
			log.Panic(errors.Wrapf(err, "could not unmarshal snapshot %s", snapshot))
		}
		for k, v := range temp {
			kv.Data[k] = v
		}
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

// Snapshot returns a json encoded snapshot of the current state of the test
// backend.
func (t *TestKV) Snapshot(test *testing.T) string {
	actual, err := json.MarshalIndent(t.Data, "", "\t")
	if err != nil {
		test.Fatal(err)
	}
	return string(actual)
}

// SaveSnapshot saves the current state of the test backend to a file. The
// state can be restored with LoadSnapshot. The data is stored as json.
func (t *TestKV) SaveSnapshot(test *testing.T, filename string) {
	test.Helper()
	data := t.Snapshot(test)
	if err := ioutil.WriteFile(filename, []byte(data), 0644); err != nil {
		test.Fatal(err)
	}
}

// TestString creates a determenistic human readable test string from the state
// of the TestKV store.
func (t *TestKV) TestString() string {
	var buf bytes.Buffer
	t.mu.Lock()
	keys := []string{}
	klen := 0
	for k := range t.Data {
		if len(k) > klen {
			klen = len(k)
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := t.Data[k]
		processed := false
		buf.WriteString(k)
		buf.WriteString(strings.Repeat(" ", klen-len(k)))
		buf.WriteString(" : ")

		// try json indenting the value
		var pretty bytes.Buffer
		err := json.Indent(&pretty, []byte(v), strings.Repeat(" ", klen+3), "    ")
		if err == nil {
			buf.Write(pretty.Bytes())
			processed = true
		}

		if !processed {
			buf.WriteString(v)
		}

		buf.WriteString("\n")
	}
	t.mu.Unlock()
	return buf.String()
}

// Lock locks a key on the test kv. The key is locked for concurrent access
// until unlocked by calling the returned function.
func (t *TestKV) Lock(ctx context.Context, key string) (func(), error) {
	if ctx.Done() != nil {
		return func() {}, ctx.Err()
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
