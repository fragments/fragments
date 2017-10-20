package client

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChecksum(t *testing.T) {
	tests := []struct {
		TestName string
		Files    []string
		Ignore   []string
		Error    bool
		Expected string
	}{
		{
			TestName: "Nonexisting file",
			Files:    []string{"nonexisting"},
			Error:    true,
		},
		{
			TestName: "Empty list",
			Error:    true,
		},
		{
			TestName: "Ignored all",
			Files: []string{
				"testdata/checksum/a.txt",
			},
			Ignore: []string{
				"a.txt",
			},
			Error: true,
		},
		{
			TestName: "Files",
			Files: []string{
				"testdata/checksum/a.txt",
				"testdata/checksum/b.txt",
			},
			// cat $(find $DIR -type f | sort) | shasum
			Expected: "928f1abc474008fb586d62bcc3e043c5e3acb2aa",
		},
		{
			TestName: "Files",
			Files: []string{
				"testdata/checksum/a.txt",
				"testdata/checksum/b.txt",
			},
			Ignore: []string{
				"b.txt",
			},
			// cat $(find . -type f | grep -v b.txt | sort) | shasum
			Expected: "e242ed3bffccdf271b7fbaf34ed72d089537b42f",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := Checksum(test.Files, test.Ignore)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			hex := hex.EncodeToString(actual)
			assert.Equal(t, test.Expected, hex)
		})
	}
}
