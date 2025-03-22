package image

import (
	"math"
	"testing"

	m "face_recognition/matrix"
)

const EPSILON = 1e-10

func TestLoadPgmImage(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		wantErr  error
	}{
		{
			name:     "Invalid filepath",
			filepath: "nonexistent.pgm",
			wantErr:  errFileOpening,
		},
		{
			name:     "Valid pgm file",
			filepath: "../data/s1/1.pgm",
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPgmImage(tt.filepath)
			if err != tt.wantErr {
				t.Errorf("LoadPgmImage(): %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlattenImage(t *testing.T) {
	tests := []struct {
		name  string
		image m.Matrix
		want  m.Matrix
	}{
		{
			name: "valid matrix with integers",
			image: m.Matrix{
				Rows: 2,
				Cols: 4,
				Data: []float64{4, 3, 2, 1, 8, 7, 6, 5},
			},
			want: m.Matrix{
				Rows: 8,
				Cols: 1,
				Data: []float64{4, 3, 2, 1, 8, 7, 6, 5},
			},
		},
		{
			name: "valid matrix with floating point numbers",
			image: m.Matrix{
				Rows: 2,
				Cols: 4,
				Data: []float64{0.1, 0.43, 3.42, 0, 0, 7, -0.21, 100.5},
			},
			want: m.Matrix{
				Rows: 8,
				Cols: 1,
				Data: []float64{0.1, 0.43, 3.42, 0, 0, 7, -0.21, 100.5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FlattenImage(tt.image)
			if result.Cols != tt.want.Cols {
				t.Errorf("FlattenImage(): returned incorrect amount of colums")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("FlattenImage(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && result.Data[i] != tt.want.Data[i] {
					t.Errorf("FlattenImage(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}

func TestMeanOfImages(t *testing.T) {
	tests := []struct {
		name    string
		faces   []m.Matrix
		want    m.Matrix
		wantErr error
	}{
		{
			name: "output is correct with valid matrices",
			faces: []m.Matrix{
				{
					Rows: 3,
					Cols: 3,
					Data: []float64{2, 3, 1, 2, 4, 0, 0, 0, 5},
				},
				{
					Rows: 3,
					Cols: 3,
					Data: []float64{7, 3, 5, 2, 1, 3, 2, 4, 0},
				},
				{
					Rows: 3,
					Cols: 3,
					Data: []float64{3, 3, 1, 5, 0, 0, 7, 2, 1},
				},
			},
			want: m.Matrix{
				Rows: 3,
				Cols: 3,
				Data: []float64{4, 3, 2.3333333333, 3, 1.6666666667, 1, 3, 2, 2},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MeanOfImages(tt.faces)
			if err != tt.wantErr {
				t.Errorf("MeanOfImages(): returned incorrect error: %v, want: %v", err, tt.wantErr)
			}
			if result.Cols != tt.want.Cols {
				t.Errorf("MeanOfImages(): returned incorrect amount of colums")
			}
			if result.Rows != tt.want.Rows {
				t.Errorf("MeanOfImages(): returned incorrect amount of rows")
			}
			for i := range result.Data {
				if i < len(tt.want.Data) && math.Abs(result.Data[i]-tt.want.Data[i]) > EPSILON {
					t.Errorf("MeanOfImages(): at index %d, got %f, want %f", i, result.Data[i], tt.want.Data[i])
				}
			}
		})
	}
}
