package qr

import (
	m "face_recognition/matrix"
	"math"
	"testing"
)

const EPSILON = 1e-6

func TestHouseholderVector(t *testing.T) {
	tests := []struct {
		name  string
		input m.Matrix
		col   int
		n     int
		want  []float64
	}{
		{
			name: "3x3 matrix column 0",
			input: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					1, 4, 1,
					-2, 1, 4,
				},
			},
			col:  0,
			n:    3,
			want: []float64{4 - math.Sqrt(21), 1, -2},
		},
		{
			name: "3x3 matrix column 1",
			input: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					1, 4, 1,
					-2, 1, 4,
				},
			},
			col:  1,
			n:    3,
			want: []float64{4 - math.Sqrt(17), 1},
		},
		{
			name: "3x3 matrix with negative column 0",
			input: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					-4, 1, -2,
					-1, 4, 1,
					-2, 1, 4,
				},
			},
			col:  0,
			n:    3,
			want: []float64{-4 + math.Sqrt(21), -1, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := householderVector(tt.input, tt.col, tt.n)
			if len(result) != len(tt.want) {
				t.Errorf("householderVector(): length should be %d, got %d", len(tt.want), len(result))
			}
			for i := range result {
				if math.Abs(result[i]-tt.want[i]) > 1e-10 {
					t.Errorf("householderVector(): at index %d: got %f, want %f", i, result[i], tt.want[i])
				}
			}
		})
	}
}

func TestNormalizeVector(t *testing.T) {
	tests := []struct {
		name string
		v    []float64
		n    int
		want []float64
	}{
		{
			name: "3x1 vector",
			v:    []float64{3, 4, 0},
			n:    3,
			want: []float64{0.6, 0.8, 0},
		},
		{
			name: "2x1 vector",
			v:    []float64{1, 1},
			n:    2,
			want: []float64{1 / math.Sqrt(2), 1 / math.Sqrt(2)},
		},
		{
			name: "output is correct with zero vector",
			v:    []float64{0, 0, 0},
			n:    3,
			want: []float64{0, 0, 0},
		},
		{
			name: "vector with negative components works",
			v:    []float64{-2, 2},
			n:    2,
			want: []float64{-1 / math.Sqrt(2), 1 / math.Sqrt(2)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeVector(tt.v, tt.n)
			if len(result) != len(tt.want) {
				t.Errorf("normalizeVector(): length should be %d, got %d", len(tt.want), len(result))
			}
			for i := range result {
				if math.Abs(result[i]-tt.want[i]) > EPSILON {
					t.Errorf("normalizeVector(): at index %d: got %f, want %f", i, result[i], tt.want[i])
				}
			}
		})
	}
}

func TestHouseholderMatrix(t *testing.T) {
	tests := []struct {
		name string
		u    []float64
		k    int
		n    int
		want m.Matrix
	}{
		{
			name: "3x3 matrix with k=0",
			u:    []float64{1 / math.Sqrt(2), 1 / math.Sqrt(2), 0},
			k:    0,
			n:    3,
			want: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					0, -1, 0,
					-1, 0, 0,
					0, 0, 1,
				},
			},
		},
		{
			name: "3x3 matrix with k=1",
			u:    []float64{1 / math.Sqrt(2), 1 / math.Sqrt(2)},
			k:    1,
			n:    3,
			want: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					1, 0, 0,
					0, 0, -1,
					0, -1, 0,
				},
			},
		},
		{
			name: "2x2 matrix with k=0",
			u:    []float64{1 / math.Sqrt(2), 1 / math.Sqrt(2)},
			k:    0,
			n:    2,
			want: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{
					0, -1,
					-1, 0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HouseholderMatrix(tt.u, tt.k, tt.n)
			if result.Cols != tt.want.Cols {
				t.Errorf("householderMatrix(): returned incorrect amount of cols")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("householderMatrix(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if math.Abs(result.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("householderMatrix(): at index %d: got %f, want %f",
						i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestQR_Householder(t *testing.T) {
}

func TestQR_algorithm(t *testing.T) {

}

func TestHasConverged(t *testing.T) {
	tests := []struct {
		name string
		prev m.Matrix
		curr m.Matrix
		want bool
	}{
		{
			name: "2x2 matrices that have converged",
			prev: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{
					1.0, 0.0,
					0.0, 2.0,
				},
			},
			curr: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{
					1.0, 0.0,
					0.0, 2.0,
				},
			},
			want: true,
		},
		{
			name: "2x2 matrices that have not converged",
			prev: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{
					1.0, 0.5,
					0.5, 2.0,
				},
			},
			curr: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{
					1.2, 0.1,
					0.1, 1.8,
				},
			},
			want: false,
		},
		{
			name: "3x3 matrices with small differences within tolerance",
			prev: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					1.0, 0.0, 0.0,
					0.0, 2.0, 0.0,
					0.0, 0.0, 3.0,
				},
			},
			curr: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					1.000000000001, 0.0, 0.0,
					0.0, 2.000000000001, 0.0,
					0.0, 0.0, 3.000000000001,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasConverged(tt.prev, tt.curr)
			if result != tt.want {
				t.Errorf("HasConverged(): %v, want %v", result, tt.want)
			}
		})
	}
}
