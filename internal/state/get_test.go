package state

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/internal/backend/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetFunction(t *testing.T) {
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	function := &Function{}
	functionData, _ := json.Marshal(function)

	matchCtx := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() == nil })
	matchCtxCancel := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() != nil })

	mockKV := &mocks.KV{}
	mockKV.
		On("Get", matchCtxCancel, mock.AnythingOfType("string")).
		Return("", errors.New("context canceled"))
	mockKV.
		On("Get", matchCtx, resourcePath(ResourceTypeFunction, "corrupted")).
		Return("invalid json", nil)
	mockKV.
		On("Get", matchCtx, resourcePath(ResourceTypeFunction, "notfound")).
		Return("", &backend.ErrNotFound{Key: "notfound"})
	mockKV.
		On("Get", matchCtx, resourcePath(ResourceTypeFunction, "found")).
		Return(string(functionData), nil)

	tests := []struct {
		TestName string
		Ctx      context.Context
		Name     string
		Error    bool
		Expected *Function
	}{
		{
			TestName: "Context canceled",
			Ctx:      ctxCanceled,
			Error:    true,
		},
		{
			TestName: "Unmarshal err",
			Ctx:      ctx,
			Name:     "corrupted",
			Error:    true,
		},
		{
			TestName: "Not found",
			Ctx:      ctx,
			Name:     "notfound",
		},
		{
			TestName: "Ok",
			Ctx:      ctx,
			Name:     "found",
			Expected: function,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := GetFunction(test.Ctx, mockKV, test.Name)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, test.Expected, actual)
		})
	}
}

func TestGetPendingUpload(t *testing.T) {
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	pendinUpload := &PendingUpload{}
	uploadData, _ := json.Marshal(pendinUpload)

	matchCtx := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() == nil })
	matchCtxCancel := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() != nil })

	mockKV := &mocks.KV{}
	mockKV.
		On("Get", matchCtxCancel, mock.AnythingOfType("string")).
		Return("", errors.New("context canceled"))
	mockKV.
		On("Get", matchCtx, uploadPath("corrupted")).
		Return("invalid json", nil)
	mockKV.
		On("Get", matchCtx, uploadPath("notfound")).
		Return("", &backend.ErrNotFound{Key: "notfound"})
	mockKV.
		On("Get", matchCtx, uploadPath("token")).
		Return(string(uploadData), nil)

	tests := []struct {
		TestName string
		Ctx      context.Context
		Token    string
		Error    bool
		Expected *PendingUpload
	}{
		{
			TestName: "Context canceled",
			Ctx:      ctxCanceled,
			Error:    true,
		},
		{
			TestName: "Unmarshal err",
			Ctx:      ctx,
			Token:    "corrupted",
			Error:    true,
		},
		{
			TestName: "Not found",
			Ctx:      ctx,
			Token:    "notfound",
		},
		{
			TestName: "Ok",
			Ctx:      ctx,
			Token:    "token",
			Expected: pendinUpload,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := GetPendingUpload(test.Ctx, mockKV, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, test.Expected, actual)
		})
	}
}
