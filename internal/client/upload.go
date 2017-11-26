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

	req, err := http.NewRequest(http.MethodPut, url, data)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "upload request failed")
	}

	if _, err = io.Copy(ioutil.Discard, res.Body); err != nil {
		return err
	}
	if err := res.Body.Close(); err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received unexpected status %v", res.StatusCode)
	}

	return nil
}
