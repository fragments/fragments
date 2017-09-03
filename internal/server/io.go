package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// resourcePath constructs a path to store resources under in the backend.
func resourcePath(resType ResourceType, name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	return fmt.Sprintf("/resources/%s/%s", resType, name), nil
}

// uploadPath builds a path that's used for storing pending uploads.
func uploadPath(token string) (string, error) {
	if token == "" {
		return "", errors.New("token is required")
	}
	return fmt.Sprintf("/uploads/%s", token), nil
}

// getFunction reads a function from the backend. Returns nil if the function
// does not exist.
func getFunction(ctx context.Context, kv backend.KV, name string) (*Function, error) {
	key, err := resourcePath(ResourceTypeFunction, name)
	if err != nil {
		return nil, errors.Wrap(err, "could not get function key")
	}
	raw, err := kv.Get(ctx, key)
	if err != nil {
		if backend.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "could not get function from backend")
	}

	var function Function
	if err := json.Unmarshal([]byte(raw), &function); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal function")
	}

	return &function, nil
}

// source must not be considered valid. Returns nil if the pending upload does
// not exist.
func getPendingUpload(ctx context.Context, kv backend.KV, token string) (*PendingUpload, error) {
	key, err := uploadPath(token)
	if err != nil {
		return nil, errors.Wrap(err, "could not get upload key")
	}
	raw, err := kv.Get(ctx, key)
	if err != nil {
		if backend.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "could not get pending upload from backend")
	}

	var pendingUpload PendingUpload
	if err := json.Unmarshal([]byte(raw), &pendingUpload); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal pending upload")
	}

	return &pendingUpload, nil
}

// getEnvironment returns an environment. Returns nil if the environment does
// not exist.
func getEnvironment(ctx context.Context, kv backend.KV, name string) (*Environment, error) {
	key, err := resourcePath(ResourceTypeEnvironment, name)
	if err != nil {
		return nil, errors.Wrap(err, "could not get environment key")
	}
	raw, err := kv.Get(ctx, key)
	if err != nil {
		if backend.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "could not get environment from backend")
	}

	var env Environment
	if err := json.Unmarshal([]byte(raw), &env); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal environment")
	}

	return &env, nil
}

// getDeployment returns a deployment. Returns nil if the deployment does not
// exist.
func getDeployment(ctx context.Context, kv backend.KV, name string) (*Deployment, error) {
	key, err := resourcePath(ResourceTypeDeployment, name)
	if err != nil {
		return nil, errors.Wrap(err, "could not get deployment key")
	}
	raw, err := kv.Get(ctx, key)
	if err != nil {
		if backend.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "could not get deployment from backend")
	}

	var deploy Deployment
	if err := json.Unmarshal([]byte(raw), &deploy); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal deployment")
	}

	return &deploy, nil
}

// putResourcecreates or updates a generic resource.
func putResource(ctx context.Context, kv backend.KV, resourceType ResourceType, resource Resource) error {
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

// putPendingUpload stores an upload request in the backend. It is used for
// confirming the upload.
func putPendingUpload(ctx context.Context, kv backend.KV, token string, upload *PendingUpload) error {
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

// deletePendingUpload deletes a pending upload. After this the upload cannot
// be confirmed any more.
func deletePendingUpload(ctx context.Context, kv backend.KV, token string) error {
	key, err := uploadPath(token)
	if err != nil {
		return errors.Wrap(err, "could not get upload key")
	}

	if err := kv.Delete(ctx, key); err != nil {
		return errors.Wrap(err, "could not delete pending upload")
	}
	return nil
}
