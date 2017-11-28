package testutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMockClock(t *testing.T) {
	m := NewMockClock()
	assert.Equal(t, "2017-11-22 12:34:56.123456789 +0000 UTC", m.Now().String())
	m.Add(24 * time.Hour)
	assert.Equal(t, "2017-11-23 12:34:56.123456789 +0000 UTC", m.Now().String())
}
