package mocks

import "github.com/stretchr/testify/mock"

import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/request"
import "github.com/aws/aws-sdk-go/service/s3"

type S3API struct {
	mock.Mock
}

// AbortMultipartUpload provides a mock function with given fields: _a0
func (_m *S3API) AbortMultipartUpload(_a0 *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.AbortMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(*s3.AbortMultipartUploadInput) *s3.AbortMultipartUploadOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.AbortMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.AbortMultipartUploadInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbortMultipartUploadWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) AbortMultipartUploadWithContext(_a0 aws.Context, _a1 *s3.AbortMultipartUploadInput, _a2 ...request.Option) (*s3.AbortMultipartUploadOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.AbortMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.AbortMultipartUploadInput, ...request.Option) *s3.AbortMultipartUploadOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.AbortMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.AbortMultipartUploadInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbortMultipartUploadRequest provides a mock function with given fields: _a0
func (_m *S3API) AbortMultipartUploadRequest(_a0 *s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.AbortMultipartUploadInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.AbortMultipartUploadOutput
	if rf, ok := ret.Get(1).(func(*s3.AbortMultipartUploadInput) *s3.AbortMultipartUploadOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.AbortMultipartUploadOutput)
		}
	}

	return r0, r1
}

// CompleteMultipartUpload provides a mock function with given fields: _a0
func (_m *S3API) CompleteMultipartUpload(_a0 *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.CompleteMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(*s3.CompleteMultipartUploadInput) *s3.CompleteMultipartUploadOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CompleteMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.CompleteMultipartUploadInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteMultipartUploadWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) CompleteMultipartUploadWithContext(_a0 aws.Context, _a1 *s3.CompleteMultipartUploadInput, _a2 ...request.Option) (*s3.CompleteMultipartUploadOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.CompleteMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.CompleteMultipartUploadInput, ...request.Option) *s3.CompleteMultipartUploadOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CompleteMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.CompleteMultipartUploadInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteMultipartUploadRequest provides a mock function with given fields: _a0
func (_m *S3API) CompleteMultipartUploadRequest(_a0 *s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.CompleteMultipartUploadInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.CompleteMultipartUploadOutput
	if rf, ok := ret.Get(1).(func(*s3.CompleteMultipartUploadInput) *s3.CompleteMultipartUploadOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.CompleteMultipartUploadOutput)
		}
	}

	return r0, r1
}

// CopyObject provides a mock function with given fields: _a0
func (_m *S3API) CopyObject(_a0 *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.CopyObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.CopyObjectInput) *s3.CopyObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CopyObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.CopyObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyObjectWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) CopyObjectWithContext(_a0 aws.Context, _a1 *s3.CopyObjectInput, _a2 ...request.Option) (*s3.CopyObjectOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.CopyObjectOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.CopyObjectInput, ...request.Option) *s3.CopyObjectOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CopyObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.CopyObjectInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyObjectRequest provides a mock function with given fields: _a0
func (_m *S3API) CopyObjectRequest(_a0 *s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.CopyObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.CopyObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.CopyObjectInput) *s3.CopyObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.CopyObjectOutput)
		}
	}

	return r0, r1
}

// CreateBucket provides a mock function with given fields: _a0
func (_m *S3API) CreateBucket(_a0 *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.CreateBucketOutput
	if rf, ok := ret.Get(0).(func(*s3.CreateBucketInput) *s3.CreateBucketOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.CreateBucketInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBucketWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) CreateBucketWithContext(_a0 aws.Context, _a1 *s3.CreateBucketInput, _a2 ...request.Option) (*s3.CreateBucketOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.CreateBucketOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.CreateBucketInput, ...request.Option) *s3.CreateBucketOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.CreateBucketInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBucketRequest provides a mock function with given fields: _a0
func (_m *S3API) CreateBucketRequest(_a0 *s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.CreateBucketInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.CreateBucketOutput
	if rf, ok := ret.Get(1).(func(*s3.CreateBucketInput) *s3.CreateBucketOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.CreateBucketOutput)
		}
	}

	return r0, r1
}

// CreateMultipartUpload provides a mock function with given fields: _a0
func (_m *S3API) CreateMultipartUpload(_a0 *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.CreateMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(*s3.CreateMultipartUploadInput) *s3.CreateMultipartUploadOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.CreateMultipartUploadInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMultipartUploadWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) CreateMultipartUploadWithContext(_a0 aws.Context, _a1 *s3.CreateMultipartUploadInput, _a2 ...request.Option) (*s3.CreateMultipartUploadOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.CreateMultipartUploadOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.CreateMultipartUploadInput, ...request.Option) *s3.CreateMultipartUploadOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateMultipartUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.CreateMultipartUploadInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMultipartUploadRequest provides a mock function with given fields: _a0
func (_m *S3API) CreateMultipartUploadRequest(_a0 *s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.CreateMultipartUploadInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.CreateMultipartUploadOutput
	if rf, ok := ret.Get(1).(func(*s3.CreateMultipartUploadInput) *s3.CreateMultipartUploadOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.CreateMultipartUploadOutput)
		}
	}

	return r0, r1
}

// DeleteBucket provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucket(_a0 *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketInput) *s3.DeleteBucketOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketInput, _a2 ...request.Option) (*s3.DeleteBucketOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketInput, ...request.Option) *s3.DeleteBucketOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketRequest(_a0 *s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketInput) *s3.DeleteBucketOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketOutput)
		}
	}

	return r0, r1
}

// DeleteBucketAnalyticsConfiguration provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketAnalyticsConfiguration(_a0 *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketAnalyticsConfigurationInput) *s3.DeleteBucketAnalyticsConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketAnalyticsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketAnalyticsConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketAnalyticsConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketAnalyticsConfigurationWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketAnalyticsConfigurationInput, _a2 ...request.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketAnalyticsConfigurationInput, ...request.Option) *s3.DeleteBucketAnalyticsConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketAnalyticsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketAnalyticsConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketAnalyticsConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketAnalyticsConfigurationRequest(_a0 *s3.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *s3.DeleteBucketAnalyticsConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketAnalyticsConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketAnalyticsConfigurationInput) *s3.DeleteBucketAnalyticsConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketAnalyticsConfigurationOutput)
		}
	}

	return r0, r1
}

// DeleteBucketCors provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketCors(_a0 *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketCorsOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketCorsInput) *s3.DeleteBucketCorsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketCorsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketCorsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketCorsWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketCorsInput, _a2 ...request.Option) (*s3.DeleteBucketCorsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketCorsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketCorsInput, ...request.Option) *s3.DeleteBucketCorsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketCorsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketCorsRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketCorsRequest(_a0 *s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketCorsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketCorsOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketCorsInput) *s3.DeleteBucketCorsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketCorsOutput)
		}
	}

	return r0, r1
}

// DeleteBucketInventoryConfiguration provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketInventoryConfiguration(_a0 *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketInventoryConfigurationInput) *s3.DeleteBucketInventoryConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketInventoryConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketInventoryConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketInventoryConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketInventoryConfigurationWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketInventoryConfigurationInput, _a2 ...request.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketInventoryConfigurationInput, ...request.Option) *s3.DeleteBucketInventoryConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketInventoryConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketInventoryConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketInventoryConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketInventoryConfigurationRequest(_a0 *s3.DeleteBucketInventoryConfigurationInput) (*request.Request, *s3.DeleteBucketInventoryConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketInventoryConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketInventoryConfigurationInput) *s3.DeleteBucketInventoryConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketInventoryConfigurationOutput)
		}
	}

	return r0, r1
}

// DeleteBucketLifecycle provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketLifecycle(_a0 *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketLifecycleInput) *s3.DeleteBucketLifecycleOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketLifecycleInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketLifecycleWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketLifecycleWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketLifecycleInput, _a2 ...request.Option) (*s3.DeleteBucketLifecycleOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketLifecycleInput, ...request.Option) *s3.DeleteBucketLifecycleOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketLifecycleInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketLifecycleRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketLifecycleRequest(_a0 *s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketLifecycleInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketLifecycleOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketLifecycleInput) *s3.DeleteBucketLifecycleOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketLifecycleOutput)
		}
	}

	return r0, r1
}

// DeleteBucketMetricsConfiguration provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketMetricsConfiguration(_a0 *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketMetricsConfigurationInput) *s3.DeleteBucketMetricsConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketMetricsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketMetricsConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketMetricsConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketMetricsConfigurationWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketMetricsConfigurationInput, _a2 ...request.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketMetricsConfigurationInput, ...request.Option) *s3.DeleteBucketMetricsConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketMetricsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketMetricsConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketMetricsConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketMetricsConfigurationRequest(_a0 *s3.DeleteBucketMetricsConfigurationInput) (*request.Request, *s3.DeleteBucketMetricsConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketMetricsConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketMetricsConfigurationInput) *s3.DeleteBucketMetricsConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketMetricsConfigurationOutput)
		}
	}

	return r0, r1
}

// DeleteBucketPolicy provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketPolicy(_a0 *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketPolicyInput) *s3.DeleteBucketPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketPolicyWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketPolicyWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketPolicyInput, _a2 ...request.Option) (*s3.DeleteBucketPolicyOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketPolicyInput, ...request.Option) *s3.DeleteBucketPolicyOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketPolicyInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketPolicyRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketPolicyRequest(_a0 *s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketPolicyOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketPolicyInput) *s3.DeleteBucketPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketPolicyOutput)
		}
	}

	return r0, r1
}

// DeleteBucketReplication provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketReplication(_a0 *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketReplicationInput) *s3.DeleteBucketReplicationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketReplicationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketReplicationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketReplicationWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketReplicationInput, _a2 ...request.Option) (*s3.DeleteBucketReplicationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketReplicationInput, ...request.Option) *s3.DeleteBucketReplicationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketReplicationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketReplicationRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketReplicationRequest(_a0 *s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketReplicationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketReplicationOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketReplicationInput) *s3.DeleteBucketReplicationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketReplicationOutput)
		}
	}

	return r0, r1
}

// DeleteBucketTagging provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketTagging(_a0 *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketTaggingInput) *s3.DeleteBucketTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketTaggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketTaggingWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketTaggingInput, _a2 ...request.Option) (*s3.DeleteBucketTaggingOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketTaggingInput, ...request.Option) *s3.DeleteBucketTaggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketTaggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketTaggingRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketTaggingRequest(_a0 *s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketTaggingInput) *s3.DeleteBucketTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketTaggingOutput)
		}
	}

	return r0, r1
}

// DeleteBucketWebsite provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketWebsite(_a0 *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketWebsiteInput) *s3.DeleteBucketWebsiteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketWebsiteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketWebsiteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteBucketWebsiteWithContext(_a0 aws.Context, _a1 *s3.DeleteBucketWebsiteInput, _a2 ...request.Option) (*s3.DeleteBucketWebsiteOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteBucketWebsiteInput, ...request.Option) *s3.DeleteBucketWebsiteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteBucketWebsiteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBucketWebsiteRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteBucketWebsiteRequest(_a0 *s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteBucketWebsiteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteBucketWebsiteOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteBucketWebsiteInput) *s3.DeleteBucketWebsiteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteBucketWebsiteOutput)
		}
	}

	return r0, r1
}

// DeleteObject provides a mock function with given fields: _a0
func (_m *S3API) DeleteObject(_a0 *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectInput) *s3.DeleteObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteObjectWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteObjectWithContext(_a0 aws.Context, _a1 *s3.DeleteObjectInput, _a2 ...request.Option) (*s3.DeleteObjectOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteObjectOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteObjectInput, ...request.Option) *s3.DeleteObjectOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteObjectInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteObjectRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteObjectRequest(_a0 *s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectInput) *s3.DeleteObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteObjectOutput)
		}
	}

	return r0, r1
}

// DeleteObjectTagging provides a mock function with given fields: _a0
func (_m *S3API) DeleteObjectTagging(_a0 *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteObjectTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectTaggingInput) *s3.DeleteObjectTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteObjectTaggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteObjectTaggingWithContext(_a0 aws.Context, _a1 *s3.DeleteObjectTaggingInput, _a2 ...request.Option) (*s3.DeleteObjectTaggingOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteObjectTaggingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteObjectTaggingInput, ...request.Option) *s3.DeleteObjectTaggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteObjectTaggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteObjectTaggingRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteObjectTaggingRequest(_a0 *s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteObjectTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectTaggingInput) *s3.DeleteObjectTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteObjectTaggingOutput)
		}
	}

	return r0, r1
}

// DeleteObjects provides a mock function with given fields: _a0
func (_m *S3API) DeleteObjects(_a0 *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.DeleteObjectsOutput
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectsInput) *s3.DeleteObjectsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteObjectsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) DeleteObjectsWithContext(_a0 aws.Context, _a1 *s3.DeleteObjectsInput, _a2 ...request.Option) (*s3.DeleteObjectsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.DeleteObjectsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.DeleteObjectsInput, ...request.Option) *s3.DeleteObjectsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.DeleteObjectsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteObjectsRequest provides a mock function with given fields: _a0
func (_m *S3API) DeleteObjectsRequest(_a0 *s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.DeleteObjectsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.DeleteObjectsOutput
	if rf, ok := ret.Get(1).(func(*s3.DeleteObjectsInput) *s3.DeleteObjectsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.DeleteObjectsOutput)
		}
	}

	return r0, r1
}

// GetBucketAccelerateConfiguration provides a mock function with given fields: _a0
func (_m *S3API) GetBucketAccelerateConfiguration(_a0 *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketAccelerateConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketAccelerateConfigurationInput) *s3.GetBucketAccelerateConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketAccelerateConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketAccelerateConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketAccelerateConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketAccelerateConfigurationWithContext(_a0 aws.Context, _a1 *s3.GetBucketAccelerateConfigurationInput, _a2 ...request.Option) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketAccelerateConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketAccelerateConfigurationInput, ...request.Option) *s3.GetBucketAccelerateConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketAccelerateConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketAccelerateConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketAccelerateConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketAccelerateConfigurationRequest(_a0 *s3.GetBucketAccelerateConfigurationInput) (*request.Request, *s3.GetBucketAccelerateConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketAccelerateConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketAccelerateConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketAccelerateConfigurationInput) *s3.GetBucketAccelerateConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketAccelerateConfigurationOutput)
		}
	}

	return r0, r1
}

// GetBucketAcl provides a mock function with given fields: _a0
func (_m *S3API) GetBucketAcl(_a0 *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketAclOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketAclInput) *s3.GetBucketAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketAclWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketAclWithContext(_a0 aws.Context, _a1 *s3.GetBucketAclInput, _a2 ...request.Option) (*s3.GetBucketAclOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketAclOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketAclInput, ...request.Option) *s3.GetBucketAclOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketAclInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketAclRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketAclRequest(_a0 *s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketAclOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketAclInput) *s3.GetBucketAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketAclOutput)
		}
	}

	return r0, r1
}

// GetBucketAnalyticsConfiguration provides a mock function with given fields: _a0
func (_m *S3API) GetBucketAnalyticsConfiguration(_a0 *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketAnalyticsConfigurationInput) *s3.GetBucketAnalyticsConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketAnalyticsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketAnalyticsConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketAnalyticsConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketAnalyticsConfigurationWithContext(_a0 aws.Context, _a1 *s3.GetBucketAnalyticsConfigurationInput, _a2 ...request.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketAnalyticsConfigurationInput, ...request.Option) *s3.GetBucketAnalyticsConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketAnalyticsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketAnalyticsConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketAnalyticsConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketAnalyticsConfigurationRequest(_a0 *s3.GetBucketAnalyticsConfigurationInput) (*request.Request, *s3.GetBucketAnalyticsConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketAnalyticsConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketAnalyticsConfigurationInput) *s3.GetBucketAnalyticsConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketAnalyticsConfigurationOutput)
		}
	}

	return r0, r1
}

// GetBucketCors provides a mock function with given fields: _a0
func (_m *S3API) GetBucketCors(_a0 *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketCorsOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketCorsInput) *s3.GetBucketCorsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketCorsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketCorsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketCorsWithContext(_a0 aws.Context, _a1 *s3.GetBucketCorsInput, _a2 ...request.Option) (*s3.GetBucketCorsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketCorsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketCorsInput, ...request.Option) *s3.GetBucketCorsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketCorsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketCorsRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketCorsRequest(_a0 *s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketCorsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketCorsOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketCorsInput) *s3.GetBucketCorsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketCorsOutput)
		}
	}

	return r0, r1
}

// GetBucketInventoryConfiguration provides a mock function with given fields: _a0
func (_m *S3API) GetBucketInventoryConfiguration(_a0 *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketInventoryConfigurationInput) *s3.GetBucketInventoryConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketInventoryConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketInventoryConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketInventoryConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketInventoryConfigurationWithContext(_a0 aws.Context, _a1 *s3.GetBucketInventoryConfigurationInput, _a2 ...request.Option) (*s3.GetBucketInventoryConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketInventoryConfigurationInput, ...request.Option) *s3.GetBucketInventoryConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketInventoryConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketInventoryConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketInventoryConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketInventoryConfigurationRequest(_a0 *s3.GetBucketInventoryConfigurationInput) (*request.Request, *s3.GetBucketInventoryConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketInventoryConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketInventoryConfigurationInput) *s3.GetBucketInventoryConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketInventoryConfigurationOutput)
		}
	}

	return r0, r1
}

// GetBucketLifecycle provides a mock function with given fields: _a0
func (_m *S3API) GetBucketLifecycle(_a0 *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLifecycleInput) *s3.GetBucketLifecycleOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLifecycleInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLifecycleWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketLifecycleWithContext(_a0 aws.Context, _a1 *s3.GetBucketLifecycleInput, _a2 ...request.Option) (*s3.GetBucketLifecycleOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketLifecycleInput, ...request.Option) *s3.GetBucketLifecycleOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketLifecycleInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLifecycleRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketLifecycleRequest(_a0 *s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLifecycleInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketLifecycleOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLifecycleInput) *s3.GetBucketLifecycleOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketLifecycleOutput)
		}
	}

	return r0, r1
}

// GetBucketLifecycleConfiguration provides a mock function with given fields: _a0
func (_m *S3API) GetBucketLifecycleConfiguration(_a0 *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLifecycleConfigurationInput) *s3.GetBucketLifecycleConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLifecycleConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLifecycleConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLifecycleConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketLifecycleConfigurationWithContext(_a0 aws.Context, _a1 *s3.GetBucketLifecycleConfigurationInput, _a2 ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketLifecycleConfigurationInput, ...request.Option) *s3.GetBucketLifecycleConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLifecycleConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketLifecycleConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLifecycleConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketLifecycleConfigurationRequest(_a0 *s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLifecycleConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLifecycleConfigurationInput) *s3.GetBucketLifecycleConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketLifecycleConfigurationOutput)
		}
	}

	return r0, r1
}

// GetBucketLocation provides a mock function with given fields: _a0
func (_m *S3API) GetBucketLocation(_a0 *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketLocationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLocationInput) *s3.GetBucketLocationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLocationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLocationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLocationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketLocationWithContext(_a0 aws.Context, _a1 *s3.GetBucketLocationInput, _a2 ...request.Option) (*s3.GetBucketLocationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketLocationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketLocationInput, ...request.Option) *s3.GetBucketLocationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLocationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketLocationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLocationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketLocationRequest(_a0 *s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLocationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketLocationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLocationInput) *s3.GetBucketLocationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketLocationOutput)
		}
	}

	return r0, r1
}

// GetBucketLogging provides a mock function with given fields: _a0
func (_m *S3API) GetBucketLogging(_a0 *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketLoggingOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLoggingInput) *s3.GetBucketLoggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLoggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLoggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketLoggingWithContext(_a0 aws.Context, _a1 *s3.GetBucketLoggingInput, _a2 ...request.Option) (*s3.GetBucketLoggingOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketLoggingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketLoggingInput, ...request.Option) *s3.GetBucketLoggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketLoggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLoggingRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketLoggingRequest(_a0 *s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketLoggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketLoggingOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketLoggingInput) *s3.GetBucketLoggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketLoggingOutput)
		}
	}

	return r0, r1
}

// GetBucketMetricsConfiguration provides a mock function with given fields: _a0
func (_m *S3API) GetBucketMetricsConfiguration(_a0 *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketMetricsConfigurationInput) *s3.GetBucketMetricsConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketMetricsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketMetricsConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketMetricsConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketMetricsConfigurationWithContext(_a0 aws.Context, _a1 *s3.GetBucketMetricsConfigurationInput, _a2 ...request.Option) (*s3.GetBucketMetricsConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketMetricsConfigurationInput, ...request.Option) *s3.GetBucketMetricsConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketMetricsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketMetricsConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketMetricsConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketMetricsConfigurationRequest(_a0 *s3.GetBucketMetricsConfigurationInput) (*request.Request, *s3.GetBucketMetricsConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketMetricsConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketMetricsConfigurationInput) *s3.GetBucketMetricsConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketMetricsConfigurationOutput)
		}
	}

	return r0, r1
}

// GetBucketNotification provides a mock function with given fields: _a0
func (_m *S3API) GetBucketNotification(_a0 *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	ret := _m.Called(_a0)

	var r0 *s3.NotificationConfigurationDeprecated
	if rf, ok := ret.Get(0).(func(*s3.GetBucketNotificationConfigurationRequest) *s3.NotificationConfigurationDeprecated); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.NotificationConfigurationDeprecated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketNotificationConfigurationRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketNotificationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketNotificationWithContext(_a0 aws.Context, _a1 *s3.GetBucketNotificationConfigurationRequest, _a2 ...request.Option) (*s3.NotificationConfigurationDeprecated, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.NotificationConfigurationDeprecated
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) *s3.NotificationConfigurationDeprecated); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.NotificationConfigurationDeprecated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketNotificationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketNotificationRequest(_a0 *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketNotificationConfigurationRequest) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.NotificationConfigurationDeprecated
	if rf, ok := ret.Get(1).(func(*s3.GetBucketNotificationConfigurationRequest) *s3.NotificationConfigurationDeprecated); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.NotificationConfigurationDeprecated)
		}
	}

	return r0, r1
}

// GetBucketNotificationConfiguration provides a mock function with given fields: _a0
func (_m *S3API) GetBucketNotificationConfiguration(_a0 *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	ret := _m.Called(_a0)

	var r0 *s3.NotificationConfiguration
	if rf, ok := ret.Get(0).(func(*s3.GetBucketNotificationConfigurationRequest) *s3.NotificationConfiguration); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.NotificationConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketNotificationConfigurationRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketNotificationConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketNotificationConfigurationWithContext(_a0 aws.Context, _a1 *s3.GetBucketNotificationConfigurationRequest, _a2 ...request.Option) (*s3.NotificationConfiguration, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.NotificationConfiguration
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) *s3.NotificationConfiguration); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.NotificationConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketNotificationConfigurationRequest, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketNotificationConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketNotificationConfigurationRequest(_a0 *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketNotificationConfigurationRequest) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.NotificationConfiguration
	if rf, ok := ret.Get(1).(func(*s3.GetBucketNotificationConfigurationRequest) *s3.NotificationConfiguration); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.NotificationConfiguration)
		}
	}

	return r0, r1
}

// GetBucketPolicy provides a mock function with given fields: _a0
func (_m *S3API) GetBucketPolicy(_a0 *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketPolicyInput) *s3.GetBucketPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketPolicyWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketPolicyWithContext(_a0 aws.Context, _a1 *s3.GetBucketPolicyInput, _a2 ...request.Option) (*s3.GetBucketPolicyOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketPolicyInput, ...request.Option) *s3.GetBucketPolicyOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketPolicyInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketPolicyRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketPolicyRequest(_a0 *s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketPolicyOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketPolicyInput) *s3.GetBucketPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketPolicyOutput)
		}
	}

	return r0, r1
}

// GetBucketReplication provides a mock function with given fields: _a0
func (_m *S3API) GetBucketReplication(_a0 *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketReplicationInput) *s3.GetBucketReplicationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketReplicationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketReplicationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketReplicationWithContext(_a0 aws.Context, _a1 *s3.GetBucketReplicationInput, _a2 ...request.Option) (*s3.GetBucketReplicationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketReplicationInput, ...request.Option) *s3.GetBucketReplicationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketReplicationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketReplicationRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketReplicationRequest(_a0 *s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketReplicationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketReplicationOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketReplicationInput) *s3.GetBucketReplicationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketReplicationOutput)
		}
	}

	return r0, r1
}

// GetBucketRequestPayment provides a mock function with given fields: _a0
func (_m *S3API) GetBucketRequestPayment(_a0 *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketRequestPaymentOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketRequestPaymentInput) *s3.GetBucketRequestPaymentOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketRequestPaymentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketRequestPaymentInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketRequestPaymentWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketRequestPaymentWithContext(_a0 aws.Context, _a1 *s3.GetBucketRequestPaymentInput, _a2 ...request.Option) (*s3.GetBucketRequestPaymentOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketRequestPaymentOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketRequestPaymentInput, ...request.Option) *s3.GetBucketRequestPaymentOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketRequestPaymentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketRequestPaymentInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketRequestPaymentRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketRequestPaymentRequest(_a0 *s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketRequestPaymentInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketRequestPaymentOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketRequestPaymentInput) *s3.GetBucketRequestPaymentOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketRequestPaymentOutput)
		}
	}

	return r0, r1
}

// GetBucketTagging provides a mock function with given fields: _a0
func (_m *S3API) GetBucketTagging(_a0 *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketTaggingInput) *s3.GetBucketTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketTaggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketTaggingWithContext(_a0 aws.Context, _a1 *s3.GetBucketTaggingInput, _a2 ...request.Option) (*s3.GetBucketTaggingOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketTaggingInput, ...request.Option) *s3.GetBucketTaggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketTaggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketTaggingRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketTaggingRequest(_a0 *s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketTaggingInput) *s3.GetBucketTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketTaggingOutput)
		}
	}

	return r0, r1
}

// GetBucketVersioning provides a mock function with given fields: _a0
func (_m *S3API) GetBucketVersioning(_a0 *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketVersioningOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketVersioningInput) *s3.GetBucketVersioningOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketVersioningOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketVersioningInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketVersioningWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketVersioningWithContext(_a0 aws.Context, _a1 *s3.GetBucketVersioningInput, _a2 ...request.Option) (*s3.GetBucketVersioningOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketVersioningOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketVersioningInput, ...request.Option) *s3.GetBucketVersioningOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketVersioningOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketVersioningInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketVersioningRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketVersioningRequest(_a0 *s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketVersioningInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketVersioningOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketVersioningInput) *s3.GetBucketVersioningOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketVersioningOutput)
		}
	}

	return r0, r1
}

// GetBucketWebsite provides a mock function with given fields: _a0
func (_m *S3API) GetBucketWebsite(_a0 *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(*s3.GetBucketWebsiteInput) *s3.GetBucketWebsiteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetBucketWebsiteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketWebsiteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetBucketWebsiteWithContext(_a0 aws.Context, _a1 *s3.GetBucketWebsiteInput, _a2 ...request.Option) (*s3.GetBucketWebsiteOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetBucketWebsiteInput, ...request.Option) *s3.GetBucketWebsiteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetBucketWebsiteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketWebsiteRequest provides a mock function with given fields: _a0
func (_m *S3API) GetBucketWebsiteRequest(_a0 *s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetBucketWebsiteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetBucketWebsiteOutput
	if rf, ok := ret.Get(1).(func(*s3.GetBucketWebsiteInput) *s3.GetBucketWebsiteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetBucketWebsiteOutput)
		}
	}

	return r0, r1
}

// GetObject provides a mock function with given fields: _a0
func (_m *S3API) GetObject(_a0 *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.GetObjectInput) *s3.GetObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetObjectWithContext(_a0 aws.Context, _a1 *s3.GetObjectInput, _a2 ...request.Option) (*s3.GetObjectOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetObjectOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetObjectInput, ...request.Option) *s3.GetObjectOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetObjectInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectRequest provides a mock function with given fields: _a0
func (_m *S3API) GetObjectRequest(_a0 *s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.GetObjectInput) *s3.GetObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetObjectOutput)
		}
	}

	return r0, r1
}

// GetObjectAcl provides a mock function with given fields: _a0
func (_m *S3API) GetObjectAcl(_a0 *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetObjectAclOutput
	if rf, ok := ret.Get(0).(func(*s3.GetObjectAclInput) *s3.GetObjectAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetObjectAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectAclWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetObjectAclWithContext(_a0 aws.Context, _a1 *s3.GetObjectAclInput, _a2 ...request.Option) (*s3.GetObjectAclOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetObjectAclOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetObjectAclInput, ...request.Option) *s3.GetObjectAclOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetObjectAclInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectAclRequest provides a mock function with given fields: _a0
func (_m *S3API) GetObjectAclRequest(_a0 *s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetObjectAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetObjectAclOutput
	if rf, ok := ret.Get(1).(func(*s3.GetObjectAclInput) *s3.GetObjectAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetObjectAclOutput)
		}
	}

	return r0, r1
}

// GetObjectTagging provides a mock function with given fields: _a0
func (_m *S3API) GetObjectTagging(_a0 *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetObjectTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.GetObjectTaggingInput) *s3.GetObjectTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetObjectTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectTaggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetObjectTaggingWithContext(_a0 aws.Context, _a1 *s3.GetObjectTaggingInput, _a2 ...request.Option) (*s3.GetObjectTaggingOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetObjectTaggingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetObjectTaggingInput, ...request.Option) *s3.GetObjectTaggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetObjectTaggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectTaggingRequest provides a mock function with given fields: _a0
func (_m *S3API) GetObjectTaggingRequest(_a0 *s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetObjectTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetObjectTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.GetObjectTaggingInput) *s3.GetObjectTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetObjectTaggingOutput)
		}
	}

	return r0, r1
}

// GetObjectTorrent provides a mock function with given fields: _a0
func (_m *S3API) GetObjectTorrent(_a0 *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.GetObjectTorrentOutput
	if rf, ok := ret.Get(0).(func(*s3.GetObjectTorrentInput) *s3.GetObjectTorrentOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectTorrentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.GetObjectTorrentInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectTorrentWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) GetObjectTorrentWithContext(_a0 aws.Context, _a1 *s3.GetObjectTorrentInput, _a2 ...request.Option) (*s3.GetObjectTorrentOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.GetObjectTorrentOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.GetObjectTorrentInput, ...request.Option) *s3.GetObjectTorrentOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectTorrentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.GetObjectTorrentInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectTorrentRequest provides a mock function with given fields: _a0
func (_m *S3API) GetObjectTorrentRequest(_a0 *s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.GetObjectTorrentInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.GetObjectTorrentOutput
	if rf, ok := ret.Get(1).(func(*s3.GetObjectTorrentInput) *s3.GetObjectTorrentOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.GetObjectTorrentOutput)
		}
	}

	return r0, r1
}

// HeadBucket provides a mock function with given fields: _a0
func (_m *S3API) HeadBucket(_a0 *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.HeadBucketOutput
	if rf, ok := ret.Get(0).(func(*s3.HeadBucketInput) *s3.HeadBucketOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.HeadBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.HeadBucketInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadBucketWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) HeadBucketWithContext(_a0 aws.Context, _a1 *s3.HeadBucketInput, _a2 ...request.Option) (*s3.HeadBucketOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.HeadBucketOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.HeadBucketInput, ...request.Option) *s3.HeadBucketOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.HeadBucketOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.HeadBucketInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadBucketRequest provides a mock function with given fields: _a0
func (_m *S3API) HeadBucketRequest(_a0 *s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.HeadBucketInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.HeadBucketOutput
	if rf, ok := ret.Get(1).(func(*s3.HeadBucketInput) *s3.HeadBucketOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.HeadBucketOutput)
		}
	}

	return r0, r1
}

// HeadObject provides a mock function with given fields: _a0
func (_m *S3API) HeadObject(_a0 *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.HeadObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.HeadObjectInput) *s3.HeadObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.HeadObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.HeadObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadObjectWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) HeadObjectWithContext(_a0 aws.Context, _a1 *s3.HeadObjectInput, _a2 ...request.Option) (*s3.HeadObjectOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.HeadObjectOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.HeadObjectInput, ...request.Option) *s3.HeadObjectOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.HeadObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.HeadObjectInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadObjectRequest provides a mock function with given fields: _a0
func (_m *S3API) HeadObjectRequest(_a0 *s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.HeadObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.HeadObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.HeadObjectInput) *s3.HeadObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.HeadObjectOutput)
		}
	}

	return r0, r1
}

// ListBucketAnalyticsConfigurations provides a mock function with given fields: _a0
func (_m *S3API) ListBucketAnalyticsConfigurations(_a0 *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListBucketAnalyticsConfigurationsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListBucketAnalyticsConfigurationsInput) *s3.ListBucketAnalyticsConfigurationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketAnalyticsConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListBucketAnalyticsConfigurationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketAnalyticsConfigurationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListBucketAnalyticsConfigurationsWithContext(_a0 aws.Context, _a1 *s3.ListBucketAnalyticsConfigurationsInput, _a2 ...request.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListBucketAnalyticsConfigurationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListBucketAnalyticsConfigurationsInput, ...request.Option) *s3.ListBucketAnalyticsConfigurationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketAnalyticsConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListBucketAnalyticsConfigurationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketAnalyticsConfigurationsRequest provides a mock function with given fields: _a0
func (_m *S3API) ListBucketAnalyticsConfigurationsRequest(_a0 *s3.ListBucketAnalyticsConfigurationsInput) (*request.Request, *s3.ListBucketAnalyticsConfigurationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListBucketAnalyticsConfigurationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListBucketAnalyticsConfigurationsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListBucketAnalyticsConfigurationsInput) *s3.ListBucketAnalyticsConfigurationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListBucketAnalyticsConfigurationsOutput)
		}
	}

	return r0, r1
}

// ListBucketInventoryConfigurations provides a mock function with given fields: _a0
func (_m *S3API) ListBucketInventoryConfigurations(_a0 *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListBucketInventoryConfigurationsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListBucketInventoryConfigurationsInput) *s3.ListBucketInventoryConfigurationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketInventoryConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListBucketInventoryConfigurationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketInventoryConfigurationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListBucketInventoryConfigurationsWithContext(_a0 aws.Context, _a1 *s3.ListBucketInventoryConfigurationsInput, _a2 ...request.Option) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListBucketInventoryConfigurationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListBucketInventoryConfigurationsInput, ...request.Option) *s3.ListBucketInventoryConfigurationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketInventoryConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListBucketInventoryConfigurationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketInventoryConfigurationsRequest provides a mock function with given fields: _a0
func (_m *S3API) ListBucketInventoryConfigurationsRequest(_a0 *s3.ListBucketInventoryConfigurationsInput) (*request.Request, *s3.ListBucketInventoryConfigurationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListBucketInventoryConfigurationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListBucketInventoryConfigurationsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListBucketInventoryConfigurationsInput) *s3.ListBucketInventoryConfigurationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListBucketInventoryConfigurationsOutput)
		}
	}

	return r0, r1
}

// ListBucketMetricsConfigurations provides a mock function with given fields: _a0
func (_m *S3API) ListBucketMetricsConfigurations(_a0 *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListBucketMetricsConfigurationsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListBucketMetricsConfigurationsInput) *s3.ListBucketMetricsConfigurationsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketMetricsConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListBucketMetricsConfigurationsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketMetricsConfigurationsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListBucketMetricsConfigurationsWithContext(_a0 aws.Context, _a1 *s3.ListBucketMetricsConfigurationsInput, _a2 ...request.Option) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListBucketMetricsConfigurationsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListBucketMetricsConfigurationsInput, ...request.Option) *s3.ListBucketMetricsConfigurationsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketMetricsConfigurationsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListBucketMetricsConfigurationsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketMetricsConfigurationsRequest provides a mock function with given fields: _a0
func (_m *S3API) ListBucketMetricsConfigurationsRequest(_a0 *s3.ListBucketMetricsConfigurationsInput) (*request.Request, *s3.ListBucketMetricsConfigurationsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListBucketMetricsConfigurationsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListBucketMetricsConfigurationsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListBucketMetricsConfigurationsInput) *s3.ListBucketMetricsConfigurationsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListBucketMetricsConfigurationsOutput)
		}
	}

	return r0, r1
}

// ListBuckets provides a mock function with given fields: _a0
func (_m *S3API) ListBuckets(_a0 *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListBucketsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListBucketsInput) *s3.ListBucketsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListBucketsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListBucketsWithContext(_a0 aws.Context, _a1 *s3.ListBucketsInput, _a2 ...request.Option) (*s3.ListBucketsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListBucketsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListBucketsInput, ...request.Option) *s3.ListBucketsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListBucketsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListBucketsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBucketsRequest provides a mock function with given fields: _a0
func (_m *S3API) ListBucketsRequest(_a0 *s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListBucketsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListBucketsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListBucketsInput) *s3.ListBucketsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListBucketsOutput)
		}
	}

	return r0, r1
}

// ListMultipartUploads provides a mock function with given fields: _a0
func (_m *S3API) ListMultipartUploads(_a0 *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListMultipartUploadsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListMultipartUploadsInput) *s3.ListMultipartUploadsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListMultipartUploadsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListMultipartUploadsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMultipartUploadsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListMultipartUploadsWithContext(_a0 aws.Context, _a1 *s3.ListMultipartUploadsInput, _a2 ...request.Option) (*s3.ListMultipartUploadsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListMultipartUploadsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListMultipartUploadsInput, ...request.Option) *s3.ListMultipartUploadsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListMultipartUploadsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListMultipartUploadsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMultipartUploadsRequest provides a mock function with given fields: _a0
func (_m *S3API) ListMultipartUploadsRequest(_a0 *s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListMultipartUploadsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListMultipartUploadsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListMultipartUploadsInput) *s3.ListMultipartUploadsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListMultipartUploadsOutput)
		}
	}

	return r0, r1
}

// ListMultipartUploadsPages provides a mock function with given fields: _a0, _a1
func (_m *S3API) ListMultipartUploadsPages(_a0 *s3.ListMultipartUploadsInput, _a1 func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListMultipartUploadsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *S3API) ListMultipartUploadsPagesWithContext(_a0 aws.Context, _a1 *s3.ListMultipartUploadsInput, _a2 func(*s3.ListMultipartUploadsOutput, bool) bool, _a3 ...request.Option) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListObjectVersions provides a mock function with given fields: _a0
func (_m *S3API) ListObjectVersions(_a0 *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListObjectVersionsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListObjectVersionsInput) *s3.ListObjectVersionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectVersionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListObjectVersionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjectVersionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListObjectVersionsWithContext(_a0 aws.Context, _a1 *s3.ListObjectVersionsInput, _a2 ...request.Option) (*s3.ListObjectVersionsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListObjectVersionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListObjectVersionsInput, ...request.Option) *s3.ListObjectVersionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectVersionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListObjectVersionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjectVersionsRequest provides a mock function with given fields: _a0
func (_m *S3API) ListObjectVersionsRequest(_a0 *s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListObjectVersionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListObjectVersionsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListObjectVersionsInput) *s3.ListObjectVersionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListObjectVersionsOutput)
		}
	}

	return r0, r1
}

// ListObjectVersionsPages provides a mock function with given fields: _a0, _a1
func (_m *S3API) ListObjectVersionsPages(_a0 *s3.ListObjectVersionsInput, _a1 func(*s3.ListObjectVersionsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListObjectVersionsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *S3API) ListObjectVersionsPagesWithContext(_a0 aws.Context, _a1 *s3.ListObjectVersionsInput, _a2 func(*s3.ListObjectVersionsOutput, bool) bool, _a3 ...request.Option) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListObjects provides a mock function with given fields: _a0
func (_m *S3API) ListObjects(_a0 *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListObjectsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsInput) *s3.ListObjectsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListObjectsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjectsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListObjectsWithContext(_a0 aws.Context, _a1 *s3.ListObjectsInput, _a2 ...request.Option) (*s3.ListObjectsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListObjectsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListObjectsInput, ...request.Option) *s3.ListObjectsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListObjectsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjectsRequest provides a mock function with given fields: _a0
func (_m *S3API) ListObjectsRequest(_a0 *s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListObjectsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListObjectsInput) *s3.ListObjectsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListObjectsOutput)
		}
	}

	return r0, r1
}

// ListObjectsPages provides a mock function with given fields: _a0, _a1
func (_m *S3API) ListObjectsPages(_a0 *s3.ListObjectsInput, _a1 func(*s3.ListObjectsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListObjectsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *S3API) ListObjectsPagesWithContext(_a0 aws.Context, _a1 *s3.ListObjectsInput, _a2 func(*s3.ListObjectsOutput, bool) bool, _a3 ...request.Option) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListObjectsV2 provides a mock function with given fields: _a0
func (_m *S3API) ListObjectsV2(_a0 *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListObjectsV2Output
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsV2Input) *s3.ListObjectsV2Output); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectsV2Output)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListObjectsV2Input) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjectsV2WithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListObjectsV2WithContext(_a0 aws.Context, _a1 *s3.ListObjectsV2Input, _a2 ...request.Option) (*s3.ListObjectsV2Output, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListObjectsV2Output
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListObjectsV2Input, ...request.Option) *s3.ListObjectsV2Output); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectsV2Output)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListObjectsV2Input, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjectsV2Request provides a mock function with given fields: _a0
func (_m *S3API) ListObjectsV2Request(_a0 *s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsV2Input) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListObjectsV2Output
	if rf, ok := ret.Get(1).(func(*s3.ListObjectsV2Input) *s3.ListObjectsV2Output); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListObjectsV2Output)
		}
	}

	return r0, r1
}

// ListObjectsV2Pages provides a mock function with given fields: _a0, _a1
func (_m *S3API) ListObjectsV2Pages(_a0 *s3.ListObjectsV2Input, _a1 func(*s3.ListObjectsV2Output, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListObjectsV2PagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *S3API) ListObjectsV2PagesWithContext(_a0 aws.Context, _a1 *s3.ListObjectsV2Input, _a2 func(*s3.ListObjectsV2Output, bool) bool, _a3 ...request.Option) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListParts provides a mock function with given fields: _a0
func (_m *S3API) ListParts(_a0 *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.ListPartsOutput
	if rf, ok := ret.Get(0).(func(*s3.ListPartsInput) *s3.ListPartsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListPartsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.ListPartsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPartsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) ListPartsWithContext(_a0 aws.Context, _a1 *s3.ListPartsInput, _a2 ...request.Option) (*s3.ListPartsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.ListPartsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListPartsInput, ...request.Option) *s3.ListPartsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListPartsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.ListPartsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPartsRequest provides a mock function with given fields: _a0
func (_m *S3API) ListPartsRequest(_a0 *s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.ListPartsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.ListPartsOutput
	if rf, ok := ret.Get(1).(func(*s3.ListPartsInput) *s3.ListPartsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.ListPartsOutput)
		}
	}

	return r0, r1
}

// ListPartsPages provides a mock function with given fields: _a0, _a1
func (_m *S3API) ListPartsPages(_a0 *s3.ListPartsInput, _a1 func(*s3.ListPartsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListPartsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *S3API) ListPartsPagesWithContext(_a0 aws.Context, _a1 *s3.ListPartsInput, _a2 func(*s3.ListPartsOutput, bool) bool, _a3 ...request.Option) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutBucketAccelerateConfiguration provides a mock function with given fields: _a0
func (_m *S3API) PutBucketAccelerateConfiguration(_a0 *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketAccelerateConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAccelerateConfigurationInput) *s3.PutBucketAccelerateConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketAccelerateConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketAccelerateConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketAccelerateConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketAccelerateConfigurationWithContext(_a0 aws.Context, _a1 *s3.PutBucketAccelerateConfigurationInput, _a2 ...request.Option) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketAccelerateConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketAccelerateConfigurationInput, ...request.Option) *s3.PutBucketAccelerateConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketAccelerateConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketAccelerateConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketAccelerateConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketAccelerateConfigurationRequest(_a0 *s3.PutBucketAccelerateConfigurationInput) (*request.Request, *s3.PutBucketAccelerateConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAccelerateConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketAccelerateConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketAccelerateConfigurationInput) *s3.PutBucketAccelerateConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketAccelerateConfigurationOutput)
		}
	}

	return r0, r1
}

// PutBucketAcl provides a mock function with given fields: _a0
func (_m *S3API) PutBucketAcl(_a0 *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketAclOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAclInput) *s3.PutBucketAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketAclWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketAclWithContext(_a0 aws.Context, _a1 *s3.PutBucketAclInput, _a2 ...request.Option) (*s3.PutBucketAclOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketAclOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketAclInput, ...request.Option) *s3.PutBucketAclOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketAclInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketAclRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketAclRequest(_a0 *s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketAclOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketAclInput) *s3.PutBucketAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketAclOutput)
		}
	}

	return r0, r1
}

// PutBucketAnalyticsConfiguration provides a mock function with given fields: _a0
func (_m *S3API) PutBucketAnalyticsConfiguration(_a0 *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAnalyticsConfigurationInput) *s3.PutBucketAnalyticsConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketAnalyticsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketAnalyticsConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketAnalyticsConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketAnalyticsConfigurationWithContext(_a0 aws.Context, _a1 *s3.PutBucketAnalyticsConfigurationInput, _a2 ...request.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketAnalyticsConfigurationInput, ...request.Option) *s3.PutBucketAnalyticsConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketAnalyticsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketAnalyticsConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketAnalyticsConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketAnalyticsConfigurationRequest(_a0 *s3.PutBucketAnalyticsConfigurationInput) (*request.Request, *s3.PutBucketAnalyticsConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketAnalyticsConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketAnalyticsConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketAnalyticsConfigurationInput) *s3.PutBucketAnalyticsConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketAnalyticsConfigurationOutput)
		}
	}

	return r0, r1
}

// PutBucketCors provides a mock function with given fields: _a0
func (_m *S3API) PutBucketCors(_a0 *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketCorsOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketCorsInput) *s3.PutBucketCorsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketCorsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketCorsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketCorsWithContext(_a0 aws.Context, _a1 *s3.PutBucketCorsInput, _a2 ...request.Option) (*s3.PutBucketCorsOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketCorsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketCorsInput, ...request.Option) *s3.PutBucketCorsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketCorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketCorsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketCorsRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketCorsRequest(_a0 *s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketCorsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketCorsOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketCorsInput) *s3.PutBucketCorsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketCorsOutput)
		}
	}

	return r0, r1
}

// PutBucketInventoryConfiguration provides a mock function with given fields: _a0
func (_m *S3API) PutBucketInventoryConfiguration(_a0 *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketInventoryConfigurationInput) *s3.PutBucketInventoryConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketInventoryConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketInventoryConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketInventoryConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketInventoryConfigurationWithContext(_a0 aws.Context, _a1 *s3.PutBucketInventoryConfigurationInput, _a2 ...request.Option) (*s3.PutBucketInventoryConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketInventoryConfigurationInput, ...request.Option) *s3.PutBucketInventoryConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketInventoryConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketInventoryConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketInventoryConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketInventoryConfigurationRequest(_a0 *s3.PutBucketInventoryConfigurationInput) (*request.Request, *s3.PutBucketInventoryConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketInventoryConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketInventoryConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketInventoryConfigurationInput) *s3.PutBucketInventoryConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketInventoryConfigurationOutput)
		}
	}

	return r0, r1
}

// PutBucketLifecycle provides a mock function with given fields: _a0
func (_m *S3API) PutBucketLifecycle(_a0 *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLifecycleInput) *s3.PutBucketLifecycleOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLifecycleInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketLifecycleWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketLifecycleWithContext(_a0 aws.Context, _a1 *s3.PutBucketLifecycleInput, _a2 ...request.Option) (*s3.PutBucketLifecycleOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketLifecycleOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketLifecycleInput, ...request.Option) *s3.PutBucketLifecycleOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLifecycleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketLifecycleInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketLifecycleRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketLifecycleRequest(_a0 *s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLifecycleInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketLifecycleOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLifecycleInput) *s3.PutBucketLifecycleOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketLifecycleOutput)
		}
	}

	return r0, r1
}

// PutBucketLifecycleConfiguration provides a mock function with given fields: _a0
func (_m *S3API) PutBucketLifecycleConfiguration(_a0 *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLifecycleConfigurationInput) *s3.PutBucketLifecycleConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLifecycleConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLifecycleConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketLifecycleConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketLifecycleConfigurationWithContext(_a0 aws.Context, _a1 *s3.PutBucketLifecycleConfigurationInput, _a2 ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketLifecycleConfigurationInput, ...request.Option) *s3.PutBucketLifecycleConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLifecycleConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketLifecycleConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketLifecycleConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketLifecycleConfigurationRequest(_a0 *s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLifecycleConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketLifecycleConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLifecycleConfigurationInput) *s3.PutBucketLifecycleConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketLifecycleConfigurationOutput)
		}
	}

	return r0, r1
}

// PutBucketLogging provides a mock function with given fields: _a0
func (_m *S3API) PutBucketLogging(_a0 *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketLoggingOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLoggingInput) *s3.PutBucketLoggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLoggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketLoggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketLoggingWithContext(_a0 aws.Context, _a1 *s3.PutBucketLoggingInput, _a2 ...request.Option) (*s3.PutBucketLoggingOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketLoggingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketLoggingInput, ...request.Option) *s3.PutBucketLoggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketLoggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketLoggingRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketLoggingRequest(_a0 *s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketLoggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketLoggingOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketLoggingInput) *s3.PutBucketLoggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketLoggingOutput)
		}
	}

	return r0, r1
}

// PutBucketMetricsConfiguration provides a mock function with given fields: _a0
func (_m *S3API) PutBucketMetricsConfiguration(_a0 *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketMetricsConfigurationInput) *s3.PutBucketMetricsConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketMetricsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketMetricsConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketMetricsConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketMetricsConfigurationWithContext(_a0 aws.Context, _a1 *s3.PutBucketMetricsConfigurationInput, _a2 ...request.Option) (*s3.PutBucketMetricsConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketMetricsConfigurationInput, ...request.Option) *s3.PutBucketMetricsConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketMetricsConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketMetricsConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketMetricsConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketMetricsConfigurationRequest(_a0 *s3.PutBucketMetricsConfigurationInput) (*request.Request, *s3.PutBucketMetricsConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketMetricsConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketMetricsConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketMetricsConfigurationInput) *s3.PutBucketMetricsConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketMetricsConfigurationOutput)
		}
	}

	return r0, r1
}

// PutBucketNotification provides a mock function with given fields: _a0
func (_m *S3API) PutBucketNotification(_a0 *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketNotificationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketNotificationInput) *s3.PutBucketNotificationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketNotificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketNotificationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketNotificationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketNotificationWithContext(_a0 aws.Context, _a1 *s3.PutBucketNotificationInput, _a2 ...request.Option) (*s3.PutBucketNotificationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketNotificationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketNotificationInput, ...request.Option) *s3.PutBucketNotificationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketNotificationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketNotificationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketNotificationRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketNotificationRequest(_a0 *s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketNotificationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketNotificationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketNotificationInput) *s3.PutBucketNotificationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketNotificationOutput)
		}
	}

	return r0, r1
}

// PutBucketNotificationConfiguration provides a mock function with given fields: _a0
func (_m *S3API) PutBucketNotificationConfiguration(_a0 *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketNotificationConfigurationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketNotificationConfigurationInput) *s3.PutBucketNotificationConfigurationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketNotificationConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketNotificationConfigurationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketNotificationConfigurationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketNotificationConfigurationWithContext(_a0 aws.Context, _a1 *s3.PutBucketNotificationConfigurationInput, _a2 ...request.Option) (*s3.PutBucketNotificationConfigurationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketNotificationConfigurationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketNotificationConfigurationInput, ...request.Option) *s3.PutBucketNotificationConfigurationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketNotificationConfigurationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketNotificationConfigurationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketNotificationConfigurationRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketNotificationConfigurationRequest(_a0 *s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketNotificationConfigurationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketNotificationConfigurationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketNotificationConfigurationInput) *s3.PutBucketNotificationConfigurationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketNotificationConfigurationOutput)
		}
	}

	return r0, r1
}

// PutBucketPolicy provides a mock function with given fields: _a0
func (_m *S3API) PutBucketPolicy(_a0 *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketPolicyInput) *s3.PutBucketPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketPolicyWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketPolicyWithContext(_a0 aws.Context, _a1 *s3.PutBucketPolicyInput, _a2 ...request.Option) (*s3.PutBucketPolicyOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketPolicyOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketPolicyInput, ...request.Option) *s3.PutBucketPolicyOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketPolicyInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketPolicyRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketPolicyRequest(_a0 *s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketPolicyOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketPolicyInput) *s3.PutBucketPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketPolicyOutput)
		}
	}

	return r0, r1
}

// PutBucketReplication provides a mock function with given fields: _a0
func (_m *S3API) PutBucketReplication(_a0 *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketReplicationInput) *s3.PutBucketReplicationOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketReplicationInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketReplicationWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketReplicationWithContext(_a0 aws.Context, _a1 *s3.PutBucketReplicationInput, _a2 ...request.Option) (*s3.PutBucketReplicationOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketReplicationOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketReplicationInput, ...request.Option) *s3.PutBucketReplicationOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketReplicationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketReplicationInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketReplicationRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketReplicationRequest(_a0 *s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketReplicationInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketReplicationOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketReplicationInput) *s3.PutBucketReplicationOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketReplicationOutput)
		}
	}

	return r0, r1
}

// PutBucketRequestPayment provides a mock function with given fields: _a0
func (_m *S3API) PutBucketRequestPayment(_a0 *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketRequestPaymentOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketRequestPaymentInput) *s3.PutBucketRequestPaymentOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketRequestPaymentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketRequestPaymentInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketRequestPaymentWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketRequestPaymentWithContext(_a0 aws.Context, _a1 *s3.PutBucketRequestPaymentInput, _a2 ...request.Option) (*s3.PutBucketRequestPaymentOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketRequestPaymentOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketRequestPaymentInput, ...request.Option) *s3.PutBucketRequestPaymentOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketRequestPaymentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketRequestPaymentInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketRequestPaymentRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketRequestPaymentRequest(_a0 *s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketRequestPaymentInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketRequestPaymentOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketRequestPaymentInput) *s3.PutBucketRequestPaymentOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketRequestPaymentOutput)
		}
	}

	return r0, r1
}

// PutBucketTagging provides a mock function with given fields: _a0
func (_m *S3API) PutBucketTagging(_a0 *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketTaggingInput) *s3.PutBucketTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketTaggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketTaggingWithContext(_a0 aws.Context, _a1 *s3.PutBucketTaggingInput, _a2 ...request.Option) (*s3.PutBucketTaggingOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketTaggingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketTaggingInput, ...request.Option) *s3.PutBucketTaggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketTaggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketTaggingRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketTaggingRequest(_a0 *s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketTaggingInput) *s3.PutBucketTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketTaggingOutput)
		}
	}

	return r0, r1
}

// PutBucketVersioning provides a mock function with given fields: _a0
func (_m *S3API) PutBucketVersioning(_a0 *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketVersioningOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketVersioningInput) *s3.PutBucketVersioningOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketVersioningOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketVersioningInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketVersioningWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketVersioningWithContext(_a0 aws.Context, _a1 *s3.PutBucketVersioningInput, _a2 ...request.Option) (*s3.PutBucketVersioningOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketVersioningOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketVersioningInput, ...request.Option) *s3.PutBucketVersioningOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketVersioningOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketVersioningInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketVersioningRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketVersioningRequest(_a0 *s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketVersioningInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketVersioningOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketVersioningInput) *s3.PutBucketVersioningOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketVersioningOutput)
		}
	}

	return r0, r1
}

// PutBucketWebsite provides a mock function with given fields: _a0
func (_m *S3API) PutBucketWebsite(_a0 *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(*s3.PutBucketWebsiteInput) *s3.PutBucketWebsiteOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutBucketWebsiteInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketWebsiteWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutBucketWebsiteWithContext(_a0 aws.Context, _a1 *s3.PutBucketWebsiteInput, _a2 ...request.Option) (*s3.PutBucketWebsiteOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutBucketWebsiteOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutBucketWebsiteInput, ...request.Option) *s3.PutBucketWebsiteOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutBucketWebsiteOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutBucketWebsiteInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBucketWebsiteRequest provides a mock function with given fields: _a0
func (_m *S3API) PutBucketWebsiteRequest(_a0 *s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutBucketWebsiteInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutBucketWebsiteOutput
	if rf, ok := ret.Get(1).(func(*s3.PutBucketWebsiteInput) *s3.PutBucketWebsiteOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutBucketWebsiteOutput)
		}
	}

	return r0, r1
}

// PutObject provides a mock function with given fields: _a0
func (_m *S3API) PutObject(_a0 *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.PutObjectInput) *s3.PutObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutObjectWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutObjectWithContext(_a0 aws.Context, _a1 *s3.PutObjectInput, _a2 ...request.Option) (*s3.PutObjectOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutObjectOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutObjectInput, ...request.Option) *s3.PutObjectOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutObjectInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutObjectRequest provides a mock function with given fields: _a0
func (_m *S3API) PutObjectRequest(_a0 *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.PutObjectInput) *s3.PutObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutObjectOutput)
		}
	}

	return r0, r1
}

// PutObjectAcl provides a mock function with given fields: _a0
func (_m *S3API) PutObjectAcl(_a0 *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutObjectAclOutput
	if rf, ok := ret.Get(0).(func(*s3.PutObjectAclInput) *s3.PutObjectAclOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutObjectAclInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutObjectAclWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutObjectAclWithContext(_a0 aws.Context, _a1 *s3.PutObjectAclInput, _a2 ...request.Option) (*s3.PutObjectAclOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutObjectAclOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutObjectAclInput, ...request.Option) *s3.PutObjectAclOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectAclOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutObjectAclInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutObjectAclRequest provides a mock function with given fields: _a0
func (_m *S3API) PutObjectAclRequest(_a0 *s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutObjectAclInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutObjectAclOutput
	if rf, ok := ret.Get(1).(func(*s3.PutObjectAclInput) *s3.PutObjectAclOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutObjectAclOutput)
		}
	}

	return r0, r1
}

// PutObjectTagging provides a mock function with given fields: _a0
func (_m *S3API) PutObjectTagging(_a0 *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.PutObjectTaggingOutput
	if rf, ok := ret.Get(0).(func(*s3.PutObjectTaggingInput) *s3.PutObjectTaggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.PutObjectTaggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutObjectTaggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) PutObjectTaggingWithContext(_a0 aws.Context, _a1 *s3.PutObjectTaggingInput, _a2 ...request.Option) (*s3.PutObjectTaggingOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.PutObjectTaggingOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.PutObjectTaggingInput, ...request.Option) *s3.PutObjectTaggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectTaggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.PutObjectTaggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutObjectTaggingRequest provides a mock function with given fields: _a0
func (_m *S3API) PutObjectTaggingRequest(_a0 *s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.PutObjectTaggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.PutObjectTaggingOutput
	if rf, ok := ret.Get(1).(func(*s3.PutObjectTaggingInput) *s3.PutObjectTaggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.PutObjectTaggingOutput)
		}
	}

	return r0, r1
}

// RestoreObject provides a mock function with given fields: _a0
func (_m *S3API) RestoreObject(_a0 *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.RestoreObjectOutput
	if rf, ok := ret.Get(0).(func(*s3.RestoreObjectInput) *s3.RestoreObjectOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.RestoreObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.RestoreObjectInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreObjectWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) RestoreObjectWithContext(_a0 aws.Context, _a1 *s3.RestoreObjectInput, _a2 ...request.Option) (*s3.RestoreObjectOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.RestoreObjectOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.RestoreObjectInput, ...request.Option) *s3.RestoreObjectOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.RestoreObjectOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.RestoreObjectInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreObjectRequest provides a mock function with given fields: _a0
func (_m *S3API) RestoreObjectRequest(_a0 *s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.RestoreObjectInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.RestoreObjectOutput
	if rf, ok := ret.Get(1).(func(*s3.RestoreObjectInput) *s3.RestoreObjectOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.RestoreObjectOutput)
		}
	}

	return r0, r1
}

// UploadPart provides a mock function with given fields: _a0
func (_m *S3API) UploadPart(_a0 *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.UploadPartOutput
	if rf, ok := ret.Get(0).(func(*s3.UploadPartInput) *s3.UploadPartOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.UploadPartOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.UploadPartInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadPartWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) UploadPartWithContext(_a0 aws.Context, _a1 *s3.UploadPartInput, _a2 ...request.Option) (*s3.UploadPartOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.UploadPartOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.UploadPartInput, ...request.Option) *s3.UploadPartOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.UploadPartOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.UploadPartInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadPartRequest provides a mock function with given fields: _a0
func (_m *S3API) UploadPartRequest(_a0 *s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.UploadPartInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.UploadPartOutput
	if rf, ok := ret.Get(1).(func(*s3.UploadPartInput) *s3.UploadPartOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.UploadPartOutput)
		}
	}

	return r0, r1
}

// UploadPartCopy provides a mock function with given fields: _a0
func (_m *S3API) UploadPartCopy(_a0 *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *s3.UploadPartCopyOutput
	if rf, ok := ret.Get(0).(func(*s3.UploadPartCopyInput) *s3.UploadPartCopyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.UploadPartCopyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.UploadPartCopyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadPartCopyWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) UploadPartCopyWithContext(_a0 aws.Context, _a1 *s3.UploadPartCopyInput, _a2 ...request.Option) (*s3.UploadPartCopyOutput, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *s3.UploadPartCopyOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.UploadPartCopyInput, ...request.Option) *s3.UploadPartCopyOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.UploadPartCopyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *s3.UploadPartCopyInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadPartCopyRequest provides a mock function with given fields: _a0
func (_m *S3API) UploadPartCopyRequest(_a0 *s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*s3.UploadPartCopyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *s3.UploadPartCopyOutput
	if rf, ok := ret.Get(1).(func(*s3.UploadPartCopyInput) *s3.UploadPartCopyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s3.UploadPartCopyOutput)
		}
	}

	return r0, r1
}

// WaitUntilBucketExists provides a mock function with given fields: _a0
func (_m *S3API) WaitUntilBucketExists(_a0 *s3.HeadBucketInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.HeadBucketInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilBucketExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) WaitUntilBucketExistsWithContext(_a0 aws.Context, _a1 *s3.HeadBucketInput, _a2 ...request.WaiterOption) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.HeadBucketInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilBucketNotExists provides a mock function with given fields: _a0
func (_m *S3API) WaitUntilBucketNotExists(_a0 *s3.HeadBucketInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.HeadBucketInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilBucketNotExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) WaitUntilBucketNotExistsWithContext(_a0 aws.Context, _a1 *s3.HeadBucketInput, _a2 ...request.WaiterOption) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.HeadBucketInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilObjectExists provides a mock function with given fields: _a0
func (_m *S3API) WaitUntilObjectExists(_a0 *s3.HeadObjectInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.HeadObjectInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilObjectExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) WaitUntilObjectExistsWithContext(_a0 aws.Context, _a1 *s3.HeadObjectInput, _a2 ...request.WaiterOption) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.HeadObjectInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilObjectNotExists provides a mock function with given fields: _a0
func (_m *S3API) WaitUntilObjectNotExists(_a0 *s3.HeadObjectInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*s3.HeadObjectInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilObjectNotExistsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *S3API) WaitUntilObjectNotExistsWithContext(_a0 aws.Context, _a1 *s3.HeadObjectInput, _a2 ...request.WaiterOption) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *s3.HeadObjectInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
