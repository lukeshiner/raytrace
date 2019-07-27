package light

import (
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/vector"
)

// Point holds point light data.
type Point struct {
	Intensity colour.Colour
	Position  vector.Vector
}

// NewPoint creates a new point light
func NewPoint(intensity colour.Colour, position vector.Vector) Point {
	return Point{Intensity: intensity, Position: position}
}
