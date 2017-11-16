package reconciler

import (
	"context"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/internal/filestore"
	"github.com/fragments/fragments/internal/state"
	"github.com/golang/sync/errgroup"
	"github.com/pkg/errors"
)

type statestore interface {
	backend.Reader
	backend.Writer
	backend.Lister
	backend.Locker
}

// Infrastructure is the implemented by structs that can reconcile resources
// for a specific infrastructure.
type Infrastructure interface {
	Reconcile(ctx context.Context, environment *state.Environment, function *state.Function) error
}

// Reconciler reads all current models and reconciles the desired state by
// trigger sub reconcilers for different infrastructures.
type Reconciler struct {
	StateStore  statestore
	SecretStore backend.Reader
	infra       map[state.InfraType]Infrastructure
}

// New creates a new reconciler.
func New(stateStore statestore, secretStore backend.Reader, sourceRepo filestore.SourceReader) *Reconciler {
	infra := make(map[state.InfraType]Infrastructure)

	return &Reconciler{
		StateStore:  stateStore,
		SecretStore: secretStore,
		infra:       infra,
	}
}

// Run runs a single reconciliation loop.
func (r *Reconciler) Run(ctx context.Context) error {
	deployments, err := state.ListDeployments(ctx, r.StateStore)
	if err != nil {
		return errors.Wrap(err, "could not list deployments")
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, d := range deployments {
		deployment := d
		g.Go(func() error {
			return r.reconcile(ctx, deployment)
		})
	}

	if err := g.Wait(); err != nil {
		if err != nil {
			return errors.Wrap(err, "failed to reconcile deployments")
		}
	}

	return nil
}

// reconcile reconciles a single deployment.
func (r *Reconciler) reconcile(ctx context.Context, deployment *state.Deployment) error {
	list, err := r.resolveModels(ctx, deployment)
	if err != nil {
		return errors.Wrap(err, "could not resolve resources")
	}

	g, ctx := errgroup.WithContext(ctx)

	// Loop through all environments
	for _, env := range list.environments {
		env := env
		targetInfra, ok := r.infra[env.Infrastructure]
		if !ok {
			return errors.Errorf("unsupported infrastructure %q", env.Infrastructure)
		}

		// Loop through all functions in the environment
		for _, function := range list.functions {
			function := function
			g.Go(func() error {
				if err := targetInfra.Reconcile(ctx, env, function); err != nil {
					return errors.Wrapf(err, "env: %s, func: %s", env.Infrastructure, function.Name())
				}
				return nil
			})
		}
	}

	return g.Wait()
}

// modelList contains the list of environments and functions that are part of a
// deployment.
type modelList struct {
	environments []*state.Environment
	functions    []*state.Function
}

// resolveModels resolves the list of environments and functions that are part
// of a deployment, based on the deployments label selectors.
func (r *Reconciler) resolveModels(ctx context.Context, deployment *state.Deployment) (*modelList, error) {
	g, ctx := errgroup.WithContext(ctx)
	out := &modelList{}

	g.Go(func() error {
		list, err := state.ListEnvironments(ctx, r.StateStore, &state.LabelMatcher{
			Labels: deployment.EnvironmentLabels,
		})
		if err != nil {
			return errors.Wrap(err, "error listing environments")
		}
		out.environments = list
		return nil
	})

	g.Go(func() error {
		list, err := state.ListFunctions(ctx, r.StateStore, &state.LabelMatcher{
			Labels: deployment.FunctionLabels,
		})
		if err != nil {
			return errors.Wrap(err, "error listing functions")
		}
		out.functions = list
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, errors.Wrap(err, "could not resolve deployment resources")
	}

	return out, nil
}
