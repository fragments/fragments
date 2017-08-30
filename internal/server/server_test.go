package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	fsmocks "github.com/fragments/fragments/internal/filestore/mocks"
	"github.com/fragments/fragments/internal/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutFunction(t *testing.T) {
	foo := &state.Function{
		Meta:     state.Meta{Name: "foo"},
		AWS:      &state.FunctionAWS{Memory: 256},
		Checksum: "foo",
	}
	fooCode := &state.Function{
		Meta:     state.Meta{Name: "foo"},
		AWS:      &state.FunctionAWS{Memory: 256},
		Checksum: "foobar",
	}
	fooConfig := &state.Function{
		Meta:     state.Meta{Name: "foo"},
		AWS:      &state.FunctionAWS{Memory: 512},
		Checksum: "foo",
	}

	tests := []struct {
		TestName string
		Input    *state.Function
		Existing *state.Function
		Token    string
		Result   *state.Function
		Expected *UploadRequest
		Error    bool
	}{
		{
			TestName: "No input",
			Input:    nil,
			Error:    true,
		},
		{
			TestName: "No name",
			Input:    &state.Function{},
			Error:    true,
		},
		{
			TestName: "No existing",
			Input:    foo,
			Token:    "token",
			Expected: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "Update code",
			Input:    fooCode,
			Existing: foo,
			Token:    "token",
			Expected: &UploadRequest{
				Token: "token",
				URL:   "https://token",
			},
		},
		{
			TestName: "Update config",
			Input:    fooConfig,
			Existing: foo,
			Result:   fooConfig,
		},
		{
			TestName: "No change",
			Input:    foo,
			Existing: foo,
			Expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			mockKV := backend.NewMemoryKV()
			if test.Existing != nil {
				_ = state.PutFunction(ctx, mockKV, test.Existing)
			}

			mockSourceStore := &fsmocks.SourceTarget{}
			mockSourceStore.
				On("NewUploadURL", test.Token).
				Return(fmt.Sprintf("https://%s", test.Token), nil)

			s := &Server{
				StateStore:    mockKV,
				SourceStore:   mockSourceStore,
				GenerateToken: func() string { return test.Token },
			}

			actual, err := s.PutFunction(ctx, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.Expected, actual)

			if test.Result != nil {
				current, _ := state.GetFunction(ctx, mockKV, test.Input.Meta.Name)
				assert.Equal(t, current, test.Input)
			}
		})
	}
}

func TestConfirmUpload(t *testing.T) {
	foo := &state.Function{
		Meta:     state.Meta{Name: "foo"},
		AWS:      &state.FunctionAWS{Memory: 256},
		Checksum: "foo",
	}
	uploadNew := &state.PendingUpload{
		Filename: "foo.tar.gz",
		Function: foo,
	}
	uploadUpdate := &state.PendingUpload{
		Filename:         "bar.tar.gz",
		PreviousFilename: "foo.tar.gz",
		Function:         foo,
	}

	tests := []struct {
		TestName string
		Token    string
		Pending  *state.PendingUpload
		Error    bool
	}{
		{
			TestName: "No token",
			Token:    "",
			Error:    true,
		},
		{
			TestName: "No pending upload",
			Token:    "token",
			Error:    true,
		},
		{
			TestName: "New upload",
			Token:    "token",
			Pending:  uploadNew,
		},
		{
			TestName: "Update code",
			Token:    "token",
			Pending:  uploadUpdate,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			ctx := context.Background()
			mockKV := backend.NewMemoryKV()
			if test.Pending != nil {
				_ = state.PutPendingUpload(ctx, mockKV, test.Token, test.Pending)
			}

			mockSourceStore := &fsmocks.SourceTarget{}
			mockSourceStore.
				On("Persist", ctx, test.Token).
				Return(nil)

			s := &Server{
				StateStore:    mockKV,
				SourceStore:   mockSourceStore,
				GenerateToken: func() string { return test.Token },
			}

			err := s.ConfirmUpload(ctx, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}

			mockSourceStore.AssertExpectations(t)

			require.NoError(t, err)

			u, _ := state.GetPendingUpload(ctx, mockKV, test.Token)
			assert.Nil(t, u)

			expected := test.Pending.Function
			expected.SourceFilename = test.Pending.Filename

			actual, _ := state.GetFunction(ctx, mockKV, expected.Meta.Name)
			assert.EqualValues(t, expected, actual)
		})
	}
}
