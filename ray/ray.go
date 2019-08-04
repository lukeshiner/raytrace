package ray

import (
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/vector"
)

// Ray is the struct for raytracer rays.
type Ray struct {
	Origin, Direction vector.Vector
}

// Position returns the position of the ray at time t.
func (r *Ray) Position(t float64) vector.Vector {
	return vector.Add(r.Origin, r.Direction.ScalarMultiply(t))
}

// Transform transforms a ray by a transform matrix.
func (r *Ray) Transform(m matrix.Matrix) Ray {
	origin := vector.MultiplyMatrixByVector(m, r.Origin)
	direction := vector.MultiplyMatrixByVector(m, r.Direction)
	return New(origin, direction)
}

// New creates a new Ray struct.
func New(origin, direction vector.Vector) Ray {
	return Ray{Origin: origin, Direction: direction}
}
