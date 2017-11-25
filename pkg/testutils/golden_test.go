package testutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssertGolden(t *testing.T) {
	if *updateGolden {
		t.Skip("cannot run golden test with -test.update-golden")
	}
	mockT := &testing.T{}
	data := "foo"
	filename := "testdata/golden.test"

	// Initially this should fail since the snapshot cannot be found.
	AssertGolden(mockT, data, filename)
	assert.True(t, mockT.Failed())

	// Update snapshot without filename, should fail.
	mockT = &testing.T{}
	AssertGolden(mockT, data, "")
	assert.True(t, mockT.Failed())

	// Update golden file
	boolT := true
	boolF := false
	mockT = &testing.T{}
	updateGolden = &boolT
	AssertGolden(mockT, data, filename)
	assert.False(t, mockT.Failed())
	updateGolden = &boolF

	// Snapshot should match now.
	mockT = &testing.T{}
	AssertGolden(mockT, data, filename)
	assert.False(t, mockT.Failed())

	// Snapshot should fail if data doesn't match.
	mockT = &testing.T{}
	AssertGolden(mockT, "bar", filename)
	assert.True(t, mockT.Failed())

	// Clean up
	err := os.Remove(filename)
	require.NoError(t, err)
}
