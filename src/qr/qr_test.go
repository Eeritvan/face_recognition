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
			result := householderMatrix(tt.u, tt.k, tt.n)
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
	tests := []struct {
		name    string
		A       m.Matrix
		wantQ   m.Matrix
		wantR   m.Matrix
		wantErr error
	}{
		{
			name: "output is correct with integer matrices",
			A: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{4, 2, 0, 1, 0, 3, 2, 6, 2},
			},
			wantQ: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					0.872871, 0.395318, -0.286038,
					0.218217, 0.208062, 0.953462,
					0.436435, -0.894669, 0.095346,
				},
			},
			wantR: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4.582575, 4.364357, 1.527525,
					0, -4.577377, -1.165150,
					0, 0, 3.051080,
				},
			},
			wantErr: nil,
		},
		// todo: testing some errors idk
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Q, R, err := qr_Householder(tt.A)

			if err != tt.wantErr {
				t.Errorf("%v", err)
			}

			// checking Q
			if Q.Rows != tt.wantQ.Rows {
				t.Errorf("qr_Householder(): Q returned incorrect amount of rows")
			}
			if Q.Cols != tt.wantQ.Cols {
				t.Errorf("qr_Householder(): Q returned incorrect amount of cols")
			}
			for i := range Q.Data {
				if i < len(tt.wantQ.Data) && math.Abs(Q.Data[i]-tt.wantQ.Data[i]) > EPSILON {
					t.Errorf("qr_Householder(): Q at index %d, got %f, want %f", i, Q.Data[i], tt.wantQ.Data[i])
				}
			}

			// checking R
			if R.Rows != tt.wantR.Rows {
				t.Errorf("qr_Householder(): R returned incorrect amount of rows")
			}
			if R.Cols != tt.wantR.Cols {
				t.Errorf("qr_Householder(): R returned incorrect amount of cols")
			}
			for i := range R.Data {
				if i < len(tt.wantR.Data) && math.Abs(R.Data[i]-tt.wantR.Data[i]) > EPSILON {
					t.Errorf("qr_Householder(): R at index %d, got %f, want %f", i, R.Data[i], tt.wantR.Data[i])
				}
			}

			// check if Q * QT or QT * Q == Identity matrix
			identity := m.Identity(Q.Rows)
			B, _ := m.Multiplication(m.Transpose(Q), Q)
			for i := range B.Data {
				if i < len(identity.Data) && math.Abs(B.Data[i]-identity.Data[i]) > EPSILON {
					t.Errorf("QR_algoqr_Householderrithm(): QT*Q at index %d, got %f, want %f", i, B.Data[i], identity.Data[i])
				}
			}
			C, _ := m.Multiplication(m.Transpose(Q), Q)
			for i := range C.Data {
				if i < len(identity.Data) && math.Abs(C.Data[i]-identity.Data[i]) > EPSILON {
					t.Errorf("qr_Householder(): Q*QTs at index %d, got %f, want %f", i, C.Data[i], identity.Data[i])
				}
			}

			// Q * R = original matrix
			D, _ := m.Multiplication(Q, R)
			for i := range D.Data {
				if i < len(tt.A.Data) && math.Abs(D.Data[i]-tt.A.Data[i]) > EPSILON {
					t.Errorf("qr_Householder(): Q*R at index %d, got %f, want %f", i, D.Data[i], tt.A.Data[i])
				}
			}

		})
	}
}

func TestQR_algorithm(t *testing.T) {
	tests := []struct {
		name        string
		A           m.Matrix
		wantValues  []float64
		wantVectors m.Matrix
		wantErr     error
	}{
		{
			name: "output is correct with integers",
			A: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{4, 2, 0, 1, 2, 3, 5, 6, 7},
			},
			wantValues: []float64{10, 3, 0},
			wantVectors: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					0.120580, 0.959658, -0.254000,
					0.361740, -0.280751, -0.889000,
					0.924448, -0.015313, 0.381000,
				},
			},
			wantErr: nil,
		},
		{
			name: "output is correct with integers",
			A: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{4.5, 2.2, 0, 1.1, 0.02, 0.03, 0, 6, 0},
			},
			wantValues: []float64{4.990437, -0.701736, 0.231298},
			wantVectors: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					0.944265, 0.216936, -0.247592,
					0.210501, 0.180336, 0.960816,
					0.253085, -0.959383, 0.124620,
				},
			},
			wantErr: nil,
		},
		// fix needed for non-square matrices !!!
		// {
		// 	name: "",
		// 	A: m.Matrix{
		// 		Rows: 3,
		// 		Cols: 4,
		// 		Data: []float64{4, 2, 0, 1, 2, 3, 5, 6, 7, 1, 2},
		// 	},
		// 	wantValues:  []float64{},
		// 	wantVectors: m.Matrix{},
		// 	wantErr:     nil,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, vectors, err := QR_algorithm(tt.A)

			if err != tt.wantErr {
				t.Errorf("%v", err)
			}

			for i := range values {
				if i < len(tt.wantValues) && math.Abs(values[i]-tt.wantValues[i]) > EPSILON {
					t.Errorf("QR_algorithm(): at index %d, got %f, want %f", i, values[i], tt.wantValues[i])
				}
			}

			if vectors.Rows != tt.wantVectors.Rows {
				t.Errorf("QR_algorithm(): returned incorrect amount of rows")
			}

			if vectors.Cols != tt.wantVectors.Cols {
				t.Errorf("QR_algorithm(): returned incorrect amount of cols")
			}

			for i := range vectors.Data {
				if i < len(tt.wantVectors.Data) && math.Abs(vectors.Data[i]-tt.wantVectors.Data[i]) > EPSILON {
					t.Errorf("QR_algorithm(): at index %d, got %f, want %f", i, values[i], tt.wantValues[i])
				}
			}
		})
	}
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
			result := hasConverged(tt.prev, tt.curr)
			if result != tt.want {
				t.Errorf("hasConverged(): %v, want %v", result, tt.want)
			}
		})
	}
}
