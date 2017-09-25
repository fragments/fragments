package state

import (
	"context"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListDeployments(t *testing.T) {
	foo := &Deployment{
		Meta: Meta{
			Name: "foo",
			Labels: map[string]string{
				"foo": "foo",
				"bar": "bar",
			},
		},
	}

	bar := &Deployment{
		Meta: Meta{
			Name: "bar",
			Labels: map[string]string{
				"bar": "bar",
				"baz": "baz",
			},
		},
	}

	baz := &Deployment{
		Meta: Meta{
			Name: "baz",
		},
	}

	ctx := context.Background()
	kv := backend.NewMemoryKV()
	err := PutModel(ctx, kv, ModelTypeDeployment, foo)
	require.NoError(t, err)
	err = PutModel(ctx, kv, ModelTypeDeployment, bar)
	require.NoError(t, err)
	err = PutModel(ctx, kv, ModelTypeDeployment, baz)
	require.NoError(t, err)

	tests := []struct {
		TestName string
		Matchers []matcher
		Expected []*Deployment
		Error    bool
	}{
		{
			TestName: "No matchers",
			Expected: []*Deployment{foo, bar, baz},
		},
		{
			TestName: "Label matcher (match all)",
			Expected: []*Deployment{foo, bar},
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"bar": "bar",
					},
				},
			},
		},
		{
			TestName: "Label matcher (match one)",
			Expected: []*Deployment{foo},
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"foo": "foo",
					},
				},
			},
		},
		{
			TestName: "Label matcher (match none)",
			Expected: []*Deployment{},
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"foo": "bar",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := ListDeployments(ctx, kv, test.Matchers...)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Len(t, actual, len(test.Expected))
		})
	}
}
