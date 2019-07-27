package material

import "github.com/lukeshiner/raytrace/colour"

// Material holds data for materials.
type Material struct {
	Colour                                colour.Colour
	Ambient, Diffuse, Specular, Shininess float64
}

// New returns a new material
func New() Material {
	return Material{
		Colour: colour.New(1, 1, 1), Ambient: 0.1, Diffuse: 0.9, Specular: 0.9,
		Shininess: 200.0,
	}
}
