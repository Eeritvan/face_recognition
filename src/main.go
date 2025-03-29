package main

import (
	"face_recognition/qr"
	"fmt"

	// "strconv"
	// "face_recognition/image"
	m "face_recognition/matrix"
)

func main() {
	// var faces []m.Matrix

	// for i := range 10 {
	// 	matrix, err := image.LoadPgmImage("data/s1/" + strconv.Itoa(i+1) + ".pgm")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	flattened := image.FlattenImage(*matrix)
	// 	faces = append(faces, flattened)
	// }

	// mean, err := image.MeanOfImages(faces)
	// fmt.Println(mean)
	// if err != nil {
	// 	panic(err)
	// }

	// diffMatrix, err := m.DifferenceMatrix(faces, mean)
	// if err != nil {
	// 	panic(err)
	// }

	// cov, err := m.Covariance(diffMatrix)
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(cov)

	// A, B, err := qr.Householder(cov)
	// if err != nil {
	// 	panic(err)
	// }

	// B, err := m.Householder(A)
	// if err != nil {
	// 	panic(err)
	// }

	// M := m.Matrix{
	// 	Rows: 3,
	// 	Cols: 3,
	// 	Data: []float64{4, 2, 2, 3, 3, 1, 0, 5, 4},
	// }

	M := m.Matrix{
		Rows: 3,
		Cols: 3,
		Data: []float64{2, 0, 0, 0, 4, 5, 0, 4, 3},
	}

	eigenValues, Q, _ := qr.QR_algorithm(M)

	// fmt.Println(eigenValues, Q)

	sortedValues, sortedVectors := m.SortEigenvectors(eigenValues, Q)

	fmt.Println(sortedValues)
	fmt.Println(sortedVectors)

	// fmt.Println(Q)

	// fmt.Println("OG", M)
	// fmt.Println("Q ", Q)
	// fmt.Println("R ", R)

	// fmt.Println()

	// Mult, _ := m.Multiplication(Q, R)
	// fmt.Println("RE", Mult)

	// fmt.Println()

	// fmt.Println(m.Multiplication(Q, m.Transpose(Q)))
	// fmt.Println(m.Multiplication(m.Transpose(Q), Q))
}
