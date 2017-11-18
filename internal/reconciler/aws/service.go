package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
)

type serviceProvider interface {
	session() (*session.Session, error)
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
