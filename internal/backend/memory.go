package backend

import (
	"context"

	"github.com/pkg/errors"
)

// MemoryKV keeps data in memory. It should only be used for unit tests.
type MemoryKV struct {
	Data map[string]string
}

// NewMemoryKV creates a new in memory KV backend
func NewMemoryKV() *MemoryKV {
	return &MemoryKV{
		Data: make(map[string]string),
	}
}

// Put adds or overwrites a key to the in memory store.
func (m *MemoryKV) Put(ctx context.Context, key, value string) error {
	if key == "" {
		return errors.New("key is empty")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	m.Data[key] = value
	return nil
}

// Get returns a value from the in memory store.
func (m *MemoryKV) Get(ctx context.Context, key string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	v, ok := m.Data[key]
	if !ok {
		return "", &ErrNotFound{
			Key: key,
		}
	}
	return v, nil
}

// Delete removes a key from the in memory store.
func (m *MemoryKV) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	_, ok := m.Data[key]
	if !ok {
		return &ErrNotFound{
			Key: key,
		}
	}
	delete(m.Data, key)
	return nil
}
