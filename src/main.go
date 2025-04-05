package main

import (
	"fmt"

	m "face_recognition/matrix"
	"face_recognition/qr"
)

func main() {
	A := m.Matrix{
		Rows: 3,
		Cols: 3,
		Data: []float64{4, 2, 0, 1, 0, 3, 2, 6, 2},
	}

	Q, R, err := qr.QR_Householder(A)

	fmt.Println(Q)
	fmt.Println(R)
	fmt.Println(err)

	fmt.Println()

	B, _ := m.Multiplication(Q, R)
	// B, _ := m.Multiplication(m.Transpose(Q), Q)

	fmt.Println(B)

	// start := time.Now()

	// B, C := m.SortEigenvectors(eigenValues, vectors)

	// fmt.Println(time.Since(start))

	// fmt.Println()
	// fmt.Println(B)
	// fmt.Println(C)
}
