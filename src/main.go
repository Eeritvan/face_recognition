package main

import (
	"fmt"

	m "face_recognition/matrix"
)

func main() {
	A := m.Matrix{
		Rows: 2,
		Cols: 3,
		Data: []int{1, 2, 3, 4, 5, 6},
	}

	B := m.Matrix{
		Rows: 3,
		Cols: 1,
		Data: []int{-1, -1, 0},
	}

	C := m.MultiplicationByScalar(A, 4)
	D := m.MultiplicationByScalar(B, 0)
	E, err := m.Multiplication(A, B)
	if err != nil {
		panic(err)
	}

	F, err := m.Addition(E, E)
	if err != nil {
		panic(err)
	}

	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
