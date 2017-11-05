package backend

import (
	"context"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("test.update", false, "update test snapshots")

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

func TestTestKVSnapshot(t *testing.T) {
	values := make(map[string]string)
	values["simple"] = "foo"
	values["multiline"] = "bar\nbaz\n\nbaz"
	values["json"] = `{"foo":"foo","bar":123}`
	values["jsonNested"] = `{"foo":"foo","bar":{"foo":"foo","baz":"baz"}}`

	snapshotFile := "testdata/TestKVSnapshot.yaml"

	if *update {
		kv := NewTestKV()
		for k, v := range values {
			err := kv.Put(context.Background(), k, v)
			require.NoError(t, err)
		}
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

	kv := NewTestKV()
	kv.LoadSnapshot(snapshotFile)
	for k, expected := range values {
		t.Run(k, func(t *testing.T) {
			actual, err := kv.Get(context.Background(), k)
			require.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}

	expected, err := ioutil.ReadFile(snapshotFile)
	require.NoError(t, err)
	actual := kv.Snapshot()
	assert.Equal(t, string(expected), actual)
}
