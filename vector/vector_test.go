package vector

import (
	"math"
	"testing"
)

func TestEqual(t *testing.T) {
	var a float64 = 0.1
	var b float64 = 0.1
	if Equal(a, b) != true {
		t.Error("Equality check failed.")
	}
}

func TestPoint(t *testing.T) {
	point := MakePoint(4.3, -4.2, 3.1)
	if point.x != 4.3 {
		t.Error("Failed to retrieve the Point.x")
	}
	if point.y != -4.2 {
		t.Error("Failed to retrieve the Point.y")
	}
	if point.z != 3.1 {
		t.Error("Failed to retrieve the Point.z")
	}
	if point.w != 1.0 {
		t.Error("Failed to retrieve the Point.w")
	}
	if point.IsPoint() != true {
		t.Error("Point.IsPoint did not return true")
	}
	if point.IsVector() != false {
		t.Error("Point.IsPoint did not return false")
	}
}

func TestVector(t *testing.T) {
	vector := MakeVector(4.3, -4.2, 3.1)
	if vector.x != 4.3 {
		t.Error("Failed to retrieve the Vector.x")
	}
	if vector.y != -4.2 {
		t.Error("Failed to retrieve the Vector.y")
	}
	if vector.z != 3.1 {
		t.Error("Failed to retrieve the Vector.z")
	}
	if vector.w != 0.0 {
		t.Error("Failed to retrieve the Vector.w")
	}
	if vector.IsPoint() != false {
		t.Error("Vector.IsPoint did not return false")
	}
	if vector.IsVector() != true {
		t.Error("Vector.IsPoint did not return true")
	}
}

func TestEqualVectors(t *testing.T) {
	var tests = []struct {
		vectorA  Vector
		vectorB  Vector
		expected bool
	}{
		{Vector{0, 0, 0, 0}, Vector{0, 0, 0, 0}, true},
		{Vector{0, 0, 0, 0}, Vector{0.25, 0, 0, 0}, false},
		{Vector{0, 0, 0, 0}, Vector{0, 1.0, 0, 0}, false},
		{Vector{0, 0, 0, 0}, Vector{0, 0, -2.4, 0}, false},
		{Vector{0, 0, 0, 0}, Vector{0, 0, 0, 1.0}, false},
		{Vector{1, -2.4, -3.2, 1.0}, Vector{1, -2.4, -3.2, 1.0}, true},
	}

	for _, test := range tests {
		output := EqualVectors(&test.vectorA, &test.vectorB)
		if output != test.expected {
			t.Error("Vector.Equal failed.")
		}
	}
}

func TestVectorEqualMethod(t *testing.T) {
	var tests = []struct {
		vectorA  Vector
		vectorB  Vector
		expected bool
	}{
		{Vector{0, 0, 0, 0}, Vector{0, 0, 0, 0}, true},
		{Vector{0, 0, 0, 0}, Vector{0.25, 0, 0, 0}, false},
		{Vector{0, 0, 0, 0}, Vector{0, 1.0, 0, 0}, false},
		{Vector{0, 0, 0, 0}, Vector{0, 0, -2.4, 0}, false},
		{Vector{0, 0, 0, 0}, Vector{0, 0, 0, 1.0}, false},
		{Vector{1, -2.4, -3.2, 1.0}, Vector{1, -2.4, -3.2, 1.0}, true},
	}

	for _, test := range tests {
		output := test.vectorA.Equal(&test.vectorB)
		if output != test.expected {
			t.Error("Vector.Equal failed.")
		}
	}
}

func TestVectorAddMethod(t *testing.T) {
	var tests = []struct {
		vectorA  Vector
		vectorB  Vector
		expected Vector
	}{
		{
			Vector{0, 0, 0, 0},
			Vector{0, 0, 0, 0},
			Vector{0, 0, 0, 0},
		},
		{
			Vector{3, -2, 5, 1},
			Vector{-2, 3, 1, 0},
			Vector{1, 1, 6, 1},
		},
	}

	for _, test := range tests {
		output := test.vectorA.Add(&test.vectorB)
		if EqualVectors(&output, &test.expected) != true {
			t.Errorf(
				"Failed adding vectors (%+v + %+v): expected %+v, recieved %+v",
				test.vectorA,
				test.vectorB,
				test.expected,
				output,
			)
		}
	}
}

func TestVectorSubMethod(t *testing.T) {
	var tests = []struct {
		vectorA  Vector
		vectorB  Vector
		expected Vector
	}{
		{
			Vector{0, 0, 0, 0},
			Vector{0, 0, 0, 0},
			Vector{0, 0, 0, 0},
		},
		{
			Vector{3, 2, 1, 1},
			Vector{5, 6, 7, 1},
			Vector{-2, -4, -6, 0},
		},
		{
			Vector{3, 2, 1, 1},
			Vector{5, 6, 7, 0},
			Vector{-2, -4, -6, 1},
		},
		{
			Vector{3, 2, 1, 0},
			Vector{5, 6, 7, 0},
			Vector{-2, -4, -6, 0},
		},
	}

	for _, test := range tests {
		output := test.vectorA.Sub(&test.vectorB)
		if EqualVectors(&output, &test.expected) != true {
			t.Errorf(
				"Failed adding vectors (%+v + %+v): expected %+v, recieved %+v",
				test.vectorA,
				test.vectorB,
				test.expected,
				output,
			)
		}
	}
}

func TestVectorNegateMethod(t *testing.T) {
	var tests = []struct {
		vector   Vector
		expected Vector
	}{
		{
			Vector{0, 0, 0, 0},
			Vector{0, 0, 0, 0},
		},
		{
			Vector{1, -2, 3, 0},
			Vector{-1, 2, -3, 0},
		},
	}

	for _, test := range tests {
		output := test.vector.Negate()
		if EqualVectors(&output, &test.expected) != true {
			t.Errorf(
				"Failed negating vector %+v: expected %+v, recieved %+v",
				test.vector,
				test.expected,
				output,
			)
		}
	}
}

func TestVectorScalarMultMethod(t *testing.T) {
	var tests = []struct {
		vector   Vector
		scalar   float64
		expected Vector
	}{
		{
			Vector{0, 0, 0, 0},
			1,
			Vector{0, 0, 0, 0},
		},
		{
			Vector{1, -2, 3, -4},
			3.5,
			Vector{3.5, -7, 10.5, -14},
		},
		{
			Vector{1, -2, 3, -4},
			0.5,
			Vector{0.5, -1, 1.5, -2},
		},
	}

	for _, test := range tests {
		output := test.vector.ScalarMult(test.scalar)
		if EqualVectors(&output, &test.expected) != true {
			t.Errorf(
				"Failed scaling vector (%+v * %v): expected %+v, recieved %+v",
				test.vector,
				test.scalar,
				test.expected,
				output,
			)
		}
	}
}

func TestVectorScalarDivMethod(t *testing.T) {
	var tests = []struct {
		vector   Vector
		scalar   float64
		expected Vector
	}{
		{
			Vector{0, 0, 0, 0},
			1,
			Vector{0, 0, 0, 0},
		},
		{
			Vector{1, -2, 3, -4},
			2,
			Vector{0.5, -1, 1.5, -2},
		},
	}

	for _, test := range tests {
		output := test.vector.ScalarDiv(test.scalar)
		if EqualVectors(&output, &test.expected) != true {
			t.Errorf(
				"Failed scaling vector (%+v * %v): expected %+v, recieved %+v",
				test.vector,
				test.scalar,
				test.expected,
				output,
			)
		}
	}
}

func TestVectorMagMethod(t *testing.T) {
	var tests = []struct {
		vector   Vector
		expected float64
	}{
		{Vector{0, 0, 0, 0}, 0},
		{Vector{1, 0, 0, 0}, 1},
		{Vector{0, 1, 0, 0}, 1},
		{Vector{0, 0, 1, 0}, 1},
		{Vector{1, 2, 3, 0}, math.Sqrt(14)},
		{Vector{-1, -2, -3, 0}, math.Sqrt(14)},
	}

	for _, test := range tests {
		output := test.vector.Mag()
		if Equal(output, test.expected) != true {
			t.Errorf(
				"Failed calculating magnitude of vector %+v: expected %+v, recieved %+v",
				test.vector,
				test.expected,
				output,
			)
		}
	}
}

func TestVectorNormalizeMethod(t *testing.T) {
	var tests = []struct {
		vector   Vector
		expected Vector
	}{
		{Vector{4, 0, 0, 0}, Vector{1, 0, 0, 0}},
		{Vector{1, 2, 3, 0}, Vector{0.26726, 0.53452, 0.80178, 0}},
	}

	for _, test := range tests {
		output := test.vector.Normalize()
		if EqualVectors(&output, &test.expected) != true {
			t.Errorf(
				"Failed normalizing vector %+v: expected %+v, recieved %+v",
				test.vector,
				test.expected,
				output,
			)
		}
		magnitude := output.Mag()
		if magnitude != 1 {
			t.Errorf(
				"The magnitude of normalized vector %+v was not 1, it was %+v",
				output,
				magnitude,
			)
		}
	}
}

func TestVectorDotProduct(t *testing.T) {
	var tests = []struct {
		a        Vector
		b        Vector
		expected float64
	}{
		{
			Vector{1, 2, 3, 0},
			Vector{2, 3, 4, 0},
			20,
		},
	}

	for _, test := range tests {
		output := DotProduct(&test.a, &test.b)
		if Equal(output, test.expected) != true {
			t.Errorf(
				"Failed calculating dot product of vectors %+v and %+v: expected %+v, recieved %+v",
				test.a,
				test.b,
				test.expected,
				output,
			)
		}
	}
}

func TestVectorCrossProduct(t *testing.T) {
	var tests = []struct {
		a          Vector
		b          Vector
		expectedAB Vector
		expectedBA Vector
	}{
		{
			Vector{1, 2, 3, 0},
			Vector{2, 3, 4, 0},
			Vector{-1, 2, -1, 0},
			Vector{1, -2, 1, 0},
		},
	}

	for _, test := range tests {
		error_message := "Failed calculating cross product of vectors %+v and %+v: expected %+v, recieved %+v"
		outputAB := CrossProduct(&test.a, &test.b)
		outputBA := CrossProduct(&test.b, &test.a)
		if EqualVectors(&outputAB, &test.expectedAB) != true {
			t.Errorf(
				error_message,
				test.a,
				test.b,
				test.expectedAB,
				outputAB,
			)
		}
		if EqualVectors(&outputBA, &test.expectedBA) != true {
			t.Errorf(
				error_message,
				test.b,
				test.a,
				test.expectedBA,
				outputBA,
			)
		}
	}
}
