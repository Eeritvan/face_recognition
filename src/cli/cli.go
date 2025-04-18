package cli

import (
	"fmt"
	"os"

	r "face_recognition/recognition"
)

func Help() {
	fmt.Println(`
usage:
    ./face_recognition [options]

	./face_recognition     # without any options this will use interactive cli mode

options:
    -h             shows this help message and terminates
    -k <num>       sets the number of eigenfaces to use. The default value is 9. This must be less than the amount of training images.
    -t             display time taken to execute each step of the algorithm
    -s <num num>   specify the test	image to be used. Given as tuple <number number> where the first number is the set being used (1-40) and the second number which image is used (1-10)		   
    -i             specify how many images are loaded from each set. Each set has 10 images.
    -d <num>       specify training datasets to use (e.g., 1 2 3). By default two random sets are used.

note 1: Using too high a value for k can actually reduce accuracy due to overfitting and noise. Lower k values often generalize better.
note 2: Using too many training images / sets will lead to slow performance. I recommend using less than 4 full data sets / 40 images in total.
	
examples:
    ./face_recognition                     # Run interactive mode 
    ./face_recognition -k 15   	           # Use 15 eigenfaces
    ./face_recognition -d 1 2 3            # Use datasets 1, 2 and 3
    ./face_recognition -s 5 5              # Use set 5 image 5 as the test image
    ./face_recognition -k 15 -d 1 2 3 4    # Use 20 eigenfaces with datasets 1-4

	`)
}

func Interactive(dataSets, testImage []int, k, imagesFromEachSet int, timing bool) {
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
		fmt.Println("  quit - terminate program")

		fmt.Print("\nenter command: ")
		var cmd string
		if _, err := fmt.Scan(&cmd); err != nil {
			panic(err)
		}

		switch cmd {
		case "k":
			fmt.Print("  enter number of eigenfaces to use: ")
			if _, err := fmt.Scan(&k); err != nil {
				panic(err)
			}
		case "t":
			timing = !timing
			fmt.Print("timing set to: ", timing)
		case "d":
			fmt.Print("  enter datasets to use (1-40) (0 to break): ")

			var newDataSets []int
			for {
				var val int
				if _, err := fmt.Scan(&val); err != nil {
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
					panic(err)
				}
				if num >= 1 && num <= 10 {
					imagesFromEachSet = num
					break
				}
				fmt.Println("  invalid number")
			}
		case "run":
			fmt.Print("\n###############################\n\n")
			matchIndex, similarity, err := r.Run(timing, dataSets, testImage[:2], k, imagesFromEachSet, "./")
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Data used:", dataSets)
			fmt.Println("Test Image: set", testImage[0], "| image", testImage[1])
			fmt.Println("closest match with: set", matchIndex/10, "| image", matchIndex%10+1)
			fmt.Printf("similarity: %.1f%% \n", similarity)

			fmt.Print("\n###############################\n")
		case "quit":
			os.Exit(0)
		}
	}
}
