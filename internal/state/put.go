package state

import (
	"context"
	"encoding/json"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// PutResource creates or updates a generic resource.
func PutResource(ctx context.Context, kv backend.KV, resourceType ResourceType, resource Resource) error {
	if resource == nil {
		return errors.New("resource is nil")
	}
	name := resource.Name()
	key, err := resourcePath(resourceType, name)
	if err != nil {
		return errors.Wrap(err, "could not construct resource key")
	}

	raw, err := json.Marshal(resource)
	if err != nil {
		return errors.Wrap(err, "could not marshal resource")
	}

	if err := kv.Put(ctx, key, string(raw)); err != nil {
		return errors.Wrap(err, "could not store resource")
	}

	return nil
}

// PutPendingUpload stores an upload request in the backend. It is used for
// confirming the upload.
func PutPendingUpload(ctx context.Context, kv backend.KV, token string, upload *PendingUpload) error {
	if upload == nil {
		return errors.New("pending upload is nil")
	}
	if upload.Function == nil {
		return errors.New("pending upload must have a function")
	}
	if upload.Filename == "" {
		return errors.New("pending upload must have a filename")
	}

	key, err := uploadPath(token)
	if err != nil {
		return errors.Wrap(err, "could not get upload key")
	}

	raw, err := json.Marshal(upload)
	if err != nil {
		return errors.Wrap(err, "could not marshal pending upload")
	}

	if err := kv.Put(ctx, key, string(raw)); err != nil {
		return errors.Wrap(err, "could not store pending upload")
	}

	return nil
}
