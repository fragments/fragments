package backend

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTestKV(t *testing.T) {
	tests := []struct {
		TestName  string
		Snapshots []string
		Assert    func(data map[string]string)
	}{
		{
			TestName: "No snapshots",
			Assert: func(data map[string]string) {
				assert.Len(t, data, 0)
			},
		},
		{
			TestName: "Single snapshot",
			Snapshots: []string{
				"testdata/test_foobar.json",
			},
			Assert: func(data map[string]string) {
				assert.Len(t, data, 2)
				assert.Equal(t, data["foo"], "foo")
				assert.Equal(t, data["bar"], "bar")
			},
		},
		{
			TestName: "Multiple snapshots",
			Snapshots: []string{
				"testdata/test_foobar.json",
				"testdata/test_baz.json",
			},
			Assert: func(data map[string]string) {
				assert.Len(t, data, 3)
				assert.Equal(t, data["foo"], "foo")
				assert.Equal(t, data["bar"], "bar")
				assert.Equal(t, data["baz"], "baz")
			},
		},
		{
			TestName: "Overwrite key",
			Snapshots: []string{
				"testdata/test_foobar.json",
				"testdata/test_baz.json",
				"testdata/test_foobarbaz.json",
			},
			Assert: func(data map[string]string) {
				assert.Len(t, data, 3)
				assert.Equal(t, data["foo"], "foo")
				assert.Equal(t, data["bar"], "bar")
				assert.Equal(t, data["baz"], "foobarbaz")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			kv := NewTestKV(test.Snapshots...)
			test.Assert(kv.Data)
		})
	}
}

func TestNewTestKVInvalid(t *testing.T) {
	assert.Panics(t, func() {
		NewTestKV("invalid")
	})
}

func TestTestKVString(t *testing.T) {
	kv := NewTestKV("testdata/testkv-string.json")
	actual := kv.TestString()
	expected, err := ioutil.ReadFile("testdata/testkv-string.golden")
	require.NoError(t, err)
	assert.Equal(t, string(expected), actual)
}
