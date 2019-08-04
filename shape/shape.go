package shape

import (
	"math"

	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/ray"
	"github.com/lukeshiner/raytrace/vector"
)

// Shape is an interface for objects
type Shape interface {
	ID() int
	Material() material.Material
	SetMaterial(m material.Material)
	Transform() matrix.Matrix
	SetTransform(m matrix.Matrix)
	InverseTransform() matrix.Matrix
	LocalIntersect(r ray.Ray) Intersections
	LocalNormalAt(p vector.Vector) vector.Vector
	SavedRay() ray.Ray
	SaveRay(ray.Ray)
}

// Sphere is the struct for spheres
type shape struct {
	id        int
	material  material.Material
	transform matrix.Matrix
	savedRay  ray.Ray
}

// ID returns the ID of the object
func (s shape) ID() int {
	return s.id
}

// Material returns the material of the shape.
func (s shape) Material() material.Material {
	return s.material
}

// SetMaterial sets a shape's material.
func (s *shape) SetMaterial(m material.Material) {
	s.material = m
}

// Transform returns the shape's transform matrix.
func (s shape) Transform() matrix.Matrix {
	return s.transform
}

// SetTransform sets the transform matrix for the shape.
func (s *shape) SetTransform(m matrix.Matrix) {
	s.transform = m
}

// InverseTransform returns the inverse of the shapes transform matrix.
func (s shape) InverseTransform() matrix.Matrix {
	t, _ := s.Transform().Invert()
	return t
}

func (s *shape) LocalIntersect(r ray.Ray) Intersections {
	s.SaveRay(r)
	return Intersections{}
}

// Transform returns the sphere's transform matrix.
func (s shape) SavedRay() ray.Ray {
	return s.savedRay
}

// SetTransform sets the transform matrix for a sphere.
func (s *shape) SaveRay(r ray.Ray) {
	s.savedRay = r
}

// LocalNormalAt returns the normal vector of the sphere at the given point in local space.
func (s *shape) LocalNormalAt(p vector.Vector) vector.Vector {
	return vector.NewPoint(p.X, p.Y, p.Z)
}

func newShape() shape {
	return shape{
		id: getID(), material: material.New(), transform: matrix.IdentityMatrix(4),
	}
}

// Sphere is the type for spheres.
type Sphere struct {
	shape
}

// LocalIntersect returns a list of intersectioins between ray and the shape in local space.
func (s Sphere) LocalIntersect(r ray.Ray) Intersections {
	sphereToRay := vector.Subtract(r.Origin, vector.NewPoint(0, 0, 0))
	a := vector.DotProduct(r.Direction, r.Direction)
	b := 2 * vector.DotProduct(r.Direction, sphereToRay)
	c := vector.DotProduct(sphereToRay, sphereToRay) - 1
	discriminant := (b * b) - 4*a*c
	if discriminant < 0 {
		return Intersections{}
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	i1 := NewIntersection(t1, &s)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	i2 := NewIntersection(t2, &s)
	return NewIntersections(i1, i2)
}

// LocalNormalAt returns the normal vector of the sphere at the given point in local space.
func (s Sphere) LocalNormalAt(p vector.Vector) vector.Vector {
	return vector.Subtract(p, vector.NewPoint(0, 0, 0))
}

// NewSphere returns a unit sphere at the origin
func NewSphere() Shape {
	return &Sphere{newShape()}
}

// Intersect returns a list of intersectioins between ray and the shape.
func Intersect(s Shape, r ray.Ray) Intersections {
	localRay := r.Transform(s.InverseTransform())
	return s.LocalIntersect(localRay)
}

// NormalAt returns the normal vector of a shape at the given point.
func NormalAt(s Shape, p vector.Vector) vector.Vector {
	localPoint := vector.MultiplyMatrixByVector(s.InverseTransform(), p)
	localNormal := s.LocalNormalAt(localPoint)
	worldNormal := vector.MultiplyMatrixByVector(
		s.InverseTransform().Transpose(), localNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}

var nextID = 0

func getID() int {
	nextID++
	return nextID - 1
}
