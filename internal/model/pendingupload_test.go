// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package model

import (
	"io/ioutil"
	"testing"

	"github.com/fragments/fragments/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalPendingUpload(t *testing.T) {
	// Marshal
	_, err := MarshalPendingUpload(nil)
	require.Error(t, err)
	s, err := MarshalPendingUpload(mockPendingUpload)
	require.NoError(t, err)
	testutils.AssertGolden(t, string(s), "testdata/GoldenPendingUpload.json")

	// Unmarshal
	var m PendingUpload
	s, err = ioutil.ReadFile("testdata/GoldenPendingUpload.json")
	require.NoError(t, err)
	err = UnmarshalPendingUpload(nil, nil)
	require.Error(t, err)
	err = UnmarshalPendingUpload(s, nil)
	require.Error(t, err)
	err = UnmarshalPendingUpload(nil, &m)
	require.Error(t, err)
	err = UnmarshalPendingUpload(s, &m)
	require.NoError(t, err)
	assert.EqualValues(t, *mockPendingUpload, m)
}