// Package hamming is for calculating the Hamming difference between two DNA strands.
package hamming

import "errors"

const testVersion = 4

// Distance determines the Hamming difference between two DNA sequesnces.
//
// The distance returned will be non-negative if successful. Sequences of inequal
// length will result in an error.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("sequences not of equal length")
	}

	var distance int
	for index, value := range a {
		if value != rune(b[index]) {
			distance++
		}
	}
	return distance, nil
}
