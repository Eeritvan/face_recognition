package matrix

import (
	"fmt"
)

type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

var (
	errIncorrectSize = fmt.Errorf("incorrect size")
)

func MultiplicationByScalar(A Matrix, scalar float64) Matrix {
	result := Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: make([]float64, A.Rows*A.Cols),
	}

	for i, num := range A.Data {
		result.Data[i] = num * scalar
	}

	return result
}

func Multiplication(A Matrix, B Matrix) (Matrix, error) {
	if A.Cols != B.Rows {
		return Matrix{}, errIncorrectSize
	}

	result := Matrix{
		Rows: A.Rows,
		Cols: B.Cols,
		Data: make([]float64, A.Rows*B.Cols),
	}

	for n := range A.Rows {
		rowOffset := n * A.Cols
		resultRowOffset := n * B.Cols
		for m := range B.Cols {
			var sum float64 = 0
			for k := range A.Cols {
				sum += A.Data[rowOffset+k] * B.Data[k*B.Cols+m]
			}
			result.Data[resultRowOffset+m] = sum
		}
	}

	return result, nil
}

func Addition(A Matrix, B Matrix) (Matrix, error) {
	if A.Rows != B.Rows || A.Cols != B.Cols {
		return Matrix{}, errIncorrectSize
	}

	result := Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: make([]float64, A.Rows*A.Cols),
	}

	for n := range A.Data {
		result.Data[n] = A.Data[n] + B.Data[n]
	}

	return result, nil
}

func Transpose(A Matrix) Matrix {
	result := Matrix{
		Rows: A.Cols,
		Cols: A.Rows,
		Data: make([]float64, A.Rows*A.Cols),
	}

	for n := range A.Rows {
		for m := range A.Cols {
			result.Data[m*A.Rows+n] = A.Data[n*A.Cols+m]
		}
	}

	return result
}

func Identity(n int) Matrix {
	result := Matrix{
		Rows: n,
		Cols: n,
		Data: make([]float64, n*n),
	}

	for i := range n {
		result.Data[i*n+i] = 1
	}

	return result
}

func DifferenceMatrix(vectors []Matrix, mean Matrix) (Matrix, error) {
	result := Matrix{
		Rows: vectors[0].Rows,
		Cols: len(vectors),
		Data: make([]float64, vectors[0].Rows*len(vectors)),
	}

	negativeMean := MultiplicationByScalar(mean, -1)

	for i := range vectors {
		diff, err := Addition(vectors[i], negativeMean)
		if err != nil {
			return Matrix{}, err
		}
		for j := range vectors[i].Rows {
			result.Data[j*result.Cols+i] = diff.Data[j]
		}
	}

	return result, nil
}

func Covariance(A Matrix) (Matrix, error) {
	AT := Transpose(A)

	result, err := Multiplication(AT, A)
	if err != nil {
		return Matrix{}, err
	}

	return result, nil
}

// todo: tests
// todo: explore possibility for faster sorting. Just relatively simple solution for now.
func SortEigenvectors(eigenvalues []float64, eigenvectors Matrix) ([]float64, Matrix) {
	indices := make([]int, len(eigenvalues))
	for i := range indices {
		indices[i] = i
	}

	for i := range len(indices) - 1 {
		for j := i + 1; j < len(indices); j++ {
			if eigenvalues[indices[i]] < eigenvalues[indices[j]] {
				indices[i], indices[j] = indices[j], indices[i]
				// eigenvalues[i], eigenvalues[j] = eigenvalues[j], eigenvalues[i]
			}
		}
	}

	sortedValues := make([]float64, len(eigenvalues))
	result := Matrix{
		Rows: eigenvectors.Rows,
		Cols: eigenvectors.Cols,
		Data: make([]float64, len(eigenvectors.Data)),
	}

	for value, idx := range indices {
		sortedValues[value] = eigenvalues[idx]
		for j := range eigenvectors.Rows {
			result.Data[j*eigenvectors.Cols+value] = eigenvectors.Data[j*eigenvectors.Cols+idx]
		}
	}

	return sortedValues, result
}
