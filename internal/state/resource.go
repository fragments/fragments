package state

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// Now returns the current timestamp. Normally this is time.Now, but it can be
// overridden for deterministic tests.
var Now = func() time.Time {
	return time.Now()
}

// ResourceType defines the type of a resource.
type ResourceType string

// ResPointer a pointer to a resource.
type ResPointer struct {
	// InfraType is the infrastructure type for the resource.
	InfraType InfraType
	// ResourceType specifies what type of resource this is.
	ResourceType ResourceType
	// Name is the unique name for the resource.
	Name string
}

// The envelope is the resource data, wrapped with additional information. It
// is stored in the backend using the ResPointer.
type envelope struct {
	Data    json.RawMessage `json:"data"`
	Created time.Time       `json:"created"`
	Updated time.Time       `json:"updated"`
}

func (r *ResPointer) key() string {
	return fmt.Sprintf("/resources/%s/%s/%s", r.InfraType, r.ResourceType, r.Name)
}

// Get reads and returns the underlying data for a resource.
// Returns false if the resource does not exist.
// The resource data is unmarshalled to data, the type of data should match the
// type that was used to write the resource data in Put.
func (r *ResPointer) Get(ctx context.Context, kv backend.Reader, data interface{}) (bool, error) {
	e, err := r.getEnvelope(ctx, kv)
	if err != nil {
		return false, errors.Wrap(err, "read resource")
	}
	if e == nil {
		return false, nil
	}
	if err := json.Unmarshal(e.Data, &data); err != nil {
		return true, errors.Wrap(err, "could not unmarshal resource data")
	}
	return true, nil
}

// Put creates or updates a resource. The data is specific to the resource
// implementation, and is wrapped in a Resource type, along with meta data.
// If the resource doesn't exist, it is created.
// If the resource exists, the data and updated timestamp are updated.
func (r *ResPointer) Put(ctx context.Context, kv backend.ReaderWriter, data interface{}) error {
	e, err := r.getEnvelope(ctx, kv)
	if err != nil {
		return errors.Wrap(err, "check existing resource")
	}
	now := Now()
	if e == nil {
		e = &envelope{
			Created: now,
		}
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "could not marshal resource data")
	}
	e.Data = dataJSON
	e.Updated = now
	raw, err := json.Marshal(e)
	if err != nil {
		return errors.Wrap(err, "could not marshal resource")
	}
	if err := kv.Put(ctx, r.key(), string(raw)); err != nil {
		return errors.Wrap(err, "could not store resource")
	}
	return nil
}

// Lock locks a resource, preventing concurrent modifications to it. Until the
// resource is unlocked, further locks block until the resource is unlocked.
// When done, the called must call the returned function to unlock the
// resource.
func (r *ResPointer) Lock(ctx context.Context, locker backend.Locker) (func(), error) {
	return locker.Lock(ctx, r.key())
}

func (r *ResPointer) getEnvelope(ctx context.Context, kv backend.Reader) (*envelope, error) {
	raw, err := kv.Get(ctx, r.key())
	if err != nil {
		if backend.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "could not get resource")
	}
	var res envelope
	if err := json.Unmarshal([]byte(raw), &res); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal resource")
	}
	return &res, nil
}
