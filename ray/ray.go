package ray

import (
	"fmt"
	"math"
	"sort"

	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/object"
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

// Reflect returns the reflection of a vector around a normal.
func Reflect(in, normal vector.Vector) vector.Vector {
	var v vector.Vector
	v = normal.ScalarMultiply(2)
	v = v.ScalarMultiply(vector.DotProduct(in, normal))
	return vector.Subtract(in, v)
}

// Intersection holds an intersection.
type Intersection struct {
	T      float64
	Object object.Sphere
}

// Intersections holds a slice of Intersection.
type Intersections struct {
	Intersections []Intersection
}

// Count returns the number of intersections.
func (i *Intersections) Count() int {
	return len(i.Intersections)
}

// Get returns an intersection by index
func (i *Intersections) Get(index int) Intersection {
	return i.Intersections[index]
}

// TSlice returns a slice of t values for the intersections.
func (i *Intersections) TSlice() []float64 {
	var ts []float64
	for x := 0; x < len(i.Intersections); x++ {
		ts = append(ts, i.Get(x).T)
	}
	return ts
}

func (i Intersections) sort() []Intersection {
	sort.Slice(i.Intersections, func(j, k int) bool {
		return i.Intersections[j].T < i.Intersections[k].T
	})
	return i.Intersections
}

// Hit returns the first hit in the intersections or an error if there are none.
func (i *Intersections) Hit() (Intersection, error) {
	i.sort()
	for _, intersection := range i.Intersections {
		if intersection.T >= 0 {
			return intersection, nil
		}
	}
	return NewIntersection(0, object.NewSphere()), fmt.Errorf("No intersections hit")
}

// NewIntersection returns an Intersection instance.
func NewIntersection(t float64, obj object.Sphere) Intersection {
	return Intersection{T: t, Object: obj}
}

// NewIntersections returns a new Intersections list.
func NewIntersections(i ...Intersection) Intersections {
	intersections := Intersections{Intersections: i}
	intersections.sort()
	return intersections
}

// New creates a new Ray struct.
func New(origin, direction vector.Vector) Ray {
	return Ray{Origin: origin, Direction: direction}
}

// Intersect returns a list of intersectioins between ray and sphere.
func Intersect(sphere object.Sphere, ray Ray) Intersections {
	transform, _ := sphere.Transform.Invert()
	tRay := ray.Transform(transform)
	sphereToRay := vector.Subtract(tRay.Origin, vector.NewPoint(0, 0, 0))
	a := vector.DotProduct(tRay.Direction, tRay.Direction)
	b := 2 * vector.DotProduct(tRay.Direction, sphereToRay)
	c := vector.DotProduct(sphereToRay, sphereToRay) - 1
	discriminant := (b * b) - 4*a*c
	if discriminant < 0 {
		return Intersections{}
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	i1 := NewIntersection(t1, sphere)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	i2 := NewIntersection(t2, sphere)
	return NewIntersections(i1, i2)
}
