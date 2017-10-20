// +build integration

package backend

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var writers = []struct {
	Name         string
	New          func(t *testing.T) Writer
	NewValidator func(t *testing.T, client Writer) validator
}{
	{
		Name: "ETCD",
		New: func(t *testing.T) Writer {
			etcd, err := NewETCDClient([]string{testETCDEndpoint}, 1*time.Second)
			require.NoError(t, err)
			return etcd
		},
		NewValidator: func(t *testing.T, _ Writer) validator {
			return newETCDValidator(t, testETCDEndpoint)
		},
	},
	{
		Name: "Vault",
		New: func(t *testing.T) Writer {
			vault, err := NewVaultClient(testVaultEndpoint)
			vault.client.SetToken(os.Getenv("VAULT_TEST_ROOT_TOKEN"))
			require.NoError(t, err)
			return vault
		},
		NewValidator: func(t *testing.T, _ Writer) validator {
			return newVaultValidator(t, testVaultEndpoint)
		},
	},
	{
		Name: "Test",
		New: func(t *testing.T) Writer {
			return NewTestKV()
		},
		NewValidator: func(t *testing.T, client Writer) validator {
			testkv := client.(*TestKV)
			return &testkvValidator{data: testkv.Data}
		},
	},
}

func TestWriterWrite(t *testing.T) {
	for _, target := range writers {
		t.Run(target.Name, func(t *testing.T) {
			client := target.New(t)
			validator := target.NewValidator(t, client)

			ctx, cancel := context.WithCancel(context.Background())

			_, exists := validator.get(t, "foo")
			assert.False(t, exists)

			err := client.Put(ctx, "foo", "foo")
			require.NoError(t, err)
			val, exists := validator.get(t, "foo")
			assert.True(t, exists)
			assert.Equal(t, "foo", val)

			// overwrite
			err = client.Put(ctx, "foo", "bar")
			require.NoError(t, err)
			val, exists = validator.get(t, "foo")
			assert.True(t, exists)
			assert.Equal(t, "bar", val)

			cancel()

			err = client.Put(ctx, "foo", "baz")
			require.Error(t, err)

			val, exists = validator.get(t, "foo")
			assert.True(t, exists)
			assert.Equal(t, "bar", val)
		})
	}
}

func TestWriterDelete(t *testing.T) {
	for _, target := range writers {
		t.Run(target.Name, func(t *testing.T) {
			client := target.New(t)
			v := target.NewValidator(t, client)

			ctx, cancel := context.WithCancel(context.Background())

			err := client.Delete(ctx, "foo")
			require.Error(t, err)
			assert.True(t, IsNotFound(err), "error is not NotFoundError")

			v.put(t, "foo", "foo")
			_, exists := v.get(t, "foo")
			assert.True(t, exists)

			err = client.Delete(ctx, "foo")
			require.NoError(t, err)

			_, exists = v.get(t, "foo")
			assert.False(t, exists)

			cancel()

			err = client.Delete(ctx, "foo")
			require.Error(t, err)
		})
	}
}
