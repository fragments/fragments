package client

import (
	"strings"

	"github.com/pkg/errors"
)

// CheckDuplicates checks that the list of resources doesn't have one or more
// resources with the same type and name
func CheckDuplicates(resources []Resource) error {
	// resMap is a map of type -> name -> files
	resMap := make(map[ResourceType]map[string][]string)

	for _, r := range resources {
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
