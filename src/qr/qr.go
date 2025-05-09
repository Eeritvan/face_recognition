package qr

import (
	"math"
	"slices"

	m "face_recognition/matrix"
)

// note: gemini 2.5 pro helped me optimize and make the Householder reflection and QR way faster

// computes the householder reflection vector and its corresponding beta value for a given column of matrix R
// returns the beta value used in the Householder transformation
func calculateHouseholderVector(R m.Matrix, colIdx, size int, householderVector []float64) (float64, error) {
	norm := 0.0
	for i := range size {
		val := R.Data[(colIdx+i)*R.Cols+colIdx]
		householderVector[i] = val
		norm += val * val
	}

	norm = math.Sqrt(norm)

	if householderVector[0] >= 0 {
		householderVector[0] += norm
	} else {
		householderVector[0] -= norm
	}

	norm_sq := 0.0
	for i := range size {
		norm_sq += householderVector[i] * householderVector[i]
	}

	return 2.0 / norm_sq, nil
}

// applies the Householder transformation to update the R matrix
// using the computed Householdr vector and beta value
func updateRMatrix(R m.Matrix, colIdx, size int, householderVector []float64, beta float64) {
	for col := colIdx; col < R.Cols; col++ {
		dotProduct := 0.0
		for i := range size {
			dotProduct += householderVector[i] * R.Data[(colIdx+i)*R.Cols+col]
		}
		scaledProduct := dotProduct * beta
		for i := range size {
			R.Data[(colIdx+i)*R.Cols+col] -= scaledProduct * householderVector[i]
		}
	}
}

// applies the Householder transformation to update the Q matrix
// using the computed Householder vector and beta value
func updateQMatrix(Q m.Matrix, colIdx, size int, householderVector []float64, beta float64) {
	for rowIndex := range Q.Rows {
		dotProduct := 0.0
		for i := range size {
			dotProduct += Q.Data[rowIndex*Q.Rows+(colIdx+i)] * householderVector[i]
		}
		scaledProduct := dotProduct * beta
		for i := range size {
			Q.Data[rowIndex*Q.Rows+(colIdx+i)] -= scaledProduct * householderVector[i]
		}
	}
}

// performs QR decomposition using Householder reflections
// returns orthogonal matrix Q and upper triangular matrix R such that A = QR
func qr_Householder(A m.Matrix) (m.Matrix, m.Matrix, error) {
	rows := A.Rows
	cols := A.Cols

	Q := m.Identity(rows)
	R := m.Matrix{
		Rows: rows,
		Cols: cols,
		Data: slices.Clone(A.Data),
	}

	householderStorage := make([]float64, rows)

	for colIdx := range cols {
		size := rows - colIdx
		householderVector := householderStorage[:size]

		beta, err := calculateHouseholderVector(R, colIdx, size, householderVector)
		if err != nil {
			return m.Matrix{}, m.Matrix{}, err
		}

		updateRMatrix(R, colIdx, size, householderVector, beta)
		updateQMatrix(Q, colIdx, size, householderVector, beta)
	}
	return Q, R, nil
}

// QR algorithm to find eigenvalues and eigenvectors of a matrix
// returns a slice of eigenvalues and a matrix of the corresponding eigenvectors
func QR_algorithm(A m.Matrix) ([]float64, m.Matrix, error) {
	currentMatrix := m.Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: slices.Clone(A.Data),
	}

	eigenvectorMatrix := m.Identity(A.Rows)

	maxIter := 1000
	for range maxIter {
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

	n := A.Rows
	eigenValues := make([]float64, n)
	for i := range n {
		eigenValues[i] = currentMatrix.Data[i*currentMatrix.Cols+i]
	}
	return eigenValues, eigenvectorMatrix, nil
}

// checks if the QR algorithm has converged by comparing the diagonal elements of two consecutive iterations
func hasConverged(prev, curr m.Matrix) bool {
	tol := 10e-8
	for i := range prev.Rows {
		if math.Abs(prev.Data[i*prev.Cols+i]-curr.Data[i*curr.Cols+i]) > tol {
			return false
		}
	}
	return true
}
