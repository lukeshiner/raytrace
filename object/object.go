package object

import (
	"math/rand"

	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/vector"
)

// Sphere is the struct for spheres
type Sphere struct {
	ID        string
	Transform matrix.Matrix
}

// SetTransform sets the transform matrix for a sphere.
func (s *Sphere) SetTransform(m matrix.Matrix) {
	s.Transform = m
}

// NormalAt returns the normal vector of the sphere at the given point.
func (s *Sphere) NormalAt(p vector.Vector) vector.Vector {
	transform, _ := s.Transform.Invert()
	objectPoint := vector.MultiplyMatrixByVector(transform, p)
	normal := vector.Subtract(objectPoint, vector.NewPoint(0, 0, 0))
	worldNormal := vector.MultiplyMatrixByVector(transform.Transpose(), normal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}

// NewSphere returns a unit sphere at the origin
func NewSphere() Sphere {
	return Sphere{ID: generateID(), Transform: matrix.IdentityMatrix(4)}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateID() string {
	b := make([]rune, 100)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
