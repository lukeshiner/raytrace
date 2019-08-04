package ray

import (
	"math"
	"testing"

	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/comparison"
	"github.com/lukeshiner/raytrace/light"
	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/shape"
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
		object   shape.Shape
		expected []float64
	}{
		{
			ray:      New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			object:   shape.NewSphere(),
			expected: []float64{4.0, 6.0},
		},
		{
			ray:      New(vector.NewPoint(0, 1, -5), vector.NewVector(0, 0, 1)),
			object:   shape.NewSphere(),
			expected: []float64{5.0, 5.0},
		},
		{
			ray:      New(vector.NewPoint(0, 2, -5), vector.NewVector(0, 0, 1)),
			object:   shape.NewSphere(),
			expected: []float64{},
		},
		{
			ray:      New(vector.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1)),
			object:   shape.NewSphere(),
			expected: []float64{-1.0, 1.0},
		},
		{
			ray:      New(vector.NewPoint(0, 0, 5), vector.NewVector(0, 0, 1)),
			object:   shape.NewSphere(),
			expected: []float64{-6.0, -4.0},
		},
	}
	for _, test := range tests {
		intersections := Intersect(test.object, test.ray)
		xs := intersections.TSlice()
		if comparison.EqualSlice(xs, test.expected) != true {
			t.Errorf(
				"Intersection of Ray %+v and Shpere %+v was %+v, expected %+v.",
				test.ray, test.object, xs, test.expected,
			)
		}
	}
}

func TestIntersection(t *testing.T) {
	var tests = []struct {
		t float64
		o shape.Shape
	}{
		{
			t: 3.5,
			o: shape.NewSphere(),
		},
	}
	for _, test := range tests {
		intersection := NewIntersection(test.t, test.o)
		if intersection.T != test.t || intersection.Object.ID() != test.o.ID() {
			t.Error("Error creating Intersection")
		}
	}
}

func TestIntersections(t *testing.T) {
	s := shape.NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := NewIntersections(i1, i2)
	if xs.Count() != 2 || xs.Get(0).T != 1 || xs.Get(1).T != 2 {
		t.Error("Intersect did not set shape.")
	}
}

func TestIntersectionsSetsObject(t *testing.T) {
	ray := New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	sphere := shape.NewSphere()
	xs := Intersect(sphere, ray)
	if xs.Get(0).Object.ID() != sphere.ID() || xs.Get(1).Object.ID() != sphere.ID() {
		t.Error("Intersect did not set shape.")
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
				NewIntersection(1, shape.NewSphere()),
				NewIntersection(2, shape.NewSphere()),
			),
			expectNil: false,
			expected:  1,
		},
		{
			intersections: NewIntersections(
				NewIntersection(-1, shape.NewSphere()),
				NewIntersection(1, shape.NewSphere()),
			),
			expectNil: false,
			expected:  1,
		},
		{
			intersections: NewIntersections(
				NewIntersection(-2, shape.NewSphere()),
				NewIntersection(-1, shape.NewSphere()),
			),
			expectNil: true,
		},
		{
			intersections: NewIntersections(
				NewIntersection(5, shape.NewSphere()),
				NewIntersection(7, shape.NewSphere()),
				NewIntersection(-3, shape.NewSphere()),
				NewIntersection(2, shape.NewSphere()),
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
		s := shape.NewSphere()
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

func TestLighting(t *testing.T) {
	var tests = []struct {
		material              material.Material
		light                 light.Light
		position, eye, normal vector.Vector
		inShadow              bool
		expected              colour.Colour
	}{
		{
			// Lighting with the eye between the light and the surface.
			material: material.New(),
			position: vector.NewPoint(0, 0, 0),
			light:    light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0, -10)),
			normal:   vector.NewVector(0, 0, -1),
			eye:      vector.NewVector(0, 0, -1),
			inShadow: false,
			expected: colour.New(1.9, 1.9, 1.9),
		},
		{
			// Lighting with the eye between light and suface, eye offset 45 degrees.
			material: material.New(),
			position: vector.NewPoint(0, 0, 0),
			light:    light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0, -10)),
			normal:   vector.NewVector(0, 0, -1),
			eye:      vector.NewVector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			inShadow: false,
			expected: colour.New(1.0, 1.0, 1.0),
		},
		{
			// Lighting with the eye opposite surface, light offset 45 degrees.
			material: material.New(),
			position: vector.NewPoint(0, 0, 0),
			light:    light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 10, -10)),
			normal:   vector.NewVector(0, 0, -1),
			eye:      vector.NewVector(0, 0, -1),
			inShadow: false,
			expected: colour.New(0.7364, 0.7364, 0.7364),
		},
		{
			// Lighting with the eye in the path of the reflection vector.
			material: material.New(),
			position: vector.NewPoint(0, 0, 0),
			light:    light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 10, -10)),
			normal:   vector.NewVector(0, 0, -1),
			eye:      vector.NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2),
			inShadow: false,
			expected: colour.New(1.6364, 1.6364, 1.6364),
		},
		{
			// Lighting with the light behind the surface.
			material: material.New(),
			position: vector.NewPoint(0, 0, 0),
			light:    light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0, 10)),
			normal:   vector.NewVector(0, 0, -1),
			eye:      vector.NewVector(0, 0, -1),
			inShadow: false,
			expected: colour.New(0.1, 0.1, 0.1),
		},
		{
			// Lighting with reflection away from eye.
			material: material.New(),
			position: vector.NewPoint(0, 10, 10),
			light:    light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0, 10)),
			normal:   vector.NewVector(0, 0, -1),
			eye:      vector.NewVector(0, 0, -1),
			inShadow: false,
			expected: colour.New(0.1, 0.1, 0.1),
		},
		{
			// Lighting with the surface in shadow.
			material: material.New(),
			position: vector.NewPoint(0, 0, 0),
			light:    light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0, -10)),
			normal:   vector.NewVector(0, 0, -1),
			eye:      vector.NewVector(0, 0, -1),
			inShadow: true,
			expected: colour.New(0.1, 0.1, 0.1),
		},
	}
	for _, test := range tests {
		result := Lighting(
			test.material, test.light, test.position, test.eye, test.normal, test.inShadow)
		if result.Equal(test.expected) != true {
			t.Errorf(
				"Lighting with material %+v, light %+v, position %+v, eye %+v and normal %+v "+
					"resulted with %+v, expected %+v",
				test.material, test.light, test.position, test.eye, test.normal, result,
				test.expected,
			)
		}
	}
}

func TestAddIntersections(t *testing.T) {
	o1 := shape.NewSphere()
	o2 := shape.NewSphere()
	i1 := NewIntersections(NewIntersection(9.8, o1), NewIntersection(5.6, o1))
	i2 := NewIntersections(NewIntersection(2.4, o2), NewIntersection(10.5, o2))
	ins := CombineIntersections(i1, i2)
	if len(ins.Intersections) != 4 {
		t.Error("Wrong number of intersections.")
	}
	values := ins.TSlice()
	expected := []float64{2.4, 5.6, 9.8, 10.5}
	if comparison.EqualSlice(values, expected) != true {
		t.Errorf("Intersection combination returned %+v, expected %+v.", values, expected)
	}
}
