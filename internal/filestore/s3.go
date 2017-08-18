package filestore

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/pkg/errors"
)

// S3 stores files in AWS S3
type S3 struct {
	Client       s3iface.S3API
	UploadBucket string
	UploadExpiry time.Duration
	SourceBucket string
}

// NewS3 creates a new S3 client
func NewS3(conf *aws.Config, uploadBucket string, uploadExpiry time.Duration, sourceBucket string) (*S3, error) {
	if conf == nil {
		return nil, errors.New("config not set")
	}

	if uploadBucket == "" {
		return nil, errors.New("upload bucket not set")
	}

	if uploadExpiry.Seconds() == 0 {
		return nil, errors.New("upload expiry not set")
	}

	if sourceBucket == "" {
		return nil, errors.New("source bucket not set")
	}

	ses, err := session.NewSession(conf)
	if err != nil {
		return nil, errors.Wrap(err, "could not create session")
	}
	cli := s3.New(ses)

	return &S3{
		Client:       cli,
		UploadBucket: uploadBucket,
		UploadExpiry: uploadExpiry,
		SourceBucket: sourceBucket,
	}, nil
}

// NewUploadURL generates a new presigned URL on S3 for upload. The url is
// signed with the credentials provided to the client. The upload expires
// automatically.
// The url points to the upload bucket, the name argument specified the file
// name.
func (s *S3) NewUploadURL(name string) (string, error) {
	req, _ := s.Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s.UploadBucket),
		Key:    aws.String(name),
	})

	presigned, err := req.Presign(s.UploadExpiry)
	if err != nil {
		return "", errors.Wrap(err, "unable to presign upload url")
	}

	return presigned, nil
}

// Persist moves an uploaded file to a permanent bucket. Files that are not
// persisted might be cleaned up.
func (s *S3) Persist(ctx context.Context, name string) error {
	if name == "" {
		return errors.New("name not set")
	}

	// Copy file to source bucket
	_, err := s.Client.CopyObjectWithContext(ctx, &s3.CopyObjectInput{
		CopySource: aws.String(fmt.Sprintf("%s/%s", s.UploadBucket, name)),
		Bucket:     aws.String(s.SourceBucket),
		Key:        aws.String(name),
	})
	if err != nil {
		return errors.Wrapf(err, "could not copy uploaded file %s from bucket %s to %s", name, s.UploadBucket, s.SourceBucket)
	}

	// Delete uploaded file
	_, err = s.Client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.SourceBucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return errors.Wrap(err, "could not delete uploaded source after copy")
	}

	return nil
}
