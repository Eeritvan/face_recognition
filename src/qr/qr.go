package qr

import (
	m "face_recognition/matrix"
	"fmt"
	"math"
)

func householderVector(R m.Matrix, k, n int) []float64 {
	v := make([]float64, n-k)
	alpha := 0.0
	for i := k; i < n; i++ {
		v[i-k] = R.Data[i*n+k]
		alpha += R.Data[i*n+k] * R.Data[i*n+k]
	}
	alpha = math.Sqrt(alpha)
	if v[0] > 0 {
		alpha = -alpha
	}
	v[0] = v[0] + alpha
	return v
}

func normalizeVector(v []float64, n int) []float64 {
	u := make([]float64, n)
	copy(u, v)

	uNorm := 0.0
	for j := range u {
		uNorm += u[j] * u[j]
	}
	uNorm = math.Sqrt(uNorm)

	if math.Abs(uNorm) < 1e-10 {
		return u
	}

	for j := range u {
		u[j] /= uNorm
	}
	return u
}

func householderMatrix(u []float64, k, n int) m.Matrix {
	Hk := m.Identity(n)
	for i := k; i < n; i++ {
		for j := k; j < n; j++ {
			Hk.Data[i*n+j] -= 2 * u[i-k] * u[j-k]
		}
	}
	return Hk
}

// https://www.youtube.com/watch?v=n0zDgkbFyQk
func QR_Householder(A m.Matrix) (m.Matrix, m.Matrix, error) {
	n := A.Rows

	Q := m.Identity(n)
	R := m.Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: append([]float64(nil), A.Data...),
	}
	for k := range n - 1 {
		v := householderVector(R, k, n)
		u := normalizeVector(v, n)
		Hk := householderMatrix(u, k, n)

		newR, err := m.Multiplication(Hk, R)
		if err != nil {
			return m.Matrix{}, m.Matrix{}, err
		}
		R = newR

		Hk_T := m.Transpose(Hk)
		newQ, err := m.Multiplication(Q, Hk_T)
		if err != nil {
			return m.Matrix{}, m.Matrix{}, err
		}
		Q = newQ
	}
	return Q, R, nil
}

// https://www.youtube.com/watch?v=McHW221J3UM
func QR_algorithm(A m.Matrix) ([]float64, m.Matrix, error) {
	Ak := m.Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: append([]float64(nil), A.Data...),
	}

	V := m.Identity(A.Rows)

	for range 50 {
		Q, R, err := QR_Householder(Ak)
		if err != nil {
			return nil, m.Matrix{}, err
		}

		newAk, err := m.Multiplication(R, Q)
		if err != nil {
			return nil, m.Matrix{}, err
		}

		// Update V by multiplying with Q
		newV, err := m.Multiplication(V, Q)
		if err != nil {
			return nil, m.Matrix{}, err
		}
		V = newV

		if hasConverged(Ak, newAk) {
			n := A.Rows
			eigenValues := make([]float64, n)
			for i := range n {
				eigenValues[i] = newAk.Data[i*newAk.Cols+i]
			}
			return eigenValues, V, nil
		}

		Ak = newAk
	}

	// todo better error message here
	return nil, m.Matrix{}, fmt.Errorf("something went wrong")
}

func hasConverged(prev, curr m.Matrix) bool {
	tol := 10e-10
	for i := range prev.Rows {
		if math.Abs(prev.Data[i*prev.Cols+i]-curr.Data[i*curr.Cols+i]) > tol {
			return false
		}
	}
	return true
}
