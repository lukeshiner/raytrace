package comparison

import (
	"math"
)

// EPSLION is the margin of error used to compare floats for equality.
const EPSLION = 0.00001

// EpsilonEqual compares two float64 for equality. It returns true if they are
// within EPSILON of each other.
func EpsilonEqual(a float64, b float64) bool {
	return math.Abs(a-b) < EPSLION
}

// EqualSlice returns true if a and b are the same, otherwise false.
func EqualSlice(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if EpsilonEqual(a[0], b[0]) != true {
			return false
		}
	}
	return true
}
