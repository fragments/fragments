package server

import (
	"context"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/internal/state"
	"github.com/pkg/errors"
)

// Server is the fragments server that accepts resources and keeps them in the
// store.
type Server struct {
	StateStore backend.KV
}

// UploadRequest is a request for source code, returned from the server.
type UploadRequest struct {
	// Token is the token to use when confirming the upload.
	Token string
	// URL is where to upload the source code. The source must be uploaded as a
	// tar.gz. After completion the upload should be confirmed.
	URL string
}

// PutFunction creates or updates a function. In case the function already
// exists it is updated. If not, source upload is requested.
func (s *Server) PutFunction(ctx context.Context, input *state.Function) (*UploadRequest, error) {
	// Get existing function
	existing, err := state.GetFunction(ctx, s.StateStore, input.Meta.Name)
	if err != nil {
		return nil, errors.Wrap(err, "could not check existing function")
	}

	if existing == nil || existing.Checksum != input.Checksum {
		// nolint: vetshadow
		res, err := s.requestUpload(ctx, input, existing)
		if err != nil {
			return nil, errors.Wrap(err, "could not request source upload")
		}
		return res, nil
	}

	if err = s.updateFunctionConfiguration(ctx, input); err != nil {
		return nil, errors.Wrap(err, "could not update function configuration")
	}

	return nil, nil
}

// requestUpload creates a url the client can upload source code to. The upload
// request is stored as a PendingUpload in the store so it can be retrieved
// when the client confirms the upload.
func (s *Server) requestUpload(ctx context.Context, input *state.Function, existing *state.Function) (*UploadRequest, error) {
	return nil, errors.New("not implemented")
}

// updateFunctionConfiguration updates the function's configuration
// (performance specs, environment variables etc) without updating the source
// code.
func (s *Server) updateFunctionConfiguration(ctx context.Context, input *state.Function) error {
	return errors.New("not implemented")
}
