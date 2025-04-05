package qr

import (
	"fmt"
	"math"
	"slices"

	m "face_recognition/matrix"
)

// todo: come up with better names for all the variables.

func householderVector(R m.Matrix, col, n int) []float64 {
	v := make([]float64, n-col)
	alpha := 0.0
	for i := col; i < n; i++ {
		v[i-col] = R.Data[i*n+col]
		alpha += R.Data[i*n+col] * R.Data[i*n+col]
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

func HouseholderMatrix(u []float64, k, n int) m.Matrix {
	Hk := m.Identity(n)
	for i := k; i < n; i++ {
		for j := k; j < n; j++ {
			Hk.Data[i*n+j] -= 2 * u[i-k] * u[j-k]
		}
	}
	return Hk
}

// https://www.youtube.com/watch?v=n0zDgkbFyQk
// todo: fix for non-square matrices
func QR_Householder(A m.Matrix) (m.Matrix, m.Matrix, error) {
	n := A.Rows

	Q := m.Identity(n)
	R := m.Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: slices.Clone(A.Data),
	}
	for currCol := range n - 1 {
		v := householderVector(R, currCol, n)
		u := normalizeVector(v, n)
		Hk := HouseholderMatrix(u, currCol, n)

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
		Data: slices.Clone(A.Data),
	}

	V := m.Identity(A.Rows)

	for range 200 {
		Q, R, err := QR_Householder(Ak)
		if err != nil {
			return nil, m.Matrix{}, err
		}

		newAk, err := m.Multiplication(R, Q)
		if err != nil {
			return nil, m.Matrix{}, err
		}

		newV, err := m.Multiplication(V, Q)
		if err != nil {
			return nil, m.Matrix{}, err
		}
		V = newV

		if HasConverged(Ak, newAk) {
			n := A.Rows
			eigenValues := make([]float64, n)
			for i := range n {
				eigenValues[i] = newAk.Data[i*newAk.Cols+i]
			}
			return eigenValues, V, nil
		}

		Ak = newAk
	}

	return nil, m.Matrix{}, fmt.Errorf("something went wrong")
}

func HasConverged(prev, curr m.Matrix) bool {
	tol := 10e-8
	for i := range prev.Rows {
		if math.Abs(prev.Data[i*prev.Cols+i]-curr.Data[i*curr.Cols+i]) > tol {
			return false
		}
	}
	return true
}
