package vector

import (
	"math"
	"testing"

	"github.com/lukeshiner/raytrace/comparison"
)

func TestPoint(t *testing.T) {
	point := MakePoint(4.3, -4.2, 3.1)
	if point.X != 4.3 {
		t.Error("Failed to retrieve the Point.X")
	}
	if point.Y != -4.2 {
		t.Error("Failed to retrieve the Point.Y")
	}
	if point.Z != 3.1 {
		t.Error("Failed to retrieve the Point.Z")
	}
	if point.W != 1.0 {
		t.Error("Failed to retrieve the Point.W")
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
	if vector.X != 4.3 {
		t.Error("Failed to retrieve the Vector.X")
	}
	if vector.Y != -4.2 {
		t.Error("Failed to retrieve the Vector.Y")
	}
	if vector.Z != 3.1 {
		t.Error("Failed to retrieve the Vector.Z")
	}
	if vector.W != 0.0 {
		t.Error("Failed to retrieve the Vector.W")
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
		if comparison.EpsilonEqual(output, test.expected) != true {
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
		if comparison.EpsilonEqual(output, test.expected) != true {
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
		errorMessage := "Failed calculating cross product of vectors %+v and %+v: expected %+v, recieved %+v"
		outputAB := CrossProduct(&test.a, &test.b)
		outputBA := CrossProduct(&test.b, &test.a)
		if EqualVectors(&outputAB, &test.expectedAB) != true {
			t.Errorf(
				errorMessage,
				test.a,
				test.b,
				test.expectedAB,
				outputAB,
			)
		}
		if EqualVectors(&outputBA, &test.expectedBA) != true {
			t.Errorf(
				errorMessage,
				test.b,
				test.a,
				test.expectedBA,
				outputBA,
			)
		}
	}
}

func TestTuple(t *testing.T) {
	var tests = []struct {
		tuple []float64
	}{
		{
			[]float64{1, 2, 3, 0},
		}, {
			[]float64{2, 3, 4, 0},
		}, {
			[]float64{-1, 2, -1, 0},
		}, {
			[]float64{1, -2, 1, 0},
		},
	}

	for _, test := range tests {
		p := MakePoint(test.tuple[0], test.tuple[1], test.tuple[2])
		pt := p.Tuple()
		if pt[0] != p.X || pt[1] != p.Y || pt[2] != p.Z || pt[3] != 1 {
			t.Errorf(
				"The Tuple of point %+v was %+v, expected %+v.", p, pt, test.tuple,
			)
		}
		v := MakeVector(test.tuple[0], test.tuple[1], test.tuple[2])
		vt := v.Tuple()
		if vt[0] != p.X || vt[1] != p.Y || vt[2] != p.Z || vt[3] != 0 {
			t.Errorf(
				"The Tuple of vector %+v was %+v, expected %+v.", v, vt, test.tuple,
			)
		}
	}
}

func TestFromSlice(t *testing.T) {
	var x, y, z, w float64
	x = 5
	y = -7
	z = 2
	w = 1
	s := []float64{x, y, z, w}
	v := FromSlice(s)
	if v.X != x || v.Y != y || v.Z != z || v.W != w {
		t.Errorf("FromSlice(%+v) produced %+v.", s, v)
	}
}
