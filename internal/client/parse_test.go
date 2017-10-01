package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		TestName string
		File     string
		Error    bool
		Models   []Model
	}{
		{
			TestName: "No file",
			File:     "",
			Error:    true,
		},
		{
			TestName: "File not found",
			File:     "nonexisting-file.yml",
			Error:    true,
		},
		{
			TestName: "Invalid name",
			File:     "_testdata/load/function-empty.yml",
			Error:    true,
		},
		{
			TestName: "Malformed (yml)",
			File:     "_testdata/load/malformed.yml",
			Error:    true,
		},
		{
			TestName: "Malformed (json)",
			File:     "_testdata/load/malformed.json",
			Error:    true,
		},
		{
			TestName: "Malformed (multiple json)",
			File:     "_testdata/load/malformed-multiple.json",
			Error:    true,
		},
		{
			TestName: "Skip",
			File:     "_testdata/load/skip.yml",
			Models:   nil,
		},
		{
			TestName: "Invalid (no meta)",
			File:     "_testdata/load/function-invalid-nometa.yml",
			Error:    true,
		},
		{
			TestName: "Invalid (no name)",
			File:     "_testdata/load/function-invalid-noname.yml",
			Error:    true,
		},
		{
			TestName: "Invalid (type)",
			File:     "_testdata/load/function-invalid-type.yml",
			Error:    true,
		},
		{
			TestName: "Valid function (yml)",
			File:     "_testdata/load/function.yml",
			Models: []Model{
				&functionModel{
					file: "_testdata/load/function.yml",
					meta: &Meta{
						Name: "test",
						Labels: map[string]string{
							"test1": "abc",
							"test2": "true",
						},
					},
					spec: &FunctionSpec{
						Runtime: "go",
						AWS: &FunctionAWSSpec{
							Timeout: 3,
							Memory:  256,
						},
					},
				},
			},
		},
		{
			TestName: "Valid deployment (yml)",
			File:     "_testdata/load/deployment.yml",
			Models: []Model{
				&deploymentModel{
					file: "_testdata/load/deployment.yml",
					meta: &Meta{
						Name: "test",
						Labels: map[string]string{
							"foo": "foobar",
						},
					},
					spec: &DeploymentSpec{
						EnvironmentLabels: map[string]string{
							"foo": "foo",
							"bar": "bar",
						},
						FunctionLabels: map[string]string{
							"bar": "bar",
							"baz": "baz",
						},
					},
				},
			},
		},
		{
			TestName: "Valid function (json)",
			File:     "_testdata/load/function.json",
			Models: []Model{
				&functionModel{
					file: "_testdata/load/function.json",
					meta: &Meta{
						Name: "test-json",
						Labels: map[string]string{
							"test1": "abc",
							"test2": "true",
						},
					},
					spec: &FunctionSpec{
						Runtime: "go",
						AWS: &FunctionAWSSpec{
							Timeout: 3,
							Memory:  256,
						},
					},
				},
			},
		},
		{
			TestName: "Multiple (yaml)",
			File:     "_testdata/load/function-multiple.yml",
			Models: []Model{
				&functionModel{
					file: "_testdata/load/function-multiple.yml",
					meta: &Meta{
						Name: "test1",
					},
					spec: &FunctionSpec{
						Runtime: "go",
					},
				},
				&functionModel{
					file: "_testdata/load/function-multiple.yml",
					meta: &Meta{
						Name: "test2",
					},
					spec: &FunctionSpec{
						Runtime: "nodejs",
					},
				},
			},
		},
		{
			TestName: "Multiple (json)",
			File:     "_testdata/load/function-multiple.json",
			Models: []Model{
				&functionModel{
					file: "_testdata/load/function-multiple.json",
					meta: &Meta{
						Name: "test1-json",
					},
					spec: &FunctionSpec{
						Runtime: "go",
						AWS: &FunctionAWSSpec{
							Timeout: 10,
						},
					},
				},
				&functionModel{
					file: "_testdata/load/function-multiple.json",
					meta: &Meta{
						Name: "test2-json",
					},
					spec: &FunctionSpec{
						Runtime: "nodejs",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			models, err := Load(test.File)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Len(t, models, len(test.Models), "number of models returned does not match")
			missing := false
			for _, expected := range test.Models {
				found := false
				for _, actual := range models {
					if expected.testString() == actual.testString() {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("model not found in results:\n\n%s\n\n", expected.testString())
					missing = true
				}
			}
			if missing {
				str := []string{}
				for _, r := range models {
					str = append(str, r.testString())
				}
				t.Logf("\nParsed the following models:\n\n%s", strings.Join(str, "---\n"))
			}
		})
	}
}
