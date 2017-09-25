package state

import (
	"fmt"

	"github.com/pkg/errors"
)

// resourcePath constructs a path to store resources under in the backend.
func resourcePath(resType ResourceType, name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	return fmt.Sprintf("/resources/%s/%s", resType, name), nil
}

// uploadPath builds a path that's used for storing pending uploads.
func uploadPath(token string) (string, error) {
	if token == "" {
		return "", errors.New("token is required")
	}
	return fmt.Sprintf("/uploads/%s", token), nil
}
