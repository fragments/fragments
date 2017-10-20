package state

import (
	"context"
	"encoding/json"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// PutModel creates or updates a generic model.
func PutModel(ctx context.Context, kv backend.Writer, modelType ModelType, model Model) error {
	if model == nil {
		return errors.New("model is nil")
	}
	name := model.Name()
	key, err := modelPath(modelType, name)
	if err != nil {
		return errors.Wrap(err, "could not construct model key")
	}

	raw, err := json.Marshal(model)
	if err != nil {
		return errors.Wrap(err, "could not marshal model")
	}

	if err := kv.Put(ctx, key, string(raw)); err != nil {
		return errors.Wrap(err, "could not store model")
	}

	return nil
}

// PutPendingUpload stores an upload request in the backend. It is used for
// confirming the upload.
func PutPendingUpload(ctx context.Context, kv backend.Writer, token string, upload *PendingUpload) error {
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
