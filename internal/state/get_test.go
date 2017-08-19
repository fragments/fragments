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

func TestGet(t *testing.T) {
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	function := &Function{}
	functionData, _ := json.Marshal(function)

	mockKV := &mocks.KV{}
	mockKV.
		On("Get", ctxCanceled, mock.AnythingOfType("string")).
		Return("", errors.New("context canceled"))
	mockKV.
		On("Get", ctx, resourcePath(ResourceTypeFunction, "corrupted")).
		Return("invalid json", nil)
	mockKV.
		On("Get", ctx, resourcePath(ResourceTypeFunction, "notfound")).
		Return("", &backend.ErrNotFound{Key: "notfound"})
	mockKV.
		On("Get", ctx, resourcePath(ResourceTypeFunction, "found")).
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
