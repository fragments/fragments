package server

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

func putUserCredentials(ctx context.Context, kv backend.KV, name, user, pass string) error {
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
