package object

import (
	"math"
	"testing"

	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/vector"
)

// func TestNewShpere(t testing.T) {
//
// }

func TestDefaultSphereTransform(t *testing.T) {
	s := NewSphere()
	if matrix.Equal(s.Transform, matrix.IdentityMatrix(4)) != true {
		t.Error("Sphere default transform was not the identity matrix.")
	}
}

func TestSetTransform(t *testing.T) {
	s := NewSphere()
	transform := matrix.TranslationMatrix(2, 3, 4)
	s.SetTransform(transform)
	if matrix.Equal(s.Transform, transform) != true {
		t.Error("Did not set transform on sphere.")
	}
}

func TestNormalAt(t *testing.T) {
	var tests = []struct {
		sphere    Sphere
		transform matrix.Matrix
		point     vector.Vector
		expected  vector.Vector
	}{
		{
			sphere:    NewSphere(),
			transform: matrix.IdentityMatrix(4),
			point:     vector.NewPoint(1, 0, 0),
			expected:  vector.NewVector(1, 0, 0),
		},
		{
			sphere:    NewSphere(),
			transform: matrix.IdentityMatrix(4),
			point:     vector.NewPoint(0, 1, 0),
			expected:  vector.NewVector(0, 1, 0),
		},
		{
			sphere:    NewSphere(),
			transform: matrix.IdentityMatrix(4),
			point:     vector.NewPoint(0, 0, 1),
			expected:  vector.NewVector(0, 0, 1),
		},
		{
			sphere:    NewSphere(),
			transform: matrix.IdentityMatrix(4),
			point:     vector.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			expected:  vector.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
		{
			sphere:    NewSphere(),
			transform: matrix.TranslationMatrix(0, 1, 0),
			point:     vector.NewPoint(0, 1.70711, -0.70711),
			expected:  vector.NewVector(0, 0.70711, -0.70711),
		},
		{
			sphere: NewSphere(),
			transform: matrix.Multiply(
				matrix.ScalingMatrix(1, 0.5, 1),
				matrix.RotationZMatrix(math.Pi/5),
			),
			point:    vector.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			expected: vector.NewVector(0, 0.97014, -0.24254),
		},
	}
	for _, test := range tests {
		test.sphere.SetTransform(test.transform)
		result := test.sphere.NormalAt(test.point)
		if vector.Equal(result, test.expected) != true {
			t.Errorf(
				"The normal of sphere %+v at point %+v was %+v, expected %+v.",
				test.sphere, test.point, result, test.expected,
			)
		}
		if vector.Equal(result, result.Normalize()) != true {
			t.Errorf(
				"The normal of sphere %+v at point %+v was not normalized.",
				test.sphere, test.point,
			)
		}
	}
}
