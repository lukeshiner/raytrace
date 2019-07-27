package material

import (
	"testing"

	"github.com/lukeshiner/raytrace/colour"
)

func TestDefaultMaterial(t *testing.T) {
	var tests = []struct {
		colour                                colour.Colour
		ambient, diffuse, specular, shininess float64
	}{
		{
			colour:    colour.New(1, 1, 1),
			ambient:   0.1,
			diffuse:   0.9,
			specular:  0.9,
			shininess: 200.0,
		},
	}
	for _, test := range tests {
		m := New()
		if m.Colour != test.colour || m.Ambient != test.ambient ||
			m.Diffuse != test.diffuse || m.Specular != test.specular ||
			m.Shininess != test.shininess {
			t.Error("Error creating material.")
		}
	}
}
