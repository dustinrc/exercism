// Package clock is for storing and displaying time of day (hour, minute)
// data. It has no concern for dates nor timezones.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock holds the time of day data.
type Clock struct {
	hour, minute int
}

// New returns a properly bounded Clock.
//
// The hour and minute values may be positive or negative, which will
// increase or decrease the stored time of day, respectively.
func New(hour, minute int) Clock {
	totalMinutes := (hour*60 + minute) % 1440
	if totalMinutes < 0 {
		totalMinutes = totalMinutes + 1440
	}

	m := totalMinutes % 60
	h := totalMinutes / 60 % 24

	return Clock{hour: h, minute: m}
}

// String returns the stored time of day data in 24 hour format (hh:mm).
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add returns a new Clock increased by the given number of minutes. A
// negative value will reduce the time of day for the new Clock.
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
