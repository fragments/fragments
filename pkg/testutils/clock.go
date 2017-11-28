package testutils

import "time"

const (
	// MockYear is the initial year for the MockClock.
	MockYear = 2017
	// MockMonth is the initial month for the MockClock.
	MockMonth = 11
	// MockDay is the initial day for the MockClock.
	MockDay = 22
	// MockHour is the initial hour for the MockClock.
	MockHour = 12
	// MockMin is the initial minute for the MockClock.
	MockMin = 34
	// MockSec is the initial second for the MockClock.
	MockSec = 56
	// MockNSec is the initial nanosecond for the MockClock.
	MockNSec = 123456789
)

// MockLoc is the initial location for the MockClock.
var MockLoc = time.UTC

// MockClock is a mocked clock.
type MockClock struct {
	Time time.Time
}

// NewMockClock creates a new mock clock with default values.
func NewMockClock() *MockClock {
	return &MockClock{
		Time: time.Date(
			MockYear,
			MockMonth,
			MockDay,
			MockHour,
			MockMin,
			MockSec,
			MockNSec,
			MockLoc,
		),
	}
}

// Now returns the current mocked time
func (m *MockClock) Now() time.Time {
	return m.Time
}

// Add modifies the mocked time
func (m *MockClock) Add(t time.Duration) *MockClock {
	m.Time = m.Time.Add(t)
	return m
}
