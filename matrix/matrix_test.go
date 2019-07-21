package matrix

import (
	"math"
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

func TestMutliplyTuple(t *testing.T) {
	var tests = []struct {
		matrix   Matrix
		tuple    []float64
		expected []float64
	}{
		{
			matrix: New(
				[]float64{1, 2, 3, 4},
				[]float64{2, 4, 4, 2},
				[]float64{8, 6, 4, 1},
				[]float64{0, 0, 0, 1},
			),
			tuple:    []float64{1, 2, 3, 1},
			expected: []float64{18, 24, 33, 1},
		},
	}
	for _, test := range tests {
		value := MultiplyTuple(test.matrix, test.tuple)
		for i := 0; i < len(test.expected); i++ {
			if value[i] != test.expected[i] {
				t.Errorf(
					"Multiply Matrix %+v with Tuple %+v got %+v, expected %+v.",
					test.matrix, test.tuple, value, test.expected,
				)
			}
		}
	}
}

func TestIdentityMatrix(t *testing.T) {
	expected := New(
		[]float64{1, 0, 0, 0},
		[]float64{0, 1, 0, 0},
		[]float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1},
	)
	value := IdentityMatrix(4)
	if Equal(value, expected) != true {
		t.Errorf("Creating Identity Matrix of size 4 returned %+v", value)
	}
}

func TestMultiplyByIdentity(t *testing.T) {
	matrix := New(
		[]float64{0, 1, 2, 4},
		[]float64{1, 2, 4, 8},
		[]float64{2, 4, 8, 16},
		[]float64{4, 8, 16, 32},
	)
	identity := IdentityMatrix(4)
	value := Multiply(matrix, identity)
	if Equal(matrix, value) != true {
		t.Errorf(
			"Multiplying matrix %+v with the identity matrix returned %+v.",
			matrix, value,
		)
	}
}

func TestMultiplyTupleByIdentity(t *testing.T) {
	tuple := []float64{1, 2, 3, 4}
	identity := IdentityMatrix(4)
	value := MultiplyTuple(identity, tuple)
	for i := 0; i < len(tuple); i++ {
		if tuple[i] != value[i] {
			t.Errorf(
				"Multiplying tuple %+v with the identity matrix returned %+v.",
				tuple, value,
			)
		}
	}
}

func TestTransposeMethod(t *testing.T) {
	var tests = []struct {
		matrix   Matrix
		expected Matrix
	}{
		{
			matrix: New(
				[]float64{0, 9, 3, 0},
				[]float64{9, 8, 0, 8},
				[]float64{1, 8, 5, 3},
				[]float64{0, 0, 5, 8},
			),
			expected: New(
				[]float64{0, 9, 1, 0},
				[]float64{9, 8, 8, 0},
				[]float64{3, 0, 5, 5},
				[]float64{0, 8, 3, 8},
			),
		},
		{
			matrix:   IdentityMatrix(4),
			expected: IdentityMatrix(4),
		},
	}
	for _, test := range tests {
		value := test.matrix.Transpose()
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Transposed Matrix %+v got %+v, expected %+v.",
				test.matrix, value, test.expected,
			)
		}
	}
}

func TestDeterminant2x2(t *testing.T) {
	var tests = []struct {
		matrix   Matrix
		expected float64
	}{
		{
			matrix: New(
				[]float64{1, 5},
				[]float64{-3, 2},
			),
			expected: 17,
		},
	}
	for _, test := range tests {
		value := test.matrix.determinant()
		if value != test.expected {
			t.Errorf(
				"Determinant of %+v: got %+v, expected %+v",
				test.matrix, value, test.expected)
		}
	}
}

func TestSubmatrix(t *testing.T) {
	var tests = []struct {
		matrix, expected Matrix
		row, column      int
	}{
		{
			matrix: New(
				[]float64{1, 5, 0},
				[]float64{-3, 2, 7},
				[]float64{0, 6, -3},
			),
			row:    0,
			column: 2,
			expected: New(
				[]float64{-3, 2},
				[]float64{0, 6},
			),
		},
	}
	for _, test := range tests {
		value := test.matrix.Submatrix(test.row, test.column)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Submatrix (%d, %d) of matrix %+v was %+v, expected %+v",
				test.row, test.column, test.matrix, value, test.expected,
			)
		}
	}
}

func TestMinor3x3(t *testing.T) {
	var tests = []struct {
		matrix      Matrix
		row, column int
		expected    float64
	}{
		{
			matrix: New(
				[]float64{3, 5, 0},
				[]float64{2, -1, 7},
				[]float64{6, -1, 5},
			),
			row:      1,
			column:   0,
			expected: 25,
		},
	}
	for _, test := range tests {
		value := test.matrix.minor(test.row, test.column)
		if value != test.expected {
			t.Errorf(
				"Minor of matrix %+v (%d, %d): got %+v, expected %+v.",
				test.matrix, test.row, test.column, value, test.expected,
			)
		}
	}
}

func TestCofactor3x3(t *testing.T) {
	var tests = []struct {
		matrix      Matrix
		row, column int
		expected    float64
	}{
		{
			matrix: New(
				[]float64{3, 5, 0},
				[]float64{2, -1, -7},
				[]float64{6, -1, 5},
			),
			row:      1,
			column:   0,
			expected: -25,
		},
	}
	for _, test := range tests {
		value := test.matrix.cofactor(test.row, test.column)
		if value != test.expected {
			t.Errorf(
				"Cofactor of matrix %+v (%d, %d): got %+v, expected %+v.",
				test.matrix, test.row, test.column, value, test.expected,
			)
		}
	}
}

func TestDeterminant(t *testing.T) {
	type determinat3x3test struct {
		row, column int
		expected    float64
	}
	var tests = []struct {
		matrix              Matrix
		row, column         int
		cofactorTests       []determinat3x3test
		expectedDeterminant float64
	}{
		{
			matrix: New(
				[]float64{1, 2, 6},
				[]float64{-5, 8, -4},
				[]float64{2, 6, 4},
			),
			cofactorTests: []determinat3x3test{
				{row: 0, column: 0, expected: 56},
				{row: 0, column: 1, expected: 12},
				{row: 0, column: 2, expected: -46},
			},
			expectedDeterminant: -196,
		},
		{
			matrix: New(
				[]float64{-2, -8, 3, 5},
				[]float64{-3, 1, 7, 3},
				[]float64{1, 2, -9, 6},
				[]float64{-6, 7, 7, -9},
			),
			cofactorTests: []determinat3x3test{
				{row: 0, column: 0, expected: 690},
				{row: 0, column: 1, expected: 447},
				{row: 0, column: 2, expected: 210},
				{row: 0, column: 3, expected: 51},
			},
			expectedDeterminant: -4071,
		},
	}
	for _, test := range tests {
		for _, cofactorTest := range test.cofactorTests {
			cValue := test.matrix.cofactor(cofactorTest.row, cofactorTest.column)
			if cValue != cofactorTest.expected {
				t.Errorf(
					"Matrix %+v cofactor (%d, %d) was %+v, expected %+v.",
					test.matrix, cofactorTest.row, cofactorTest.column,
					cValue, cofactorTest.expected,
				)
			}
		}
		value := test.matrix.determinant()
		if value != test.expectedDeterminant {
			t.Errorf(
				"Matrix %+v determinant was %+v, expected %+v.",
				test.matrix, value, test.expectedDeterminant,
			)
		}
	}
}

// Inverse

func TestInvertable(t *testing.T) {
	var tests = []struct {
		matrix      Matrix
		determinant float64
		invertable  bool
	}{
		{
			matrix: New(
				[]float64{6, 4, 4, 4},
				[]float64{5, 5, 7, 6},
				[]float64{4, -9, 3, -7},
				[]float64{9, 1, 7, -6},
			),
			determinant: -2120,
			invertable:  true,
		},
		{
			matrix: New(
				[]float64{-4, 2, -2, -3},
				[]float64{9, 6, 2, 6},
				[]float64{0, -5, 1, -5},
				[]float64{0, 0, 0, 0},
			),
			determinant: 0,
			invertable:  false,
		},
	}
	for _, test := range tests {
		determinantValue := test.matrix.determinant()
		if determinantValue != test.determinant {
			t.Errorf(
				"The determinant of matrix %+v was %+v, expected %+v",
				test.matrix, determinantValue, test.determinant,
			)
		}
		invertableValue := test.matrix.Invertable()
		if invertableValue != test.invertable {
			t.Errorf(
				"Matrix.Invertable for matrix %+v was %+v, expected %+v",
				test.matrix, invertableValue, test.invertable,
			)
		}
	}
}

func TestInvert(t *testing.T) {
	var tests = []struct {
		matrix  Matrix
		inverse Matrix
	}{
		{
			matrix: New(
				[]float64{-5, 2, 6, -8},
				[]float64{1, -5, 1, 8},
				[]float64{7, 7, -6, -7},
				[]float64{1, -3, 7, 4},
			),
			inverse: New(
				[]float64{0.21805, 0.45113, 0.24060, -0.04511},
				[]float64{-0.80827, -1.45677, -0.44361, 0.52068},
				[]float64{-0.07895, -0.22368, -0.05263, 0.19737},
				[]float64{-0.52256, -0.81391, -0.30075, 0.30639},
			),
		},
		{
			matrix: New(
				[]float64{8, -5, 9, 2},
				[]float64{7, 5, 6, 1},
				[]float64{-6, 0, 9, 6},
				[]float64{-3, 0, -9, -4},
			),
			inverse: New(
				[]float64{-0.15385, -0.15385, -0.28205, -0.53846},
				[]float64{-0.07692, 0.12308, 0.02564, 0.03077},
				[]float64{0.35897, 0.35897, 0.43590, 0.92308},
				[]float64{-0.69231, -0.69231, -0.76923, -1.92308},
			),
		},
		{
			matrix: New(
				[]float64{9, 3, 0, 9},
				[]float64{-5, -2, -6, -3},
				[]float64{-4, 9, 6, 4},
				[]float64{-7, 6, 6, 2},
			),
			inverse: New(
				[]float64{-0.04074, -0.07778, 0.14444, -0.22222},
				[]float64{-0.07778, 0.03333, 0.36667, -0.33333},
				[]float64{-0.02901, -0.14630, -0.10926, 0.12963},
				[]float64{0.17778, 0.06667, -0.26667, 0.33333},
			),
		},
		{
			matrix: New(
				[]float64{-4, 2, -2, -3},
				[]float64{9, 6, 2, 6},
				[]float64{0, -5, 1, -5},
				[]float64{0, 0, 0, 0},
			),
			inverse: Matrix{},
		},
	}
	for _, test := range tests {
		value, _ := test.matrix.Invert()
		if Equal(value, test.inverse) != true {
			t.Errorf(
				"Inverse of matrix %+v was %+v, expected %+v.",
				test.matrix, value, test.inverse,
			)
		}
	}
}

func TestMultiplyByInverse(t *testing.T) {
	var tests = []struct {
		a, b Matrix
	}{
		{
			a: New(
				[]float64{3, -9, 7, 3},
				[]float64{3, -8, 2, -9},
				[]float64{-4, 4, 4, 1},
				[]float64{-6, 5, -1, 1},
			),
			b: New(
				[]float64{8, 2, 2, 2},
				[]float64{3, -1, 7, 0},
				[]float64{7, 0, 5, 4},
				[]float64{6, -2, 0, 5},
			),
		},
	}
	for _, test := range tests {
		c := Multiply(test.a, test.b)
		inverse, _ := test.b.Invert()
		value := Multiply(c, inverse)
		if Equal(test.a, value) != true {
			t.Errorf("Failed get back original matrix by multiplying with the inverse.")
		}
	}
}

// Translation

func TestTranslationMatrix(t *testing.T) {
	var tests = []struct {
		x, y, z  float64
		expected Matrix
	}{
		{
			x: 3, y: -4, z: -7,
			expected: New(
				[]float64{1, 0, 0, 3},
				[]float64{0, 1, 0, -4},
				[]float64{0, 0, 1, -7},
				[]float64{0, 0, 0, 1},
			),
		},
	}
	for _, test := range tests {
		value := TranslationMatrix(test.x, test.y, test.z)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Translation(%+v, %+v, %+v) produced %+v, expected %+v.",
				test.x, test.y, test.z, value, test.expected,
			)
		}
	}
}

// Scaling

func TestScalingMatrix(t *testing.T) {
	var tests = []struct {
		x, y, z  float64
		expected Matrix
	}{
		{
			x: 3, y: -4, z: -7,
			expected: New(
				[]float64{3, 0, 0, 0},
				[]float64{0, -4, 0, 0},
				[]float64{0, 0, -7, 0},
				[]float64{0, 0, 0, 1},
			),
		},
	}
	for _, test := range tests {
		value := ScalingMatrix(test.x, test.y, test.z)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Scaling(%+v, %+v, %+v) produced %+v, expected %+v.",
				test.x, test.y, test.z, value, test.expected,
			)
		}
	}
}

// Rotation
func TestRotationXMatrix(t *testing.T) {
	var tests = []struct {
		radians  float64
		expected Matrix
	}{
		{
			radians: math.Pi / 2,
			expected: New(
				[]float64{1, 0, 0, 0},
				[]float64{0, math.Cos(math.Pi / 2), -math.Sin(math.Pi / 2), 0},
				[]float64{0, math.Sin(math.Pi / 2), math.Cos(math.Pi / 2), 0},
				[]float64{0, 0, 0, 1},
			),
		},
	}
	for _, test := range tests {
		value := RotationXMatrix(test.radians)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"RotationX(%v) produced %+v, expected %+v.",
				test.radians, value, test.expected,
			)
		}
	}
}

func TestRotationYMatrix(t *testing.T) {
	var tests = []struct {
		radians  float64
		expected Matrix
	}{
		{
			radians: math.Pi / 2,
			expected: New(
				[]float64{math.Cos(math.Pi / 2), 0, math.Sin(math.Pi / 2), 0},
				[]float64{0, 1, 0, 0},
				[]float64{-math.Sin(math.Pi / 2), 0, math.Cos(math.Pi / 2), 0},
				[]float64{0, 0, 0, 1},
			),
		},
	}
	for _, test := range tests {
		value := RotationYMatrix(test.radians)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"RotationY(%v) produced %+v, expected %+v.",
				test.radians, value, test.expected,
			)
		}
	}
}

func TestRotationZMatrix(t *testing.T) {
	var tests = []struct {
		radians  float64
		expected Matrix
	}{
		{
			radians: math.Pi / 2,
			expected: New(
				[]float64{math.Cos(math.Pi / 2), -math.Sin(math.Pi / 2), 0, 0},
				[]float64{math.Sin(math.Pi / 2), math.Cos(math.Pi / 2), 0, 0},
				[]float64{0, 0, 1, 0},
				[]float64{0, 0, 0, 1},
			),
		},
	}
	for _, test := range tests {
		value := RotationZMatrix(test.radians)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"RotationZ(%v) produced %+v, expected %+v.",
				test.radians, value, test.expected,
			)
		}
	}
}

// Shearing

func TestShearingMartrix(t *testing.T) {
	var tests = []struct {
		Xy, Xz, Yx, Yz, Zx, Zy float64
		expected               Matrix
	}{
		{
			Xy: 1, Xz: 2, Yx: 3, Yz: 4, Zx: 5, Zy: 6,
			expected: New(
				[]float64{1, 1, 2, 0},
				[]float64{3, 1, 4, 0},
				[]float64{5, 6, 1, 0},
				[]float64{0, 0, 0, 1},
			),
		},
	}
	for _, test := range tests {
		value := ShearingMatrix(test.Xy, test.Xz, test.Yx, test.Yz, test.Zx, test.Zy)
		if Equal(value, test.expected) != true {
			t.Errorf(
				"Shearing(%v, %v, %v, %v, %v, %v) produced %+v, expected %+v.",
				test.Xy, test.Xz, test.Yx, test.Yz, test.Zx, test.Zy, value, test.expected,
			)
		}
	}
}
