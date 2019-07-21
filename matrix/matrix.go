package matrix

import (
	"errors"
	"math"

	"github.com/lukeshiner/raytrace/comparison"
)

// Matrix is a type for matricies.
type Matrix struct {
	Width, Height int
	Cells         [][]float64
}

// New creates a new Matrix.
func New(rows ...[]float64) Matrix {
	width := len(rows[0])
	height := len(rows)
	matrix := Matrix{width, height, rows}
	return matrix
}

// Get returns the value at postition (x,y) in the Matrix.
func (m *Matrix) Get(x, y int) float64 {
	return m.Cells[x][y]
}

// Row returns the values in row[index] of the matrix.
func (m *Matrix) Row(index int) []float64 {
	return m.Cells[index]
}

// Column returns the values in column[index] of the matrix.
func (m *Matrix) Column(index int) []float64 {
	var column []float64
	for y := 0; y < m.Height; y++ {
		column = append(column, m.Cells[y][index])
	}
	return column
}

// Transpose returns the matrix created by transposing m.
func (m Matrix) Transpose() Matrix {
	var cells [][]float64
	for i := 0; i < m.Height; i++ {
		cells = append(cells, m.Column(i))
	}
	return New(cells...)
}

// Submatrix returns a Submatrix of m with row rowNumber and column columnNumber removed.
func (m Matrix) Submatrix(rowNumber, columnNumber int) Matrix {
	var cells [][]float64
	var row []float64
	for x := 0; x < m.Width; x++ {
		if x != rowNumber {
			for y := 0; y < m.Height; y++ {
				if y != columnNumber {
					row = append(row, m.Get(x, y))
				}
			}
			cells = append(cells, make([]float64, len(row)))
			copy(cells[len(cells)-1], row)
			row = row[:0]
		}
	}
	return New(cells...)
}

func (m Matrix) determinant() float64 {
	if m.Width == 2 && m.Height == 2 {
		return (m.Get(0, 0) * m.Get(1, 1)) - (m.Get(0, 1) * m.Get(1, 0))
	}
	const rowNumber = 0
	var value float64
	row := m.Row(rowNumber)
	for i := 0; i < m.Width; i++ {
		cofactor := m.cofactor(rowNumber, i)
		value += cofactor * row[i]
	}
	return value
}

func (m Matrix) cofactor(rowNumber, columnNumber int) float64 {
	minor := m.minor(rowNumber, columnNumber)
	if (rowNumber+columnNumber)%2 != 0 {
		minor = -minor
	}
	return minor
}

func (m Matrix) minor(rowNumber, columnNumber int) float64 {
	submatrix := m.Submatrix(rowNumber, columnNumber)
	return submatrix.determinant()
}

// Invertable returns true if the matrix is invertable, otherwise returns false.
func (m Matrix) Invertable() bool {
	return m.determinant() != 0
}

// Invert returns the inverse of the matrix
func (m Matrix) Invert() (Matrix, error) {
	if m.Invertable() == false {
		err := errors.New("matrix is not invertable")
		return Matrix{}, err
	}
	var cells [][]float64
	var row []float64
	var cofactor, value float64
	determinant := m.determinant()
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			cofactor = m.cofactor(x, y)
			value = cofactor / determinant
			row = append(row, value)
		}
		cells = append(cells, make([]float64, len(row)))
		copy(cells[len(cells)-1], row)
		row = row[:0]
	}
	matrix := New(cells...)
	return matrix.Transpose(), nil
}

// Equal compares two instances of Matrix for equality.
func Equal(a, b Matrix) bool {
	if a.Width != b.Width || a.Height != b.Height {
		return false
	}
	for x := 0; x < a.Height; x++ {
		for y := 0; y < a.Height; y++ {
			if comparison.EpsilonEqual(a.Get(x, y), b.Get(x, y)) != true {
				return false
			}
		}
	}
	return true
}

func multiplyCell(a, b Matrix, x, y int) float64 {
	var value float64
	row := a.Row(x)
	column := b.Column(y)
	for i := 0; i < a.Width; i++ {
		value += row[i] * column[i]
	}
	return value
}

// Multiply returns the product of two matricies.
func Multiply(a, b Matrix) Matrix {
	var values [][]float64
	var row []float64
	for x := 0; x < a.Height; x++ {
		for y := 0; y < a.Width; y++ {
			row = append(row, multiplyCell(a, b, x, y))
		}
		values = append(values, make([]float64, len(row)))
		copy(values[x], row)
		row = row[:0]
	}
	m := Matrix{Width: a.Width, Height: a.Height, Cells: values}
	return m
}

func multiplyTupleCell(row, tuple []float64) float64 {
	var value float64
	for i := 0; i < len(row); i++ {
		value += row[i] * tuple[i]
	}
	return value
}

// MultiplyTuple multiplies a Matrix by a tuple.
func MultiplyTuple(m Matrix, tuple []float64) []float64 {
	var value []float64
	for x := 0; x < m.Height; x++ {
		value = append(value, multiplyTupleCell(m.Row(x), tuple))
	}
	return value
}

// IdentityMatrix returns the identity matrix of size size.
func IdentityMatrix(size int) Matrix {
	var values [][]float64
	var row []float64
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if x == y {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		values = append(values, make([]float64, len(row)))
		copy(values[x], row)
		row = row[:0]
	}
	return New(values...)
}

// Translation returns a traslation transform matrix.
func Translation(x, y, z float64) Matrix {
	m := IdentityMatrix(4)
	m.Cells[0][3] = x
	m.Cells[1][3] = y
	m.Cells[2][3] = z
	return m
}

// Scaling returns a scaling transform matrix.
func Scaling(x, y, z float64) Matrix {
	m := IdentityMatrix(4)
	m.Cells[0][0] = x
	m.Cells[1][1] = y
	m.Cells[2][2] = z
	return m
}

// RotationX returns a rotation matrix for the x axis
func RotationX(radians float64) Matrix {
	m := IdentityMatrix(4)
	m.Cells[1][1] = math.Cos(radians)
	m.Cells[1][2] = -math.Sin(radians)
	m.Cells[2][1] = math.Sin(radians)
	m.Cells[2][2] = math.Cos(radians)
	return m
}

// RotationY returns a rotation matrix for the x axis
func RotationY(radians float64) Matrix {
	m := IdentityMatrix(4)
	m.Cells[0][0] = math.Cos(radians)
	m.Cells[0][2] = math.Sin(radians)
	m.Cells[2][0] = -math.Sin(radians)
	m.Cells[2][2] = math.Cos(radians)
	return m
}

// RotationZ returns a rotation matrix for the x axis
func RotationZ(radians float64) Matrix {
	m := IdentityMatrix(4)
	m.Cells[0][0] = math.Cos(radians)
	m.Cells[0][1] = -math.Sin(radians)
	m.Cells[1][0] = math.Sin(radians)
	m.Cells[1][1] = math.Cos(radians)
	return m
}
