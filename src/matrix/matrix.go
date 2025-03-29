package matrix

import "fmt"

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
