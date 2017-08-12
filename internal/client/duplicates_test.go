package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckDuplicates(t *testing.T) {
	functionA := &functionResource{
		meta: &Meta{
			Name: "A",
		},
	}

	functionB := &functionResource{
		meta: &Meta{
			Name: "B",
		},
	}

	other := &testResource{
		meta: &Meta{
			Name: "A",
		},
	}

	tests := []struct {
		TestName  string
		Resources []Resource
		Error     bool
	}{
		{
			TestName: "Duplicates",
			Resources: []Resource{
				functionA,
				functionA,
				functionB,
			},
			Error: true,
		},
		{
			TestName: "No duplicates",
			Resources: []Resource{
				functionA,
				functionB,
			},
		},
		{
			TestName: "Other type",
			Resources: []Resource{
				functionA,
				other,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := CheckDuplicates(test.Resources)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

type testResource struct {
	meta *Meta
}

func (t *testResource) File() string       { return "test" }
func (t *testResource) Meta() *Meta        { return t.meta }
func (t *testResource) Type() ResourceType { return "test" }
func (t *testResource) testString() string { return "test" }
