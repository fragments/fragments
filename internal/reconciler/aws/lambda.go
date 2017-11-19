package aws

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
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
	defaultLambdaMemory          = 128
	defaultLambdaTimeout         = 3
)

type lambdaData struct {
	FunctionConfiguration *lambda.FunctionConfiguration
	CodeChecksum          string
}

type lambdaService interface {
	CreateFunctionWithContext(aws.Context, *lambda.CreateFunctionInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	UpdateFunctionConfigurationWithContext(ctx aws.Context, input *lambda.UpdateFunctionConfigurationInput, opts ...request.Option) (*lambda.FunctionConfiguration, error)
	UpdateFunctionCodeWithContext(ctx aws.Context, input *lambda.UpdateFunctionCodeInput, opts ...request.Option) (*lambda.FunctionConfiguration, error)
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

	var existing lambdaData
	exists, err := res.Get(ctx, l.store, &existing)
	if err != nil {
		return nil, errors.Wrap(err, "could not check existing lambda")
	}

	if exists {
		f, err := l.update(ctx, &existing, input)
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
	res := state.Resource(state.InfrastructureTypeAWS, lambdaResource, input.Meta.Name)
	svc, err := l.svcProvider.lambda()
	if err != nil {
		return nil, err
	}
	zip, err := l.getSourceZip(input)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get lambda source")
	}
	funcConfig, err := svc.CreateFunctionWithContext(
		ctx,
		&lambda.CreateFunctionInput{
			Code: &lambda.FunctionCode{
				ZipFile: zip,
			},
			// DeadLetterConfig
			Description: aws.String("test fragments function"),
			// Environment (for env vars)
			FunctionName: aws.String(input.Meta.Name),
			Handler:      aws.String(l.getHandler(input)),
			// KMSKeyArn
			MemorySize: aws.Int64(l.getMemory(input)),
			Publish:    aws.Bool(true),
			Role:       role.Arn,
			Runtime:    aws.String(l.getRuntime(input)),
			Tags:       l.getTags(input),
			Timeout:    aws.Int64(l.getTimeout(input)),
			// TracingConfig,
			// VpcConfig
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not create lambda")
	}
	err = res.Put(ctx, l.store, &lambdaData{
		FunctionConfiguration: funcConfig,
		CodeChecksum:          input.Checksum,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not store created lambda")
	}
	return funcConfig, nil
}

func (l *lambdaReconciler) update(ctx context.Context, existing *lambdaData, input *state.Function) (*lambda.FunctionConfiguration, error) {
	// Update config first
	if *existing.FunctionConfiguration.Timeout != input.AWS.Timeout ||
		*existing.FunctionConfiguration.MemorySize != input.AWS.Memory {
		if err := l.updateConfig(ctx, existing, input); err != nil {
			return nil, errors.Wrap(err, "could not update function config")
		}
	}
	// Update code
	if existing.CodeChecksum != input.Checksum {
		if err := l.updateCode(ctx, existing, input); err != nil {
			return nil, errors.Wrap(err, "could not update function code")
		}
	}
	return existing.FunctionConfiguration, nil
}

// updateConfig updates the configuration for a lambda function. The underlying
// lambdaData is updated in place with the new config. The updated
// configuration is stored in the backend store before returning.
func (l *lambdaReconciler) updateConfig(ctx context.Context, data *lambdaData, input *state.Function) error {
	res := state.Resource(state.InfrastructureTypeAWS, lambdaResource, input.Meta.Name)
	svc, err := l.svcProvider.lambda()
	if err != nil {
		return err
	}
	funcConfig, err := svc.UpdateFunctionConfigurationWithContext(
		ctx,
		&lambda.UpdateFunctionConfigurationInput{
			// DeadLetterConfig
			Description: aws.String("test fragments function"),
			// Environment (for env vars)
			FunctionName: data.FunctionConfiguration.FunctionName,
			Handler:      aws.String(l.getHandler(input)),
			// KMSKeyArn
			MemorySize: aws.Int64(l.getMemory(input)),
			Role:       data.FunctionConfiguration.Role,
			Runtime:    aws.String(l.getRuntime(input)),
			Timeout:    aws.Int64(l.getTimeout(input)),
			// TracingConfig,
			// VpcConfig
		},
	)
	if err != nil {
		return errors.Wrap(err, "could not update lambda config")
	}
	data.FunctionConfiguration = funcConfig
	if err := res.Put(ctx, l.store, data); err != nil {
		return errors.Wrap(err, "could not store update lambda config")
	}
	return nil
}

func (l *lambdaReconciler) updateCode(ctx context.Context, data *lambdaData, input *state.Function) error {
	res := state.Resource(state.InfrastructureTypeAWS, lambdaResource, input.Meta.Name)
	svc, err := l.svcProvider.lambda()
	if err != nil {
		return err
	}
	zip, err := l.getSourceZip(input)
	if err != nil {
		return errors.Wrapf(err, "could not get lambda source")
	}
	funcConfig, err := svc.UpdateFunctionCodeWithContext(
		ctx,
		&lambda.UpdateFunctionCodeInput{
			DryRun:       aws.Bool(false),
			FunctionName: data.FunctionConfiguration.FunctionName,
			Publish:      aws.Bool(true),
			ZipFile:      zip,
		},
	)
	if err != nil {
		return errors.Wrap(err, "could not update lambda code")
	}
	data.FunctionConfiguration = funcConfig
	data.CodeChecksum = input.Checksum
	if err := res.Put(ctx, l.store, data); err != nil {
		return errors.Wrap(err, "could not store updated lambda code")
	}
	return nil
}

func (l *lambdaReconciler) getSourceZip(input *state.Function) ([]byte, error) {
	f, err := l.source.GetFile(input.SourceFilename)
	if err != nil {
		return nil, errors.Wrap(err, "error reading source")
	}
	gzf, err := gzip.NewReader(f)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create gzip reader")
	}
	tf := tar.NewReader(gzf)
	b := &bytes.Buffer{}
	if err := filestore.Zip(tf, b); err != nil {
		return nil, errors.Wrap(err, "could not re-compress tar to zip")
	}
	if err := f.Close(); err != nil {
		return nil, errors.Wrap(err, "could not close file stream")
	}
	if err := gzf.Close(); err != nil {
		return nil, errors.Wrap(err, "could not close gzip stream")
	}
	return b.Bytes(), nil
}

func (l *lambdaReconciler) getTimeout(input *state.Function) int64 {
	if input.AWS != nil && input.AWS.Timeout != 0 {
		return input.AWS.Timeout
	}
	return defaultLambdaTimeout
}

func (l *lambdaReconciler) getMemory(input *state.Function) int64 {
	if input.AWS != nil && input.AWS.Memory != 0 {
		return input.AWS.Memory
	}
	return defaultLambdaMemory
}

func (l *lambdaReconciler) getHandler(input *state.Function) string {
	// TODO(akupila): set handler in model
	return "index.handler"
}

func (l *lambdaReconciler) getRuntime(input *state.Function) string {
	runtime := input.Runtime
	if runtime == "nodejs" {
		runtime = "nodejs6.10"
	}
	return runtime
}

func (l *lambdaReconciler) getTags(input *state.Function) map[string]*string {
	out := make(map[string]*string)
	for k, v := range input.Meta.Labels {
		out[k] = aws.String(v)
	}
	return out
}
