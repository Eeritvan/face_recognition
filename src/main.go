package main

import (
	"fmt"

	"face_recognition/image"
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

	matrix, err := image.LoadPgmImage()
	if err != nil {
		panic(err)
	}

	flattened := image.FlattenImage(*matrix)

	fmt.Println(matrix)
	fmt.Println(flattened)
}
