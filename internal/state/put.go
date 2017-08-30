package state

import (
	"context"
	"encoding/json"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// PutPendingUpload stores an upload request in the backend. It is used for
// confirming the upload.
func PutPendingUpload(ctx context.Context, kv backend.KV, token string, upload *PendingUpload) error {
	if token == "" {
		return errors.New("token not set")
	}

	if upload == nil {
		return errors.New("pending upload is nil")
	}

	raw, err := json.Marshal(upload)
	if err != nil {
		return errors.Wrap(err, "could not marshal pending upload")
	}

	key := uploadPath(token)
	if err := kv.Put(ctx, key, string(raw)); err != nil {
		return errors.Wrap(err, "could not store pending upload")
	}

	return nil
}

// PutFunction creates or updates a function.
func PutFunction(ctx context.Context, kv backend.KV, function *Function) error {
	if function == nil {
		return errors.New("function is nil")
	}

	name := function.Meta.Name
	if name == "" {
		return errors.New("name not set")
	}

	raw, err := json.Marshal(function)
	if err != nil {
		return errors.Wrap(err, "could not marshal function")
	}

	key := resourcePath(ResourceTypeFunction, name)
	if err := kv.Put(ctx, key, string(raw)); err != nil {
		return errors.Wrap(err, "could not store function")
	}

	return nil
}

// PutEnvironment creates or updates an environment.
func PutEnvironment(ctx context.Context, kv backend.KV, env *Environment) error {
	if env == nil {
		return errors.New("environment is nil")
	}

	name := env.Meta.Name
	if name == "" {
		return errors.New("name not set")
	}

	raw, err := json.Marshal(env)
	if err != nil {
		return errors.Wrap(err, "could not marshal environment")
	}

	key := resourcePath(ResourceTypeEnvironment, name)
	if err := kv.Put(ctx, key, string(raw)); err != nil {
		return errors.Wrap(err, "could not store environment")
	}

	return nil
}
