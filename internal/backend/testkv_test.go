package backend

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NOTE(akupila): megacheck doesn't seem to see that this is used in the tests
// and reports U1000. It is used, so we'll disable the linter.
// nolint: megacheck
type testkvValidator struct {
	data map[string]string
}

func (kv *testkvValidator) put(t *testing.T, key, value string) {
	t.Helper()
	kv.data[key] = value
}

func (kv *testkvValidator) get(t *testing.T, key string) (string, bool) {
	t.Helper()
	val, ok := kv.data[key]
	return val, ok
}

func TestTestKVCopy(t *testing.T) {
	initial := NewTestKV()
	ctx := context.Background()
	_ = initial.Put(ctx, "foo", "foo")
	assert.Len(t, initial.Data, 1)
	copied := initial.Copy()
	assert.Len(t, copied.Data, 1)
	_ = copied.Put(ctx, "bar", "bar")
	assert.Len(t, initial.Data, 1)
	assert.Len(t, copied.Data, 2)
	_ = initial.Put(ctx, "baz", "baz")
	_ = initial.Put(ctx, "foobarbaz", "foobarbaz")
	assert.Len(t, initial.Data, 3)
	assert.Len(t, copied.Data, 2)
}
