package main

import (
	"fmt"
	"log"
	"math/rand"
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

options:
    -h             shows this help message and terminates
    -k <num>       sets the number of eigenfaces to use. Higher values will provide better accuracy. At the moment the default value is 9.
    -t             display time taken to execute each step of the algorithm
    -s <num num>   specify the test image to be used. Given as tuple <number number> where the first number is the set being used (1-40) and the second number which image is used (1-10)		   
    -i             specify how many images are loaded from each set.
    -d <num>       specify training datasets to use (e.g., 1,2,3). By default two random sets are used.

examples:
    ./face_recognition                     # Run with default settings (k=9)
    ./face_r"os/exec"ecognition -k 15      # Use 15 eigenfaces
    ./face_recognition -d 1,2,3            # Use datasets 1, 2 and 3
    ./face_recognition -k 20 -d 1 2 3 4    # Use 20 eigenfaces with datasets 1-4
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

func run(timing bool, dataSets, testImage []int, k, imagesFromEachSet int) {
	var faces []m.Matrix
	var eigenfaces, mean m.Matrix
	var projectedFaces []m.Matrix
	var projectedTest m.Matrix
	var matchIndex int
	var minDistance float64
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
		projectedTest, err = r.LoadTestImage(eigenfaces, mean, testImage)
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := timeExecution("find closest match", timing, func() error {
		matchIndex, minDistance = r.FindClosestMatch(projectedTest, projectedFaces)
		similarity = r.GetSimilarity(minDistance)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if timing {
		fmt.Print("Total time:", time.Since(totalStart), "\n\n")
	}

	fmt.Println("Data used:", dataSets)
	fmt.Println("Test Image: set", testImage[0], "| image", testImage[1])
	fmt.Println("closest match with: set", matchIndex/10, "| image", matchIndex%10+1)
	fmt.Printf("similarity: %.1f%% \n", similarity)
}

// todo: refactor
func main() {
	k := 9
	imagesFromEachSet := 10
	timing := false
	interactive := true
	args := os.Args[1:]
	var dataSets []int
	var testImage []int

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
			interactive = false
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
			interactive = false
		case "-t":
			timing = true
			interactive = false
		case "-i":
			num, err := strconv.Atoi(args[i+1])
			if err != nil {
				panic(err)
			}
			if num < 1 || num > 10 {
				// todo: proper erroe message
				panic("failed")
			}

			imagesFromEachSet = num
			interactive = false
		case "-s":
			j := i + 1
			for j < len(args) && !strings.HasPrefix(args[j], "-") {
				value, err := strconv.Atoi(args[j])
				if err != nil {
					panic(err)
				}
				testImage = append(testImage, value)
				j++
			}
			interactive = false
		}
	}

	// validate test image input

	if len(dataSets) == 0 {
		num1 := rand.Intn(40) + 1
		num2 := rand.Intn(40) + 1
		for num2 == num1 {
			num2 = rand.Intn(40) + 1
		}
		dataSets = append(dataSets, num1)
		dataSets = append(dataSets, num2)
	}

	if len(testImage) == 0 {
		testImage = append(testImage, rand.Intn(40)+1)
		testImage = append(testImage, rand.Intn(10)+1)
	} else {
		if testImage[0] < 1 || testImage[0] > 40 {
			panic("incorrect set number")
		}
		if testImage[1] < 1 || testImage[1] > 10 {
			panic("incorrect image number")
		}
	}

	if !interactive {
		if k < 0 || k > len(dataSets)*imagesFromEachSet {
			log.Fatal("invalid -k value. It must be positive and less than the size of the training data")
		}

		if len(dataSets) > 4 {
			fmt.Print("loading many datasets may be super slow. Continue? (Y/n) ")

			var response string
			if _, err := fmt.Scan(&response); err != nil {
				panic(err)
			}

			response = strings.ToLower(response)
			if response == "n" || response == "no" {
				os.Exit(0)
			}
		}

		run(timing, dataSets, testImage[:2], k, imagesFromEachSet)
	} else {
		for {
			fmt.Println("\ncurrent settings:")
			fmt.Println("-----------------------------------")
			fmt.Println("  eigenfaces (k):        ", k)
			fmt.Println("  data sets (d):         ", dataSets)
			fmt.Println("  data sets (s):         ", testImage)
			fmt.Println("  images per set:        ", imagesFromEachSet)
			fmt.Println("  time algorithm steps:  ", timing)
			fmt.Println("-----------------------------------")
			fmt.Println("\navailable commands:")
			fmt.Println("  k    - change number of eigenfaces")
			fmt.Println("  d    - select data sets")
			fmt.Println("  s    - select test image")
			fmt.Println("  t    - toggle timing")
			fmt.Println("  i    - specify amount of images to use from each set")
			fmt.Println("  run  - run the algoritm")
			fmt.Println("  quit - exit")

			fmt.Print("\nenter command: ")
			var cmd string
			_, err := fmt.Scan(&cmd)
			if err != nil {
				// todo: proper error message
				panic(err)
			}

			switch cmd {
			case "k":
				fmt.Print("  enter number of eigenfaces to use: ")
				if _, err := fmt.Scan(&k); err != nil {
					// todo: proper error message
					panic(err)
				}
			case "t":
				timing = !timing
				fmt.Print("timing set to: ", timing)
			case "d":
				fmt.Print("  enter dataset numbers (1-40) (0 to break): ")

				var newDataSets []int
				for {
					var val int
					if _, err := fmt.Scan(&val); err != nil {
						// todo: proper error message
						panic(err)
					}
					if val == 0 {
						break
					}
					if val < 1 || val > 40 {
						fmt.Println("  invalid number")
						continue
					}

					newDataSets = append(newDataSets, val)
				}
				dataSets = newDataSets
			case "s":
				var newTestImage []int

				fmt.Print("  enter set number (1-40) ")

				for {
					var set int
					if _, err := fmt.Scan(&set); err != nil {
						// todo: proper error message
						panic(err)
					}
					if set >= 1 && set <= 40 {
						newTestImage = append(newTestImage, set)
						break
					}
					fmt.Println("  invalid set number")
				}

				fmt.Print("  enter image number (1-10) ")
				for {
					var num int
					if _, err := fmt.Scan(&num); err != nil {
						// todo: proper error message
						panic(err)
					}
					if num >= 1 && num <= 10 {
						newTestImage = append(newTestImage, num)
						break
					}
					fmt.Println("  invalid number")
				}
				testImage = newTestImage
			case "i":
				fmt.Print("  enter amount of images to use (1-10): ")

				for {
					var num int
					if _, err := fmt.Scan(&num); err != nil {
						// todo: proper error message
						panic(err)
					}
					if num >= 1 && num <= 10 {
						imagesFromEachSet = num
						break
					}
					fmt.Println("  invalid number")
				}
			case "run":
				if k < 0 || k > len(dataSets)*imagesFromEachSet {
					fmt.Println("  invalid -k value. It must be positive and less than the size of the training data")
					continue
				}

				if len(dataSets) > 4 {
					fmt.Print("  loading many datasets may be super slow. I recommend using at most 4 sets. Continue? (Y/n) ")

					var response string
					if _, err := fmt.Scan(&response); err != nil {
						// todo: proper error message
						panic(err)
					}

					response = strings.ToLower(response)
					if response == "n" || response == "no" {
						os.Exit(0)
					}
				}

				fmt.Print("\n###############################\n\n")
				run(timing, dataSets, testImage[:2], k, imagesFromEachSet)
				fmt.Print("\n###############################\n")
			case "quit":
				os.Exit(0)
			}
		}
	}
}
