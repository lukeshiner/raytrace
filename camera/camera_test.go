package camera

import (
	"math"
	"testing"

	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/vector"
	"github.com/lukeshiner/raytrace/world"
)

func TestVeiwTransform(t *testing.T) {
	var tests = []struct {
		from, to, up vector.Vector
		expected     matrix.Matrix
	}{
		{
			//The transformation matrix for the default orientation.
			from:     vector.NewPoint(0, 0, 0),
			to:       vector.NewPoint(0, 0, -1),
			up:       vector.NewVector(0, 1, 0),
			expected: matrix.IdentityMatrix(4),
		},
		{
			// A view transformation matrix looking in positive z direction.
			from:     vector.NewPoint(0, 0, 0),
			to:       vector.NewPoint(0, 0, 1),
			up:       vector.NewVector(0, 1, 0),
			expected: matrix.ScalingMatrix(-1, 1, -1),
		},
		{
			// The view transformation moves the world.
			from:     vector.NewPoint(0, 0, 8),
			to:       vector.NewPoint(0, 0, 0),
			up:       vector.NewVector(0, 1, 0),
			expected: matrix.TranslationMatrix(0, 0, -8),
		},
		{
			//An arbitrary view transformation.
			from: vector.NewPoint(1, 3, 2),
			to:   vector.NewPoint(4, -2, 8),
			up:   vector.NewVector(1, 1, 0),
			expected: matrix.New(
				[]float64{-0.50709, 0.50709, 0.67612, -2.36643},
				[]float64{0.76772, 0.60609, 0.12122, -2.82843},
				[]float64{-0.35857, 0.59761, -0.71714, 0.00000},
				[]float64{0.00000, 0.00000, 0.00000, 1.00000},
			),
		},
	}
	for _, test := range tests {
		result := ViewTransform(test.from, test.to, test.up)
		if matrix.Equal(result, test.expected) != true {
			t.Errorf(
				"View transform(%v, %v, %v) was %+v, expected %+v.",
				test.from, test.to, test.up, result, test.expected,
			)
		}
	}
}

func TestNew(t *testing.T) {
	var tests = []struct {
		HSize, VSize int
		FOV          float64
		Transform    matrix.Matrix
	}{
		{
			HSize:     160,
			VSize:     120,
			FOV:       math.Pi / 2,
			Transform: matrix.IdentityMatrix(4),
		},
	}
	for _, test := range tests {
		c := New(test.HSize, test.VSize, test.FOV)
		if c.HSize != test.HSize || c.VSize != test.VSize || c.FOV != test.FOV ||
			matrix.Equal(c.Transform, test.Transform) != true {
			t.Errorf("Camera (%v, %v, %v) produced %+v.", test.HSize, test.VSize, test.FOV, c)
		}
	}
}

func TestPixelSize(t *testing.T) {
	var tests = []struct {
		hSize, vSize  int
		FOV, expected float64
	}{
		{
			// The pixel size for a horizontal canvas.
			hSize:    200,
			vSize:    125,
			FOV:      math.Pi / 2,
			expected: 0.01,
		},
		{
			// The pixel size for a vertical canvas.
			hSize:    125,
			vSize:    200,
			FOV:      math.Pi / 2,
			expected: 0.01,
		},
	}
	for _, test := range tests {
		c := New(test.hSize, test.vSize, test.FOV)
		result := c.PixelSize
		if result != test.expected {
			t.Errorf(
				"Camera with size %dx%d and FOV %v has a pixel size of %v, expected %v.",
				test.hSize, test.vSize, test.FOV, result, test.expected,
			)
		}
	}
}

func TestRayForPixel(t *testing.T) {
	var tests = []struct {
		camera                            Camera
		transform                         matrix.Matrix
		x, y                              int
		expectedOrigin, expectedDirection vector.Vector
	}{
		{
			// Constructing a ray through the center of the canvas.
			camera:            New(201, 101, math.Pi/2),
			transform:         matrix.IdentityMatrix(4),
			x:                 100,
			y:                 50,
			expectedOrigin:    vector.NewPoint(0, 0, 0),
			expectedDirection: vector.NewVector(0, 0, -1),
		},
		{
			// Constructing a ray through a corner of the canvas.
			camera:            New(201, 101, math.Pi/2),
			transform:         matrix.IdentityMatrix(4),
			x:                 0,
			y:                 0,
			expectedOrigin:    vector.NewPoint(0, 0, 0),
			expectedDirection: vector.NewVector(0.66519, 0.33259, -0.66851),
		},
		{
			// Constructing a ray when the camera is transformed.
			camera: New(201, 101, math.Pi/2),
			transform: matrix.Multiply(
				matrix.RotationYMatrix(math.Pi/4), matrix.TranslationMatrix(0, -2, 5),
			),
			x:                 100,
			y:                 50,
			expectedOrigin:    vector.NewPoint(0, 2, -5),
			expectedDirection: vector.NewVector(math.Sqrt(2)/2, 0, -(math.Sqrt(2) / 2)),
		},
	}
	for _, test := range tests {
		test.camera.SetTransform(test.transform)
		r := RayForPixel(test.camera, test.x, test.y)
		if !vector.Equal(r.Origin, test.expectedOrigin) ||
			!vector.Equal(r.Direction, test.expectedDirection) {
			t.Errorf(
				"RayForPixel(%+v, %v, %v) returned %+v, expected ray(%+v, %+v)).",
				test.camera, test.x, test.y, r, test.expectedOrigin, test.expectedDirection,
			)
		}
	}
}

func TestRender(t *testing.T) {
	w := world.Default()
	c := New(11, 11, math.Pi/2)
	from := vector.NewPoint(0, 0, -5)
	to := vector.NewPoint(0, 0, 0)
	up := vector.NewVector(0, 1, 0)
	c.SetTransform(ViewTransform(from, to, up))
	image := Render(c, w)
	expected := colour.New(0.38066, 0.47583, 0.2855)
	result := image.Pixel(5, 5)
	if !result.Equal(expected) {
		t.Errorf("Render returned %+v, expected %+v.", result, expected)
	}
}
