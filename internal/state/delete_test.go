package state

import (
	"context"
	"errors"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/internal/backend/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDeletePendingUpload(t *testing.T) {
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	matchCtx := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() == nil })
	matchCtxCancel := mock.MatchedBy(func(ctx context.Context) bool { return ctx.Err() != nil })

	mockKV := &mocks.KV{}
	mockKV.
		On("Delete", matchCtxCancel, mock.AnythingOfType("string")).
		Return(errors.New("context canceled"))
	mockKV.
		On("Delete", matchCtx, uploadPath("notfound")).
		Return(&backend.ErrNotFound{Key: "notfound"})
	mockKV.
		On("Delete", matchCtx, uploadPath("found")).
		Return(nil)

	tests := []struct {
		TestName string
		Ctx      context.Context
		Token    string
		Error    bool
	}{
		{
			TestName: "Context canceled",
			Ctx:      ctxCanceled,
			Error:    true,
		},
		{
			TestName: "Not found",
			Ctx:      ctx,
			Token:    "notfound",
			Error:    true,
		},
		{
			TestName: "Ok",
			Ctx:      ctx,
			Token:    "found",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := DeletePendingUpload(test.Ctx, mockKV, test.Token)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
