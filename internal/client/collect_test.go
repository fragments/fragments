package client

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectSource(t *testing.T) {
	tests := []struct {
		TestName string
		Dir      string
		Ignore   []string
		Error    bool
		Expected []string
	}{
		{
			TestName: "Nonexisting",
			Dir:      "nonexisting",
			Error:    true,
		},
		{
			TestName: "Valid",
			Dir:      "testdata/collect",
			Expected: []string{
				"testdata/collect/file.js",
				"testdata/collect/.hidden",
				"testdata/collect/sub/file.js",
				"testdata/collect/node_modules/module.js",
			},
		},
		{
			TestName: "Ignore",
			Dir:      "testdata/collect",
			Ignore:   []string{"node_modules"},
			Expected: []string{
				"testdata/collect/file.js",
				"testdata/collect/.hidden",
				"testdata/collect/sub/file.js",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := CollectSource(test.Dir, test.Ignore)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			sort.Strings(test.Expected)
			sort.Strings(actual)
			assert.Equal(t, test.Expected, actual)
		})
	}
}
