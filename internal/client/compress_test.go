package client

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		TestName string
		Files    []string
		Error    bool
		Content  map[string]string
	}{
		{
			TestName: "No files",
			Error:    true,
		},
		{
			TestName: "File not found",
			Files:    []string{"nonexisting.txt"},
			Error:    true,
		},
		{
			TestName: "Ok",
			Files: []string{
				"testdata/compress/file1.txt",
				"testdata/compress/file2.txt",
			},
			Content: map[string]string{
				"file1.txt": "foo\n",
				"file2.txt": "bar\n",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			// Compress archive
			r, err := Compress(test.Files)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// Verify contents
			gzf, err := gzip.NewReader(r)
			require.NoError(t, err)

			tar := tar.NewReader(gzf)
			n := 0
			for {
				tf, err := tar.Next()
				if err == io.EOF {
					break
				}

				expected, ok := test.Content[tf.Name]
				require.True(t, ok, "expected %s but not found in archive", tf.Name)

				actual, err := ioutil.ReadAll(tar)
				require.NoError(t, err)

				assert.Equal(t, expected, string(actual), "contents do not match for %s", tf.Name)

				n++
			}
			assert.Equal(t, len(test.Content), n, "number of files do not match")
		})
	}

}
