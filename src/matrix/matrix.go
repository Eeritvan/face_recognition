package matrix

import (
	"fmt"
	"sort"
)

type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

// define possible errors
var (
	errIncorrectSize = fmt.Errorf("incorrect size")
)

// multiplies all elements of given matrix by a scalar value
// returns a new matrix containing the result
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

// multiplies two matrices together
// returns a new matrix containing the result or an error if dimensions are incompatible
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

// adds two matrices together
// returns a new matrix containing the result or an error if dimensions are incompatible
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

// Subracts two matrices
// returns a new matrix containing the result or an error if dimensions are incompatible
func Subraction(A Matrix, B Matrix) (Matrix, error) {
	if A.Rows != B.Rows || A.Cols != B.Cols {
		return Matrix{}, errIncorrectSize
	}

	result := Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: make([]float64, A.Rows*A.Cols),
	}

	for n := range A.Data {
		result.Data[n] = A.Data[n] - B.Data[n]
	}

	return result, nil
}

// constructs by swapping rows and columns
// returns a new matrix containign the result
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

// constructs n * n identity matrix for given n value
// returns a new matrix containing the result
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

// computes A * AT of a given matrix
// returns a new matrix containing the result or an error
// however, the error shouldn't be possible since the multiplication is always valid
func Covariance(A Matrix) (Matrix, error) {
	AT := Transpose(A)

	result, err := Multiplication(AT, A)
	if err != nil {
		return Matrix{}, err
	}

	return result, nil
}

// computes the difference between each input matrix and the mean matrix
// returns a new matrix where each column represents the difference vector for a corresponding input
func DifferenceMatrix(vectors []Matrix, mean Matrix) (Matrix, error) {
	result := Matrix{
		Rows: vectors[0].Rows,
		Cols: len(vectors),
		Data: make([]float64, vectors[0].Rows*len(vectors)),
	}

	for i := range vectors {
		diff, err := Subraction(vectors[i], mean)
		if err != nil {
			return Matrix{}, err
		}
		for j := range vectors[i].Rows {
			result.Data[j*result.Cols+i] = diff.Data[j]
		}
	}

	return result, nil
}

// sorts the eigenvalues and eigenvectors in descending order and rearranges the columns
// of the eigenvectors matrix accordingly. Returns sorted eigenvectors
func SortEigenvectors(eigenvalues []float64, eigenvectors Matrix) Matrix {
	indices := make([]int, len(eigenvalues))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return eigenvalues[indices[i]] > eigenvalues[indices[j]]
	})

	result := Matrix{
		Rows: eigenvectors.Rows,
		Cols: eigenvectors.Cols,
		Data: make([]float64, len(eigenvectors.Data)),
	}

	for newCol, idx := range indices {
		for row := range eigenvectors.Rows {
			result.Data[row*eigenvectors.Cols+newCol] = eigenvectors.Data[row*eigenvectors.Cols+idx]
		}
	}

	return result
}
