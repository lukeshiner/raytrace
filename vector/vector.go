package vector

import "math"

const EPSLION = 0.00001

func main() {
}

func Equal(a float64, b float64) bool {
	return math.Abs(a-b) < EPSLION
}

type Vector struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (v *Vector) IsPoint() bool {
	return v.W == 1.0
}

func (v *Vector) IsVector() bool {
	return v.W == 0.0
}

func (v *Vector) Equal(other *Vector) bool {
	return EqualVectors(v, other)
}

func (v *Vector) Add(other *Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y, v.Z + other.Z, v.W + other.W}
}

func (v *Vector) Sub(other *Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y, v.Z - other.Z, v.W - other.W}
}

func (v *Vector) Negate() Vector {
	return Vector{-v.X, -v.Y, -v.Z, -v.W}
}

func (v *Vector) ScalarMult(scalar float64) Vector {
	return Vector{v.X * scalar, v.Y * scalar, v.Z * scalar, v.W * scalar}
}

func (v *Vector) ScalarDiv(scalar float64) Vector {
	return Vector{v.X / scalar, v.Y / scalar, v.Z / scalar, v.W / scalar}
}

func (v *Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

func (v *Vector) Normalize() Vector {
	mag := v.Mag()
	return Vector{v.X / mag, v.Y / mag, v.Z / mag, v.W / mag}
}

func EqualVectors(vectorA *Vector, vectorB *Vector) bool {
	if Equal(vectorA.X, vectorB.X) == false {
		return false
	}
	if Equal(vectorA.Y, vectorB.Y) == false {
		return false
	}
	if Equal(vectorA.Z, vectorB.Z) == false {
		return false
	}
	if Equal(vectorA.W, vectorB.W) == false {
		return false
	}
	return true
}

func DotProduct(a *Vector, b *Vector) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z) + (a.W * b.W)
}

func CrossProduct(a *Vector, b *Vector) Vector {
	return Vector{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
		0,
	}
}

func MakePoint(x float64, y float64, z float64) Vector {
	return Vector{x, y, z, 1.0}
}

func MakeVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z, 0.0}
}
