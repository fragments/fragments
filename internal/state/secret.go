package state

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
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
func PutUserCredentials(ctx context.Context, kv backend.Writer, name, user, pass string) error {
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

// UserAWSCredentials returns AWS credentials for a user. The username assigned
// to credentials is used as the access key, the password is the secret.
func UserAWSCredentials(ctx context.Context, kv backend.Reader, name string) (*credentials.Credentials, error) {
	userPath := fmt.Sprintf("user/%s/%s", name, keySecretUser)
	passPath := fmt.Sprintf("user/%s/%s", name, keySecretPass)

	var id string
	var secret string

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		v, err := kv.Get(ctx, userPath)
		if err != nil {
			return errors.Wrap(err, "user")
		}
		id = v
		return nil
	})
	g.Go(func() error {
		v, err := kv.Get(ctx, passPath)
		if err != nil {
			return errors.Wrap(err, "pass")
		}
		secret = v
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, errors.Wrap(err, "could not read credentials")
	}

	return credentials.NewStaticCredentials(id, secret, ""), nil
}
