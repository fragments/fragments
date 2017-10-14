package snapshot

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssertString(t *testing.T) {
	mockT := new(testing.T)
	data := "foo"

	// Initially this should fail since the snapshot cannot be found.
	AssertString(mockT, data, "snapshot-test", false)
	assert.True(t, mockT.Failed())

	// Update snapshot without filename, should fail.
	mockT = new(testing.T)
	AssertString(mockT, data, "", true)
	assert.True(t, mockT.Failed())

	// Update snapshot, should not fail.
	mockT = new(testing.T)
	AssertString(mockT, data, "snapshot-test", true)
	assert.False(t, mockT.Failed())

	// Snapshot should match now.
	mockT = new(testing.T)
	AssertString(mockT, data, "snapshot-test", false)
	assert.False(t, mockT.Failed())

	// Snapshot should fail if data doesn't match.
	mockT = new(testing.T)
	AssertString(mockT, "bar", "snapshot-test", false)
	assert.True(t, mockT.Failed())

	// Clean up
	err := os.Remove("snapshot-test")
	require.NoError(t, err)
}
