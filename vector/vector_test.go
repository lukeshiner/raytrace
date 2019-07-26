package vector

import (
	"math"
	"testing"

	"github.com/lukeshiner/raytrace/comparison"
	"github.com/lukeshiner/raytrace/matrix"
)

func TestPoint(t *testing.T) {
	point := NewPoint(4.3, -4.2, 3.1)
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
	vector := NewVector(4.3, -4.2, 3.1)
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

func TestEqual(t *testing.T) {
	var tests = []struct {
		a        Vector
		b        Vector
		expected bool
	}{
		{NewVector(0, 0, 0), NewVector(0, 0, 0), true},
		{NewVector(0, 0, 0), NewVector(0.25, 0, 0), false},
		{NewVector(0, 0, 0), NewVector(0, 1.0, 0), false},
		{NewVector(0, 0, 0), NewVector(0, 0, -2.4), false},
		{NewVector(0, 0, 0), NewPoint(0, 0, 0), false},
		{NewPoint(1, -2.4, -3.2), NewPoint(1, -2.4, -3.2), true},
	}
	for _, test := range tests {
		output := Equal(test.a, test.b)
		if output != test.expected {
			t.Error("Vector.Equal failed.")
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b, expected Vector
	}{
		{
			NewVector(0, 0, 0),
			NewVector(0, 0, 0),
			NewVector(0, 0, 0),
		},
		{
			NewPoint(3, -2, 5),
			NewVector(-2, 3, 1),
			NewPoint(1, 1, 6),
		},
	}
	for _, test := range tests {
		output := Add(test.a, test.b)
		if Equal(output, test.expected) != true {
			t.Errorf(
				"Failed adding vectors (%+v + %+v): expected %+v, recieved %+v",
				test.a, test.b, test.expected, output,
			)
		}
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		a, b, expected Vector
	}{
		{
			NewVector(0, 0, 0),
			NewVector(0, 0, 0),
			NewVector(0, 0, 0),
		},
		{
			NewPoint(3, 2, 1),
			NewPoint(5, 6, 7),
			NewVector(-2, -4, -6),
		},
		{
			NewPoint(3, 2, 1),
			NewVector(5, 6, 7),
			NewPoint(-2, -4, -6),
		},
		{
			NewVector(3, 2, 1),
			NewVector(5, 6, 7),
			NewVector(-2, -4, -6),
		},
	}
	for _, test := range tests {
		output := Subtract(test.a, test.b)
		if Equal(output, test.expected) != true {
			t.Errorf(
				"Failed subtracting vectors (%+v - %+v): expected %+v, recieved %+v",
				test.a, test.b, test.expected, output,
			)
		}
	}
}

func TestNegate(t *testing.T) {
	var tests = []struct {
		vector, expected Vector
	}{
		{
			NewVector(0, 0, 0),
			NewVector(0, 0, 0),
		},
		{
			NewVector(1, -2, 3),
			NewVector(-1, 2, -3),
		},
	}
	for _, test := range tests {
		output := test.vector.Negate()
		if Equal(output, test.expected) != true {
			t.Errorf(
				"Negating vector %+v: expected %+v, recieved %+v",
				test.vector, test.expected, output,
			)
		}
	}
}

func TestScalarMultiply(t *testing.T) {
	var tests = []struct {
		vector, expected Vector
		scalar           float64
	}{
		{
			vector:   NewVector(0, 0, 0),
			scalar:   1,
			expected: NewVector(0, 0, 0),
		},
		{
			vector:   Vector{1, -2, 3, -4},
			scalar:   3.5,
			expected: Vector{3.5, -7, 10.5, -14},
		},
		{
			vector:   Vector{1, -2, 3, -4},
			scalar:   0.5,
			expected: Vector{0.5, -1, 1.5, -2},
		},
	}
	for _, test := range tests {
		output := test.vector.ScalarMultiply(test.scalar)
		if Equal(output, test.expected) != true {
			t.Errorf(
				"Failed scaling vector (%+v * %v): expected %+v, recieved %+v",
				test.vector, test.scalar, test.expected, output,
			)
		}
	}
}

func TestScalarDivide(t *testing.T) {
	var tests = []struct {
		vector, expected Vector
		scalar           float64
	}{
		{
			vector:   Vector{0, 0, 0, 0},
			scalar:   1,
			expected: Vector{0, 0, 0, 0},
		},
		{
			vector:   Vector{1, -2, 3, -4},
			scalar:   2,
			expected: Vector{0.5, -1, 1.5, -2},
		},
	}
	for _, test := range tests {
		output := test.vector.ScalarDivide(test.scalar)
		if Equal(output, test.expected) != true {
			t.Errorf(
				"Failed scaling vector (%+v * %v): expected %+v, recieved %+v",
				test.vector, test.scalar, test.expected, output,
			)
		}
	}
}

func TestMagnitude(t *testing.T) {
	var tests = []struct {
		vector   Vector
		expected float64
	}{
		{vector: NewVector(0, 0, 0), expected: 0},
		{vector: NewVector(1, 0, 0), expected: 1},
		{vector: NewVector(0, 1, 0), expected: 1},
		{vector: NewVector(0, 0, 1), expected: 1},
		{vector: NewVector(1, 2, 3), expected: math.Sqrt(14)},
		{vector: NewVector(-1, -2, -3), expected: math.Sqrt(14)},
	}
	for _, test := range tests {
		output := test.vector.Magnitude()
		if comparison.EpsilonEqual(output, test.expected) != true {
			t.Errorf(
				"Failed calculating magnitude of vector %+v: expected %+v, recieved %+v",
				test.vector, test.expected, output,
			)
		}
	}
}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		vector, expected Vector
	}{
		{vector: NewVector(4, 0, 0), expected: NewVector(1, 0, 0)},
		{vector: NewVector(1, 2, 3), expected: NewVector(0.26726, 0.53452, 0.80178)},
	}
	for _, test := range tests {
		output := test.vector.Normalize()
		if Equal(output, test.expected) != true {
			t.Errorf(
				"Failed normalizing vector %+v: expected %+v, recieved %+v",
				test.vector, test.expected, output,
			)
		}
		magnitude := output.Magnitude()
		if magnitude != 1 {
			t.Errorf(
				"The magnitude of normalized vector %+v was not 1, it was %+v",
				output, magnitude,
			)
		}
	}
}

func TestDotProduct(t *testing.T) {
	var tests = []struct {
		a, b     Vector
		expected float64
	}{
		{
			a:        NewVector(1, 2, 3),
			b:        NewVector(2, 3, 4),
			expected: 20,
		},
	}
	for _, test := range tests {
		output := DotProduct(test.a, test.b)
		if comparison.EpsilonEqual(output, test.expected) != true {
			t.Errorf(
				"Failed calculating dot product of vectors %+v and %+v: expected %+v, recieved %+v",
				test.a, test.b, test.expected, output,
			)
		}
	}
}

func TestCrossProduct(t *testing.T) {
	var tests = []struct {
		a, b, expectedAB, expectedBA Vector
	}{
		{
			a:          NewVector(1, 2, 3),
			b:          NewVector(2, 3, 4),
			expectedAB: NewVector(-1, 2, -1),
			expectedBA: NewVector(1, -2, 1),
		},
	}
	for _, test := range tests {
		outputAB := CrossProduct(test.a, test.b)
		outputBA := CrossProduct(test.b, test.a)
		errorMessage := "Failed calculating cross product of vectors %+v and %+v: expected %+v, recieved %+v"
		if Equal(outputAB, test.expectedAB) != true {
			t.Errorf(errorMessage, test.a, test.b, test.expectedAB, outputAB)
		}
		if Equal(outputBA, test.expectedBA) != true {
			t.Errorf(errorMessage, test.b, test.a, test.expectedBA, outputBA)
		}
	}
}

func TestTuple(t *testing.T) {
	var tests = []struct {
		tuple []float64
	}{
		{[]float64{1, 2, 3, 0}},
		{[]float64{2, 3, 4, 0}},
		{[]float64{-1, 2, -1, 0}},
		{[]float64{1, -2, 1, 0}},
	}
	for _, test := range tests {
		p := NewPoint(test.tuple[0], test.tuple[1], test.tuple[2])
		pt := p.AsSlice()
		if pt[0] != p.X || pt[1] != p.Y || pt[2] != p.Z || pt[3] != 1 {
			t.Errorf(
				"The Tuple of point %+v was %+v, expected %+v.", p, pt, test.tuple,
			)
		}
		v := NewVector(test.tuple[0], test.tuple[1], test.tuple[2])
		vt := v.AsSlice()
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

func TestMultiplyMatrixByVector(t *testing.T) {
	var tests = []struct {
		matrix           matrix.Matrix
		vector, expected Vector
	}{
		{
			matrix: matrix.New(
				[]float64{1, 2, 3, 4},
				[]float64{2, 4, 4, 2},
				[]float64{8, 6, 4, 1},
				[]float64{0, 0, 0, 1},
			),
			vector:   NewPoint(1, 2, 3),
			expected: NewPoint(18, 24, 33),
		},
	}
	for _, test := range tests {
		value := MultiplyMatrixByVector(test.matrix, test.vector)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Multiplying matrix %+v with vector %+v returned %+v, expected %+v.",
				test.matrix, test.vector, value, test.expected,
			)
		}
	}
}

func TestTranslate(t *testing.T) {
	var tests = []struct {
		x, y, z         float64
		point, expected Vector
	}{
		{
			x: 5, y: -3, z: 2,
			point:    NewPoint(-3, 4, 5),
			expected: NewPoint(2, 1, 7),
		},
		{
			x: 5, y: -3, z: 2,
			point:    NewVector(-3, 4, 5),
			expected: NewVector(-3, 4, 5),
		},
	}
	for _, test := range tests {
		value := test.point.Translate(test.x, test.y, test.z)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Vector %+v translated by (%v, %v, %v) returned %+v, expected %+v",
				test.point, test.x, test.y, test.z, value, test.expected,
			)
		}
	}
}

func TestInverseTranslate(t *testing.T) {
	var tests = []struct {
		x, y, z         float64
		point, expected Vector
	}{
		{
			x: 5, y: -3, z: 2,
			point:    NewPoint(-3, 4, 5),
			expected: NewPoint(-8, 7, 3),
		},
	}
	for _, test := range tests {
		value := test.point.InverseTranslate(test.x, test.y, test.z)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Inverse Translation of point %+v by (%v, %v, %v) was %+v, expected %+v",
				test.point, test.x, test.y, test.z, value, test.expected,
			)
		}
	}
}

func TestScale(t *testing.T) {
	var tests = []struct {
		x, y, z         float64
		point, expected Vector
	}{
		{
			x: 2, y: 3, z: 4,
			point:    NewPoint(-4, 6, 8),
			expected: NewPoint(-8, 18, 32),
		},
		{
			x: -1, y: 1, z: 1,
			point:    NewPoint(2, 3, 4),
			expected: NewPoint(-2, 3, 4),
		},
		{
			x: 2, y: 3, z: 4,
			point:    NewVector(-4, 6, 8),
			expected: NewVector(-8, 18, 32),
		},
	}
	for _, test := range tests {
		value := test.point.Scale(test.x, test.y, test.z)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Scale vector %+v by (%v, %v, %v) returned %+v, expected %+v",
				test.point, test.x, test.y, test.z, value, test.expected,
			)
		}
	}
}

func TestInverseScale(t *testing.T) {
	var tests = []struct {
		x, y, z         float64
		point, expected Vector
	}{
		{
			x: 2, y: 3, z: 4,
			point:    NewPoint(-4, 6, 8),
			expected: NewPoint(-2, 2, 2),
		},
	}
	for _, test := range tests {
		value := test.point.InverseScale(test.x, test.y, test.z)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Inverse Scaling of point %+v by (%v, %v, %v) was %+v, expected %+v",
				test.point, test.x, test.y, test.z, value, test.expected,
			)
		}
	}
}

func TestRotateX(t *testing.T) {
	var tests = []struct {
		rotation        float64
		point, expected Vector
	}{
		{
			rotation: math.Pi / 4,
			point:    NewPoint(0, 1, 0),
			expected: NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
		{
			rotation: math.Pi / 2,
			point:    NewPoint(0, 1, 0),
			expected: NewPoint(0, 0, 1),
		},
	}
	for _, test := range tests {
		value := test.point.RotateX(test.rotation)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"RotateX %+v by %v returned %+v, expected %+v",
				test.point, test.rotation, value, test.expected,
			)
		}
	}
}

func TestInverseRotateX(t *testing.T) {
	var tests = []struct {
		rotation        float64
		point, expected Vector
	}{
		{
			rotation: math.Pi / 4,
			point:    NewPoint(0, 1, 0),
			expected: NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
		},
	}
	for _, test := range tests {
		value := test.point.InverseRotateX(test.rotation)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Inverse X Rotation of point %+v by %+v was %+v, expected %+v",
				test.point, test.rotation, value, test.expected,
			)
		}
	}
}

func TestRotateY(t *testing.T) {
	var tests = []struct {
		rotation        float64
		point, expected Vector
	}{
		{
			rotation: math.Pi / 4,
			point:    NewPoint(0, 0, 1),
			expected: NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2),
		},
		{
			rotation: math.Pi / 2,
			point:    NewPoint(0, 0, 1),
			expected: NewPoint(1, 0, 0),
		},
	}
	for _, test := range tests {
		value := test.point.RotateY(test.rotation)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"RotateY %+v by %v returned %+v, expected %+v",
				test.point, test.rotation, value, test.expected,
			)
		}
	}
}

func TestInverseRotateY(t *testing.T) {
	var tests = []struct {
		rotation        float64
		point, expected Vector
	}{
		{
			rotation: math.Pi / 4,
			point:    NewPoint(1, 0, 0),
			expected: NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2),
		},
	}
	for _, test := range tests {
		value := test.point.InverseRotateY(test.rotation)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Inverse Y Rotation of point %+v by %+v was %+v, expected %+v",
				test.point, test.rotation, value, test.expected,
			)
		}
	}
}

func TestRotateZ(t *testing.T) {
	var tests = []struct {
		rotation        float64
		point, expected Vector
	}{
		{
			rotation: math.Pi / 4,
			point:    NewPoint(0, 1, 0),
			expected: NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
		},
		{
			rotation: math.Pi / 2,
			point:    NewPoint(0, 1, 0),
			expected: NewPoint(-1, 0, 0),
		},
	}
	for _, test := range tests {
		value := test.point.RotateZ(test.rotation)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"RotateZ %+v by %v returned %+v, expected %+v",
				test.point, test.rotation, value, test.expected,
			)
		}
	}
}

func TestInverseRotateZ(t *testing.T) {
	var tests = []struct {
		rotation        float64
		point, expected Vector
	}{
		{
			rotation: math.Pi / 4,
			point:    NewPoint(0, 1, 0),
			expected: NewPoint(math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
		},
	}
	for _, test := range tests {
		value := test.point.InverseRotateZ(test.rotation)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Inverse Z Rotation of point %+v by %+v was %+v, expected %+v",
				test.point, test.rotation, value, test.expected,
			)
		}
	}
}

func TestShear(t *testing.T) {
	var tests = []struct {
		Xy, Xz, Yx, Yz, Zx, Zy float64
		point, expected        Vector
	}{
		{
			Xy: 1, Xz: 0, Yx: 0, Yz: 0, Zx: 0, Zy: 0,
			point:    NewPoint(2, 3, 4),
			expected: NewPoint(5, 3, 4),
		},
		{
			Xy: 0, Xz: 1, Yx: 0, Yz: 0, Zx: 0, Zy: 0,
			point:    NewPoint(2, 3, 4),
			expected: NewPoint(6, 3, 4),
		},
		{
			Xy: 0, Xz: 0, Yx: 1, Yz: 0, Zx: 0, Zy: 0,
			point:    NewPoint(2, 3, 4),
			expected: NewPoint(2, 5, 4),
		},
		{
			Xy: 0, Xz: 0, Yx: 0, Yz: 1, Zx: 0, Zy: 0,
			point:    NewPoint(2, 3, 4),
			expected: NewPoint(2, 7, 4),
		},
		{
			Xy: 0, Xz: 0, Yx: 0, Yz: 0, Zx: 1, Zy: 0,
			point:    NewPoint(2, 3, 4),
			expected: NewPoint(2, 3, 6),
		},
		{
			Xy: 0, Xz: 0, Yx: 0, Yz: 0, Zx: 0, Zy: 1,
			point:    NewPoint(2, 3, 4),
			expected: NewPoint(2, 3, 7),
		},
	}
	for _, test := range tests {
		value := test.point.Shear(test.Xy, test.Xz, test.Yx, test.Yz, test.Zx, test.Zy)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Shearing of point %+v by (%v, %v, %v, %v, %v, %v) was %+v, expected %+v",
				test.point, test.Xz, test.Yx, test.Yx, test.Yz, test.Zx, test.Zy, value, test.expected,
			)
		}
	}
}
