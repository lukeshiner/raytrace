package shape

import (
	"math"
	"testing"

	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/comparison"
	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/ray"
	"github.com/lukeshiner/raytrace/vector"
)

func TestShapeDefaultTransform(t *testing.T) {
	s := newShape()
	if matrix.Equal(s.Transform(), matrix.IdentityMatrix(4)) != true {
		t.Error("Sphere default transform was not the identity matrix.")
	}
}

func TestShapeDefaultMaterial(t *testing.T) {
	s := newShape()
	if s.Material() != material.New() {
		t.Error("Sphere default material was not correct.")
	}
}

func TestShapeSetMaterial(t *testing.T) {
	m := material.New()
	m.Colour = colour.New(0.5, 0.5, 0.5)
	m.Ambient = 0.5
	m.Diffuse = 0.3
	m.Specular = 0.8
	m.Shininess = 150.0
	s := newShape()
	s.SetMaterial(m)
	if s.Material() != m {
		t.Error("Could not set Sphere material.")
	}
}

func TestShapeSetTransform(t *testing.T) {
	s := newShape()
	transform := matrix.TranslationMatrix(2, 3, 4)
	s.SetTransform(transform)
	if matrix.Equal(s.Transform(), transform) != true {
		t.Error("Did not set transform on sphere.")
	}
}

func TestIntersetTransformedShapeWithRay(t *testing.T) {
	var tests = []struct {
		ray                               ray.Ray
		transform                         matrix.Matrix
		expectedOrigin, expectedDirection vector.Vector
	}{
		{
			// Intersecting a scaled shape with a ray.
			ray:               ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			transform:         matrix.ScalingMatrix(2, 2, 2),
			expectedOrigin:    vector.NewPoint(0, 0, -2.5),
			expectedDirection: vector.NewVector(0, 0, 0.5),
		},
		{
			// Intersecting a translated shape with a ray.
			ray:               ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			transform:         matrix.TranslationMatrix(5, 0, 0),
			expectedOrigin:    vector.NewPoint(-5, 0, -5),
			expectedDirection: vector.NewVector(0, 0, 1),
		},
	}
	for _, test := range tests {
		s := newShape()
		s.SetTransform(test.transform)
		xs := Intersect(&s, test.ray)
		_ = xs
		result := s.SavedRay()
		if !vector.Equal(result.Origin, test.expectedOrigin) ||
			!vector.Equal(result.Direction, test.expectedDirection) {
			t.Errorf(
				"Local Intersect produced %+v, expected (%v, %v).",
				result, test.expectedOrigin, test.expectedDirection,
			)
		}
	}
}

func TestNormalAt(t *testing.T) {
	var tests = []struct {
		transform       matrix.Matrix
		point, expected vector.Vector
	}{
		{
			// Computing the normal on a translated shape.
			transform: matrix.TranslationMatrix(0, 1, 0),
			point:     vector.NewPoint(0, 1.70711, -0.70711),
			expected:  vector.NewVector(0, 0.70711, -0.70711),
		},
		{
			// Computing the normal on a transformed shape.
			transform: matrix.Multiply(
				matrix.ScalingMatrix(1, 0.5, 1), matrix.RotationZMatrix(math.Pi/5)),
			point:    vector.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			expected: vector.NewVector(0, 0.97014, -0.24254),
		},
	}
	for _, test := range tests {
		s := newShape()
		s.SetTransform(test.transform)
		result := NormalAt(&s, test.point)
		if !vector.Equal(result, test.expected) {
			t.Errorf("NormalAt(%v) was %v, expected %v.", test.point, result, test.expected)
		}
	}
}

func TestNormalAtOnSphere(t *testing.T) {
	var tests = []struct {
		transform matrix.Matrix
		point     vector.Vector
		expected  vector.Vector
	}{
		{
			transform: matrix.IdentityMatrix(4),
			point:     vector.NewPoint(1, 0, 0),
			expected:  vector.NewVector(1, 0, 0),
		},
		{
			transform: matrix.IdentityMatrix(4),
			point:     vector.NewPoint(0, 1, 0),
			expected:  vector.NewVector(0, 1, 0),
		},
		{
			transform: matrix.IdentityMatrix(4),
			point:     vector.NewPoint(0, 0, 1),
			expected:  vector.NewVector(0, 0, 1),
		},
		{
			transform: matrix.IdentityMatrix(4),
			point:     vector.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			expected:  vector.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
		{
			transform: matrix.TranslationMatrix(0, 1, 0),
			point:     vector.NewPoint(0, 1.70711, -0.70711),
			expected:  vector.NewVector(0, 0.70711, -0.70711),
		},
		{
			transform: matrix.Multiply(
				matrix.ScalingMatrix(1, 0.5, 1),
				matrix.RotationZMatrix(math.Pi/5),
			),
			point:    vector.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			expected: vector.NewVector(0, 0.97014, -0.24254),
		},
	}
	for _, test := range tests {
		s := NewSphere()
		s.SetTransform(test.transform)
		result := NormalAt(s, test.point)
		if vector.Equal(result, test.expected) != true {
			t.Errorf(
				"The normal of sphere %+v at point %+v was %+v, expected %+v.",
				s, test.point, result, test.expected,
			)
		}
		if vector.Equal(result, result.Normalize()) != true {
			t.Errorf(
				"The normal of sphere %+v at point %+v was not normalized.",
				s, test.point,
			)
		}
	}
}

func TestGetID(t *testing.T) {
	expected := 0
	nextID = expected
	s := newShape()
	if s.ID() != expected {
		t.Errorf("First ID was %d, expected %d.", s.ID(), expected)
	}
	s = newShape()
	expected++
	if s.ID() != expected {
		t.Errorf("Second ID was %d, expected %d.", s.ID(), expected)
	}
	s = newShape()
	expected++
	if s.ID() != expected {
		t.Errorf("Second ID was %d, expected %d.", s.ID(), expected)
	}
}

func TestPlaneNormalAt(t *testing.T) {
	var tests = []struct {
		plane           Shape
		point, expected vector.Vector
	}{
		{
			plane:    NewPlane(),
			point:    vector.NewPoint(0, 0, 0),
			expected: vector.NewVector(0, 1, 0),
		},
		{
			plane:    NewPlane(),
			point:    vector.NewPoint(10, 0, -10),
			expected: vector.NewVector(0, 1, 0),
		},
		{
			plane:    NewPlane(),
			point:    vector.NewPoint(-5, 0, 150),
			expected: vector.NewVector(0, 1, 0),
		},
	}
	for _, test := range tests {
		result := test.plane.LocalNormalAt(test.point)
		if !vector.Equal(result, test.expected) {
			t.Errorf(
				"Plane local normal at %v was %v, expected %v.",
				test.point, result, test.expected,
			)
		}
	}
}

func TestPlaneLocalIntersect(t *testing.T) {
	var tests = []struct {
		plane    Shape
		ray      ray.Ray
		expected []float64
	}{
		{
			// Intersect with a ray parallel to the plane.
			plane:    NewPlane(),
			ray:      ray.New(vector.NewPoint(0, 10, 0), vector.NewVector(0, 0, 1)),
			expected: []float64{},
		},
		{
			// Intersect with a coplanar ray.
			plane:    NewPlane(),
			ray:      ray.New(vector.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1)),
			expected: []float64{},
		},
		{
			// A ray intersecting a plane from above.
			plane:    NewPlane(),
			ray:      ray.New(vector.NewPoint(0, 1, 0), vector.NewVector(0, -1, 0)),
			expected: []float64{1},
		},
		{
			// A ray intersecting a plane from below.
			plane:    NewPlane(),
			ray:      ray.New(vector.NewPoint(0, -1, 0), vector.NewVector(0, 1, 0)),
			expected: []float64{1},
		},
	}
	for _, test := range tests {
		intersections := test.plane.LocalIntersect(test.ray)
		result := intersections.TSlice()
		if !comparison.EqualSlice(result, test.expected) {
			t.Errorf(
				"Intersection of plane and ray (%v) was %v, expected %v",
				test.ray, result, test.expected,
			)
		}
	}
}
