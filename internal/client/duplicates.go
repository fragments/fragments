package client

import (
	"strings"

	"github.com/pkg/errors"
)

// CheckDuplicates checks that the list of models doesn't have one or more
// models with the same type and name
func CheckDuplicates(models []Model) error {
	// resMap is a map of type -> name -> files
	resMap := make(map[ModelType]map[string][]string)

	for _, r := range models {
		resType := r.Type()
		name := r.Meta().Name

		if resMap[resType] == nil {
			resMap[resType] = make(map[string][]string)
		}
		if resMap[resType][name] == nil {
			resMap[resType][name] = []string{}
		}
		resMap[resType][name] = append(resMap[resType][name], r.File())
	}

	for resType, names := range resMap {
		for name, files := range names {
			if len(files) > 1 {
				return errors.Errorf("duplicate %s definitions for %s:\n- %s", resType, name, strings.Join(files, "\n- "))
			}
		}
	}

	return nil
}
