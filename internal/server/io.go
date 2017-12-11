package server

import (
	"context"
	"fmt"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/internal/model"
	"github.com/golang/sync/errgroup"
	"github.com/pkg/errors"
)

func functionPath(name string) string {
	return fmt.Sprintf("function/%s", name)
}

func deploymentPath(name string) string {
	return fmt.Sprintf("deployment/%s", name)
}

func environmentPath(name string) string {
	return fmt.Sprintf("environment/%s", name)
}

func pendingUploadPath(token string) string {
	return fmt.Sprintf("pendingupload/%s", token)
}

func userSecretName(name string) string {
	return fmt.Sprintf("user/%s/name", name)
}

func userSecretPass(pass string) string {
	return fmt.Sprintf("user/%s/pass", pass)
}

func putFunction(ctx context.Context, kv backend.Writer, f *model.Function) error {
	raw, err := model.MarshalFunction(f)
	if err != nil {
		return err
	}
	return kv.Put(ctx, functionPath(f.Name), string(raw))
}

func getFunction(ctx context.Context, kv backend.Reader, name string) (*model.Function, error) {
	raw, err := kv.Get(ctx, functionPath(name))
	if err != nil {
		if backend.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	var f model.Function
	if err := model.UnmarshalFunction([]byte(raw), &f); err != nil {
		return nil, err
	}
	return &f, nil
}

func putPendingUpload(ctx context.Context, kv backend.Writer, p *model.PendingUpload) error {
	raw, err := model.MarshalPendingUpload(p)
	if err != nil {
		return err
	}
	return kv.Put(ctx, pendingUploadPath(p.Token), string(raw))
}

func getPendingUpload(ctx context.Context, kv backend.Reader, token string) (*model.PendingUpload, error) {
	raw, err := kv.Get(ctx, pendingUploadPath(token))
	if err != nil {
		if backend.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	var p model.PendingUpload
	if err := model.UnmarshalPendingUpload([]byte(raw), &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func putEnvironment(ctx context.Context, kv backend.Writer, p *model.Environment) error {
	raw, err := model.MarshalEnvironment(p)
	if err != nil {
		return err
	}
	return kv.Put(ctx, environmentPath(p.Name), string(raw))
}

func putDeployment(ctx context.Context, kv backend.Writer, p *model.Deployment) error {
	raw, err := model.MarshalDeployment(p)
	if err != nil {
		return err
	}
	return kv.Put(ctx, deploymentPath(p.Name), string(raw))
}

func storeUserCredentials(ctx context.Context, kv backend.Writer, name, u, p string) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := kv.Put(ctx, userSecretName(name), u); err != nil {
			return errors.Wrap(err, "user")
		}
		return nil
	})
	g.Go(func() error {
		if err := kv.Put(ctx, userSecretPass(name), p); err != nil {
			return errors.Wrap(err, "pass")
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return errors.Wrap(err, "could not store credentials")
	}
	return nil
}
