package ray

import (
	"testing"

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
