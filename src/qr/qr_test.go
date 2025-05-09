package qr

import (
	"math"
	"slices"
	"testing"

	m "face_recognition/matrix"
)

const EPSILON = 1e-6

func TestCalculateHouseholderVector(t *testing.T) {
	tests := []struct {
		name              string
		R                 m.Matrix
		colIdx            int
		size              int
		householderVector []float64
		want              float64
		wantErr           error
	}{
		{
			name: "3x3 matrix column 0",
			R: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					0, 3, 1,
					0, 0, 2,
				},
			},
			colIdx:            0,
			size:              1,
			householderVector: make([]float64, 1),
			want:              0.03125,
			wantErr:           nil,
		},
		{
			name: "3x3 matrix column 1",
			R: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					0, 3, 1,
					0, 0, 2,
				},
			},
			colIdx:            1,
			size:              2,
			householderVector: make([]float64, 2),
			want:              0.0555555,
			wantErr:           nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calculateHouseholderVector(tt.R, tt.colIdx, tt.size, tt.householderVector)
			if err != tt.wantErr {
				t.Errorf("calculateHouseholderVector(): returned incorrect error: %v, want %v", err, tt.wantErr)
			}

			if math.Abs(result-tt.want) > EPSILON {
				t.Errorf("calculateHouseholderVector(): returned incorrect result: %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateRMatrix(t *testing.T) {
	tests := []struct {
		name              string
		R                 m.Matrix
		colIdx            int
		size              int
		householderVector []float64
		beta              float64
		want              m.Matrix
	}{
		{
			name: "3x3 matrix first column",
			R: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					0, 3, 1,
					0, 0, 2,
				},
			},
			colIdx:            0,
			size:              3,
			householderVector: []float64{4, 0, 0},
			beta:              0.125,
			want: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					-4, -1, 2,
					0, 3, 1,
					0, 0, 2,
				},
			},
		},
		{
			name: "3x3 matrix second column",
			R: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					0, 3, 1,
					0, 0, 2,
				},
			},
			colIdx:            1,
			size:              2,
			householderVector: []float64{3, 0},
			beta:              0.2222222,
			want: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					0, -3, -1,
					0, 0, 2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rcopy := m.Matrix{
				Rows: tt.R.Rows,
				Cols: tt.R.Cols,
				Data: slices.Clone(tt.R.Data),
			}
			updateRMatrix(Rcopy, tt.colIdx, tt.size, slices.Clone(tt.householderVector), tt.beta)
			for i := range Rcopy.Data {
				if math.Abs(Rcopy.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("updateRMatrix(): at index %d, got %f, want %f", i, Rcopy.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestUpdateQMatrix(t *testing.T) {
	tests := []struct {
		name              string
		Q                 m.Matrix
		colIdx            int
		size              int
		householderVector []float64
		beta              float64
		want              m.Matrix
	}{
		{
			name: "3x3 matrix first column",
			Q: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					1, 3, 1,
					2, 2, 2,
				},
			},
			colIdx:            0,
			size:              3,
			householderVector: []float64{4, 0, 0},
			beta:              0.125,
			want: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					-4, 1, -2,
					-1, 3, 1,
					-2, 2, 2,
				},
			},
		},
		{
			name: "3x3 matrix second column",
			Q: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, 1, -2,
					1, 3, 1,
					2, 3, 2,
				},
			},
			colIdx:            1,
			size:              2,
			householderVector: []float64{3, 0},
			beta:              0.33333333,
			want: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					4, -2, -2,
					1, -6, 1,
					2, -6, 2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rcopy := m.Matrix{
				Rows: tt.Q.Rows,
				Cols: tt.Q.Cols,
				Data: slices.Clone(tt.Q.Data),
			}
			updateQMatrix(Rcopy, tt.colIdx, tt.size, slices.Clone(tt.householderVector), tt.beta)
			for i := range Rcopy.Data {
				if math.Abs(Rcopy.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("updateRMatrix(): at index %d, got %f, want %f", i, Rcopy.Data[i], tt.want.Data[i])
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
					-0.872871, -0.395318, 0.286038,
					-0.218217, -0.208062, -0.953462,
					-0.436435, 0.894669, -0.095346,
				},
			},
			wantR: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{
					-4.582575, -4.364357, -1.527525,
					0, 4.577377, 1.165150,
					0, 0, -3.051080,
				},
			},
			wantErr: nil,
		},
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
					t.Errorf("qr_Householder(): QT*Q at index %d, got %f, want %f", i, B.Data[i], identity.Data[i])
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
					0.120580, -0.959658, -0.254000,
					0.361740, 0.280751, -0.889000,
					0.924448, 0.015313, 0.381000,
				},
			},
			wantErr: nil,
		},
		{
			name: "output is correct with floating point numbers",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, vectors, err := QR_algorithm(tt.A)

			if err != tt.wantErr {
				t.Errorf("QR_algorithm(): returned wrong error: %v", err)
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
