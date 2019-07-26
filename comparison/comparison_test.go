package comparison

import (
	"testing"
)

func TestEpsilonEqual(t *testing.T) {
	a := 0.20000000000000007
	b := 0.2
	if EpsilonEqual(a, b) != true {
		t.Error("Equality check failed.")
	}
}

func TestEqualSlice(t *testing.T) {
	var tests = []struct {
		a, b     []float64
		expected bool
	}{
		{
			a:        []float64{},
			b:        []float64{},
			expected: true,
		},
		{
			a:        []float64{0},
			b:        []float64{0},
			expected: true,
		},
		{
			a:        []float64{0},
			b:        []float64{1},
			expected: false,
		},
		{
			a:        []float64{0},
			b:        []float64{0, 1},
			expected: false,
		},
	}
	for _, test := range tests {
		result := EqualSlice(test.a, test.b)
		if result != test.expected {
			t.Errorf("Slice equality check failed on %+v and %+v", test.a, test.b)
		}
	}
}
