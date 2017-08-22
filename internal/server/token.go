package server

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// GenerateToken generates a new ulid token. The token is lexiographically
// stortable and includes a millisecond precision timestamp.
// Example token: 01AN4Z07BY79KA1307SR9X4MV3
//
// See: https://github.com/oklog/ulid for more
func GenerateToken() string {
	t := time.Now()
	entropy := rand.Reader
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
