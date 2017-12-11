//go:generate genny -in=$GOFILE -out=deployment.go gen "Type=*Deployment typename=deployment"
//go:generate genny -in=$GOFILE -out=environment.go gen "Type=*Environment typename=environment"
//go:generate genny -in=$GOFILE -out=function.go gen "Type=*Function typename=function"
//go:generate genny -in=$GOFILE -out=pendingupload.go gen "Type=*PendingUpload typename=pending-upload"

package model

import (
	"encoding/json"

	"github.com/cheekybits/genny/generic"
	"github.com/pkg/errors"
)

// nolint: golint
type Type generic.Type

// MarshalType marshals t to a json encoded byte array.
func MarshalType(t Type) ([]byte, error) {
	if t == nil {
		return nil, errors.New("typename is nil")
	}
	s, err := json.Marshal(t)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal typename")
	}
	return s, nil
}

// UnmarshalType unmarshals a json encoded Type to t
func UnmarshalType(s []byte, t Type) error {
	if t == nil {
		return errors.New("target typename is nil")
	}
	if err := json.Unmarshal(s, t); err != nil {
		return errors.Wrap(err, "could not unmarshal typename")
	}
	return nil
}
