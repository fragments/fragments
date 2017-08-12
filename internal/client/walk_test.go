package client

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWalk(t *testing.T) {
	tests := []struct {
		TestName string
		Dir      string
		Ignore   []string
		Error    bool
		Expected []string
	}{
		{
			TestName: "Not found",
			Dir:      "nonexisting",
			Error:    true,
		},
		{
			TestName: "Invalid ignore",
			Dir:      "_testdata/walk",
			Ignore: []string{
				"[",
			},
			Error: true,
		},
		{
			TestName: "Walk",
			Dir:      "_testdata/walk",
			Expected: []string{
				"_testdata/walk/root.yml",
				"_testdata/walk/foo/foo.yml",
				"_testdata/walk/foo/baz.yml",
				"_testdata/walk/foo/bar/bar1.yml",
				"_testdata/walk/foo/bar/bar2.json",
			},
		},
		{
			TestName: "Ignore",
			Dir:      "_testdata/walk",
			Ignore: []string{
				"bar",
			},
			Expected: []string{
				"_testdata/walk/root.yml",
				"_testdata/walk/foo/foo.yml",
				"_testdata/walk/foo/baz.yml",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := Walk(test.Dir, &WalkOptions{Ignore: test.Ignore})
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			sort.Strings(test.Expected)
			sort.Strings(actual)
			assert.EqualValues(t, test.Expected, actual)
		})
	}
}
