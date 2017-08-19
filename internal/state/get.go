package state

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// ResourceType defines the type of a resource. It is used to group the same
// resources in the backend.
type ResourceType string

const (
	// ResourceTypeFunction is a function resource.
	ResourceTypeFunction ResourceType = "function"
)

// resourcePath constructs a path to store resources under in the backend.
func resourcePath(resType ResourceType, name string) string {
	return fmt.Sprintf("/resources/%s/%s", resType, name)
}

// GetFunction reads a function from the backend. Returns nil if the function
// does not exist.
func GetFunction(ctx context.Context, kv backend.KV, name string) (*Function, error) {
	key := resourcePath(ResourceTypeFunction, name)
	raw, err := kv.Get(ctx, key)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *backend.ErrNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, "could not get function from backend")
		}
	}

	var function Function
	if err := json.Unmarshal([]byte(raw), &function); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal function")
	}

	return &function, nil
}
