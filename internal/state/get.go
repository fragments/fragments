package state

import (
	"context"
	"encoding/json"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

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
