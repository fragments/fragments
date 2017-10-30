// +build integration

package backend

import (
	"context"
	"io"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type LockerReaderWriter interface {
	Locker
	Reader
	Writer
}

var lockers = []struct {
	Name string
	New  func(t *testing.T) LockerReaderWriter
}{
	{
		Name: "ETCD",
		New: func(t *testing.T) LockerReaderWriter {
			etcd, err := NewETCDClient([]string{testETCDEndpoint}, 1*time.Second)
			require.NoError(t, err)
			return etcd
		},
	},
	{
		Name: "TestKV",
		New: func(t *testing.T) LockerReaderWriter {
			return NewTestKV()
		},
	},
}

func TestLockerLock(t *testing.T) {
	for _, target := range lockers {
		t.Run(target.Name, func(t *testing.T) {
			client := target.New(t)

			var counter int64

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

			// Delete key in case it previously existed
			_ = client.Delete(ctx, "/testkey")

			// Start many goroutines that all try to acquire the same lock. One will
			// succeed, try read the value that doesn't exist, and create the value.
			// After this the lock is released and the other goroutines can proceed.
			// The value exists at this point so the goroutine can exit.
			var wg sync.WaitGroup
			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					unlock, err := client.Lock(ctx, "/testlock")
					require.NoError(t, err)
					defer unlock()

					value, err := client.Get(ctx, "/testkey")
					if IsNotFound(err) {
						atomic.AddInt64(&counter, 1)
						err = client.Put(ctx, "/testkey", "value")
						require.NoError(t, err)
						return
					}
					require.NoError(t, err)

					assert.Equal(t, "value", value)
				}()
			}

			wg.Wait()

			assert.Equal(t, counter, int64(1))

			cancel()

			_, err := client.Lock(ctx, "/lockcancel")
			require.Error(t, err)

			if closer, ok := client.(io.Closer); ok {
				err := closer.Close()
				require.NoError(t, err)
			}
		})
	}
}
