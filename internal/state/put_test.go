package state

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/fragments/fragments/internal/backend/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPutPendingUpload(t *testing.T) {
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	pendingUpload := &PendingUpload{}
	data, _ := json.Marshal(pendingUpload)

	matchCtx := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() == nil })
	matchCtxCancel := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() != nil })

	mockKV := &mocks.KV{}
	mockKV.
		On("Put", matchCtxCancel, mock.Anything, mock.Anything).
		Return(errors.New("context canceled"))
	mockKV.
		On("Put", matchCtx, "/uploads/token", string(data)).
		Return(nil)

	tests := []struct {
		TestName string
		Ctx      context.Context
		Token    string
		Payload  *PendingUpload
		Error    bool
	}{
		{
			TestName: "No token",
			Ctx:      ctx,
			Token:    "",
			Error:    true,
		},
		{
			TestName: "No payload",
			Ctx:      ctx,
			Token:    "token",
			Payload:  nil,
			Error:    true,
		},
		{
			TestName: "Context canceled",
			Ctx:      ctxCanceled,
			Token:    "token",
			Payload:  pendingUpload,
			Error:    true,
		},
		{
			TestName: "Ok",
			Ctx:      ctx,
			Token:    "token",
			Payload:  pendingUpload,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := PutPendingUpload(test.Ctx, mockKV, test.Token, test.Payload)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestPutFunction(t *testing.T) {
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	noNameFunction := &Function{}
	function := &Function{
		Meta: Meta{
			Name: "name",
		},
		Runtime: "go",
	}
	data, _ := json.Marshal(function)

	matchCtx := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() == nil })
	matchCtxCancel := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() != nil })

	mockKV := &mocks.KV{}
	mockKV.
		On("Put", matchCtxCancel, mock.Anything, mock.Anything).
		Return(errors.New("context canceled"))
	mockKV.
		On("Put", matchCtx, "/resources/function/name", string(data)).
		Return(nil)

	tests := []struct {
		TestName string
		Ctx      context.Context
		Payload  *Function
		Error    bool
	}{
		{
			TestName: "No payload",
			Ctx:      ctx,
			Payload:  nil,
			Error:    true,
		},
		{
			TestName: "No name",
			Ctx:      ctx,
			Payload:  noNameFunction,
			Error:    true,
		},
		{
			TestName: "Context canceled",
			Ctx:      ctxCanceled,
			Payload:  function,
			Error:    true,
		},
		{
			TestName: "Ok",
			Ctx:      ctx,
			Payload:  function,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := PutFunction(test.Ctx, mockKV, test.Payload)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestPutEnvironment(t *testing.T) {
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	noNameEnvironment := &Environment{}
	environment := &Environment{
		Meta: Meta{
			Name: "name",
		},
		Infrastructure: InfrastructureTypeAWS,
	}
	data, _ := json.Marshal(environment)

	matchCtx := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() == nil })
	matchCtxCancel := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() != nil })

	mockKV := &mocks.KV{}
	mockKV.
		On("Put", matchCtxCancel, mock.Anything, mock.Anything).
		Return(errors.New("context canceled"))
	mockKV.
		On("Put", matchCtx, "/resources/environment/name", string(data)).
		Return(nil)

	tests := []struct {
		TestName string
		Ctx      context.Context
		Payload  *Environment
		Error    bool
	}{
		{
			TestName: "No payload",
			Ctx:      ctx,
			Payload:  nil,
			Error:    true,
		},
		{
			TestName: "No name",
			Ctx:      ctx,
			Payload:  noNameEnvironment,
			Error:    true,
		},
		{
			TestName: "Context canceled",
			Ctx:      ctxCanceled,
			Payload:  environment,
			Error:    true,
		},
		{
			TestName: "Ok",
			Ctx:      ctx,
			Payload:  environment,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := PutEnvironment(test.Ctx, mockKV, test.Payload)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
