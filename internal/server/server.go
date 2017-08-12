package server

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

// Server is the fragments server that accepts resources and keeps them in the
// store.
type Server struct {
}

// PutFunctionInput contains parameters for creating or updating function.
type PutFunctionInput struct {
	Checksum string
	Function Function
}

// PutFunctionOutput is the result of putting a new function.
type PutFunctionOutput struct {
}

// PutFunction creates or updates a function. In case the function already
// exists it is updated. If not, source upload is requested.
func (s *Server) PutFunction(ctx context.Context, input *PutFunctionInput) (*PutFunctionOutput, error) {
	fmt.Printf("%+v\n", input)
	return nil, errors.New("put function: not implemented")
}
