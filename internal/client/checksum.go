package client

import (
	"crypto/sha1"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Checksum calculates the sha1 checksum of a the contents of a list of
// files. The files are processed in the order of the input so sort the files
// to ensure consistent hashing.
// Returns an error in case no files were provided or all files were filtered
// by the ignore.
func Checksum(files, ignore []string) ([]byte, error) {
	sha := sha1.New()

	added := 0
	for _, f := range files {
		matchedIgnore := false
		for _, i := range ignore {
			if strings.Contains(f, i) {
				matchedIgnore = true
			}
		}
		if matchedIgnore {
			continue
		}

		file, err := os.Open(f)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read file")
		}
		defer file.Close() // nolint: errcheck

		_, err = io.Copy(sha, file)
		if err != nil {
			return nil, errors.Wrap(err, "unable to add to hash")
		}

		added++
	}

	if added == 0 {
		return nil, errors.New("no files were included in checksum")
	}

	return sha.Sum(nil), nil
}
