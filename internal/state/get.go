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

// GetPendingUpload returns a pending upload. If the upload does not exist the
// source must not be considered valid. Returns nil if the pending upload does
// not exist.
func GetPendingUpload(ctx context.Context, kv backend.KV, token string) (*PendingUpload, error) {
	key := uploadPath(token)
	raw, err := kv.Get(ctx, key)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *backend.ErrNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, "could not get pending upload from backend")
		}
	}

	var pendingUpload PendingUpload
	if err := json.Unmarshal([]byte(raw), &pendingUpload); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal pending upload")
	}

	return &pendingUpload, nil
}

// GetEnvironment returns an environment. Returns nil if the environment does
// not exist.
func GetEnvironment(ctx context.Context, kv backend.KV, name string) (*Environment, error) {
	key := resourcePath(ResourceTypeEnvironment, name)
	raw, err := kv.Get(ctx, key)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *backend.ErrNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, "could not get environment from backend")
		}
	}

	var env Environment
	if err := json.Unmarshal([]byte(raw), &env); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal environment")
	}

	return &env, nil
}
