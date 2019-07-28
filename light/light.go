package light

import (
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/vector"
)

// Light is the interface for lights
type Light interface {
	Intensity() colour.Colour
	SetIntensity(i colour.Colour)
	Position() vector.Vector
	SetPosition(p vector.Vector)
}

// Point holds point light data.
type Point struct {
	intensity colour.Colour
	position  vector.Vector
}

// Intensity returns the intensity of the light.
func (p Point) Intensity() colour.Colour {
	return p.intensity
}

// SetIntensity sets the intensity of the light
func (p *Point) SetIntensity(i colour.Colour) {
	p.intensity = i
}

// Position returns the position of the light.
func (p Point) Position() vector.Vector {
	return p.position
}

// SetPosition sets the intensity of the light
func (p *Point) SetPosition(pos vector.Vector) {
	p.position = pos
}

// NewPoint creates a new point light
func NewPoint(intensity colour.Colour, position vector.Vector) Point {
	return Point{intensity: intensity, position: position}
}
