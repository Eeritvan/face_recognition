package qr

import (
	"math"
	"slices"

	m "face_recognition/matrix"
)

func householderVector(R m.Matrix, col, n int) []float64 {
	reflVector := make([]float64, n-col)
	vectorNorm := 0.0
	for i := col; i < n; i++ {
		reflVector[i-col] = R.Data[i*n+col]
		vectorNorm += R.Data[i*n+col] * R.Data[i*n+col]
	}

	vectorNorm = math.Sqrt(vectorNorm)
	if reflVector[0] > 0 {
		vectorNorm = -vectorNorm
	}

	reflVector[0] = reflVector[0] + vectorNorm
	return reflVector
}

func normalizeVector(vector []float64, dimension int) []float64 {
	normVector := make([]float64, dimension)
	copy(normVector, vector)

	length := 0.0
	for j := range normVector {
		length += normVector[j] * normVector[j]
	}
	length = math.Sqrt(length)

	if math.Abs(length) < 1e-10 {
		return normVector
	}

	for j := range normVector {
		normVector[j] /= length
	}
	return normVector
}

func householderMatrix(u []float64, startIdx, dimension int) m.Matrix {
	reflMatrix := m.Identity(dimension)
	for i := startIdx; i < dimension; i++ {
		for j := startIdx; j < dimension; j++ {
			reflMatrix.Data[i*dimension+j] -= 2 * u[i-startIdx] * u[j-startIdx]
		}
	}
	return reflMatrix
}

func qr_Householder(A m.Matrix) (m.Matrix, m.Matrix, error) {
	dimension := A.Rows

	Q := m.Identity(dimension)
	R := m.Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: slices.Clone(A.Data),
	}
	for columnIdx := range dimension - 1 {
		reflVector := householderVector(R, columnIdx, dimension)
		unitVector := normalizeVector(reflVector, dimension)
		reflMatrix := householderMatrix(unitVector, columnIdx, dimension)

		newR, err := m.Multiplication(reflMatrix, R)
		if err != nil {
			return m.Matrix{}, m.Matrix{}, err
		}
		R = newR

		reflMatrix_T := m.Transpose(reflMatrix)
		newQ, err := m.Multiplication(Q, reflMatrix_T)
		if err != nil {
			return m.Matrix{}, m.Matrix{}, err
		}
		Q = newQ
	}
	return Q, R, nil
}

func QR_algorithm(A m.Matrix) ([]float64, m.Matrix, error) {
	currentMatrix := m.Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: slices.Clone(A.Data),
	}

	eigenvectorMatrix := m.Identity(A.Rows)

	for {
		Q, R, err := qr_Householder(currentMatrix)
		if err != nil {
			return nil, m.Matrix{}, err
		}

		nextMatrix, err := m.Multiplication(R, Q)
		if err != nil {
			return nil, m.Matrix{}, err
		}

		newEigenvectorMatrix, err := m.Multiplication(eigenvectorMatrix, Q)
		if err != nil {
			return nil, m.Matrix{}, err
		}
		eigenvectorMatrix = newEigenvectorMatrix

		if hasConverged(currentMatrix, nextMatrix) {
			n := A.Rows
			eigenValues := make([]float64, n)
			for i := range n {
				eigenValues[i] = nextMatrix.Data[i*nextMatrix.Cols+i]
			}
			return eigenValues, eigenvectorMatrix, nil
		}

		currentMatrix = nextMatrix
	}
}

func hasConverged(prev, curr m.Matrix) bool {
	tol := 10e-8
	for i := range prev.Rows {
		if math.Abs(prev.Data[i*prev.Cols+i]-curr.Data[i*curr.Cols+i]) > tol {
			return false
		}
	}
	return true
}
