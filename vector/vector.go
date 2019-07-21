package transformation

import (
	"math"

	"github.com/lukeshiner/raytrace/comparison"
	"github.com/lukeshiner/raytrace/matrix"
)

// Vector is a translatable vector
type Vector struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewPoint returns a point
func NewPoint(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z, W: 1}
}

// NewVector returns a vector
func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z, W: 0}
}

// FromSlice returns a Vector from a slice of X, Y, Z, W
func FromSlice(s []float64) Vector {
	return Vector{X: s[0], Y: s[1], Z: s[2], W: s[3]}
}

// IsPoint returns true if the Vector is a point, otherwise false.
func (v *Vector) IsPoint() bool {
	return v.W == 1.0
}

// IsVector returns true if the Vector is a vector, otherwise false.
func (v *Vector) IsVector() bool {
	return v.W == 0.0
}

// AsSlice returns the values of the vector as a []float64.
func (v *Vector) AsSlice() []float64 {
	return []float64{v.X, v.Y, v.Z, v.W}
}

// Negate retruns the negated vector of v.
func (v *Vector) Negate() Vector {
	return Vector{-v.X, -v.Y, -v.Z, -v.W}
}

// ScalarMultiply returns the Vector created by performing a scalar multiplication
// of scalar on v.
func (v *Vector) ScalarMultiply(scalar float64) Vector {
	return Vector{v.X * scalar, v.Y * scalar, v.Z * scalar, v.W * scalar}
}

// ScalarDivide returns the Vector created by performing a scalar division of v by
// scalar.
func (v *Vector) ScalarDivide(scalar float64) Vector {
	return Vector{v.X / scalar, v.Y / scalar, v.Z / scalar, v.W / scalar}
}

// Magnitude returns the magnitude of v.
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Normalize returns the nomalized vector of v.
func (v *Vector) Normalize() Vector {
	mag := v.Magnitude()
	return Vector{v.X / mag, v.Y / mag, v.Z / mag, v.W / mag}
}

// Translate returns a translated Vector
func (v Vector) Translate(x, y, z float64) Vector {
	m := matrix.TranslationMatrix(x, y, z)
	return MultiplyMatrixByVector(m, v)
}

// InverseTranslate returns a translated Vector
func (v Vector) InverseTranslate(x, y, z float64) Vector {
	m, _ := matrix.TranslationMatrix(x, y, z).Invert()
	return MultiplyMatrixByVector(m, v)
}

// Scale returns a sacled Vector
func (v Vector) Scale(x, y, z float64) Vector {
	m := matrix.ScalingMatrix(x, y, z)
	return MultiplyMatrixByVector(m, v)
}

// InverseScale returns a translated Vector
func (v Vector) InverseScale(x, y, z float64) Vector {
	m, _ := matrix.ScalingMatrix(x, y, z).Invert()
	return MultiplyMatrixByVector(m, v)
}

// RotateX returns a Vector rotated around the X axis
func (v Vector) RotateX(radians float64) Vector {
	m := matrix.RotationXMatrix(radians)
	return MultiplyMatrixByVector(m, v)
}

// InverseRotateX returns a Vector rotated around the X axis
func (v Vector) InverseRotateX(radians float64) Vector {
	m, _ := matrix.RotationXMatrix(radians).Invert()
	return MultiplyMatrixByVector(m, v)
}

// RotateY returns a Vector rotated around the X axis
func (v Vector) RotateY(radians float64) Vector {
	m := matrix.RotationYMatrix(radians)
	return MultiplyMatrixByVector(m, v)
}

// InverseRotateY returns a Vector rotated around the X axis
func (v Vector) InverseRotateY(radians float64) Vector {
	m, _ := matrix.RotationYMatrix(radians).Invert()
	return MultiplyMatrixByVector(m, v)
}

// RotateZ returns a Vector rotated around the X axis
func (v Vector) RotateZ(radians float64) Vector {
	m := matrix.RotationZMatrix(radians)
	return MultiplyMatrixByVector(m, v)
}

// InverseRotateZ returns a Vector rotated around the X axis
func (v Vector) InverseRotateZ(radians float64) Vector {
	m, _ := matrix.RotationZMatrix(radians).Invert()
	return MultiplyMatrixByVector(m, v)
}

// Shear returns a sheared Vector
func (v Vector) Shear(Xy, Xz, Yx, Yz, Zx, Zy float64) Vector {
	m := matrix.ShearingMatrix(Xy, Xz, Yx, Yz, Zx, Zy)
	return MultiplyMatrixByVector(m, v)
}

// Equal returns True if two Vectors are equal, otherwise false
func Equal(a, b Vector) bool {
	if comparison.EpsilonEqual(a.X, b.X) == false {
		return false
	}
	if comparison.EpsilonEqual(a.Y, b.Y) == false {
		return false
	}
	if comparison.EpsilonEqual(a.Z, b.Z) == false {
		return false
	}
	if comparison.EpsilonEqual(a.W, b.W) == false {
		return false
	}
	return true
}

// Add returns the Vector created by adding other to v.
func Add(a, b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

// Subtract returns the Vector created by subtracting other from v.
func Subtract(a, b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

// DotProduct returns the dot product of Vectors a and b.
func DotProduct(a, b Vector) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z) + (a.W * b.W)
}

// CrossProduct returns the cross product of Vectors a and b.
func CrossProduct(a, b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
		W: 0,
	}
}

// MultiplyMatrixByVector multiplies a matrix by a point or vector
func MultiplyMatrixByVector(m matrix.Matrix, v Vector) Vector {
	tuple := matrix.MultiplyTuple(m, v.AsSlice())
	return FromSlice(tuple)
}
