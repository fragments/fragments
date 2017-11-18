package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/pkg/errors"
)

type serviceProvider interface {
	lambda() (lambdaService, error)
	iam() (iamService, error)
}

type defaultProvider struct {
	credentials *credentials.Credentials
	region      string
}

func (d *defaultProvider) session() (*session.Session, error) {
	ses, err := session.NewSession(
		&aws.Config{
			Credentials: d.credentials,
			Region:      aws.String(d.region),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating aws session")
	}
	return ses, nil
}

func (d *defaultProvider) lambda() (lambdaService, error) {
	ses, err := d.session()
	if err != nil {
		return nil, err
	}
	return lambda.New(ses), nil
}

func (d *defaultProvider) iam() (iamService, error) {
	ses, err := d.session()
	if err != nil {
		return nil, err
	}
	return iam.New(ses), nil
}
