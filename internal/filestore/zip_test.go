package filestore

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZip(t *testing.T) {
	f, err := os.Open("testdata/ziptest.tar")
	require.NoError(t, err)
	defer func() {
		_ = f.Close()
	}()

	tf := tar.NewReader(f)

	buffer := &bytes.Buffer{}
	err = Zip(tf, buffer)
	require.NoError(t, err)

	data := bytes.NewReader(buffer.Bytes())

	zip, err := zip.NewReader(data, int64(buffer.Len()))
	require.NoError(t, err)

	// foo.txt
	// foo/
	//   bar.txt
	require.Len(t, zip.File, 3)
}
