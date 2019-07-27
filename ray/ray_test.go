package ray

import (
	"math"
	"testing"

	"github.com/lukeshiner/raytrace/comparison"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/vector"
)

func TestRay(t *testing.T) {
	var tests = []struct {
		origin, direction vector.Vector
	}{
		{
			origin:    vector.NewPoint(1, 2, 3),
			direction: vector.NewVector(4, 5, 6),
		},
	}
	for _, test := range tests {
		ray := New(test.origin, test.direction)
		if vector.Equal(test.origin, ray.Origin) != true {
			t.Errorf("Ray origin was %+v, expected %+v.", ray.Origin, test.origin)
		}
		if vector.Equal(test.direction, ray.Direction) != true {
			t.Errorf("Ray direction was %+v, expected %+v.", ray.Direction, test.direction)
		}
	}
}

func TestPosition(t *testing.T) {
	var tests = []struct {
		ray      Ray
		t        float64
		expected vector.Vector
	}{
		{
			ray:      New(vector.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0)),
			t:        0,
			expected: vector.NewPoint(2, 3, 4),
		},
		{
			ray:      New(vector.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0)),
			t:        1,
			expected: vector.NewPoint(3, 3, 4),
		},
		{
			ray:      New(vector.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0)),
			t:        -1,
			expected: vector.NewPoint(1, 3, 4),
		},
		{
			ray:      New(vector.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0)),
			t:        2.5,
			expected: vector.NewPoint(4.5, 3, 4),
		},
	}
	for _, test := range tests {
		result := test.ray.Position(test.t)
		if vector.Equal(result, test.expected) != true {
			t.Errorf(
				"Ray %+v at position %+v was %+v, expected %+v.",
				test.ray, test.t, result, test.expected,
			)
		}
	}
}

func TestIntersect(t *testing.T) {
	var tests = []struct {
		ray      Ray
		sphere   Sphere
		expected []float64
	}{
		{
			ray:      New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			sphere:   NewSphere(),
			expected: []float64{4.0, 6.0},
		},
		{
			ray:      New(vector.NewPoint(0, 1, -5), vector.NewVector(0, 0, 1)),
			sphere:   NewSphere(),
			expected: []float64{5.0, 5.0},
		},
		{
			ray:      New(vector.NewPoint(0, 2, -5), vector.NewVector(0, 0, 1)),
			sphere:   NewSphere(),
			expected: []float64{},
		},
		{
			ray:      New(vector.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1)),
			sphere:   NewSphere(),
			expected: []float64{-1.0, 1.0},
		},
		{
			ray:      New(vector.NewPoint(0, 0, 5), vector.NewVector(0, 0, 1)),
			sphere:   NewSphere(),
			expected: []float64{-6.0, -4.0},
		},
	}
	for _, test := range tests {
		intersections := Intersect(test.sphere, test.ray)
		xs := intersections.TSlice()
		if comparison.EqualSlice(xs, test.expected) != true {
			t.Errorf(
				"Intersection of Ray %+v and Shpere %+v was %+v, expected %+v.",
				test.ray, test.sphere, xs, test.expected,
			)
		}
	}
}

func TestIntersection(t *testing.T) {
	var tests = []struct {
		t float64
		o Sphere
	}{
		{
			t: 3.5,
			o: NewSphere(),
		},
	}
	for _, test := range tests {
		intersection := Intersection{test.t, test.o}
		if intersection.T != test.t || intersection.Object.ID != test.o.ID {
			t.Error("Error creating Intersection")
		}
	}
}

func TestIntersections(t *testing.T) {
	s := NewSphere()
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections{[]Intersection{i1, i2}}
	if xs.Count() != 2 || xs.Get(0).T != 1 || xs.Get(1).T != 2 {
		t.Error("Intersect did not set object.")
	}
}

func TestIntersectionsSetsObject(t *testing.T) {
	ray := New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	sphere := NewSphere()
	xs := Intersect(sphere, ray)
	if xs.Get(0).Object.ID != sphere.ID || xs.Get(1).Object.ID != sphere.ID {
		t.Error("Intersect did not set object.")
	}
}

func TestHit(t *testing.T) {
	var tests = []struct {
		intersections Intersections
		expectNil     bool
		expected      float64
	}{
		{
			intersections: NewIntersections(
				NewIntersection(1, NewSphere()),
				NewIntersection(2, NewSphere()),
			),
			expectNil: false,
			expected:  1,
		},
		{
			intersections: NewIntersections(
				NewIntersection(-1, NewSphere()),
				NewIntersection(1, NewSphere()),
			),
			expectNil: false,
			expected:  1,
		},
		{
			intersections: NewIntersections(
				NewIntersection(-2, NewSphere()),
				NewIntersection(-1, NewSphere()),
			),
			expectNil: true,
		},
		{
			intersections: NewIntersections(
				NewIntersection(5, NewSphere()),
				NewIntersection(7, NewSphere()),
				NewIntersection(-3, NewSphere()),
				NewIntersection(2, NewSphere()),
			),
			expectNil: false,
			expected:  2,
		},
	}
	for _, test := range tests {
		result, err := test.intersections.Hit()
		if err != nil && test.expectNil != true {
			t.Errorf(
				"Intersections %+v had a hit when none was expected",
				test.intersections,
			)
		} else {
			if result.T != test.expected {
				t.Errorf(
					"Intersections %+v hit was %v, expected %v.",
					test.intersections, result.T, test.expected,
				)
			}
		}
	}
}

func TestTransformRay(t *testing.T) {
	var tests = []struct {
		ray      Ray
		m        matrix.Matrix
		expected Ray
	}{
		{
			ray:      New(vector.NewPoint(1, 2, 3), vector.NewVector(0, 1, 0)),
			m:        matrix.TranslationMatrix(3, 4, 5),
			expected: New(vector.NewPoint(4, 6, 8), vector.NewVector(0, 1, 0)),
		},
		{
			ray:      New(vector.NewPoint(1, 2, 3), vector.NewVector(0, 1, 0)),
			m:        matrix.ScalingMatrix(2, 3, 4),
			expected: New(vector.NewPoint(2, 6, 12), vector.NewVector(0, 3, 0)),
		},
	}
	for _, test := range tests {
		result := test.ray.Transform(test.m)
		if vector.Equal(result.Origin, test.expected.Origin) != true ||
			vector.Equal(result.Direction, test.expected.Direction) != true {
			t.Errorf(
				"Transforming ray %+v with matrix %+v results in %+v, expected %+v.",
				test.ray, test.m, result, test.expected,
			)
		}
	}
}

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

func TestIntersectionWithTransformedSphere(t *testing.T) {
	var tests = []struct {
		ray       Ray
		transform matrix.Matrix
		expected  []float64
	}{
		{
			ray:       New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			transform: matrix.ScalingMatrix(2, 2, 2),
			expected:  []float64{3, 7},
		},
		{
			ray:       New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			transform: matrix.TranslationMatrix(5, 0, 0),
			expected:  []float64{},
		},
	}
	for _, test := range tests {
		s := NewSphere()
		s.SetTransform(test.transform)
		xs := Intersect(s, test.ray)
		if comparison.EqualSlice(xs.TSlice(), test.expected) != true {
			t.Errorf(
				"Intersection of ray %+v with sphere transformed by %+v gave %+v, expected %+v.",
				test.ray, test.transform, xs.TSlice(), test.expected,
			)
		}
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

func TestReflect(t *testing.T) {
	var tests = []struct {
		vector, normal, expected vector.Vector
	}{
		{
			vector:   vector.NewVector(1, -1, 0),
			normal:   vector.NewVector(0, 1, 0),
			expected: vector.NewVector(1, 1, 0),
		},
		{
			vector:   vector.NewVector(0, -1, 0),
			normal:   vector.NewVector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
			expected: vector.NewVector(1, 0, 0),
		},
	}
	for _, test := range tests {
		result := Reflect(test.vector, test.normal)
		if vector.Equal(result, test.expected) != true {
			t.Errorf(
				"The reflection of vector %+v around normal %+v was %+v, expected %+v.",
				test.vector, test.normal, result, test.expected,
			)
		}
	}
}