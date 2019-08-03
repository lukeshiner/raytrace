package camera

import (
	"math"

	"github.com/lukeshiner/raytrace/canvas"
	"github.com/lukeshiner/raytrace/colour"
	"github.com/lukeshiner/raytrace/matrix"
	"github.com/lukeshiner/raytrace/ray"
	"github.com/lukeshiner/raytrace/vector"
	"github.com/lukeshiner/raytrace/world"
)

// Camera holds camera data.
type Camera struct {
	HSize, VSize                          int
	FOV, HalfWidth, HalfHeight, PixelSize float64
	Transform                             matrix.Matrix
}

// SetTransform sets the tranform matrix for the camera.
func (c *Camera) SetTransform(t matrix.Matrix) {
	c.Transform = t
}

// New returns a new camera instance.
func New(hSize, vSize int, FOV float64) Camera {
	var halfWidth, halfHeight float64
	halfView := math.Tan(FOV / 2)
	aspect := float64(hSize) / float64(vSize)
	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}
	pixelSize := (halfWidth * 2) / float64(hSize)
	return Camera{
		HSize: hSize, VSize: vSize, FOV: FOV, Transform: matrix.IdentityMatrix(4),
		HalfHeight: halfHeight, HalfWidth: halfWidth, PixelSize: pixelSize,
	}
}

// ViewTransform returns the transformation matrix for a camera position.
func ViewTransform(from, to, up vector.Vector) matrix.Matrix {
	forward := vector.Subtract(to, from)
	forward = forward.Normalize()
	normalUp := up.Normalize()
	left := vector.CrossProduct(forward, normalUp)
	trueUp := vector.CrossProduct(left, forward)
	orientation := matrix.New(
		[]float64{left.X, left.Y, left.Z, 0},
		[]float64{trueUp.X, trueUp.Y, trueUp.Z, 0},
		[]float64{-forward.X, -forward.Y, -forward.Z, 0},
		[]float64{0, 0, 0, 1},
	)
	return matrix.Multiply(orientation, matrix.TranslationMatrix(-from.X, -from.Y, -from.Z))
}

// RayForPixel returns the ray from camera to pixel (pX, pY).
func RayForPixel(c Camera, pX, pY int) ray.Ray {
	xOffset := (float64(pX) + 0.5) * c.PixelSize
	yOffset := (float64(pY) + 0.5) * c.PixelSize
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	transform, _ := c.Transform.Invert()
	pixel := vector.MultiplyMatrixByVector(
		transform, vector.NewPoint(worldX, worldY, -1))
	origin := vector.MultiplyMatrixByVector(transform, vector.NewPoint(0, 0, 0))
	direction := vector.Subtract(pixel, origin)
	return ray.New(origin, direction.Normalize())
}

// Render returns a rendered canvas.Canvas for a camera and a world.
func Render(c Camera, w world.World) canvas.Canvas {
	var r ray.Ray
	var col colour.Colour
	img := canvas.New(c.HSize, c.VSize)
	for y := 0; y < c.VSize; y++ {
		for x := 0; x < c.HSize; x++ {
			r = RayForPixel(c, x, y)
			col = world.ColourAt(w, r)
			img.WritePixel(x, y, col)
		}
	}
	return img
}
