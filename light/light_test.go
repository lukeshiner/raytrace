package light

import (
	"testing"

	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/vector"
)

func TestPoint(t *testing.T) {
	var tests = []struct {
		intensity colour.Colour
		position  vector.Vector
	}{
		{
			intensity: colour.New(1, 1, 1),
			position:  vector.NewPoint(0, 0, 0),
		},
	}
	for _, test := range tests {
		l := NewPoint(test.intensity, test.position)
		if vector.Equal(l.Position(), test.position) != true ||
			l.Intensity() != test.intensity {
			t.Errorf(
				"Creating Point Light with intensity %+v and position %+v produced %+v.",
				test.intensity, test.position, l,
			)
		}
	}
}

func TestSetPointIntensity(t *testing.T) {
	l := NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0, 0))
	newIntensity := colour.New(0.5, 0.5, 0.5)
	l.SetIntensity(newIntensity)
	if l.Intensity().Equal(newIntensity) != true {
		t.Error("Could not set point light intensity.")
	}
}

func TestPointPosition(t *testing.T) {
	l := NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0, 0))
	newPosition := vector.NewPoint(10, 10, 10)
	l.SetPosition(newPosition)
	if vector.Equal(l.Position(), newPosition) != true {
		t.Error("Could not set point light position.")
	}
}
