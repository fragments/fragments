package aws

import (
	"os"
	"testing"
	"time"

	"github.com/fragments/fragments/internal/state"
)

func TestMain(m *testing.M) {
	state.Now = func() time.Time {
		mockTime, _ := time.Parse(time.RFC3339, "2017-10-31T12:34:56+02:00")
		return mockTime
	}

	code := m.Run()

	state.Now = func() time.Time {
		return time.Now()
	}

	os.Exit(code)
}
