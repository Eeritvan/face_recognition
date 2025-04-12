package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	m "face_recognition/matrix"
	r "face_recognition/recognition"
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
	imagesFromEachSet := 10
	timing := false
	args := os.Args[1:]
	var dataSets []int

	for i, flag := range args {
		switch flag {
		case "-h":
			help()
			os.Exit(0)
		case "-k":
			value, err := strconv.Atoi(args[i+1])
			if err != nil {
				panic(err)
			}
			k = value
		case "-d":
			j := i + 1
			for j < len(args) && !strings.HasPrefix(args[j], "-") {
				value, err := strconv.Atoi(args[j])
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

	if k < 0 || k > len(dataSets)*imagesFromEachSet {
		log.Fatal("invalid -k value. It must be positive and less than the size of the training data")
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
		faces, err = r.LoadTrainingFaces(dataSets, imagesFromEachSet)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("compute eigenfaces", timing, func() error {
		var err error
		eigenfaces, mean, err = r.ComputeEigenfaces(faces, k)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("project eigenfaces", timing, func() error {
		var err error
		projectedFaces, err = r.ProjectFaces(faces, eigenfaces, mean)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("load test image", timing, func() error {
		var err error
		projectedTest, err = r.LoadTestImage(eigenfaces, mean)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("find closest match", timing, func() error {
		matchIndex, similarity = r.FindClosestMatch(projectedTest, projectedFaces)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if timing {
		fmt.Println("Total time:", time.Since(totalStart))
	}

	fmt.Printf("closest match with: %v\nsimilarity: %.1f %% \n", matchIndex, similarity)
}
