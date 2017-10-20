package filestore

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocal(t *testing.T) {
	base, err := ioutil.TempDir("", "fragments-test")
	require.NoError(t, err)
	uploads := filepath.Join(base, "uploads")
	source := filepath.Join(base, "source")

	// Invalid arguments
	_, err = NewLocal("", source)
	require.Error(t, err)
	_, err = NewLocal(uploads, "")
	require.Error(t, err)

	// Create local filestore
	local, err := NewLocal(uploads, source)
	require.NoError(t, err)

	// Generate url
	url, err := local.NewUploadURL("test")
	require.NoError(t, err)

	// Upload to url
	fixture, err := ioutil.ReadFile("_testdata/upload.txt")
	require.NoError(t, err)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(fixture))
	require.NoError(t, err)
	_, err = http.DefaultClient.Do(req)
	require.NoError(t, err)

	// Upload should be set
	_, err = os.Stat(filepath.Join(uploads, "test"))
	require.NoError(t, err)

	// Persist
	err = local.Persist(context.Background(), "test")
	require.NoError(t, err)

	// Assert
	actual, err := ioutil.ReadFile(filepath.Join(source, "test"))
	require.NoError(t, err)
	assert.Equal(t, string(fixture), string(actual))

	// Get file
	_, err = local.GetFile("nonexisting")
	require.Error(t, err)

	file, err := local.GetFile("test")
	require.NoError(t, err)
	data, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	assert.Equal(t, fixture, data)
	file.Close()

	err = local.Shutdown()
	require.NoError(t, err)
}
