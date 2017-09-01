package backend

import (
	"context"
	"sync"

	"github.com/pkg/errors"
)

// MemoryKV keeps data in memory. It should only be used for unit tests.
type MemoryKV struct {
	Data map[string]string
	mu   sync.Mutex
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
	m.mu.Lock()
	m.Data[key] = value
	m.mu.Unlock()
	return nil
}

// Get returns a value from the in memory store.
func (m *MemoryKV) Get(ctx context.Context, key string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	m.mu.Lock()
	v, ok := m.Data[key]
	m.mu.Unlock()
	if !ok {
		return "", &NotFoundError{key}
	}
	return v, nil
}

// Delete removes a key from the in memory store.
func (m *MemoryKV) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	m.mu.Lock()
	_, ok := m.Data[key]
	m.mu.Unlock()
	if !ok {
		return &NotFoundError{key}
	}
	m.mu.Lock()
	delete(m.Data, key)
	m.mu.Unlock()
	return nil
}
