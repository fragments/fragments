// +build integration

package backend

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var listers = []struct {
	Name         string
	New          func(t *testing.T) Lister
	NewValidator func(t *testing.T, client Lister) validator
}{
	{
		Name: "ETCD",
		New: func(t *testing.T) Lister {
			etcd, err := NewETCDClient([]string{testETCDEndpoint}, 1*time.Second)
			require.NoError(t, err)
			return etcd
		},
		NewValidator: func(t *testing.T, _ Lister) validator {
			return newETCDValidator(t, testETCDEndpoint)
		},
	},
	{
		Name: "Test",
		New: func(t *testing.T) Lister {
			return NewTestKV()
		},
		NewValidator: func(t *testing.T, client Lister) validator {
			testkv := client.(*TestKV)
			return &testkvValidator{data: testkv.Data}
		},
	},
}

func TestListerList(t *testing.T) {
	for _, target := range listers {
		t.Run(target.Name, func(t *testing.T) {
			client := target.New(t)
			v := target.NewValidator(t, client)

			ctx, cancel := context.WithCancel(context.Background())

			values, err := client.List(ctx, "/foo")
			require.NoError(t, err)
			assert.Empty(t, values)

			v.put(t, "/foo/foo", "foo")
			v.put(t, "/foo/bar", "bar")
			v.put(t, "/bar/foo", "baz")

			values, err = client.List(ctx, "/foo")
			require.NoError(t, err)
			assert.Len(t, values, 2)
			assert.Equal(t, "foo", values["foo"])
			assert.Equal(t, "bar", values["bar"])

			cancel()

			_, err = client.List(ctx, "foo")
			require.Error(t, err)
		})
	}
}
