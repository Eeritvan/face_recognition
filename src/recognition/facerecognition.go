package recognition

import (
	"math"
	"strconv"

	"face_recognition/image"
	m "face_recognition/matrix"
	"face_recognition/qr"
)

// todo: tests
func LoadTrainingFaces(dataSets []int, count int) ([]m.Matrix, error) {
	var faces []m.Matrix

	for _, set := range dataSets {
		for i := range count {
			matrix, err := image.LoadPgmImage("data/s" + strconv.Itoa(set) + "/" + strconv.Itoa(i+1) + ".pgm")
			if err != nil {
				return nil, err
			}
			flattened := image.FlattenImage(*matrix)
			faces = append(faces, flattened)
		}
	}

	return faces, nil
}

// todo: tests
func ComputeEigenfaces(faces []m.Matrix, k int) (m.Matrix, m.Matrix, error) {
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

	sortedVectors := m.SortEigenvectors(eigenvalues, eigenvectors)

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

// todo: tests
func ProjectFaces(faces []m.Matrix, eigenfaces, mean m.Matrix) ([]m.Matrix, error) {
	projectedFaces := make([]m.Matrix, len(faces))

	for i, face := range faces {
		centeredFace, err := m.Subraction(face, mean)
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

// todo: tests
func LoadTestImage(eigenfaces, mean m.Matrix) (m.Matrix, error) {
	testImage, err := image.LoadPgmImage("data/s15/10.pgm")
	if err != nil {
		return m.Matrix{}, err
	}
	flattenedTest := image.FlattenImage(*testImage)

	centeredTest, err := m.Subraction(flattenedTest, mean)
	if err != nil {
		return m.Matrix{}, err
	}

	projectedTest, err := m.Multiplication(m.Transpose(eigenfaces), centeredTest)
	if err != nil {
		return m.Matrix{}, err
	}

	return projectedTest, nil
}

// todo: tests
// todo: better function for approximations
// todo: fix the image possibly being in the set
// https://stats.stackexchange.com/questions/53068/euclidean-distance-score-and-similarity
func FindClosestMatch(projectedTest m.Matrix, projectedFaces []m.Matrix) (int, float64) {
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

	similarity := 1.0 / (1 + minDistance) * 100.0

	return matchIndex + 1, similarity
}
