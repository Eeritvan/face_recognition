package matrix

import (
	"testing"
)

func createBigMatrix(rows int, cols int, value int) Matrix {
	data := make([]int, rows*cols)
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
		scalar int
		want   Matrix
	}{
		{
			name: "Output is correct with positive scalar",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []int{1, 2, 3, 4, 5, 6},
			},
			scalar: 3,
			want: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []int{3, 6, 9, 12, 15, 18},
			},
		},
		{
			name: "Output is correct with negative scalar",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []int{1, 2, 3, 4, 5, 6},
			},
			scalar: -2,
			want: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []int{-2, -4, -6, -8, -10, -12},
			},
		},
		{
			name: "Output is 0 with scalar 0",
			A: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []int{1, 2, 3, 4, 5, 6},
			},
			scalar: 0,
			want: Matrix{
				Rows: 2,
				Cols: 3,
				Data: []int{0, 0, 0, 0, 0, 0},
			},
		},
		{
			name:   "algoritm works with 'big' matrices",
			A:      createBigMatrix(92, 112, 2),
			scalar: 4,
			want:   createBigMatrix(92, 112, 8),
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
				if i < len(tt.want.Data) && result.Data[i] != tt.want.Data[i] {
					t.Errorf("MultiplicationByScalar(): at index %d, got %d, want %d", i, result.Data[i], tt.want.Data[i])
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
				Data: []int{2, 4, 1, 2, 3, 5},
			},
			B: Matrix{
				Rows: 2,
				Cols: 4,
				Data: []int{2, 0, 1, 3, 0, 7, 8, 1},
			},
			want: Matrix{
				Rows: 3,
				Cols: 4,
				Data: []int{4, 28, 34, 10, 2, 14, 17, 5, 6, 35, 43, 14},
			},
			wantErr: nil,
		},
		{
			name: "fails with invalid matrices",
			A: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []int{2, 4, 1, 2, 3, 5},
			},
			B: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []int{2, 4, 1, 2, 3, 5},
			},
			want:    Matrix{},
			wantErr: errIncorrectSize,
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
				Data: []int{2, 4, 1, 2, 3, 5},
			},
			B: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []int{0, 0, 0, 0},
			},
			want: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []int{0, 0, 0, 0, 0, 0},
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
				if i < len(tt.want.Data) && result.Data[i] != tt.want.Data[i] {
					t.Errorf("Multiplication(): at index %d, got %d, want %d", i, result.Data[i], tt.want.Data[i])
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
				Data: []int{1, 2, 3, 4},
			},
			B: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []int{1, 2, 3, 4},
			},
			want: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []int{2, 4, 6, 8},
			},
			wantErr: nil,
		},
		{
			name: "fails with matrices of different dimensions",
			A: Matrix{
				Rows: 2,
				Cols: 2,
				Data: []int{1, 2, 3, 4},
			},
			B: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []int{1, 2, 3, 4, 5, 6},
			},
			want:    Matrix{},
			wantErr: errIncorrectSize,
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
				if i < len(tt.want.Data) && result.Data[i] != tt.want.Data[i] {
					t.Errorf("Multiplication(): at index %d, got %d, want %d", i, result.Data[i], tt.want.Data[i])
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
				Data: []int{1, 2, 3, 4, 5, 6},
			},
			want: Matrix{
				Rows: 3,
				Cols: 2,
				Data: []int{1, 4, 2, 5, 3, 6},
			},
		},
		{
			name: "output is correct with square matrix",
			A: Matrix{
				Rows: 3,
				Cols: 3,
				Data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: Matrix{
				Rows: 3,
				Cols: 3,
				Data: []int{1, 4, 7, 2, 5, 8, 3, 6, 9},
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
				if i < len(tt.want.Data) && result.Data[i] != tt.want.Data[i] {
					t.Errorf("Transpose(): at index %d, got %d, want %d", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}
