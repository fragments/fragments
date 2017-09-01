//go:generate mockery -name KV

package backend

import (
	"context"
	"fmt"
)

// KV  interface is implemented by backends that persistently store key-value
// data.
type KV interface {
	// Put inserts a value under a key. In case the value already exists it is
	// overwritten. The key is case-insensitive.
	Put(ctx context.Context, key, value string) error
	// Get gets a value. In case the key does not exist, returns NotFoundError if
	// the key was not found.
	Get(ctx context.Context, key string) (string, error)
	// Delete deletes a key. In case the key does not exist, returns NotFoundError
	// if the key was not found.
	Delete(ctx context.Context, key string) error
}

// NotFoundError indicates that a key was not found.
type NotFoundError struct {
	Key string
}

// Error returns the error string for a not found error.
func (e *NotFoundError) Error() string { return fmt.Sprintf("key not found: %s", e.Key) }

// IsNotFound returns true if the error returned is for a key that was not
// found.
func IsNotFound(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}
