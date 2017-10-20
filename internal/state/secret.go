package state

import (
	"context"
	"fmt"

	"github.com/fragments/fragments/internal/backend"
	"github.com/golang/sync/errgroup"
	"github.com/pkg/errors"
)

const (
	keySecretUser = "username"
	keySecretPass = "password"
)

// PutUserCredentials stores the credentials for a user. The credentials are
// used to perform action on behalf on the user when applying changes to the
// infrastructure.
func PutUserCredentials(ctx context.Context, kv backend.KV, name, user, pass string) error {
	if name == "" {
		return errors.New("name must be set")
	}

	userPath := fmt.Sprintf("user/%s/%s", name, keySecretUser)
	passPath := fmt.Sprintf("user/%s/%s", name, keySecretPass)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := kv.Put(ctx, userPath, user); err != nil {
			return errors.Wrap(err, "user")
		}
		return nil
	})
	g.Go(func() error {
		if err := kv.Put(ctx, passPath, pass); err != nil {
			return errors.Wrap(err, "pass")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return errors.Wrap(err, "could not store credentials")
	}
	return nil
}
