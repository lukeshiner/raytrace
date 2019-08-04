package world

import (
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/comparison"
	"github.com/lukeshiner/raytrace/light"
	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/ray"
	"github.com/lukeshiner/raytrace/shape"
	"github.com/lukeshiner/raytrace/vector"
)

// World holds world data.
type World struct {
	Objects []shape.Shape
	Lights  []light.Light
}

// New returns an empty world.
func New() World {
	return World{}
}

// Default returns a default world with a light at two spheres.
func Default() World {
	w := New()
	s1 := shape.NewSphere()
	m := material.New()
	m.Colour = colour.New(0.8, 1.0, 0.6)
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.SetMaterial(m)
	s2 := shape.NewSphere()
	s2.SetTransform(matrix.ScalingMatrix(0.5, 0.5, 0.5))
	l := light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(-10, 10, -10))
	w.Objects = []shape.Shape{s1, s2}
	w.Lights = []light.Light{l}
	return w
}

// IntersectWorld returns the intersections of a ray with objects in the world.
func IntersectWorld(w World, r ray.Ray) shape.Intersections {
	var intersections = []shape.Intersections{}
	for i := 0; i < len(w.Objects); i++ {
		intersections = append(intersections, shape.Intersect(w.Objects[i], r))
	}
	return shape.CombineIntersections(intersections...)
}

// Comps holds computations for ray intersections.
type Comps struct {
	T                               float64
	Object                          shape.Shape
	Point, EyeV, NormalV, OverPoint vector.Vector
	Inside                          bool
}

// PrepareComputations returns a Comps for an intersection and a ray.
func PrepareComputations(i shape.Intersection, r ray.Ray) Comps {
	inside := false
	point := r.Position(i.T)
	eyeV := r.Direction.Negate()
	normalV := shape.NormalAt(i.Object, point)
	if vector.DotProduct(normalV, eyeV) < 0 {
		inside = true
		normalV = normalV.Negate()
	}
	overPoint := vector.Add(point, normalV.ScalarMultiply(comparison.EPSLION))
	return Comps{
		T: i.T, Object: i.Object, Point: point, EyeV: eyeV, NormalV: normalV, Inside: inside,
		OverPoint: overPoint,
	}
}

// ShadeHit returns the colour for a computed intersection.
func ShadeHit(world World, comps Comps) colour.Colour {
	var lightColour colour.Colour
	var shadowed bool
	c := colour.New(0, 0, 0)
	for i := 0; i < len(world.Lights); i++ {
		shadowed = IsShadowed(world, comps.OverPoint, world.Lights[i])
		lightColour = shape.Lighting(
			comps.Object.Material(), world.Lights[i], comps.Point, comps.EyeV, comps.NormalV,
			shadowed,
		)
		c = c.Add(lightColour)
	}
	return c
}

// ColourAt returns the colour for a given ray in a given world.
func ColourAt(w World, r ray.Ray) colour.Colour {
	intersections := IntersectWorld(w, r)
	hit, err := intersections.Hit()
	if err != nil {
		return colour.New(0, 0, 0)
	}
	comps := PrepareComputations(hit, r)
	return ShadeHit(w, comps)
}

// IsShadowed returns true if a point in the world is shadowed from light.
func IsShadowed(w World, p vector.Vector, l light.Light) bool {
	v := vector.Subtract(l.Position(), p)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := ray.New(p, direction)
	intersections := IntersectWorld(w, r)

	h, err := intersections.Hit()
	if err == nil && h.T < distance {
		return true
	}
	return false
}
