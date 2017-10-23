package mocks

import "github.com/stretchr/testify/mock"

import "io"

type SourceReader struct {
	mock.Mock
}

// GetFile provides a mock function with given fields: filename
func (_m *SourceReader) GetFile(filename string) (io.ReadCloser, error) {
	ret := _m.Called(filename)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Get(0).(io.ReadCloser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
