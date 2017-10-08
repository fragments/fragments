// +build integration

package backend

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewETCDClient(t *testing.T) {
	tests := []struct {
		TestName  string
		Endpoints []string
		Error     bool
	}{
		{
			TestName:  "No endpoints",
			Endpoints: nil,
			Error:     true,
		},
		{
			TestName:  "No listener",
			Endpoints: []string{"127.0.0.1:1"},
			Error:     true,
		},
		{
			TestName:  "OK",
			Endpoints: []string{testETCDEndpoint},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			client, err := NewETCDClient(test.Endpoints, 250*time.Millisecond)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			defer func() {
				_ = client.Close()
			}()
		})
	}
}
