package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		TestName  string
		File      string
		Error     bool
		Resources []Resource
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
			TestName:  "Skip",
			File:      "_testdata/load/skip.yml",
			Resources: nil,
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
			TestName: "Valid (yml)",
			File:     "_testdata/load/function.yml",
			Resources: []Resource{
				&functionResource{
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
			TestName: "Valid (json)",
			File:     "_testdata/load/function.json",
			Resources: []Resource{
				&functionResource{
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
			Resources: []Resource{
				&functionResource{
					file: "_testdata/load/function-multiple.yml",
					meta: &Meta{
						Name: "test1",
					},
					spec: &FunctionSpec{
						Runtime: "go",
					},
				},
				&functionResource{
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
			Resources: []Resource{
				&functionResource{
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
				&functionResource{
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
			resources, err := Load(test.File)
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Len(t, resources, len(test.Resources), "number of resources returned does not match")
			missing := false
			for _, expected := range test.Resources {
				found := false
				for _, actual := range resources {
					if expected.testString() == actual.testString() {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("resource not found in results:\n\n%s\n\n", expected.testString())
					missing = true
				}
			}
			if missing {
				str := []string{}
				for _, r := range resources {
					str = append(str, r.testString())
				}
				t.Logf("\nParsed the following resources:\n\n%s", strings.Join(str, "---\n"))
			}
		})
	}
}
