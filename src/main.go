package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"face_recognition/image"
	m "face_recognition/matrix"
	"face_recognition/qr"
)

func help() {
	fmt.Println(`
usage:
    ./face_recognition [options]

Options:
-h             shows this help message and terminates
-k <number>    sets the number of eigenfaces to use. Higher values will provide better accuracy. At the moment the default value is 9.
-t             display time taken to execute each step of the algorithm
-d <numbers>   Specify training datasets to use (e.g., 1,2,3)
	`)
	os.Exit(0)
}

// todo: consider own file(s) for all of the function below excluding the main function

// todo: tests
func loadTrainingFaces(dataSets []int, count int) ([]m.Matrix, error) {
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

// todo: tests
func projectFaces(faces []m.Matrix, eigenfaces, mean m.Matrix) ([]m.Matrix, error) {
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
func loadTestImage(eigenfaces, mean m.Matrix) (m.Matrix, error) {
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
// https://stats.stackexchange.com/questions/53068/euclidean-distance-score-and-similarity
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

	similarity := 1.0 / (1 + minDistance) * 100.0

	return matchIndex + 1, similarity
}

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

func main() {
	k := 9
	timing := false
	args := os.Args[1:]
	var dataSets []int

	for i, flag := range args {
		switch flag {
		case "-h":
			help()
		case "-k":
			// todo: check that value is valid. positive and less than the
			// 		 size of the training data
			// todo: better error message
			value, err := strconv.Atoi(args[i+1])
			if err != nil {
				panic(err)
			}
			k = value
		case "-d":
			j := i + 1
			for j < len(args) && !strings.HasPrefix(args[j], "-") {
				value, err := strconv.Atoi(args[j])
				// todo: better error message
				if err != nil {
					panic(err)
				}
				dataSets = append(dataSets, value)
				j++
			}
		case "-t":
			timing = true
		}
	}

	if len(dataSets) > 4 {
		fmt.Print("Loading many datasets may be super slow. Continue? (Y/n) ")

		var response string
		fmt.Scan(&response)

		response = strings.ToLower(response)
		if response == "n" || response == "no" {
			os.Exit(0)
		}
	}

	var faces []m.Matrix
	var eigenfaces, mean m.Matrix
	var projectedFaces []m.Matrix
	var projectedTest m.Matrix
	var matchIndex int
	var similarity float64

	totalStart := time.Now()

	if err := timeExecution("process training images", timing, func() error {
		var err error
		faces, err = loadTrainingFaces(dataSets, 10)
		return err
	}); err != nil {
		panic(err)
	}

	if err := timeExecution("compute eigenfaces", timing, func() error {
		var err error
		eigenfaces, mean, err = computeEigenfaces(faces, k)
		return err
	}); err != nil {
		panic(err)
	}

	if err := timeExecution("project eigenfaces", timing, func() error {
		var err error
		projectedFaces, err = projectFaces(faces, eigenfaces, mean)
		return err
	}); err != nil {
		panic(err)
	}

	if err := timeExecution("load test image", timing, func() error {
		var err error
		projectedTest, err = loadTestImage(eigenfaces, mean)
		return err
	}); err != nil {
		panic(err)
	}

	if err := timeExecution("find closest match", timing, func() error {
		matchIndex, similarity = findClosestMatch(projectedTest, projectedFaces)
		return nil
	}); err != nil {
		panic(err)
	}

	if timing {
		fmt.Println("Total time:", time.Since(totalStart))
	}

	fmt.Printf("closest match with: %v\nsimilarity: %.1f %% \n", matchIndex, similarity)
}
