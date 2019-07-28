package world

import (
	"testing"

	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/comparison"
	"github.com/lukeshiner/raytrace/light"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/object"
	"github.com/lukeshiner/raytrace/ray"
	"github.com/lukeshiner/raytrace/vector"
)

func TestCreateWorld(t *testing.T) {
	w := New()
	if len(w.Objects) != 0 {
		t.Errorf("World created with %d objects.", len(w.Objects))
	}
	if len(w.Lights) != 0 {
		t.Errorf("World created with %d objects.", len(w.Lights))
	}
}

func TestDefaultWorld(t *testing.T) {
	w := Default()
	l := w.Lights[0]
	if len(w.Lights) != 1 {
		t.Errorf(
			"Default world has incorrect number of lights. Had %d, expected 1.",
			len(w.Lights),
		)
	}
	expectedLight := light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(-10, 10, -10))
	if l.Intensity().Equal(expectedLight.Intensity()) != true ||
		vector.Equal(l.Position(), expectedLight.Position()) != true {
		t.Errorf("Default world lights %+v, expected %+v.", w.Lights[0], expectedLight)
	}
	if len(w.Objects) != 2 {
		t.Errorf("Default world created with %d objects, expected 2.", len(w.Objects))
	}
	switch w.Objects[0].(type) {
	case *object.Sphere:
		break
	default:
		t.Error("Default world first object was not object.Sphere.")
	}
	m := w.Objects[0].Material()
	if m.Colour.Equal(colour.New(0.8, 1.0, 0.6)) != true ||
		m.Diffuse != 0.7 || m.Specular != 0.2 {
		t.Errorf("Default world first object has incorrect material: %+v.", m)
	}
	switch w.Objects[1].(type) {
	case *object.Sphere:
		break
	default:
		t.Error("Default world second object was not object.Sphere.")
	}
	if matrix.Equal(w.Objects[1].Transform(), matrix.ScalingMatrix(0.5, 0.5, 0.5)) != true {
		t.Error("Default world second object not transformed correctly.")
	}
	w.Lights[0].SetIntensity(colour.New(0.75, 0.75, 0.75))
}

func TestIntersectWorld(t *testing.T) {
	w := Default()
	r := ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	xs := IntersectWorld(w, r)
	expected := []float64{4, 4.5, 5.5, 6}
	result := xs.TSlice()
	if comparison.EqualSlice(result, expected) != true {
		t.Errorf(
			"Intersection of default world with ray %+v was %+v, expected %+v",
			r,
			result,
			expected,
		)
	}
}

func TestPrepareComputations(t *testing.T) {
	r := ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	s := object.NewSphere()
	i := ray.NewIntersection(4, &s)
	comps := PrepareComputations(i, r)
	expectedPoint := vector.NewPoint(0, 0, -1)
	expectedEyeV := vector.NewVector(0, 0, -1)
	expectedNormalv := vector.NewVector(0, 0, -1)
	if comps.T != i.T {
		t.Errorf("Comps.T was %v, expected %v.", comps.T, i.T)
	}
	if comps.Object.ID() != s.ID() {
		t.Errorf("Comps.Object was %+v, expected %+v.", comps.Object, s)
	}
	if vector.Equal(comps.Point, expectedPoint) != true {
		t.Errorf("Comps.Point was %+v, expected %+v.", comps.Point, expectedPoint)
	}
	if vector.Equal(comps.EyeV, expectedEyeV) != true {
		t.Errorf("Comps.EyeV was %+v, expected %+v.", comps.EyeV, expectedEyeV)
	}
	if vector.Equal(comps.EyeV, expectedNormalv) != true {
		t.Errorf("Comps.NormalV was %+v, expected %+v.", comps.NormalV, expectedNormalv)
	}
}
