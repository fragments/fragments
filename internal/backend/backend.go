package backend

import (
	"context"
	"fmt"
)

// Reader reads values from a backend.
type Reader interface {
	// Get gets a value. In case the key does not exist, returns NotFoundError
	// if the key was not found.
	Get(ctx context.Context, key string) (string, error)
}

// The Writer interface is implemented by a backend that can be written to.
type Writer interface {
	// Put inserts a value under a key. In case the value already exists it is
	// overwritten. The key is case-insensitive.
	Put(ctx context.Context, key, value string) error
	// Delete deletes a key. In case the key does not exist, returns NotFoundError
	// if the key was not found.
	Delete(ctx context.Context, key string) error
}

// The Lister interface is implemented by backends that can list keys under a
// key.
type Lister interface {
	// List lists keys directly under a root key.
	List(ctx context.Context, root string) (map[string]string, error)
}

// Locker locks resources, preventing multiple clients modifying the same
// resource in the backend.
type Locker interface {
	// Locker creates a distributed lock on a key.
	Lock(ctx context.Context, key string) (func(), error)
}

// NotFoundError indicates that a key was not found.
type NotFoundError struct {
	// Key is the key that was not found.
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
