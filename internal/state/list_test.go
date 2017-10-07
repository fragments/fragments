package state

import (
	"context"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListDeployments(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutModel(ctx, kv, ModelTypeDeployment, &Deployment{
			Meta: Meta{
				Name: "foo",
				Labels: map[string]string{
					"foo": "foo",
					"bar": "bar",
				},
			},
		})
		require.NoError(t, err)
		err = PutModel(ctx, kv, ModelTypeDeployment, &Deployment{
			Meta: Meta{
				Name: "bar",
				Labels: map[string]string{
					"bar": "bar",
					"baz": "baz",
				},
			},
		})
		require.NoError(t, err)
		err = PutModel(ctx, kv, ModelTypeDeployment, &Deployment{
			Meta: Meta{
				Name: "baz",
			},
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "TestListDeployments.json")
	}

	ctx := context.Background()
	kv := backend.NewTestKV("TestListDeployments.json")

	tests := []struct {
		TestName string
		Matchers []matcher
		Results  int
		Error    bool
	}{
		{
			TestName: "No matchers",
			Results:  3,
		},
		{
			TestName: "Label matcher (match all)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"bar": "bar",
					},
				},
			},
			Results: 2,
		},
		{
			TestName: "Label matcher (match one)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"foo": "foo",
					},
				},
			},
			Results: 1,
		},
		{
			TestName: "Label matcher (match none)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"foo": "bar",
					},
				},
			},
			Results: 0,
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
			assert.Len(t, actual, test.Results)
		})
	}
}

func TestListEnvironments(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutModel(ctx, kv, ModelTypeEnvironment, &Environment{
			Meta: Meta{
				Name: "foo",
				Labels: map[string]string{
					"foo": "foo",
					"bar": "bar",
				},
			},
		})
		require.NoError(t, err)
		err = PutModel(ctx, kv, ModelTypeEnvironment, &Environment{
			Meta: Meta{
				Name: "bar",
				Labels: map[string]string{
					"bar": "bar",
					"baz": "baz",
				},
			},
		})
		require.NoError(t, err)
		err = PutModel(ctx, kv, ModelTypeEnvironment, &Environment{
			Meta: Meta{
				Name: "baz",
			},
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "TestListEnvironments.json")
	}

	ctx := context.Background()
	kv := backend.NewTestKV("TestListEnvironments.json")

	tests := []struct {
		TestName string
		Matchers []matcher
		Results  int
		Error    bool
	}{
		{
			TestName: "No matchers",
			Results:  3,
		},
		{
			TestName: "Label matcher (match all)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"bar": "bar",
					},
				},
			},
			Results: 2,
		},
		{
			TestName: "Label matcher (match one)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"foo": "foo",
					},
				},
			},
			Results: 1,
		},
		{
			TestName: "Label matcher (match none)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"foo": "bar",
					},
				},
			},
			Results: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := ListEnvironments(ctx, kv, test.Matchers...)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Len(t, actual, test.Results)
		})
	}
}

func TestListFunctions(t *testing.T) {
	if *update {
		kv := backend.NewTestKV()
		ctx := context.Background()
		err := PutModel(ctx, kv, ModelTypeFunction, &Function{
			Meta: Meta{
				Name: "foo",
				Labels: map[string]string{
					"foo": "foo",
					"bar": "bar",
				},
			},
		})
		require.NoError(t, err)
		err = PutModel(ctx, kv, ModelTypeFunction, &Function{
			Meta: Meta{
				Name: "bar",
				Labels: map[string]string{
					"bar": "bar",
					"baz": "baz",
				},
			},
		})
		require.NoError(t, err)
		err = PutModel(ctx, kv, ModelTypeFunction, &Function{
			Meta: Meta{
				Name: "baz",
			},
		})
		require.NoError(t, err)
		kv.SaveSnapshot(t, "TestListFunctions.json")
	}

	ctx := context.Background()
	kv := backend.NewTestKV("TestListFunctions.json")

	tests := []struct {
		TestName string
		Matchers []matcher
		Results  int
		Error    bool
	}{
		{
			TestName: "No matchers",
			Results:  3,
		},
		{
			TestName: "Label matcher (match all)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"bar": "bar",
					},
				},
			},
			Results: 2,
		},
		{
			TestName: "Label matcher (match one)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"foo": "foo",
					},
				},
			},
			Results: 1,
		},
		{
			TestName: "Label matcher (match none)",
			Matchers: []matcher{
				&LabelMatcher{
					Labels: map[string]string{
						"foo": "bar",
					},
				},
			},
			Results: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			actual, err := ListFunctions(ctx, kv, test.Matchers...)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Len(t, actual, test.Results)
		})
	}
}
