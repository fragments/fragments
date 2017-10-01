package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckDuplicates(t *testing.T) {
	functionA := &functionModel{
		meta: &Meta{
			Name: "A",
		},
	}

	functionB := &functionModel{
		meta: &Meta{
			Name: "B",
		},
	}

	other := &testModel{
		meta: &Meta{
			Name: "A",
		},
	}

	tests := []struct {
		TestName string
		Models   []Model
		Error    bool
	}{
		{
			TestName: "Duplicates",
			Models: []Model{
				functionA,
				functionA,
				functionB,
			},
			Error: true,
		},
		{
			TestName: "No duplicates",
			Models: []Model{
				functionA,
				functionB,
			},
		},
		{
			TestName: "Other type",
			Models: []Model{
				functionA,
				other,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			err := CheckDuplicates(test.Models)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

type testModel struct {
	meta *Meta
}

func (t *testModel) File() string       { return "test" }
func (t *testModel) Meta() *Meta        { return t.meta }
func (t *testModel) Type() ModelType    { return "test" }
func (t *testModel) testString() string { return "test" }
