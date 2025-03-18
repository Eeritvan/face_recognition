package main

import (
	"fmt"

	m "face_recognition/matrix"
)

func main() {
	A := m.Matrix{
		Rows: 3,
		Cols: 2,
		Data: []int{4, 2, 3, 2, 5, 9},
	}

	E := m.Transpose(A)

	fmt.Println(E)
}
