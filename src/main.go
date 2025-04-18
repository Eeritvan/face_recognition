package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"

	"face_recognition/cli"
	r "face_recognition/recognition"
)

func generateRandomDataset(dataSets []int) []int {
	num1 := rand.Intn(40) + 1
	num2 := rand.Intn(40) + 1
	for num2 == num1 {
		num2 = rand.Intn(40) + 1
	}
	dataSets = append(dataSets, num1)
	dataSets = append(dataSets, num2)

	return dataSets
}

func generateRandomTestImage() []int {
	return []int{rand.Intn(40) + 1, rand.Intn(10) + 1}
}

func main() {
	k := 9
	imagesFromEachSet := 10
	timing := false
	interactiveMode := true
	args := os.Args[1:]
	var dataSets []int
	var testImage []int

	for i, flag := range args {
		switch flag {
		case "-h":
			cli.Help()
			os.Exit(0)
		case "-k":
			value, err := strconv.Atoi(args[i+1])
			if err != nil {
				panic(err)
			}
			k = value
			interactiveMode = false
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
			interactiveMode = false
		case "-t":
			timing = true
			interactiveMode = false
		case "-i":
			num, err := strconv.Atoi(args[i+1])
			if err != nil {
				panic(err)
			}
			if num < 1 || num > 10 {
				panic("-i failed")
			}

			imagesFromEachSet = num
			interactiveMode = false
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
			interactiveMode = false
		}
	}

	// generate random data to be used
	if len(dataSets) == 0 {
		dataSets = generateRandomDataset(dataSets)
	}

  // generate random test image to be used or validate given test image
	if len(testImage) == 0 {
		testImage = generateRandomTestImage()
	} else {
		if testImage[0] < 1 || testImage[0] > 40 {
			panic("incorrect set number")
		}
		if testImage[1] < 1 || testImage[1] > 10 {
			panic("incorrect image number")
		}
	}

  // decide to run in interactive mode or not
	if !interactiveMode {
		r.Run(timing, dataSets, testImage[:2], k, imagesFromEachSet)
	} else {
		cli.Interactive(dataSets, testImage, k, imagesFromEachSet, timing)
	}
}
