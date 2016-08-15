package raindrops

import "strconv"

const testVersion = 2

// Convert makes it rain.
//
// An integer with prime factors of 3, 5 and/or 7 will be converted to
// a string combination of "Pling", "Plang" and/or "Plong", respectively.
// Integers containing none of those factors will simply be converted to
// a string representation of that digit.
func Convert(in int) string {
	var out string
	if in%3 == 0 {
		out += "Pling"
	}
	if in%5 == 0 {
		out += "Plang"
	}
	if in%7 == 0 {
		out += "Plong"
	}
	if len(out) == 0 {
		out = strconv.Itoa(in)
	}

	return out
}

// The test program has a benchmark too.  How fast does your Convert convert?
