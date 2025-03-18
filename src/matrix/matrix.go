package matrix

import "fmt"

type Matrix struct {
	Rows int
	Cols int
	Data []int
}

var (
	errIncorrectSize = fmt.Errorf("incorrect size")
)

func MultiplicationByScalar(A Matrix, scalar int) Matrix {
	result := Matrix{
		Rows: A.Rows,
		Cols: A.Cols,
		Data: make([]int, A.Rows*A.Cols),
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
		Data: make([]int, A.Rows*B.Cols),
	}

	for n := range A.Rows {
		rowOffset := n * A.Cols
		resultRowOffset := n * B.Cols
		for m := range B.Cols {
			sum := 0
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
		Data: make([]int, A.Rows*A.Cols),
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
		Data: make([]int, A.Rows*A.Cols),
	}

	for n := range A.Rows {
		for m := range A.Cols {
			result.Data[m*A.Rows+n] = A.Data[n*A.Cols+m]
		}
	}

	return result
}
