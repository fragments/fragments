//go:generate mockery -name SourceTarget

package filestore

import "context"

// SourceTarget is a target that accepts source code uploads.
type SourceTarget interface {
	// NewUploadURL generates a new URL that the source can be uploaded to.
	NewUploadURL(name string) (string, error)
	// Persist persists an uploaded file.
	Persist(ctx context.Context, name string) error
}

// SourceReader reads source code from the filestore.
type SourceReader interface {
	// GetFile gets a file from the filestore
	GetFile(filename string)
}
