package state

import "fmt"

// ResourceType defines the type of a resource. It is used to group the same
// resources in the backend.
type ResourceType string

const (
	// ResourceTypeFunction is a function resource.
	ResourceTypeFunction ResourceType = "function"
)

// resourcePath constructs a path to store resources under in the backend.
func resourcePath(resType ResourceType, name string) string {
	return fmt.Sprintf("/resources/%s/%s", resType, name)
}

// uploadPath builds a path that's used for storing pending uploads.
func uploadPath(token string) string {
	return fmt.Sprintf("/uploads/%s", token)
}
