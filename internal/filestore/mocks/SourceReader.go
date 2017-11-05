package mocks

import "github.com/stretchr/testify/mock"

import "os"

type SourceReader struct {
	mock.Mock
}

// GetFile provides a mock function with given fields: filename
func (_m *SourceReader) GetFile(filename string) (*os.File, error) {
	ret := _m.Called(filename)

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(string) *os.File); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
