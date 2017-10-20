// +build integration

package backend

import (
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var readers = []struct {
	Name         string
	New          func(t *testing.T) Reader
	NewValidator func(t *testing.T, client Reader) validator
}{
	{
		Name: "ETCD",
		New: func(t *testing.T) Reader {
			etcd, err := NewETCDClient([]string{testETCDEndpoint}, 1*time.Second)
			require.NoError(t, err)
			return etcd
		},
		NewValidator: func(t *testing.T, _ Reader) validator {
			return newETCDValidator(t, testETCDEndpoint)
		},
	},
	{
		Name: "Vault",
		New: func(t *testing.T) Reader {
			vault, err := NewVaultClient(testVaultEndpoint)
			vault.client.SetToken(os.Getenv("VAULT_TEST_ROOT_TOKEN"))
			require.NoError(t, err)
			return vault
		},
		NewValidator: func(t *testing.T, _ Reader) validator {
			return newVaultValidator(t, testVaultEndpoint)
		},
	},
	{
		Name: "Test",
		New: func(t *testing.T) Reader {
			return NewTestKV()
		},
		NewValidator: func(t *testing.T, client Reader) validator {
			testkv := client.(*TestKV)
			return &testkvValidator{data: testkv.Data}
		},
	},
}

func TestReaderGet(t *testing.T) {
	for _, target := range readers {
		t.Run(target.Name, func(t *testing.T) {
			client := target.New(t)
			n := target.NewValidator(t, client)

			ctx, cancel := context.WithCancel(context.Background())

			_, err := client.Get(ctx, "foo")
			require.Error(t, err)
			assert.True(t, IsNotFound(err), "error is not NotFoundError")

			n.put(t, "foo", "foo")

			val, err := client.Get(ctx, "foo")
			require.NoError(t, err)
			assert.Equal(t, "foo", val)

			cancel()

			_, err = client.Get(ctx, "foo")
			require.Error(t, err)

			if closer, ok := client.(io.Closer); ok {
				err := closer.Close()
				require.NoError(t, err)
			}
		})
	}
}
