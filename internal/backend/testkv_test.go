package backend

import (
	"os"
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
				"test_foobar.json",
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
				"test_foobar.json",
				"test_baz.json",
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
				"test_foobar.json",
				"test_baz.json",
				"test_foobarbaz.json",
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

func TestAssertSnapshot(t *testing.T) {
	mockT := new(testing.T)
	kv := NewTestKV()

	// Initially this should fail since the snapshot cannot be found.
	kv.AssertSnapshot(mockT, "snapshot-test", false)
	assert.True(t, mockT.Failed())

	// Update snapshot, should not fail
	mockT = new(testing.T)
	kv.AssertSnapshot(mockT, "snapshot-test", true)
	assert.False(t, mockT.Failed())

	// Snapshot should match now
	mockT = new(testing.T)
	kv.AssertSnapshot(mockT, "snapshot-test", false)
	assert.False(t, mockT.Failed())

	// Clean up
	filename := kv.snapshotFilename("snapshot-test")
	err := os.Remove(filename)
	require.NoError(t, err)
}
