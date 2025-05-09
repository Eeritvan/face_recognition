package recognition

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"face_recognition/image"
	m "face_recognition/matrix"
	"face_recognition/qr"
)

// define possible errors
var (
	errInvalidKValue = fmt.Errorf("invalid -k value. It must be positive and less than the size of the training data")
)

// unit tests ignored since I/O testing wasn't required
// loads and flattens training images from the data directory for the specified sets and image count per set.
// Returns a slice of matrices containing the images
func loadTrainingFaces(dataSets []int, count int, rootDir string) ([]m.Matrix, error) {
	var faces []m.Matrix

	for _, set := range dataSets {
		for i := range count {
			matrix, err := image.LoadPgmImage(rootDir + "data/s" + strconv.Itoa(set) + "/" + strconv.Itoa(i+1) + ".pgm")
			if err != nil {
				return nil, err
			}
			flattened := image.FlattenImage(*matrix)
			faces = append(faces, flattened)
		}
	}

	return faces, nil
}

// calculates the eigenfaces and mean face from the training data
// Returns the eigenfaces matrix and the mean matrix
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

// projects all training faces into the eigenspace defined by eigenfaces and mean
// Returns a slice of projected face matrices
func projectFaces(faces []m.Matrix, eigenfaces, mean m.Matrix) ([]m.Matrix, error) {
	projectedFaces := make([]m.Matrix, len(faces))
	eigenfaces_T := m.Transpose(eigenfaces)

	for i, face := range faces {
		centeredFace, err := m.Subraction(face, mean)
		if err != nil {
			return nil, err
		}

		projected, err := m.Multiplication(eigenfaces_T, centeredFace)
		if err != nil {
			return nil, err
		}
		projectedFaces[i] = projected
	}

	return projectedFaces, nil
}

// unit tests ignored since I/O testing wasn't required
// loads and projects a test image into the eigenspace using the given eigenfaces and mean
// Returns the projected test image matrix
func loadTestImage(eigenfaces, mean m.Matrix, testImageParams []int, rootDir string) (m.Matrix, error) {
	testImage, err := image.LoadPgmImage(rootDir + "data/s" + strconv.Itoa(testImageParams[0]) + "/" + strconv.Itoa(testImageParams[1]) + ".pgm")
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

// findClosestMatch finds the closest training face to the projected test image
// Returns the index of the closest match and the minimum distance
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

// converts the minimum distance to a similarity percentage (0-100) using a function (1 - 0.04x) * 100
func getSimilarity(minDistance float64) float64 {
	return max((1-(minDistance*0.04))*100.0, 0)
}

// tests ignored. Not relevant for the course or the program
// measures time if timing flag is enabled
func timeExecution(name string, timing bool, fn func() error) error {
	if !timing {
		return fn()
	}

	start := time.Now()
	err := fn()
	if err != nil {
		return err
	}

	fmt.Printf("time to %s: %v\n", name, time.Since(start))
	return nil
}

// executes the full face recognition pipeline
// loads training images, computes eigenfaces, projects faces, loads and projects test image
// finds the closest match, and returns the match index and similarity or a possible error
func Run(timing bool, dataSets, testImage []int, k, imagesFromEachSet int, rootDir string) (int, float64, error) {
	if k < 0 || k > len(dataSets)*imagesFromEachSet {
		return 0, 0.0, errInvalidKValue
	}

	var (
		faces          []m.Matrix
		eigenfaces     m.Matrix
		mean           m.Matrix
		projectedFaces []m.Matrix
		projectedTest  m.Matrix
		matchIndex     int
		minDistance    float64
		similarity     float64
	)

	totalStart := time.Now()

	if err := timeExecution("process training images", timing, func() error {
		var err error
		faces, err = loadTrainingFaces(dataSets, imagesFromEachSet, rootDir)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("compute eigenfaces", timing, func() error {
		var err error
		eigenfaces, mean, err = computeEigenfaces(faces, k)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("project eigenfaces", timing, func() error {
		var err error
		projectedFaces, err = projectFaces(faces, eigenfaces, mean)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("load test image", timing, func() error {
		var err error
		projectedTest, err = loadTestImage(eigenfaces, mean, testImage, rootDir)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("find closest match", timing, func() error {
		matchIndex, minDistance = findClosestMatch(projectedTest, projectedFaces)
		similarity = getSimilarity(minDistance)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if timing {
		fmt.Print("Total time:", time.Since(totalStart), "\n\n")
	}

	return matchIndex, similarity, nil
}
