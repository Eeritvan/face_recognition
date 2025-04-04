package main

import (
	"fmt"
	"time"

	m "face_recognition/matrix"
	"face_recognition/qr"
)

func main() {
	A := m.Matrix{
		Rows: 3,
		Cols: 3,
		Data: []float64{4, 2, 20, 1, 2, 3, 5, 6, 7},
	}

	eigenValues, vectors, err := qr.QR_algorithm(A)

	fmt.Println(eigenValues)
	fmt.Println(vectors)
	fmt.Println(err)

	start := time.Now()

	B, C := m.SortEigenvectors(eigenValues, vectors)

	fmt.Println(time.Since(start))

	fmt.Println()
	fmt.Println(B)
	fmt.Println(C)
}
