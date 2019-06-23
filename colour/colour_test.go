package colour

import (
	"testing"
)

func TestColour(t *testing.T) {
	colour := Colour{-0.5, 0.4, 1.7}
	if colour.Red != -0.5 {
		t.Error("Could not access the Red attribute of Colour.")
	}
	if colour.Green != 0.4 {
		t.Error("Could not access the Red attribute of Colour.")
	}
	if colour.Blue != 1.7 {
		t.Error("Could not access the Red attribute of Colour.")
	}
}

func TestColourEqualMethod(t *testing.T) {
	var tests = []struct {
		a        Colour
		b        Colour
		expected bool
	}{
		{Colour{0, 0, 0}, Colour{0, 0, 0}, true},
		{Colour{1, 0, 0}, Colour{0, 0, 0}, false},
		{Colour{0, 1, 0}, Colour{0, 0, 0}, false},
		{Colour{0, 0, 1}, Colour{0, 0, 0}, false},
	}

	for _, test := range tests {
		output := test.a.Equal(test.b)
		if output != test.expected {
			t.Error("Failed Colour equality test.")
		}
	}
}

func TestColourAddMethod(t *testing.T) {
	var tests = []struct {
		a        Colour
		b        Colour
		expected Colour
	}{
		{
			Colour{0, 0, 0},
			Colour{0, 0, 0},
			Colour{0, 0, 0},
		},
		{
			Colour{0.9, 0.6, 0.75},
			Colour{0.7, 0.1, 0.25},
			Colour{1.6, 0.7, 1},
		},
	}

	for _, test := range tests {
		output := test.a.Add(test.b)
		if output.Equal(test.expected) != true {
			t.Errorf(
				"Failed adding colours (%+v + %+v): expected %+v, recieved %+v",
				test.a,
				test.b,
				test.expected,
				output,
			)
		}
	}
}

func TestColourSubMethod(t *testing.T) {
	var tests = []struct {
		a        Colour
		b        Colour
		expected Colour
	}{
		{
			Colour{0, 0, 0},
			Colour{0, 0, 0},
			Colour{0, 0, 0},
		},
		{
			Colour{0.9, 0.6, 0.75},
			Colour{0.7, 0.1, 0.25},
			Colour{0.2, 0.5, 0.5},
		},
	}

	for _, test := range tests {
		output := test.a.Sub(test.b)
		if output.Equal(test.expected) != true {
			t.Errorf(
				"Failed subtracting colours (%+v - %+v): expected %+v, recieved %+v",
				test.a,
				test.b,
				test.expected,
				output,
			)
		}
	}
}

func TestColourScalarMultMethod(t *testing.T) {
	var tests = []struct {
		colour     Colour
		multiplier float64
		expected   Colour
	}{
		{
			Colour{0, 0, 0},
			1,
			Colour{0, 0, 0},
		},
		{
			Colour{0.2, 0.3, 0.4},
			2,
			Colour{0.4, 0.6, 0.8},
		},
	}

	for _, test := range tests {
		output := test.colour.ScalarMult(test.multiplier)
		if output.Equal(test.expected) != true {
			t.Errorf(
				"Failed scalar multiplying (%+v + %+v): expected %+v, recieved %+v",
				test.colour,
				test.multiplier,
				test.expected,
				output,
			)
		}
	}
}

func TestColourMultMethod(t *testing.T) {
	var tests = []struct {
		a        Colour
		b        Colour
		expected Colour
	}{
		{
			Colour{0, 0, 0},
			Colour{0, 0, 0},
			Colour{0, 0, 0},
		},
		{
			Colour{1, 0.2, 0.4},
			Colour{0.9, 1, 0.1},
			Colour{0.9, 0.2, 0.04},
		},
	}

	for _, test := range tests {
		output := test.a.Mult(test.b)
		if output.Equal(test.expected) != true {
			t.Errorf(
				"Failed multiplying colours (%+v - %+v): expected %+v, recieved %+v",
				test.a,
				test.b,
				test.expected,
				output,
			)
		}
	}
}
