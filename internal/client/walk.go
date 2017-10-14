package client

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Walk walks a target directory looking for models. Returns a list of
// potential model definitions.
// Directories can be ignored by supplying patterns to ignoreDirs. See
// https://golang.org/pkg/path/filepath/#Match for details.
func Walk(dir string, ignoreDirs []string) ([]string, error) {
	out := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// An unknown error occurred when trying to read the dir
			return err
		}

		if path == dir {
			// Skip self
			return nil
		}

		if info.IsDir() {
			dirname := filepath.Base(path)
			if strings.HasPrefix(dirname, ".") {
				// Skip hidden directory
				return filepath.SkipDir
			}

			for _, pattern := range ignoreDirs {
				match, err := filepath.Match(pattern, dirname)
				if err != nil {
					return errors.Wrap(err, pattern)
				}

				if match {
					// skip ignored directory
					return filepath.SkipDir
				}
			}

			// Continue into directory
			return nil
		}

		filename := filepath.Base(path)

		// skip hidden files
		if strings.HasPrefix(filename, ".") {
			// skip hidden file
			return nil
		}

		// Determine if file is a function config
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".json" && ext != ".yml" && ext != ".yaml" {
			// Extension doesn't match a model config's extension
			return nil
		}

		out = append(out, path)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}
