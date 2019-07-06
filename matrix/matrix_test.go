package matrix

import (
	"testing"
)

func TestMatrix(t *testing.T) {

	type matrixTestResult struct {
		x, y  int
		value float64
	}

	type matrixTest struct {
		matrix   Matrix
		expected []matrixTestResult
	}

	var tests = []matrixTest{
		matrixTest{
			matrix: New(
				[]float64{1, 2, 1, 4},
				[]float64{5.5, 6.5, 7.5, 8.5},
				[]float64{9, 10, 11, 12},
				[]float64{13.5, 14.5, 15.5, 16.5},
			),
			expected: []matrixTestResult{
				{x: 0, y: 0, value: 1},
				{x: 0, y: 3, value: 4},
				{x: 1, y: 0, value: 5.5},
				{x: 2, y: 2, value: 11},
				{x: 3, y: 0, value: 13.5},
				{x: 3, y: 2, value: 15.5},
			},
		},
		matrixTest{
			matrix: New(
				[]float64{-3, 5},
				[]float64{1, -2},
			),
			expected: []matrixTestResult{
				{x: 0, y: 0, value: -3},
				{x: 0, y: 1, value: 5},
				{x: 1, y: 0, value: 1},
				{x: 1, y: 1, value: -2},
			},
		},
		matrixTest{
			matrix: New(
				[]float64{-3, 5, 0},
				[]float64{1, -2, -7},
				[]float64{0, 1, 1},
			),
			expected: []matrixTestResult{
				{x: 0, y: 0, value: -3},
				{x: 1, y: 1, value: -2},
				{x: 2, y: 2, value: 1},
			},
		},
	}
	for _, test := range tests {
		for _, cell := range test.expected {
			if test.matrix.Get(cell.x, cell.y) != cell.value {
				t.Error("Matrix not initialized correclty.")
			}
		}
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		a, b     Matrix
		expected bool
	}{
		{
			a: New(
				[]float64{1, 2, 3, 4},
				[]float64{5, 6, 7, 8},
				[]float64{9, 8, 7, 6},
				[]float64{5, 4, 3, 2},
			),
			b: New(
				[]float64{1, 2, 3, 4},
				[]float64{5, 6, 7, 8},
				[]float64{9, 8, 7, 6},
				[]float64{5, 4, 3, 2},
			),
			expected: true,
		},
		{
			a: New(
				[]float64{1, 2, 3, 4},
				[]float64{5, 6, 7, 8},
				[]float64{9, 8, 7, 6},
				[]float64{5, 4, 3, 2},
			),
			b: New(
				[]float64{1, 2, 3, 4},
				[]float64{5, 6, 7, 8},
				[]float64{9, 8, 7, 6},
			),
			expected: false,
		},
		{
			a: New(
				[]float64{1, 2, 3, 4},
				[]float64{5, 6, 7, 8},
				[]float64{9, 8, 7, 6},
				[]float64{5, 4, 3, 2},
			),
			b: New(
				[]float64{2, 3, 4, 5},
				[]float64{6, 7, 8, 9},
				[]float64{8, 7, 6, 5},
				[]float64{4, 3, 2, 1},
			),
			expected: false,
		},
	}
	for _, test := range tests {
		value := Equal(test.a, test.b)
		if value != test.expected {
			t.Error("Equal Matricies failed.")
		}
	}
}

func TestRowMethod(t *testing.T) {
	var tests = []struct {
		matrix   Matrix
		row      int
		expected []float64
	}{
		{
			matrix: New(
				[]float64{2, 3, 4, 5},
				[]float64{6, 7, 8, 9},
				[]float64{8, 7, 6, 5},
				[]float64{4, 3, 2, 1},
			),
			row:      0,
			expected: []float64{2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		value := test.matrix.Row(test.row)
		for i := 0; i < len(value); i++ {
			if value[i] != test.expected[i] {
				t.Errorf(
					"Matrix %+v row %d got %+v expected %+v.",
					test.matrix,
					test.row,
					value,
					test.expected)
			}
		}
	}
}

func TestColumnMethod(t *testing.T) {
	var tests = []struct {
		matrix   Matrix
		column   int
		expected []float64
	}{
		{
			matrix: New(
				[]float64{2, 3, 4, 5},
				[]float64{6, 7, 8, 9},
				[]float64{8, 7, 6, 5},
				[]float64{4, 3, 2, 1},
			),
			column:   0,
			expected: []float64{2, 6, 8, 4},
		},
		{
			matrix: New(
				[]float64{2, 3, 4, 5},
				[]float64{6, 7, 8, 9},
				[]float64{8, 7, 6, 5},
				[]float64{4, 3, 2, 1},
			),
			column:   1,
			expected: []float64{3, 7, 7, 3},
		},
		{
			matrix: New(
				[]float64{2, 3, 4, 5},
				[]float64{6, 7, 8, 9},
				[]float64{8, 7, 6, 5},
				[]float64{4, 3, 2, 1},
			),
			column:   2,
			expected: []float64{4, 8, 6, 2},
		},
		{
			matrix: New(
				[]float64{2, 3, 4, 5},
				[]float64{6, 7, 8, 9},
				[]float64{8, 7, 6, 5},
				[]float64{4, 3, 2, 1},
			),
			column:   3,
			expected: []float64{5, 9, 5, 1},
		},
	}
	for _, test := range tests {
		value := test.matrix.Column(test.column)
		for i := 0; i < len(value); i++ {
			if value[i] != test.expected[i] {
				t.Errorf(
					"Matrix %+v column %d got %+v expected %+v.",
					test.matrix,
					test.column,
					value,
					test.expected)
			}
		}
	}
}

func TestMultiplyCell(t *testing.T) {
	var tests = []struct {
		a, b     Matrix
		x, y     int
		expected float64
	}{
		{
			a: New(
				[]float64{1, 2, 3, 4},
				[]float64{2, 3, 4, 5},
				[]float64{3, 4, 5, 6},
				[]float64{4, 5, 6, 7},
			),
			b: New(
				[]float64{0, 1, 2, 4},
				[]float64{1, 2, 4, 8},
				[]float64{2, 4, 8, 16},
				[]float64{4, 6, 16, 32},
			),
			x:        1,
			y:        0,
			expected: 31,
		},
		{
			a: New(
				[]float64{1, 2, 3, 4},
				[]float64{5, 6, 7, 8},
				[]float64{9, 8, 7, 6},
				[]float64{5, 4, 3, 2},
			),
			b: New(
				[]float64{-2, 1, 2, 3},
				[]float64{3, 2, 1, -1},
				[]float64{4, 3, 6, 5},
				[]float64{1, 2, 7, 8},
			),
			x:        1,
			y:        0,
			expected: 44,
		},
	}
	for _, test := range tests {
		value := multiplyCell(test.a, test.b, test.x, test.y)
		if value != test.expected {
			t.Errorf("Cell multiply got %+v, expected %+v.", value, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []struct {
		a, b, expected Matrix
	}{
		{
			a: New(
				[]float64{1, 2, 3, 4},
				[]float64{5, 6, 7, 8},
				[]float64{9, 8, 7, 6},
				[]float64{5, 4, 3, 2},
			),
			b: New(
				[]float64{-2, 1, 2, 3},
				[]float64{3, 2, 1, -1},
				[]float64{4, 3, 6, 5},
				[]float64{1, 2, 7, 8},
			),
			expected: New(
				[]float64{20, 22, 50, 48},
				[]float64{44, 54, 114, 108},
				[]float64{40, 58, 110, 102},
				[]float64{16, 26, 46, 42},
			),
		},
	}
	for _, test := range tests {
		value := Multiply(test.a, test.b)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Failed to mulitply matricies %+v x %+v. Got %+v, expected %+v.",
				test.a, test.b, value, test.expected,
			)
		}
	}
}

func TestTmp(t *testing.T) {
	a := New(
		[]float64{1, 2, 3, 4},
		[]float64{5, 6, 7, 8},
		[]float64{9, 8, 7, 6},
		[]float64{5, 4, 3, 2},
	)
	b := New(
		[]float64{-2, 1, 2, 3},
		[]float64{3, 2, 1, -1},
		[]float64{4, 3, 6, 5},
		[]float64{1, 2, 7, 8},
	)
	multiplyCell(a, b, 1, 0)
}
