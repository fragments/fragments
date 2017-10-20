package state

import (
	"context"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// DeletePendingUpload deletes a pending upload. After this the upload cannot
// be confirmed any more.
func DeletePendingUpload(ctx context.Context, kv backend.Writer, token string) error {
	key, err := uploadPath(token)
	if err != nil {
		return errors.Wrap(err, "could not get upload key")
	}

	if err := kv.Delete(ctx, key); err != nil {
		return errors.Wrap(err, "could not delete pending upload")
	}
	return nil
}
