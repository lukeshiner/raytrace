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

func TestEqualMatricies(t *testing.T) {
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
		value := EqualMatricies(test.a, test.b)
		if value != test.expected {
			t.Error("Equal Matricies failed.")
		}
	}
}
