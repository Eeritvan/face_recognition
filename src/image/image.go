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
	errFileOpening   = fmt.Errorf("error opening file")
	errFileReading   = fmt.Errorf("error reading file")
	errWrongFaceSize = fmt.Errorf("Size of the face was incorrect")
)

func LoadPgmImage(filepath string) (*m.Matrix, error) {
	file, err := os.Open(filepath)
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

func FlattenImage(img m.Matrix) m.Matrix {
	result := m.Matrix{
		Rows: img.Rows * img.Cols,
		Cols: 1,
		Data: img.Data,
	}

	return result
}

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
