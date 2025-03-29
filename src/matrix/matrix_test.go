package matrix

import (
	"math"
	"testing"
)

const EPSILON = 1e-10

func createBigMatrix(rows int, cols int, value float64) Matrix {
	data := make([]float64, rows*cols)
	for i := range data {
		data[i] = value
	}
	return Matrix{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func TestMultiplicationByScalar(t *testing.T) {
	tests := []struct {
		name   string
		A      Matrix
		scalar float64
		want   Matrix
	}{
		{
			name: "Output is correct with positive scalar",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{1, 2, 3, 4, 5, 6},
			},
			scalar: 3,
			want: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{3, 6, 9, 12, 15, 18},
			},
		},
		{
			name: "Output is correct with negative scalar",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{1, 2, 3, 4, 5, 6},
			},
			scalar: -2,
			want: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{-2, -4, -6, -8, -10, -12},
			},
		},
		{
			name: "Output is 0 with scalar 0",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{1, 2, 3, 4, 5, 6},
			},
			scalar: 0,
			want: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{0, 0, 0, 0, 0, 0},
			},
		},
		{
			name:   "algoritm works with 'big' matrices",
			A:      createBigMatrix(92, 112, 2),
			scalar: 4,
			want:   createBigMatrix(92, 112, 8),
		},
		{
			name: "Works with floating point numbers",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{2.5, 2, 3, 1.25, 0, 10},
			},
			scalar: float64(2.5),
			want: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{6.25, 5, 7.5, 3.125, 0, 25},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MultiplicationByScalar(tt.A, tt.scalar)
			if result.Cols != tt.want.Cols {
				t.Errorf("MultiplicationByScalar(): returned incorrect amount of colums")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("MultiplicationByScalar(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && math.Abs(result.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("MultiplicationByScalar(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	tests := []struct {
		name    string
		A       Matrix
		B       Matrix
		want    Matrix
		wantErr error
	}{
		{
			name: "Output is correct with valid matrices",
			A: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{2, 4, 1, 2, 3, 5},
			},
			B: Matrix{
				Rows: 2,
				Cols: 4,
				Data: []float64{2, 0, 1, 3, 0, 7, 8, 1},
			},
			want: Matrix{
				Rows: 3,
				Cols: 4,
				Data: []float64{4, 28, 34, 10, 2, 14, 17, 5, 6, 35, 43, 14},
			},
			wantErr: nil,
		},
		{
			name: "fails with invalid matrices",
			A: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{2, 4, 1, 2, 3, 5},
			},
			B: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{2, 4, 1, 2, 3, 5},
			},
			want:    Matrix{},
			wantErr: errIncorrectSize,
		},
		{
			name: "Output is correct with floating point matrices",
			A: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{2.5, 4.1, 0.8, 0.5, 2, 0.5},
			},
			B: Matrix{
				Rows: 2,
				Cols: 4,
				Data: []float64{2.5, 0, 0.1, 3, 0, 0.7, 8, 1},
			},
			want: Matrix{
				Rows: 3,
				Cols: 4,
				Data: []float64{6.25, 2.87, 33.05, 11.6, 2, 0.35, 4.08, 2.9, 5, 0.35, 4.2, 6.5},
			},
			wantErr: nil,
		},
		{
			name:    "algoritm works with 'big' matrices",
			A:       createBigMatrix(100, 100, 3),
			B:       createBigMatrix(100, 100, 2),
			want:    createBigMatrix(100, 100, 600),
			wantErr: nil,
		},
		{
			name: "Output is 0 when multiplied with 0-matrice.",
			A: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{2, 4, 1, 2, 3, 5},
			},
			B: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{0, 0, 0, 0},
			},
			want: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{0, 0, 0, 0, 0, 0},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Multiplication(tt.A, tt.B)
			if err != tt.wantErr {
				t.Errorf("Multiplication(): returned wrong error: %v", err)
			}
			if result.Cols != tt.want.Cols {
				t.Errorf("Multiplication(): returned incorrect amount of colums")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("Multiplication(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && math.Abs(result.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("Multiplication(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestAddition(t *testing.T) {
	tests := []struct {
		name    string
		A       Matrix
		B       Matrix
		want    Matrix
		wantErr error
	}{
		{
			name: "output is correct with valid matrices",
			A: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{1, 2, 3, 4},
			},
			B: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{1, 2, 3, 4},
			},
			want: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{2, 4, 6, 8},
			},
			wantErr: nil,
		},
		{
			name: "fails with matrices of different dimensions",
			A: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{1, 2, 3, 4},
			},
			B: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{1, 2, 3, 4, 5, 6},
			},
			want:    Matrix{},
			wantErr: errIncorrectSize,
		},
		{
			name: "output is correct with floating point matrices",
			A: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{1.5, 2.1, 3.68, 0.123},
			},
			B: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{0.5, 2.99, -0.321, 0},
			},
			want: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{2, 5.09, 3.359, 0.123},
			},
			wantErr: nil,
		},
		{
			name:    "algoritm works with 'big' matrices",
			A:       createBigMatrix(92, 112, 2),
			B:       createBigMatrix(92, 112, 4),
			want:    createBigMatrix(92, 112, 6),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Addition(tt.A, tt.B)
			if err != tt.wantErr {
				t.Errorf("Addition(): ...")
			}
			if result.Cols != tt.want.Cols {
				t.Errorf("Addition(): returned incorrect amount of colums")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("Addition(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && math.Abs(result.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("Multiplication(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		name string
		A    Matrix
		want Matrix
	}{
		{
			name: "output is correct with non-square matrix",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{1, 2, 3, 4, 5, 6},
			},
			want: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{1, 4, 2, 5, 3, 6},
			},
		},
		{
			name: "output is correct with square matrix",
			A: Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{1, 4, 7, 2, 5, 8, 3, 6, 9},
			},
		},
		{
			name: "output is correct with floating point matrix",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []float64{0.123, 2.7, 3, 0, 5.99, -0.6},
			},
			want: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{0.123, 0, 2.7, 5.99, 3, -0.6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Transpose(tt.A)
			if result.Cols != tt.want.Cols {
				t.Errorf("Transpose(): returned incorrect amount of cols")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("Transpose(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && math.Abs(result.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("Transpose(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestIdentity(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want Matrix
	}{
		{
			name: "Output is corrent with 3x3 matrix",
			n:    3,
			want: Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
			},
		},
		{
			name: "Output is corrent with n=0",
			n:    0,
			want: Matrix{
				Rows: 0,
				Cols: 0,
				Data: []float64{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Identity(tt.n)
			if result.Cols != tt.want.Cols {
				t.Errorf("Identity(): returned incorrect amount of cols")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("Identity(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && result.Data[i] != tt.want.Data[i] {
					t.Errorf("Identity(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestDifferenceMatrix(t *testing.T) {
	tests := []struct {
		name    string
		vectors []Matrix
		mean    Matrix
		want    Matrix
		wantErr error
	}{
		{
			name: "output is correct with valid matrices",
			vectors: []Matrix{
				{
					Rows: 4,
					Cols: 1,
					Data: []float64{1, 2, 3, 4},
				},
				{
					Rows: 4,
					Cols: 1,
					Data: []float64{9, 7, 5, 3},
				},
			},
			mean: Matrix{
				Rows: 4,
				Cols: 1,
				Data: []float64{5, 4.5, 4, 3.5},
			},
			want: Matrix{
				Rows: 4,
				Cols: 2,
				Data: []float64{-4, 4, -2.5, 2.5, -1, 1, 0.5, -0.5},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DifferenceMatrix(tt.vectors, tt.mean)
			if err != tt.wantErr {
				t.Errorf("DifferenceMatrix(): returned wrong error %v, want %v", err, tt.wantErr)
			}
			if result.Cols != tt.want.Cols {
				t.Errorf("DifferenceMatrix(): returned incorrect amount of cols")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("DifferenceMatrix(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && math.Abs(result.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("DifferenceMatrix(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestCovariance(t *testing.T) {
	tests := []struct {
		name    string
		A       Matrix
		want    Matrix
		wantErr error
	}{
		{
			name: "output is correct with valid matrices",
			A: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []float64{1, 2, 3, 4, 5, 6},
			},
			want: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{35, 44, 44, 56},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Covariance(tt.A)
			if err != tt.wantErr {
				t.Errorf("Covariance(): returned wrong error %v, want %v", err, tt.wantErr)
			}
			if result.Cols != tt.want.Cols {
				t.Errorf("Covariance(): returned incorrect amount of cols")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("Covariance(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && math.Abs(result.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("Covariance(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}
