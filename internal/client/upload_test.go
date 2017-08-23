package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpload(t *testing.T) {
	err := Upload(nil, "not a valid url")
	require.Error(t, err)

	tests := []struct {
		TestName string
		GetData  func() []byte
		Response int
		Error    bool
	}{
		{
			TestName: "Error response",
			Response: 400,
			GetData: func() []byte {
				// Empty body would return error
				return nil
			},
			Error: true,
		},
		{
			TestName: "Ok",
			Response: 200,
			GetData: func() []byte {
				data, err := ioutil.ReadFile("_testdata/upload/file.txt")
				require.NoError(t, err)
				return data
			},
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			data := test.GetData()
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.Response)
				assert.Equal(t, http.MethodPut, r.Method)
				assert.EqualValues(t, len(data), r.ContentLength)
			}))

			err := Upload(bytes.NewReader(data), ts.URL)
			ts.Close()
			if test.Error {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
