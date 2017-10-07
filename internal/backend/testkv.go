package backend

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// TestKV keeps data in memory. It should only be used for unit tests.
type TestKV struct {
	Data map[string]string
	mu   sync.Mutex
}

func (t *TestKV) snapshotFilename(key string) string {
	if !strings.HasPrefix(key, "testdata/") {
		key = "testdata/" + key
	}
	if !strings.HasSuffix(key, ".json") {
		key += ".json"
	}
	return strings.Replace(key, " ", "_", -1)
}

// NewTestKV creates a new in key-value backend for tests.
// Snapshots are loaded to set the initial state.
// The snapshots are loaded in the order they are passed in. In case snapshots
// contain duplicate keys the keys from the latter snapshots overwrite the
// earlier ones.
// Panics if a snapshot fails to load.
func NewTestKV(snapshots ...string) *TestKV {
	kv := &TestKV{
		Data: make(map[string]string),
	}

	for _, snapshot := range snapshots {
		filename := kv.snapshotFilename(snapshot)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Panic(err)
		}
		temp := make(map[string]string)
		if err := json.Unmarshal(data, &temp); err != nil {
			log.Panic(errors.Wrapf(err, "could not unmarshal snapshot %s", filename))
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
func (t *TestKV) SaveSnapshot(test *testing.T, key string) {
	test.Helper()
	data := t.Snapshot(test)
	filename := t.snapshotFilename(key)
	if err := ioutil.WriteFile(filename, []byte(data), 0644); err != nil {
		test.Fatal(err)
	}
}

// AssertSnapshot asserts that the state of the test store matches a previously
// generated snapshot.
func (t *TestKV) AssertSnapshot(test *testing.T, key string, update bool) {
	test.Helper()
	if update {
		t.SaveSnapshot(test, key)
	}
	actual := t.Snapshot(test)
	filename := t.snapshotFilename(key)
	expected, err := ioutil.ReadFile(filename)
	if err != nil {
		test.Log(err)
		test.Fail()
	}
	assert.Equal(test, string(expected), actual)
}
