package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/fragments/fragments/internal/filestore"
	"github.com/fragments/fragments/internal/state"
	"github.com/pkg/errors"
)

const lambdaResource state.ResourceType = "lambda"

const (
	defaultLambdaRoleName        = "fragments-default-lambda-role"
	defaultLambdaRoleDescription = "Default Fragments Lambda execute role"
	defaultLambdaRolePath        = "/fragments/"
)

type lambdaService interface {
	CreateFunctionWithContext(aws.Context, *lambda.CreateFunctionInput, ...request.Option) (*lambda.FunctionConfiguration, error)
}

type lambdaReconciler struct {
	store       store
	source      filestore.SourceReader
	svcProvider serviceProvider
}

func newLambda(store store, source filestore.SourceReader, svcProvider serviceProvider) *lambdaReconciler {
	return &lambdaReconciler{
		store:       store,
		source:      source,
		svcProvider: svcProvider,
	}
}

func (l *lambdaReconciler) putFunction(ctx context.Context, input *state.Function) (*lambda.FunctionConfiguration, error) {
	res := state.Resource(state.InfrastructureTypeAWS, lambdaResource, input.Meta.Name)
	unlock, err := res.Lock(ctx, l.store)
	if err != nil {
		return nil, errors.Wrap(err, "could not acquire lock for lambda")
	}
	defer unlock()

	var existing lambda.FunctionConfiguration
	exists, err := res.Get(ctx, l.store, &existing)
	if err != nil {
		return nil, errors.Wrap(err, "could not check existing lambda")
	}

	if exists {
		f, err := l.update(ctx, existing, input)
		if err != nil {
			return nil, errors.Wrap(err, "update")
		}
		return f, nil
	}

	// TODO(akupila): allow passing in/replacing default role
	iam := newIAM(l.store, l.svcProvider)
	role, err := iam.putRole(ctx, &iamRoleInput{
		assumeRolePolicyDocument: mustCompress(defaultAssumeLambdaExecPolicy),
		description:              defaultLambdaRoleDescription,
		path:                     defaultLambdaRolePath,
		roleName:                 defaultLambdaRoleName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "errot putting lambda default role")
	}

	f, err := l.create(ctx, role, input)
	if err != nil {
		return nil, errors.Wrap(err, "create")
	}
	return f, nil
}

func (l *lambdaReconciler) create(ctx context.Context, role *iam.Role, input *state.Function) (*lambda.FunctionConfiguration, error) {
	return nil, errors.New("not implemented")
}

func (l *lambdaReconciler) update(ctx context.Context, existing lambda.FunctionConfiguration, input *state.Function) (*lambda.FunctionConfiguration, error) {
	return nil, errors.New("not implemented")
}
