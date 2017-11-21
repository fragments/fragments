package aws

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAWSIAMCreate(t *testing.T) {
	kv := backend.NewTestKV()
	mockAWS := newMockAWS()
	iamReconciler := newIAM(kv, mockAWS)
	ctx := context.Background()

	_, err := iamReconciler.putRole(ctx, &iamRoleInput{
		assumeRolePolicyDocument: "doc",
		description:              "error",
		path:                     "/test/",
		roleName:                 "", // empty
	})
	require.Error(t, err)

	role, err := iamReconciler.putRole(ctx, &iamRoleInput{
		assumeRolePolicyDocument: "doc",
		description:              "description",
		path:                     "/test/",
		roleName:                 "test",
	})
	require.NoError(t, err)
	require.NotNil(t, role)

	testutils.AssertGolden(
		t,
		testutils.SnapshotJSONMap(kv.Data),
		"testdata/TestAWSIAMCreate-state.yaml",
	)
	testutils.AssertGolden(
		t,
		snapshotService(mockAWS.iamMock),
		"testdata/TestAWSIAMCreate-aws.yaml",
	)
}

func TestAWSIAMUpdate(t *testing.T) {
	createDate, _ := time.Parse(time.RFC3339, "2017-10-01T12:34:56+00:00")
	roleID := "ABCDEFGHIJKLMNOPQR123"
	existing := &iam.Role{
		Arn: aws.String("arn:aws:iam::123456789000:role/path/existing"),
		AssumeRolePolicyDocument: aws.String("doc"),
		CreateDate:               &createDate,
		Description:              aws.String("previous"),
		Path:                     aws.String("/path/"),
		RoleId:                   aws.String(roleID),
		RoleName:                 aws.String("existing"),
	}

	kv := backend.NewTestKV()
	mockAWS := newMockAWS()
	iamReconciler := newIAM(kv, mockAWS)
	ctx := context.Background()

	res := iamReconciler.pointer("existing")
	err := res.Put(ctx, kv, existing)
	require.NoError(t, err)
	mockAWS.iamMock.Roles[*existing.RoleName] = existing

	_, err = iamReconciler.putRole(ctx, &iamRoleInput{
		roleName:    "",
		description: "err",
	})
	require.Error(t, err)

	role, err := iamReconciler.putRole(ctx, &iamRoleInput{
		roleName:    "existing",
		description: "updated",
	})
	require.NoError(t, err)
	require.NotNil(t, role)

	assert.Equal(t, *existing.Arn, *role.Arn)
	assert.Equal(t, "updated", *role.Description)

	testutils.AssertGolden(
		t,
		testutils.SnapshotJSONMap(kv.Data),
		"testdata/TestAWSIAMUpdate-state.yaml",
	)
	testutils.AssertGolden(
		t,
		snapshotService(mockAWS.iamMock),
		"testdata/TestAWSIAMUpdate-aws.yaml",
	)
}

// ---

type mockIAM struct {
	Roles map[string]*iam.Role
}

func newMockIAM() *mockIAM {
	return &mockIAM{
		Roles: make(map[string]*iam.Role),
	}
}

func (m *mockIAM) CreateRoleWithContext(ctx aws.Context, input *iam.CreateRoleInput, opts ...request.Option) (*iam.CreateRoleOutput, error) {
	name := *input.RoleName
	if name == "" {
		return nil, errors.New("name must be set")
	}
	if m.Roles[name] != nil {
		return nil, errors.New("role exists")
	}
	createDate, _ := time.Parse(time.RFC3339, "2017-10-31T12:34:56+00:00")
	roleID := "ABCDEFGHIJKLMNOPQR123"
	role := &iam.Role{
		Arn: aws.String(fmt.Sprintf("arn:aws:iam::123456789000:role/path/%s", name)),
		AssumeRolePolicyDocument: input.AssumeRolePolicyDocument,
		CreateDate:               &createDate,
		Description:              input.Description,
		Path:                     input.Path,
		RoleId:                   aws.String(roleID),
		RoleName:                 input.RoleName,
	}
	m.Roles[name] = role
	return &iam.CreateRoleOutput{
		Role: role,
	}, nil
}

func (m *mockIAM) UpdateRoleDescriptionWithContext(ctx aws.Context, input *iam.UpdateRoleDescriptionInput, opts ...request.Option) (*iam.UpdateRoleDescriptionOutput, error) {
	name := *input.RoleName
	if name == "" {
		return nil, errors.New("name must be set")
	}
	if m.Roles[name] == nil {
		return nil, errors.New("role does not exist")
	}
	m.Roles[name].Description = input.Description
	return &iam.UpdateRoleDescriptionOutput{
		Role: &iam.Role{
			Description: input.Description,
		},
	}, nil
}
