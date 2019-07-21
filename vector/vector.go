package vector

import (
	"math"

	"github.com/lukeshiner/raytrace/comparison"
)

// Vector is the struct used for points and vectors.
type Vector struct {
	X float64
	Y float64
	Z float64
	W float64
}

// IsPoint returns true if the Vector is a point, otherwise false.
func (v *Vector) IsPoint() bool {
	return v.W == 1.0
}

// IsVector returns true if the Vector is a vector, otherwise false.
func (v *Vector) IsVector() bool {
	return v.W == 0.0
}

// Equal returns true if all attributes of two vectors are within EPSILON of
// each other.
func (v *Vector) Equal(other *Vector) bool {
	return EqualVectors(v, other)
}

// Add returns the Vector created by adding other to v.
func (v *Vector) Add(other *Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y, v.Z + other.Z, v.W + other.W}
}

// Sub returns the Vector created by subtracting other from v.
func (v *Vector) Sub(other *Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y, v.Z - other.Z, v.W - other.W}
}

// Negate retruns the negated vector of v.
func (v *Vector) Negate() Vector {
	return Vector{-v.X, -v.Y, -v.Z, -v.W}
}

// ScalarMult returns the Vector created by performing a scalar multiplication
// of scalar on v.
func (v *Vector) ScalarMult(scalar float64) Vector {
	return Vector{v.X * scalar, v.Y * scalar, v.Z * scalar, v.W * scalar}
}

// ScalarDiv returns the Vector created by performing a scalar division of v by
// scalar.
func (v *Vector) ScalarDiv(scalar float64) Vector {
	return Vector{v.X / scalar, v.Y / scalar, v.Z / scalar, v.W / scalar}
}

// Mag returns the magnitude of v.
func (v *Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Normalize returns the nomalized vector of v.
func (v *Vector) Normalize() Vector {
	mag := v.Mag()
	return Vector{v.X / mag, v.Y / mag, v.Z / mag, v.W / mag}
}

// Tuple returns the values of the vector as a []float64.
func (v *Vector) Tuple() []float64 {
	return []float64{v.X, v.Y, v.Z, v.W}
}

// EqualVectors returns true if all attributes of Vectors a and b are within
// EPSILON of eachother.
func EqualVectors(a *Vector, b *Vector) bool {
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

// DotProduct returns the dot product of Vectors a and b.
func DotProduct(a *Vector, b *Vector) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z) + (a.W * b.W)
}

// CrossProduct returns the cross product of Vectors a and b.
func CrossProduct(a *Vector, b *Vector) Vector {
	return Vector{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
		0,
	}
}

// MakePoint returns a point with x, y and z attributes.
func MakePoint(x float64, y float64, z float64) Vector {
	return Vector{x, y, z, 1.0}
}

// MakeVector returns a vector with x, y and z attributes.
func MakeVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z, 0.0}
}

// FromSlice returns a vector from a four element slice
func FromSlice(t []float64) Vector {
	return Vector{X: t[0], Y: t[1], Z: t[2], W: t[3]}
}
