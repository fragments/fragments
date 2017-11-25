package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// SnapshotStringMap takes a snapshot of m and encodes the result as json.
func SnapshotStringMap(m map[string]string) string {
	// Get list of keys so we can iterate over the keys in a determenistic
	// order
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b bytes.Buffer
	for _, k := range keys {
		v := m[k]
		b.WriteString(k)
		b.WriteString(": ")

		if v == "" {
			b.WriteString("<EMPTY>\n")
			continue
		}

		if strings.Contains(v, "\n") {
			b.WriteString("|\n    ")
			b.WriteString(strings.Replace(v, "\n", "\n    ", -1))
			b.WriteString("\n")
			continue
		}

		b.WriteString(v)
		b.WriteString("\n")
	}

	return b.String()
}

// SnapshotJSONMap takes a snapshot of a map where each value is json encoded.
// The values are pretty printed to a yaml output.
func SnapshotJSONMap(m map[string]string) string {
	// Get list of keys so we can iterate over the keys in a determenistic
	// order
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b bytes.Buffer
	for _, k := range keys {
		v := m[k]
		b.WriteString(k)
		b.WriteString(": |\n    ")

		if v == "" {
			b.WriteString("<EMPTY>\n")
			continue
		}

		if err := json.Indent(&b, []byte(v), "    ", "    "); err != nil {
			b.WriteString(fmt.Sprintf("<ERROR: %s>", err))
		}
		b.WriteString("\n")
	}

	return b.String()
}
