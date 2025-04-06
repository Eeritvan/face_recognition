package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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
	`)
	os.Exit(0)
}

// todo: consider own file(s) for all of the function below excluding the main function

// todo: tests
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
	testImage, err := image.LoadPgmImage("data/s20/1.pgm")
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

// todo: explore ways to make the timing things clutter the code less
func main() {
	k := 9
	timing := false
	args := os.Args[1:]

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
		case "-t":
			timing = true
		}
	}

	var start time.Time
	if timing {
		start = time.Now()
	}

	faces, err := loadTrainingFaces(10)
	if err != nil {
		panic(err)
	}

	if timing {
		fmt.Println("time to process training images:", time.Since(start))
	}

	var eigenfacesTime time.Time
	if timing {
		eigenfacesTime = time.Now()
	}

	eigenfaces, mean, err := computeEigenfaces(faces, k)
	if err != nil {
		panic(err)
	}

	if timing {
		fmt.Println("time to compute eigenfaces:", time.Since(eigenfacesTime))
	}

	var eigenfaceProjectionTime time.Time
	if timing {
		eigenfaceProjectionTime = time.Now()
	}

	projectedFaces, err := projectFaces(faces, eigenfaces, mean)
	if err != nil {
		panic(err)
	}

	if timing {
		fmt.Println("time to project eigenfaces:", time.Since(eigenfaceProjectionTime))
	}

	projectedTest, err := loadTestImage(eigenfaces, mean)
	if err != nil {
		panic(err)
	}

	var matchTime time.Time
	if timing {
		matchTime = time.Now()
	}

	matchIndex, minDistance := findClosestMatch(projectedTest, projectedFaces)

	if timing {
		fmt.Println("time to find closest match:", time.Since(matchTime))
		fmt.Println("Total time:", time.Since(start))
	}

	fmt.Println(matchIndex+1, minDistance)
}
