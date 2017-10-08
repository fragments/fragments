// +build integration

package backend

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewVaultClient(t *testing.T) {
	tests := []struct {
		TestName string
		Address  string
		Error    bool
	}{
		{
			TestName: "No address",
			Address:  "",
			Error:    true,
		},
		{
			TestName: "Invalid format (no protocol)",
			Address:  "127.0.0.1:8200",
			Error:    true,
		},
		{
			TestName: "OK",
			Address:  testVaultEndpoint,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			_, err := NewVaultClient(test.Address)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
