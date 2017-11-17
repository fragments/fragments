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
	Now = func() time.Time {
		mockTime, _ := time.Parse(time.RFC3339, "2017-10-31T12:34:56+02:00")
		return mockTime
	}

	kv := backend.NewTestKV()
	r := &ResPointer{
		InfraType:    InfraType("testinfra"),
		ResourceType: ResourceType("testres"),
		Name:         "existing",
	}
	err := r.Put(context.Background(), kv, resData{Value: "test"})
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
	getTime1 := func() time.Time {
		time1, _ := time.Parse(time.RFC3339, "2017-02-01T10:34:56+02:00")
		return time1
	}

	getTime2 := func() time.Time {
		time2, _ := time.Parse(time.RFC3339, "2017-10-10T11:34:56+02:00")
		return time2
	}

	initial := backend.NewTestKV()
	r := &ResPointer{
		InfraType:    InfraType("testinfra"),
		ResourceType: ResourceType("testres"),
		Name:         "existing",
	}
	Now = getTime1
	err := r.Put(context.Background(), initial, resData{Value: "existing"})
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
			Now = getTime2
			payload := resData{
				Value: "new",
			}
			err := r.Put(ctx, kv, &payload)
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
			err = r.Put(ctx, kv, data)
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
