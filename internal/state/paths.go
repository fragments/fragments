package state

import (
	"fmt"

	"github.com/pkg/errors"
)

func modelListPath(resType ModelType) string {
	return fmt.Sprintf("/models/%s", resType)
}

// modelPath constructs a path to store models under in the backend.
func modelPath(modelType ModelType, name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	if string(modelType) == "" {
		return "", errors.New("model type is required")
	}
	return fmt.Sprintf("%s/%s", modelListPath(modelType), name), nil
}

// uploadPath builds a path that's used for storing pending uploads.
func uploadPath(token string) (string, error) {
	if token == "" {
		return "", errors.New("token is required")
	}
	return fmt.Sprintf("/uploads/%s", token), nil
}
