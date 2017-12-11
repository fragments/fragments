//go:generate genny -in=$GOFILE -out=deployment_test.go gen "Type=Deployment typename=deployment"
//go:generate genny -in=$GOFILE -out=environment_test.go gen "Type=Environment typename=environment"
//go:generate genny -in=$GOFILE -out=function_test.go gen "Type=Function typename=function"
//go:generate genny -in=$GOFILE -out=pendingupload_test.go gen "Type=PendingUpload typename=pending-upload"

package model

import (
	"io/ioutil"
	"testing"

	"github.com/fragments/fragments/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalType(t *testing.T) {
	// Marshal
	_, err := MarshalType(nil)
	require.Error(t, err)
	s, err := MarshalType(mockType)
	require.NoError(t, err)
	testutils.AssertGolden(t, string(s), "testdata/GoldenType.json")

	// Unmarshal
	var m Type
	s, err = ioutil.ReadFile("testdata/GoldenType.json")
	require.NoError(t, err)
	err = UnmarshalType(nil, nil)
	require.Error(t, err)
	err = UnmarshalType(s, nil)
	require.Error(t, err)
	err = UnmarshalType(nil, &m)
	require.Error(t, err)
	err = UnmarshalType(s, &m)
	require.NoError(t, err)
	assert.EqualValues(t, *mockType, m)
}
