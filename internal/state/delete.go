package state

import (
	"context"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
)

// DeletePendingUpload deletes a pending upload. After this the upload cannot
// be confirmed any more.
func DeletePendingUpload(ctx context.Context, kv backend.KV, token string) error {
	key := uploadPath(token)
	err := kv.Delete(ctx, key)
	if err != nil {
		return errors.Wrap(err, "could not delete pending upload")
	}
	return nil
}
