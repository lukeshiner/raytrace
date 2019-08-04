package shape

import (
	"fmt"
	"math"
	"sort"

	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/light"
	"github.com/lukeshiner/raytrace/material"
	"github.com/lukeshiner/raytrace/vector"
)

// Intersection holds an intersection.
type Intersection struct {
	T      float64
	Object Shape
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
	return NewIntersection(0, nil), fmt.Errorf("No intersections hit")
}

// NewIntersection returns an Intersection instance.
func NewIntersection(t float64, obj Shape) Intersection {
	return Intersection{T: t, Object: obj}
}

// NewIntersections returns a new Intersections list.
func NewIntersections(i ...Intersection) Intersections {
	intersections := Intersections{Intersections: i}
	intersections.sort()
	return intersections
}

// CombineIntersections returns combined intersection lists.
func CombineIntersections(intersections ...Intersections) Intersections {
	ins := NewIntersections()
	for i := 0; i < len(intersections); i++ {
		for x := 0; x < len(intersections[i].Intersections); x++ {
			ins.Intersections = append(ins.Intersections, intersections[i].Intersections[x])
		}
	}
	ins.sort()
	return ins
}

// Lighting calculates the lighting on a surface
func Lighting(
	m material.Material, l light.Light, p, e, n vector.Vector, inShadow bool,
) colour.Colour {
	var diffuse, specular colour.Colour
	effectiveColour := m.Colour.Mult(l.Intensity())
	lightVector := vector.Subtract(l.Position(), p)
	lightVector = lightVector.Normalize()
	ambient := effectiveColour.ScalarMult(m.Ambient)
	lightDotNormal := vector.DotProduct(lightVector, n)
	if inShadow || lightDotNormal < 0 {
		// Light behind surface
		diffuse = colour.New(0, 0, 0)
		specular = colour.New(0, 0, 0)
	} else {
		diffuse = effectiveColour.ScalarMult(m.Diffuse * lightDotNormal)
		reflectVector := Reflect(lightVector.Negate(), n)
		reflectDotEye := vector.DotProduct(reflectVector, e)
		if reflectDotEye <= 0 {
			// Light reflects away from eye
			specular = colour.New(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity().ScalarMult(m.Specular * factor)
		}
	}

	return ambient.Add(diffuse.Add(specular))
}

// Reflect returns the reflection of a vector around a normal.
func Reflect(in, normal vector.Vector) vector.Vector {
	var v vector.Vector
	v = normal.ScalarMultiply(2)
	v = v.ScalarMultiply(vector.DotProduct(in, normal))
	return vector.Subtract(in, v)
}
