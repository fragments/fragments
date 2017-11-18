package aws

import (
	"context"

	"github.com/pkg/errors"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/internal/filestore"
	"github.com/fragments/fragments/internal/state"
)

const (
	defaultRegion = "us-east-1"
)

type store interface {
	backend.ReaderWriter
	backend.Lister
	backend.Locker
}

// Reconciler reconciles Amazon Web Services resources.
type Reconciler struct {
	StateStore  store
	SecretStore backend.Reader
	SourceRepo  filestore.SourceReader
}

// Reconcile reconciles function to be executed on AWS Lambda.
func (r *Reconciler) Reconcile(ctx context.Context, environment *state.Environment, function *state.Function) error {
	userCreds, err := state.UserAWSCredentials(ctx, r.SecretStore, environment.Name())
	if err != nil {
		return errors.Wrapf(err, "could not get aws credentials for %s", environment.Name())
	}

	config := r.getConfig(environment)
	svcProvider := &defaultProvider{
		credentials: userCreds,
		region:      config.Region,
	}

	_ = svcProvider

	return errors.New("aws reconciler not implemented")
}

func (r *Reconciler) getConfig(environment *state.Environment) *state.InfrastructureAWS {
	awsConfig := environment.AWS
	if awsConfig == nil {
		awsConfig = &state.InfrastructureAWS{}
	}
	if awsConfig.Region == "" {
		awsConfig.Region = defaultRegion
	}
	return awsConfig
}
