package world

import (
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/light"
	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/object"
	"github.com/lukeshiner/raytrace/ray"
	"github.com/lukeshiner/raytrace/vector"
)

// World holds world data.
type World struct {
	Objects []object.Object
	Lights  []light.Light
}

// New returns an empty world.
func New() World {
	return World{}
}

// Default returns a default world with a light at two spheres.
func Default() World {
	w := New()
	s1 := object.NewSphere()
	m := material.New()
	m.Colour = colour.New(0.8, 1.0, 0.6)
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.SetMaterial(m)
	s2 := object.NewSphere()
	s2.SetTransform(matrix.ScalingMatrix(0.5, 0.5, 0.5))
	l := light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(-10, 10, -10))
	w.Objects = []object.Object{&s1, &s2}
	w.Lights = []light.Light{&l}
	return w
}

// IntersectWorld returns the intersections of a ray with objects in the world.
func IntersectWorld(w World, r ray.Ray) ray.Intersections {
	var intersections = []ray.Intersections{}
	for i := 0; i < len(w.Objects); i++ {
		intersections = append(intersections, ray.Intersect(w.Objects[i], r))
	}
	return ray.CombineIntersections(intersections...)
}

// Comps holds computations for ray intersections.
type Comps struct {
	T                    float64
	Object               object.Object
	Point, EyeV, NormalV vector.Vector
}

// PrepareComputations returns a Comps for an intersection and a ray.
func PrepareComputations(i ray.Intersection, r ray.Ray) Comps {
	point := r.Position(i.T)
	eyeV := r.Direction.Negate()
	normalV := i.Object.NormalAt(point)
	return Comps{T: i.T, Object: i.Object, Point: point, EyeV: eyeV, NormalV: normalV}
}
