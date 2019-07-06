package matrix

import "github.com/lukeshiner/raytrace/comparison"

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
