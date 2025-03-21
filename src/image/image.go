package image

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	m "face_recognition/matrix"
)

var (
	errFileOpening = fmt.Errorf("error opening file")
	errFileReading = fmt.Errorf("error reading file")
)

func LoadPgmImage() (*m.Matrix, error) {
	file, err := os.Open("data/s1/3.pgm")
	if err != nil {
		return nil, errFileOpening
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip the p5 header
	_, err = reader.ReadString('\n')
	if err != nil {
		return nil, errFileReading
	}

	dimLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, errFileReading
	}
	dimensions := strings.Fields(dimLine)
	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return nil, errFileReading
	}
	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return nil, errFileReading
	}

	// skip the max value line
	_, err = reader.ReadString('\n')
	if err != nil {
		return nil, errFileReading
	}

	matrix := &m.Matrix{
		Rows: height,
		Cols: width,
		Data: make([]int, width*height),
	}

	rawData := make([]byte, width*height)
	n, err := io.ReadFull(reader, rawData)
	if n != width*height {
		return nil, errFileReading
	}
	if err != nil {
		return nil, errFileReading
	}

	for i := range rawData {
		matrix.Data[i] = int(rawData[i])
	}

	return matrix, nil
}

func FlattenImage(img m.Matrix) m.Matrix {
	result := m.Matrix{
		Rows: img.Rows * img.Cols,
		Cols: 1,
		Data: img.Data,
	}

	return result
}
