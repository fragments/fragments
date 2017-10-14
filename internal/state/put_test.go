package state

import (
	"context"
	"fmt"
	"testing"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/pkg/snapshot"
	"github.com/stretchr/testify/require"
)

func TestPutModel(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		TestName  string
		Input     Model
		ModelType ModelType
		Error     bool
	}{
		{
			TestName: "No model",
			Error:    true,
		},
		{
			TestName: "No name",
			Input:    &Function{},
			Error:    true,
		},
		{
			TestName: "Function",
			Input: &Function{
				Meta: Meta{
					Name: "foo",
					Labels: map[string]string{
						"foo": "foo",
					},
				},
				Runtime:        "go",
				Checksum:       "abc",
				SourceFilename: "file.tar.gz",
			},
			ModelType: ModelTypeFunction,
		},
		{
			TestName: "Environment",
			Input: &Environment{
				Meta: Meta{
					Name: "foo",
					Labels: map[string]string{
						"bar": "bar",
					},
				},
				Infrastructure: InfrastructureTypeAWS,
			},
			ModelType: ModelTypeEnvironment,
		},
		{
			TestName: "Deployment",
			Input: &Deployment{
				Meta: Meta{
					Name: "foo",
					Labels: map[string]string{
						"bar": "bar",
					},
				},
				EnvironmentLabels: map[string]string{
					"foo": "foo",
				},
				FunctionLabels: map[string]string{
					"bar": "bar",
				},
			},
			ModelType: ModelTypeDeployment,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			kv := backend.NewTestKV()
			err := PutModel(ctx, kv, test.ModelType, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			snapshot.AssertString(t, kv.TestString(), fmt.Sprintf("testdata/TestPutModel-%s.txt", test.TestName), *update)
		})
	}
}

func TestPutPendingUpload(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		TestName string
		Input    *PendingUpload
		Token    string
		Error    bool
	}{
		{
			TestName: "NoPendingUpload",
			Error:    true,
		},
		{
			TestName: "NoFilename",
			Token:    "token",
			Input: &PendingUpload{
				Filename: "",
				Function: &Function{
					Meta: Meta{
						Name: "foo",
					},
				},
			},
			Error: true,
		},
		{
			TestName: "NoFunction",
			Token:    "token",
			Input: &PendingUpload{
				Filename: "foo.tar.gz",
				Function: nil,
			},
			Error: true,
		},
		{
			TestName: "NoToken",
			Input: &PendingUpload{
				Filename: "foo.tar.gz",
				Function: &Function{
					Meta: Meta{
						Name: "foo",
					},
				},
			},
			Error: true,
		},
		{
			TestName: "Ok",
			Token:    "token",
			Input: &PendingUpload{
				Filename: "foo.tar.gz",
				Function: &Function{
					Meta: Meta{
						Name: "foo",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			kv := backend.NewTestKV()
			err := PutPendingUpload(ctx, kv, test.Token, test.Input)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			snapshot.AssertString(t, kv.TestString(), fmt.Sprintf("testdata/TestPutPendingUpload-%s.txt", test.TestName), *update)
		})
	}
}
