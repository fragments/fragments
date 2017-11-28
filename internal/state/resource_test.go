package state

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type resData struct {
	Value string
}

func TestResourceGet(t *testing.T) {
	kv := backend.NewTestKV()
	r := &ResPointer{
		InfraType:    InfraType("testinfra"),
		ResourceType: ResourceType("testres"),
		Name:         "existing",
	}
	clock := testutils.NewMockClock()
	err := r.Put(context.Background(), kv, clock, resData{Value: "test"})
	require.NoError(t, err)

	tests := []struct {
		TestName  string
		Snapshot  string
		InfraType InfraType
		ResType   ResourceType
		Name      string
		Result    *resData
		Error     bool
	}{
		{
			TestName:  "NoExisting",
			InfraType: InfraType("nonexisting"),
			ResType:   ResourceType("nonexisting"),
			Name:      "nonexisting",
			Result:    nil,
		},
		{
			TestName:  "Existing",
			InfraType: InfraType("testinfra"),
			ResType:   ResourceType("testres"),
			Name:      "existing",
			Result:    &resData{Value: "test"},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			var payload resData
			r := &ResPointer{
				InfraType:    test.InfraType,
				ResourceType: test.ResType,
				Name:         test.Name,
			}
			exists, err := r.Get(context.Background(), kv, &payload)
			if test.Error {
				require.Error(t, err)
				return
			}
			fmt.Println("exists", exists)
			require.NoError(t, err)
			assert.Equal(t, test.Result != nil, exists)
			if exists {
				assert.Equal(t, *test.Result, payload)
			}
		})
	}
}

func TestResourcePut(t *testing.T) {
	initial := backend.NewTestKV()
	r := &ResPointer{
		InfraType:    InfraType("testinfra"),
		ResourceType: ResourceType("testres"),
		Name:         "existing",
	}
	clock := testutils.NewMockClock()
	err := r.Put(context.Background(), initial, clock, resData{Value: "existing"})
	require.NoError(t, err)

	tests := []struct {
		TestName  string
		InfraType InfraType
		ResType   ResourceType
		Name      string
		Error     bool
	}{
		{
			TestName:  "Create",
			InfraType: InfraType("new"),
			ResType:   ResourceType("new"),
			Name:      "new",
		},
		{
			TestName:  "Update",
			InfraType: InfraType("testinfra"),
			ResType:   ResourceType("testres"),
			Name:      "existing",
		},
	}

	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			kv := initial.Copy()
			r := &ResPointer{
				InfraType:    test.InfraType,
				ResourceType: test.ResType,
				Name:         test.Name,
			}
			payload := resData{
				Value: "new",
			}
			updateTime := testutils.NewMockClock().Add(24 * time.Hour)
			err := r.Put(ctx, kv, updateTime, &payload)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			testutils.AssertGolden(
				t,
				testutils.SnapshotJSONMap(kv.Data),
				fmt.Sprintf("testdata/TestResourcePut-%s.yaml", test.TestName),
			)
		})
	}
}

type lockData struct {
	Counter int
}

func TestLockResource(t *testing.T) {
	infraType := InfraType("infra")
	resType := ResourceType("resType")
	name := "lock"
	ctx := context.Background()
	kv := backend.NewTestKV()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := &ResPointer{
				InfraType:    infraType,
				ResourceType: resType,
				Name:         name,
			}
			unlock, err := r.Lock(ctx, kv)
			require.NoError(t, err)
			defer unlock()
			var data lockData
			existing, err := r.Get(ctx, kv, &data)
			require.NoError(t, err)
			if !existing {
				data = lockData{
					Counter: 1,
				}
			} else {
				data.Counter++
			}
			clock := testutils.NewMockClock()
			err = r.Put(ctx, kv, clock, data)
			require.NoError(t, err)
		}()
	}

	wg.Wait()

	var final lockData
	r := &ResPointer{
		InfraType:    infraType,
		ResourceType: resType,
		Name:         name,
	}
	existing, err := r.Get(ctx, kv, &final)
	require.NoError(t, err)
	require.True(t, existing)
	assert.Equal(t, final.Counter, 100)
}
