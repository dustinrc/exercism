// Package leap provides a function for leap year determination.
package leap

// TestVersion indentifies the version of the test program that this
// code is written against.
const TestVersion = 1

// IsLeapYear determines if the given year is, in fact, a leap year.
//
// The logical operations are ordered by the greatest chance of
// disqualification.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%400 == 0 || year%100 != 0)
}