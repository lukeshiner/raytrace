package shape

import (
	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/vector"
)

// Shape is an interface for objects
type Shape interface {
	ID() int
	Material() material.Material
	SetMaterial(m material.Material)
	Transform() matrix.Matrix
	SetTransform(m matrix.Matrix)
	NormalAt(p vector.Vector) vector.Vector
}

// Sphere is the struct for spheres
type shape struct {
	id        int
	material  material.Material
	transform matrix.Matrix
}

// ID returns the ID of the object
func (s shape) ID() int {
	return s.id
}

// Material returns the material of the sphere
func (s shape) Material() material.Material {
	return s.material
}

// SetMaterial sets a sphere's material.
func (s *shape) SetMaterial(m material.Material) {
	s.material = m
}

// Transform returns the sphere's transform matrix.
func (s shape) Transform() matrix.Matrix {
	return s.transform
}

// SetTransform sets the transform matrix for a sphere.
func (s *shape) SetTransform(m matrix.Matrix) {
	s.transform = m
}

// NormalAt returns the normal vector of the sphere at the given point.
func (s *shape) NormalAt(p vector.Vector) vector.Vector {
	transform, _ := s.Transform().Invert()
	objectPoint := vector.MultiplyMatrixByVector(transform, p)
	normal := vector.Subtract(objectPoint, vector.NewPoint(0, 0, 0))
	worldNormal := vector.MultiplyMatrixByVector(transform.Transpose(), normal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}

func newShape() shape {
	return shape{
		id: getID(), material: material.New(), transform: matrix.IdentityMatrix(4),
	}
}

type testShape struct {
	shape
}

func newTestShape() Shape {
	return &testShape{newShape()}
}

// Sphere is the type for spheres.
type Sphere struct {
	shape
}

// NewSphere returns a unit sphere at the origin
func NewSphere() Shape {
	return &Sphere{newShape()}
}

var nextID = 0

func getID() int {
	nextID++
	return nextID - 1
}
