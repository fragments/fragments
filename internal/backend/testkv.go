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
	"unicode"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
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

// Snapshot returns a json encoded snapshot of the current state of the test
// backend. Panics if an error occurs.
// Nested json structures are pretty printed so they don't exactly match what's
// in the store. However, this is useful for test snapshots, which is the
// primary purpose of this anwyay.  The output is yaml compatible.
func (t *TestKV) Snapshot() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	var buf bytes.Buffer

	// Get list of keys so we can iterate over the keys in a determenistic
	// order
	keys := []string{}
	for k := range t.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(": ")

		processed := false
		v := t.Data[k]

		// Try pretty printing json
		var pretty bytes.Buffer
		err := json.Indent(&pretty, []byte(v), "    ", "    ")
		if err == nil {
			buf.WriteString("|\n    ")
			buf.Write(pretty.Bytes())
			processed = true
		}

		if !processed && strings.Contains(v, "\n") {
			buf.WriteString("|\n    ")
			buf.WriteString(strings.Join(strings.Split(v, "\n"), "\n    "))
			processed = true
		}

		if !processed {
			buf.WriteString(v)
		}

		buf.WriteString("\n")
	}

	return strings.TrimFunc(buf.String(), unicode.IsSpace) + "\n"
}

// LoadSnapshot loads a snapshot from disk and sets the state of the memory
// key-value store. This can be used for tests to set the initial state.
// Panics if the snapshot fails to load.
func (t *TestKV) LoadSnapshot(filename string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	temp := make(map[string]string)
	if err := yaml.Unmarshal(data, &temp); err != nil {
		log.Panic(errors.Wrapf(err, "could not unmarshal snapshot"))
	}
	for k, v := range temp {
		if json.Valid([]byte(v)) {
			// Try to recompress json which was pretty printed in snapshot
			var buf bytes.Buffer
			if err := json.Compact(&buf, []byte(v)); err == nil {
				t.Data[k] = buf.String()
				continue
			}
		}
		t.Data[k] = strings.TrimFunc(v, unicode.IsSpace)
	}
}
