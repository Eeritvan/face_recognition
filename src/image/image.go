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

// define possible errors
var (
	errFileOpening   = fmt.Errorf("error opening file")
	errFileReading   = fmt.Errorf("error reading file")
	errWrongFaceSize = fmt.Errorf("size of the face was incorrect")
)

// reads a PGM image file from data directory and converts it to a matrix
// the function reads the dimensions, and pixel data
// returns a pointer to Matrix containing the image data
func LoadPgmImage(filepath string) (*m.Matrix, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errFileOpening
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip the "p5"" header
	_, err = reader.ReadString('\n')
	if err != nil {
		return nil, errFileReading
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, errFileReading
	}
	dimensions := strings.Fields(line)
	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return nil, errFileReading
	}
	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return nil, errFileReading
	}

	// skip color information
	_, err = reader.ReadString('\n')
	if err != nil {
		return nil, errFileReading
	}

	matrix := &m.Matrix{
		Rows: height,
		Cols: width,
		Data: make([]float64, width*height),
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
		matrix.Data[i] = float64(rawData[i])
	}

	return matrix, nil
}

// converts a 2D image matrix into a 1D column vector
func FlattenImage(image m.Matrix) m.Matrix {
	result := m.Matrix{
		Rows: image.Rows * image.Cols,
		Cols: 1,
		Data: image.Data,
	}

	return result
}

// calculates the average face from a slice of face matrices
// all input faces must have the same dimensions
// returns a matrix representing the mean face
func MeanOfImages(faces []m.Matrix) (m.Matrix, error) {
	result := m.Matrix{
		Rows: faces[0].Rows,
		Cols: faces[0].Cols,
		Data: make([]float64, faces[0].Rows*faces[0].Cols),
	}

	for _, face := range faces {
		if face.Cols != result.Cols || face.Rows != result.Rows {
			return m.Matrix{}, errWrongFaceSize
		}

		sum, err := m.Addition(result, face)
		if err != nil {
			return m.Matrix{}, err
		}

		result = sum
	}

	result = m.MultiplicationByScalar(result, 1/float64(len(faces)))

	return result, nil
}
