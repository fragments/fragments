package testutils

import "testing"

func TestSnapshotStringMap(t *testing.T) {
	m := map[string]string{
		"b":   "foo",
		"a":   "bar",
		"c":   "",
		"err": "foo\nbar\nbaz",
	}
	AssertGolden(t, SnapshotStringMap(m), "testdata/snapshot-stringmap.yaml")
}

func TestSnapshotJSONMap(t *testing.T) {
	m := map[string]string{
		"b":   `{"a":"a","b":"b","c":"c"}`,
		"a":   `{"foo":"foo","bar":{"baz":123}}`,
		"c":   ``,
		"err": `{invalid json`,
	}
	AssertGolden(t, SnapshotJSONMap(m), "testdata/snapshot-jsonmap.yaml")
}
