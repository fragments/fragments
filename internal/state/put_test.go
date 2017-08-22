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
