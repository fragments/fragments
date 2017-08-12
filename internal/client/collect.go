package client

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// CollectSource collects source files belonging to a function.
// Directories and files that match the ignore patterns are excluded.
func CollectSource(dir string, ignore []string) ([]string, error) {
	files := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// An unknown error occurred when trying to read the dir
			return errors.Wrapf(err, "could not read %s", path)
		}

		if path == dir {
			// Skip self
			return nil
		}

		skip := false
		for _, i := range ignore {
			if strings.Contains(path, i) {
				skip = true
			}
		}

		if info.IsDir() {
			if skip {
				return filepath.SkipDir
			}
			// Continue into directory
			return nil
		}

		files = append(files, path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
