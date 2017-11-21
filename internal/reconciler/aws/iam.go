package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/fragments/fragments/internal/state"
	"github.com/pkg/errors"
)

const iamResource state.ResourceType = "iam"

type iamService interface {
	CreateRoleWithContext(aws.Context, *iam.CreateRoleInput, ...request.Option) (*iam.CreateRoleOutput, error)
	UpdateRoleDescriptionWithContext(aws.Context, *iam.UpdateRoleDescriptionInput, ...request.Option) (*iam.UpdateRoleDescriptionOutput, error)
}

type iamReconciler struct {
	store       store
	svcProvider serviceProvider
}

func newIAM(store store, svcProvider serviceProvider) *iamReconciler {
	return &iamReconciler{
		store:       store,
		svcProvider: svcProvider,
	}
}

type iamRoleInput struct {
	assumeRolePolicyDocument string
	description              string
	path                     string
	roleName                 string
}

func (i *iamReconciler) pointer(name string) *state.ResPointer {
	return &state.ResPointer{
		InfraType:    state.InfrastructureTypeAWS,
		ResourceType: iamResource,
		Name:         name,
	}
}

func (i *iamReconciler) putRole(ctx context.Context, input *iamRoleInput) (*iam.Role, error) {
	res := i.pointer(input.roleName)
	unlock, err := res.Lock(ctx, i.store)
	if err != nil {
		return nil, errors.Wrap(err, "could not acquire lock for iam role")
	}
	defer unlock()

	var existing iam.Role
	exists, err := res.Get(ctx, i.store, &existing)
	if err != nil {
		return nil, errors.Wrap(err, "could not check existing iam role")
	}
	if exists {
		r, err := i.update(ctx, res, existing, input)
		if err != nil {
			return nil, errors.Wrap(err, "update")
		}
		return r, nil
	}
	r, err := i.create(ctx, res, input)
	if err != nil {
		return nil, errors.Wrap(err, "create")
	}
	return r, err
}

func (i *iamReconciler) create(ctx context.Context, res *state.ResPointer, input *iamRoleInput) (*iam.Role, error) {
	svc, err := i.svcProvider.iam()
	if err != nil {
		return nil, err
	}
	out, err := svc.CreateRoleWithContext(
		ctx,
		&iam.CreateRoleInput{
			AssumeRolePolicyDocument: aws.String(input.assumeRolePolicyDocument),
			Description:              aws.String(input.description),
			Path:                     aws.String(input.path),
			RoleName:                 aws.String(input.roleName),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not create iam role")
	}
	if err := res.Put(ctx, i.store, out.Role); err != nil {
		return nil, errors.Wrap(err, "could not store created iam role")
	}
	return out.Role, nil
}

func (i *iamReconciler) update(ctx context.Context, res *state.ResPointer, existing iam.Role, input *iamRoleInput) (*iam.Role, error) {
	// Only description can be updated, if it hasn't changed we can return
	// early. Since it's possible the previous description is nil, we can't
	// directly derefence it.
	prevDesc := ""
	if existing.Description != nil {
		prevDesc = *existing.Description
	}
	if prevDesc == input.description {
		return &existing, nil
	}
	svc, err := i.svcProvider.iam()
	if err != nil {
		return nil, err
	}
	out, err := svc.UpdateRoleDescriptionWithContext(
		ctx,
		&iam.UpdateRoleDescriptionInput{
			Description: aws.String(input.description),
			RoleName:    aws.String(input.roleName),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not update iam role")
	}
	role := existing
	role.Description = out.Role.Description
	if err = res.Put(ctx, i.store, role); err != nil {
		return nil, errors.Wrap(err, "could not store created iam role")
	}
	return &role, nil
}
