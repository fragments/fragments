package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// Upload uploads data to a url. The upload is done through a http put.
// Has a timeout of 1 minute
func Upload(data io.Reader, url string) error {
	client := &http.Client{
		Timeout: 1 * time.Minute,
	}

	req, _ := http.NewRequest(http.MethodPut, url, data)
	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "upload request failed")
	}

	_, _ = io.Copy(ioutil.Discard, res.Body)
	_ = res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received unexpected status %v", res.StatusCode)
	}

	return nil
}
