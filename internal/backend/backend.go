//go:generate mockery -name KV

package backend

import (
	"context"
	"fmt"
)

// KV  interface is implemented by backends that persistently store key-value
// data
type KV interface {
	// Put inserts a value under a key. In case the value already exists it is
	// overwritten. The key is case-insensitive.
	Put(ctx context.Context, key, value string) error
	// Get gets a value. In case the key does not exist, returns ErrNotFound if
	// the key was not found
	Get(ctx context.Context, key string) (string, error)
}

// ErrNotFound represents an error of a key that was not found
type ErrNotFound struct {
	Key string
}

// Error returns the error string for a not found error
func (e *ErrNotFound) Error() string { return fmt.Sprintf("key not found: %s", e.Key) }
