//go:generate mockery -dir ../../vendor/github.com/aws/aws-sdk-go/service/s3/s3iface -all

package filestore

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fragments/fragments/internal/filestore/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	conf := aws.NewConfig()
	tests := []struct {
		TestName      string
		Config        *aws.Config
		UploadBucket  string
		UploadTimeout time.Duration
		SourceBucket  string
		Error         bool
	}{
		{
			TestName: "No config",
			Error:    true,
		},
		{
			TestName: "No upload bucket",
			Config:   conf,
			Error:    true,
		},
		{
			TestName:     "No source bucket",
			Config:       conf,
			UploadBucket: "uploads",
			Error:        true,
		},
		{
			TestName:     "No timeout",
			Config:       conf,
			UploadBucket: "uploads",
			SourceBucket: "source",
			Error:        true,
		},
		{
			TestName:      "Ok",
			Config:        conf,
			UploadBucket:  "uploads",
			UploadTimeout: 10 * time.Minute,
			SourceBucket:  "source",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			_, err := NewS3(test.Config, test.UploadBucket, test.UploadTimeout, test.SourceBucket)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestNewUploadURL(t *testing.T) {
	tests := []struct {
		TestName string
		Name     string
		Error    bool
		Contains string
	}{
		{
			TestName: "Empty",
			Name:     "",
			Error:    true,
		},
		{
			TestName: "Ok",
			Name:     "testfile",
			Contains: "https://uploads.s3.amazonaws.com/testfile",
		},
	}

	conf := aws.
		NewConfig().
		WithCredentials(credentials.NewStaticCredentials("id", "secret", "token")).
		WithRegion("us-east-1")
	ses, _ := session.NewSession(conf)
	s := &S3{
		Client:       s3.New(ses),
		UploadBucket: "uploads",
		UploadExpiry: 5 * time.Minute,
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := s.NewUploadURL(test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Contains(t, actual, test.Contains)
		})
	}
}

func TestPersist(t *testing.T) {
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	tests := []struct {
		TestName    string
		Name        string
		Ctx         context.Context
		ArgError    bool
		CopyError   bool
		DeleteError bool
	}{
		{
			TestName: "No name",
			Name:     "",
			ArgError: true,
		},
		{
			TestName: "Context canceled",
			Name:     "File",
			Ctx:      ctxCanceled,
			ArgError: true,
		},
		{
			TestName:  "Copy error",
			Name:      "File",
			Ctx:       ctx,
			CopyError: true,
		},
		{
			TestName:    "Delete error",
			Name:        "File",
			Ctx:         ctx,
			DeleteError: true,
		},
		{
			TestName: "Ok",
			Name:     "File",
			Ctx:      ctx,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			mockS3 := &mocks.S3API{}
			s := &S3{
				Client:       mockS3,
				UploadBucket: "uploads",
				SourceBucket: "source",
			}

			var opts []request.Option
			mockS3.
				On("CopyObjectWithContext", ctxCanceled, mock.Anything, opts).
				Return(nil, errors.New("context canceled"))

			var copyErr error
			if test.CopyError {
				copyErr = errors.New("copy error")
			}
			mockS3.
				On("CopyObjectWithContext", ctx, mock.Anything, opts).
				Return(nil, copyErr)

			var delErr error
			if test.DeleteError {
				delErr = errors.New("delete error")
			}
			mockS3.
				On("DeleteObjectWithContext", ctx, mock.Anything, opts).
				Return(nil, delErr)

			err := s.Persist(test.Ctx, test.Name)
			if test.ArgError || test.CopyError || test.DeleteError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
