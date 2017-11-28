package aws

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/fragments/fragments/internal/backend"
	mockfs "github.com/fragments/fragments/internal/filestore/mocks"
	"github.com/fragments/fragments/internal/state"
	"github.com/fragments/fragments/pkg/testutils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAWSLambdaCreate(t *testing.T) {
	kv := backend.NewTestKV()
	mockAWS := newMockAWS()
	mockFS := &mockfs.SourceReader{}
	clock := testutils.NewMockClock()
	lambdaReconciler := newLambda(kv, mockFS, mockAWS, clock)
	ctx := context.Background()

	src, err := os.Open("testdata/function.tar.gz")
	require.NoError(t, err)
	defer src.Close()
	mockFS.On("GetFile", "function.tar.gz").Return(src, nil)
	mockFS.On("GetFile", mock.Anything).Return(nil, errors.New("not found"))

	_, err = lambdaReconciler.putFunction(ctx, &state.Function{
		Meta: state.Meta{
			Name: "foo",
			Labels: map[string]string{
				"foo": "foo",
				"bar": "bar",
			},
		},
		Runtime:        "nodejs",
		Checksum:       "abc123",
		SourceFilename: "", // not found -> error
	})
	require.Error(t, err)

	_, err = lambdaReconciler.putFunction(ctx, &state.Function{
		Meta: state.Meta{
			Name: "foo",
			Labels: map[string]string{
				"foo": "foo",
				"bar": "bar",
			},
		},
		Runtime:        "nodejs",
		Checksum:       "abc123",
		SourceFilename: "function.tar.gz",
	})
	require.NoError(t, err)

	testutils.AssertGolden(
		t,
		testutils.SnapshotJSONMap(kv.Data),
		"testdata/TestAWSLambdaCreate-state.yaml",
	)
	testutils.AssertGolden(
		t,
		snapshotService(mockAWS.lambdaMock),
		"testdata/TestAWSLambdaCreate-aws.yaml",
	)
}

func TestAWSLambdaUpdate(t *testing.T) {
	kv := backend.NewTestKV()
	initial := &lambda.FunctionConfiguration{
		CodeSha256:       aws.String("20d3b32ab5658cd12984408ecb4c67b833cdbf8681b22a8cf6dc63b06ab6960a"),
		CodeSize:         aws.Int64(130),
		DeadLetterConfig: nil,
		Description:      aws.String("test fragments function"),
		Environment:      nil,
		FunctionArn:      aws.String("arn:aws:lambda::us-east-1:123456789000:function:foo"),
		FunctionName:     aws.String("foo"),
		Handler:          aws.String("index.handler"),
		KMSKeyArn:        nil,
		LastModified:     aws.String("2017-10-01T19:19:37.640+0000"),
		MasterArn:        aws.String("arn:aws:lambda::us-east-1:123456789000:function:foo"),
		MemorySize:       aws.Int64(128),
		Role:             aws.String("arn:aws:iam::123456789000:role/path/fragments-default-lambda-role"),
		Runtime:          aws.String("nodejs6.10"),
		Timeout:          aws.Int64(3),
		TracingConfig:    &lambda.TracingConfigResponse{},
		Version:          aws.String("$LATEST"),
		VpcConfig:        &lambda.VpcConfigResponse{},
	}

	mockAWS := newMockAWS()
	mockFS := &mockfs.SourceReader{}
	clock := testutils.NewMockClock()
	lambdaReconciler := newLambda(kv, mockFS, mockAWS, clock)
	ctx := context.Background()

	res := lambdaReconciler.pointer("foo")
	err := res.Put(context.Background(), kv, clock, &lambdaData{
		FunctionConfiguration: initial,
		CodeChecksum:          "abc",
	})
	require.NoError(t, err)

	src, err := os.Open("testdata/function.tar.gz")
	require.NoError(t, err)
	defer src.Close()
	mockFS.On("GetFile", "function.tar.gz").Return(src, nil)
	mockFS.On("GetFile", mock.Anything).Return(nil, errors.New("not found"))

	_, err = lambdaReconciler.putFunction(ctx, &state.Function{
		Meta: state.Meta{
			Name: "existing",
			Labels: map[string]string{
				"foo": "foo",
				"bar": "bar",
			},
		},
		Runtime:        "nodejs",
		Checksum:       "abc123",
		SourceFilename: "", // not found -> error
	})
	require.Error(t, err)

	_, err = lambdaReconciler.putFunction(ctx, &state.Function{
		Meta: state.Meta{
			Name: "existing",
			Labels: map[string]string{
				"foo": "foo",
				"bar": "bar",
			},
		},
		Runtime:        "nodejs",
		Checksum:       "abc123",
		SourceFilename: "function.tar.gz",
	})
	require.NoError(t, err)

	testutils.AssertGolden(
		t,
		testutils.SnapshotJSONMap(kv.Data),
		"testdata/TestAWSLambdaUpdate-after-state.yaml",
	)
	testutils.AssertGolden(
		t,
		snapshotService(mockAWS.lambdaMock),
		"testdata/TestAWSLambdaUpdate-after-aws.yaml",
	)
}

// ---

type mockLambda struct {
	Functions    map[string]*lambda.FunctionConfiguration
	FunctionTags map[string]map[string]string
}

func newMockLambda() *mockLambda {
	return &mockLambda{
		Functions:    make(map[string]*lambda.FunctionConfiguration),
		FunctionTags: make(map[string]map[string]string),
	}
}

func (m *mockLambda) sha256(input []byte) string {
	hasher := sha256.New()
	hasher.Write(input)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (m *mockLambda) CreateFunctionWithContext(ctx aws.Context, input *lambda.CreateFunctionInput, opts ...request.Option) (*lambda.FunctionConfiguration, error) {
	name := *input.FunctionName
	if name == "" {
		return nil, errors.New("name must be set")
	}
	if m.Functions[name] != nil {
		return nil, errors.New("function exists")
	}
	environment := &lambda.EnvironmentResponse{}
	if input.Environment != nil {
		environment.Variables = input.Environment.Variables
	}
	tracingConfig := &lambda.TracingConfigResponse{}
	if input.TracingConfig != nil {
		tracingConfig.Mode = input.TracingConfig.Mode
	}
	vpcConfig := &lambda.VpcConfigResponse{}
	if input.VpcConfig != nil {
		vpcConfig.SecurityGroupIds = input.VpcConfig.SecurityGroupIds
		vpcConfig.SubnetIds = input.VpcConfig.SubnetIds
		vpcConfig.VpcId = aws.String("vpcid")
	}
	arn := fmt.Sprintf("arn:aws:lambda::us-east-1:123456789000:function:%s", name)
	function := &lambda.FunctionConfiguration{
		CodeSha256:       aws.String(m.sha256(input.Code.ZipFile)),
		CodeSize:         aws.Int64(int64(len(input.Code.ZipFile))),
		DeadLetterConfig: input.DeadLetterConfig,
		Description:      input.Description,
		Environment:      environment,
		FunctionArn:      aws.String(arn),
		FunctionName:     input.FunctionName,
		Handler:          input.Handler,
		KMSKeyArn:        input.KMSKeyArn,
		LastModified:     aws.String("2017-11-01T19:19:37.640+0000"),
		MasterArn:        aws.String(arn),
		MemorySize:       input.MemorySize,
		Role:             input.Role,
		Runtime:          input.Runtime,
		Timeout:          input.Timeout,
		TracingConfig:    tracingConfig,
		Version:          aws.String("$LATEST"),
		VpcConfig:        vpcConfig,
	}
	m.Functions[name] = function
	m.FunctionTags[name] = make(map[string]string)
	for k, v := range input.Tags {
		m.FunctionTags[name][k] = *v
	}
	return function, nil
}

func (m *mockLambda) UpdateFunctionConfigurationWithContext(ctx aws.Context, input *lambda.UpdateFunctionConfigurationInput, opts ...request.Option) (*lambda.FunctionConfiguration, error) {
	return nil, errors.New("not implemented")
}

func (m *mockLambda) UpdateFunctionCodeWithContext(ctx aws.Context, input *lambda.UpdateFunctionCodeInput, opts ...request.Option) (*lambda.FunctionConfiguration, error) {
	return nil, errors.New("not implemented")
}
