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

// EqualMatricies compares two instances of Matrix for equality.
func EqualMatricies(a, b Matrix) bool {
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
