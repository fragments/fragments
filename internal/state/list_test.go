package state

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListDeployments(t *testing.T) {
	snapshotFile := "testdata/TestListDeployments.yaml"
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
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

	ctx := context.Background()
	kv := backend.NewTestKV()
	kv.LoadSnapshot(snapshotFile)

	tests := []struct {
		TestName string
		Matchers []matcher
		Results  int
		Error    bool
	}{
		{
			TestName: "NoMatchers",
			Results:  3,
		},
		{
			TestName: "LabelMatcherMatchAll",
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
			TestName: "LabelMatcherMatchOne",
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
			TestName: "LabelMatcherMatchNone",
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
	snapshotFile := "testdata/TestListEnvironments.yaml"
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
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

	ctx := context.Background()
	kv := backend.NewTestKV()
	kv.LoadSnapshot(snapshotFile)

	tests := []struct {
		TestName string
		Matchers []matcher
		Results  int
		Error    bool
	}{
		{
			TestName: "NoMatchers",
			Results:  3,
		},
		{
			TestName: "LabelMatcherMatchAll",
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
			TestName: "LabelMatcherMatchOne",
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
			TestName: "LabelMatcherMatchNone",
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
	snapshotFile := "testdata/TestListFunctions.yaml"
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
		data := kv.Snapshot()
		if err := ioutil.WriteFile(snapshotFile, []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}

	ctx := context.Background()
	kv := backend.NewTestKV()
	kv.LoadSnapshot(snapshotFile)

	tests := []struct {
		TestName string
		Matchers []matcher
		Results  int
		Error    bool
	}{
		{
			TestName: "NoMatchers",
			Results:  3,
		},
		{
			TestName: "LabelMatcherMatchAll",
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
			TestName: "LabelMatcherMatchOne",
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
			TestName: "LabelMatcherMatchNone",
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
