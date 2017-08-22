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
