package snapshot

import (
	"io/ioutil"
	"testing"

	"github.com/andreyvit/diff"
)

// AssertString asserts that a snapshot matches. If the update flag is passed,
// the snapshot is updated on disk.
func AssertString(t *testing.T, data, filename string, update bool) {
	t.Helper()
	if update {
		if err := ioutil.WriteFile(filename, []byte(data), 0644); err != nil {
			t.Errorf("could not write snapshot: %s", err)
		}
		return
	}
	expected, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("could not read snapshot: %s", err)
		return
	}
	if string(expected) != data {
		t.Errorf("snapshot does not match:\n%v", diff.LineDiff(string(expected), data))
		t.Fail()
	}
}
