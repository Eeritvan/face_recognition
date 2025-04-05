package main

import (
	"fmt"
	"math"
	"strconv"

	"face_recognition/image"
	m "face_recognition/matrix"
	"face_recognition/qr"
)

func loadTrainingFaces(count int) ([]m.Matrix, error) {
	var faces []m.Matrix

	for i := range count {
		matrix, err := image.LoadPgmImage("data/s1/" + strconv.Itoa(i+1) + ".pgm")
		if err != nil {
			return nil, err
		}
		flattened := image.FlattenImage(*matrix)
		faces = append(faces, flattened)
	}

	return faces, nil
}

func computeEigenfaces(faces []m.Matrix, k int) (m.Matrix, m.Matrix, error) {
	mean, err := image.MeanOfImages(faces)
	if err != nil {
		return m.Matrix{}, m.Matrix{}, err
	}

	diffMatrix, err := m.DifferenceMatrix(faces, mean)
	if err != nil {
		return m.Matrix{}, m.Matrix{}, err
	}

	covariance, err := m.Covariance(diffMatrix)
	if err != nil {
		return m.Matrix{}, m.Matrix{}, err
	}

	eigenvalues, eigenvectors, err := qr.QR_algorithm(covariance)
	if err != nil {
		return m.Matrix{}, m.Matrix{}, err
	}

	_, sortedVectors := m.SortEigenvectors(eigenvalues, eigenvectors)

	eigenfaces := m.Matrix{
		Rows: diffMatrix.Rows,
		Cols: k,
		Data: make([]float64, diffMatrix.Rows*k),
	}
	for i := range sortedVectors.Rows {
		for j := range k {
			eigenfaces.Data[i*k+j] = sortedVectors.Data[i*sortedVectors.Cols+j]
		}
	}

	return eigenfaces, mean, nil
}

func projectFaces(faces []m.Matrix, eigenfaces, mean m.Matrix) ([]m.Matrix, error) {
	projectedFaces := make([]m.Matrix, len(faces))

	for i, face := range faces {
		centeredFace, err := m.Addition(face, m.MultiplicationByScalar(mean, -1))
		if err != nil {
			return nil, err
		}

		projected, err := m.Multiplication(m.Transpose(eigenfaces), centeredFace)
		if err != nil {
			return nil, err
		}
		projectedFaces[i] = projected
	}

	return projectedFaces, nil
}

func loadTestImage(eigenfaces, mean m.Matrix) (m.Matrix, error) {
	testImage, err := image.LoadPgmImage("data/s20/10.pgm")
	if err != nil {
		return m.Matrix{}, err
	}
	flattenedTest := image.FlattenImage(*testImage)

	centeredTest, err := m.Addition(flattenedTest, m.MultiplicationByScalar(mean, -1))
	if err != nil {
		return m.Matrix{}, err
	}

	projectedTest, err := m.Multiplication(m.Transpose(eigenfaces), centeredTest)
	if err != nil {
		return m.Matrix{}, err
	}

	return projectedTest, nil
}

func findClosestMatch(projectedTest m.Matrix, projectedFaces []m.Matrix) (int, float64) {
	var minDistance float64 = math.Inf(1)
	matchIndex := -1

	for i, projectedFace := range projectedFaces {
		var distance float64
		for j := range projectedTest.Data {
			diff := projectedTest.Data[j] - projectedFace.Data[j]
			distance += diff * diff
		}
		distance = math.Sqrt(distance)

		if distance < minDistance {
			minDistance = distance
			matchIndex = i
		}
	}

	return matchIndex + 1, minDistance
}

func main() {
	faces, err := loadTrainingFaces(9)
	if err != nil {
		panic(err)
	}

	eigenfaces, mean, err := computeEigenfaces(faces, 9)
	if err != nil {
		panic(err)
	}

	projectedFaces, err := projectFaces(faces, eigenfaces, mean)
	if err != nil {
		panic(err)
	}

	projectedTest, err := loadTestImage(eigenfaces, mean)
	if err != nil {
		panic(err)
	}

	matchIndex, minDistance := findClosestMatch(projectedTest, projectedFaces)

	fmt.Println(matchIndex+1, minDistance)
}
