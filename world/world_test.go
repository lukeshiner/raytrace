package world

import (
	"testing"

	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/comparison"
	"github.com/lukeshiner/raytrace/light"
	"github.com/lukeshiner/raytrace/material"
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
	var tests = []struct {
		ray                                          ray.Ray
		intersection                                 ray.Intersection
		expectedPoint, expectedEyeV, expectedNormalV vector.Vector
		expectedInside                               bool
	}{
		{
			// The hit, when an intersection occurs on the outside.
			ray:             ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			intersection:    ray.NewIntersection(4, object.NewSphere()),
			expectedPoint:   vector.NewPoint(0, 0, -1),
			expectedEyeV:    vector.NewVector(0, 0, -1),
			expectedNormalV: vector.NewVector(0, 0, -1),
			expectedInside:  false,
		},
		{
			// The hit, when an intersection occurs on the inside.
			ray:             ray.New(vector.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1)),
			intersection:    ray.NewIntersection(1, object.NewSphere()),
			expectedPoint:   vector.NewPoint(0, 0, 1),
			expectedEyeV:    vector.NewVector(0, 0, -1),
			expectedNormalV: vector.NewVector(0, 0, -1),
			expectedInside:  true,
		},
	}
	for _, test := range tests {
		comps := PrepareComputations(test.intersection, test.ray)
		if comps.T != test.intersection.T {
			t.Errorf("Comps.T was %v, expected %v.", comps.T, test.intersection.T)
		}
		if comps.Object.ID() != test.intersection.Object.ID() {
			t.Errorf("Comps.Object was %+v, expected %+v.", comps.Object, test.intersection.Object)
		}
		if vector.Equal(comps.Point, test.expectedPoint) != true {
			t.Errorf("Comps.Point was %+v, expected %+v.", comps.Point, test.expectedPoint)
		}
		if vector.Equal(comps.EyeV, test.expectedEyeV) != true {
			t.Errorf("Comps.EyeV was %+v, expected %+v.", comps.EyeV, test.expectedEyeV)
		}
		if vector.Equal(comps.EyeV, test.expectedNormalV) != true {
			t.Errorf("Comps.NormalV was %+v, expected %+v.", comps.NormalV, test.expectedNormalV)
		}
		if comps.Inside != test.expectedInside {
			t.Errorf("Comps.Inside was %v, expected %v.", comps.Inside, test.expectedInside)
		}
	}
}

func TestShadeHit(t *testing.T) {
	var tests = []struct {
		world       World
		light       light.Light
		ray         ray.Ray
		objectIndex int
		t           float64
		expected    colour.Colour
	}{
		{
			// Shading an intersection.
			world:       Default(),
			light:       nil,
			ray:         ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			objectIndex: 0,
			t:           4,
			expected:    colour.New(0.38066, 0.47583, 0.2855),
		},
		{
			// Shading an intersection from inside.
			world:       Default(),
			light:       light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0.25, 0)),
			ray:         ray.New(vector.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1)),
			objectIndex: 1,
			t:           0.5,
			expected:    colour.New(0.90498, 0.90498, 0.90498),
		},
	}
	for _, test := range tests {
		if test.light != nil {
			test.world.Lights[0] = test.light
		}
		i := ray.NewIntersection(test.t, test.world.Objects[test.objectIndex])
		comps := PrepareComputations(i, test.ray)
		result := ShadeHit(test.world, comps)
		if result.Equal(test.expected) != true {
			t.Errorf("Shade hit returned %v, expected %v.", result, test.expected)
		}
	}
}

func TestShadeHitWithShadow(t *testing.T) {
	w := Default()
	w.Lights = []light.Light{light.NewPoint(colour.New(1, 1, 1), vector.NewPoint(0, 0, -10))}
	w.Objects = []object.Object{object.NewSphere(), object.NewSphere()}
	w.Objects[1].SetTransform(matrix.TranslationMatrix(0, 0, 10))
	r := ray.New(vector.NewPoint(0, 0, 5), vector.NewVector(0, 0, 1))
	i := ray.Intersect(w.Objects[1], r).Intersections[0]
	comps := PrepareComputations(i, r)
	expected := colour.New(0.1, 0.1, 0.1)
	result := ShadeHit(w, comps)
	if !result.Equal(expected) {
		t.Errorf("ShaderHit with shadow returned %v, expected %v.", result, expected)
	}
}

func TestColourAt(t *testing.T) {
	var tests = []struct {
		ray      ray.Ray
		expected colour.Colour
	}{
		{
			// The colour when a ray misses.
			ray:      ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 1, 0)),
			expected: colour.New(0, 0, 0),
		},
		{
			// The colour with an intersection hits.
			ray:      ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			expected: colour.New(0.38066, 0.47583, 0.2855),
		},
	}
	for _, test := range tests {
		w := Default()
		result := ColourAt(w, test.ray)
		if result.Equal(test.expected) != true {
			t.Errorf("ColourAt returned %v, expected %v.", result, test.expected)
		}
	}
}

func TestColourAtWithIntersectionBehindRay(t *testing.T) {
	w := Default()
	m := material.New()
	m.Ambient = 1
	w.Objects[0].SetMaterial(m)
	m2 := material.New()
	m2.Ambient = 1
	w.Objects[1].SetMaterial(m2)
	r := ray.New(vector.NewPoint(0, 0, 0.75), vector.NewVector(0, 0, -1))
	expected := w.Objects[1].Material().Colour
	result := ColourAt(w, r)
	if result.Equal(expected) != true {
		t.Errorf("ColourAt returned %v, expected %v.", result, expected)
	}
}

func TestInShadow(t *testing.T) {
	var tests = []struct {
		world    World
		point    vector.Vector
		expected bool
	}{
		{
			// There is no shadow when nothing is collinear with point and light.
			world:    Default(),
			point:    vector.NewPoint(0, 10, 0),
			expected: false,
		},
		{
			// The shadow when an object is between the point and the light.
			world:    Default(),
			point:    vector.NewPoint(10, -10, 10),
			expected: true,
		},
		{
			// The shadow when an object is behind the light.
			world:    Default(),
			point:    vector.NewPoint(-20, 20, -20),
			expected: false,
		},
		{
			// The shadow when an object is behind the light.
			world:    Default(),
			point:    vector.NewPoint(-2, 2, -2),
			expected: false,
		},
	}
	for _, test := range tests {
		result := IsShadowed(test.world, test.point, test.world.Lights[0])
		if result != test.expected {
			t.Errorf(
				"InShadow for point %v was %v, expected %v.", test.point, result, test.expected,
			)
		}
	}
}

func TestOverPoint(t *testing.T) {
	r := ray.New(vector.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	s := object.NewSphere()
	s.SetTransform(matrix.TranslationMatrix(0, 0, 1))
	i := ray.NewIntersection(5, s)
	comps := PrepareComputations(i, r)
	result := comps.OverPoint.Z
	if result >= -comparison.EPSLION/2 {
		t.Errorf("Over Point %v too low.", result)
	}
	if result > comps.Point.Z {
		t.Errorf("Over Point %v too high.", result)
	}
}
