package main

import (
	"fmt"
	"strconv"

	"face_recognition/image"
	m "face_recognition/matrix"
)

func main() {
	var faces []m.Matrix

	for i := range 10 {
		matrix, err := image.LoadPgmImage("data/s1/" + strconv.Itoa(i+1) + ".pgm")
		if err != nil {
			panic(err)
		}
		flattened := image.FlattenImage(*matrix)
		faces = append(faces, flattened)
	}

	image, err := image.MeanOfImages(faces)
	if err != nil {
		panic(err)
	}

	fmt.Println(image)
}
