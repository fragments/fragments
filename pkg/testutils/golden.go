package testutils

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

var updateGolden = flag.Bool("test.update-golden", false, "update golden files")

// AssertGolden checks that actual matches the contents of a golden file. The
// golden file is updated if -test.update-golden is set.
func AssertGolden(t *testing.T, actual, filename string) {
	if *updateGolden {
		if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
			t.Error(err)
			return
		}
		if err := ioutil.WriteFile(filename, []byte(actual), 0644); err != nil {
			t.Error(err)
			return
		}
		return
	}
	expected, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}
	if string(expected) == actual {
		return
	}
	diff, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(expected)),
		B:        difflib.SplitLines(actual),
		FromFile: "Expected",
		ToFile:   "Actual",
		Context:  3,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Errorf("snapshot %s does not match:\n%v", filename, diff)
	t.Fail()
}
