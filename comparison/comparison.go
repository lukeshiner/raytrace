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
