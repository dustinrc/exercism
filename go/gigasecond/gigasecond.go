// Package gigasecond calculates the date that someone will or did
// celebrate his/her 1 Gs anniversary.
package gigasecond

import "time"

const (
	// Gigasecond is one billion seconds
	Gigasecond time.Duration = time.Second * 1000000000

	testVersion = 4
)

// AddGigasecond calculates one billion seconds from the given start time.
func AddGigasecond(start time.Time) time.Time {
	return start.Add(Gigasecond)
}
