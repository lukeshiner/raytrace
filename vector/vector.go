package vector

import "math"

const EPSLION = 0.00001

func main() {
}

func Equal(a float64, b float64) bool {
	return math.Abs(a-b) < EPSLION
}

type Vector struct {
	x float64
	y float64
	z float64
	w float64
}

func (v *Vector) IsPoint() bool {
	return v.w == 1.0
}

func (v *Vector) IsVector() bool {
	return v.w == 0.0
}

func (v *Vector) Equal(other *Vector) bool {
	return EqualVectors(v, other)
}

func (v *Vector) Add(other *Vector) Vector {
	return Vector{v.x + other.x, v.y + other.y, v.z + other.z, v.w + other.w}
}

func (v *Vector) Sub(other *Vector) Vector {
	return Vector{v.x - other.x, v.y - other.y, v.z - other.z, v.w - other.w}
}

func (v *Vector) Negate() Vector {
	return Vector{-v.x, -v.y, -v.z, -v.w}
}

func (v *Vector) ScalarMult(scalar float64) Vector {
	return Vector{v.x * scalar, v.y * scalar, v.z * scalar, v.w * scalar}
}

func (v *Vector) ScalarDiv(scalar float64) Vector {
	return Vector{v.x / scalar, v.y / scalar, v.z / scalar, v.w / scalar}
}

func (v *Vector) Mag() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z + v.w*v.w)
}

func (v *Vector) Normalize() Vector {
	mag := v.Mag()
	return Vector{v.x / mag, v.y / mag, v.z / mag, v.w / mag}
}

func EqualVectors(vectorA *Vector, vectorB *Vector) bool {
	if Equal(vectorA.x, vectorB.x) == false {
		return false
	}
	if Equal(vectorA.y, vectorB.y) == false {
		return false
	}
	if Equal(vectorA.z, vectorB.z) == false {
		return false
	}
	if Equal(vectorA.w, vectorB.w) == false {
		return false
	}
	return true
}

func DotProduct(a *Vector, b *Vector) float64 {
	return (a.x * b.x) + (a.y * b.y) + (a.z * b.z) + (a.w * b.w)
}

func CrossProduct(a *Vector, b *Vector) Vector {
	return Vector{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
		0,
	}
}

func MakePoint(x float64, y float64, z float64) Vector {
	return Vector{x, y, z, 1.0}
}

func MakeVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z, 0.0}
}
