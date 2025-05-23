package recognition

import (
	"math"
	"testing"

	m "face_recognition/matrix"
)

const EPSILON = 1e-6

func TestComputeEigenfaces(t *testing.T) {
	tests := []struct {
		name           string
		faces          []m.Matrix
		k              int
		wantEigenfaces m.Matrix
		wantMean       m.Matrix
		wantErr        error
	}{
		{
			name: "output is correct with valid inputs",
			faces: []m.Matrix{
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{4, 5, 1, 2},
				},
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{4, 1, 2, 9},
				},
			},
			k: 2,
			wantEigenfaces: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{
					0.707106, -0.707106,
					-0.707106, -0.707106,
				},
			},
			wantMean: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{4, 3, 1.5, 5.5},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eigenfaces, mean, err := computeEigenfaces(tt.faces, tt.k)
			if err != tt.wantErr {
				t.Errorf("ComputeEigenfaces(): returned wrong error: %v", err)
			}

			// checking eigenfaces
			if eigenfaces.Rows != tt.wantEigenfaces.Rows {
				t.Errorf("ComputeEigenfaces(): eigenfaces returned incorrect amount of rows")
			}
			if eigenfaces.Cols != tt.wantEigenfaces.Cols {
				t.Errorf("ComputeEigenfaces(): eigenfaces returned incorrect amount of cols")
			}
			for i := range eigenfaces.Data {
				if math.Abs(eigenfaces.Data[i]-tt.wantEigenfaces.Data[i]) > EPSILON {
					t.Errorf("ComputeEigenfaces(): at index %d: got %f, want %f", i, eigenfaces.Data[i], tt.wantEigenfaces.Data[i])
				}
			}

			// checking mean
			if mean.Rows != tt.wantMean.Rows {
				t.Errorf("ComputeEigenfaces(): mean returned incorrect amount of rows")
			}
			if mean.Cols != tt.wantMean.Cols {
				t.Errorf("ComputeEigenfaces(): mean returned incorrect amount of cols")
			}
			for i := range mean.Data {
				if math.Abs(mean.Data[i]-tt.wantMean.Data[i]) > EPSILON {
					t.Errorf("ComputeEigenfaces(): at index %d: got %f, want %f", i, mean.Data[i], tt.wantMean.Data[i])
				}
			}
		})
	}
}

func TestProjectFaces(t *testing.T) {
	tests := []struct {
		name               string
		faces              []m.Matrix
		eigenfaces         m.Matrix
		mean               m.Matrix
		wantProjectedFaces []m.Matrix
		wantErr            error
	}{
		{
			name: "output is correct with valid inputs",
			faces: []m.Matrix{
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{4, 5, 1, 2},
				},
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{4, 1, 2, 9},
				},
			},
			eigenfaces: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{
					0.7071067, -0.7071067,
					-0.7071067, -0.7071067,
				},
			},
			mean: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{4, 3, 1.5, 5.5},
			},
			wantProjectedFaces: []m.Matrix{
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{
						0.353553, 3.889087,
						0.353553, 1.060660,
					},
				},
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{
						-0.353553, -3.889087,
						-0.353553, -1.060660,
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projectedFaces, err := projectFaces(tt.faces, tt.eigenfaces, tt.mean)
			if err != tt.wantErr {
				t.Errorf("ProjectFaces(): returned wrong error: %v", err)
			}

			if len(projectedFaces) != len(tt.wantProjectedFaces) {
				t.Errorf("ProjectFaces(): the length was incorrect, want %v, got %v", len(projectedFaces), len(tt.wantProjectedFaces))
			}

			for idx := range projectedFaces {
				if projectedFaces[idx].Rows != tt.wantProjectedFaces[idx].Rows {
					t.Errorf("ProjectFaces(): returned incorrect amount of rows")
				}
				if projectedFaces[idx].Cols != tt.wantProjectedFaces[idx].Cols {
					t.Errorf("ProjectFaces(): returned incorrect amount of cols")
				}
				for i := range projectedFaces[idx].Data {
					if math.Abs(projectedFaces[idx].Data[i]-tt.wantProjectedFaces[idx].Data[i]) > EPSILON {
						t.Errorf("ProjectFaces(): at index %d: got %f, want %f", i, projectedFaces[idx].Data[i], tt.wantProjectedFaces[idx].Data[i])
					}
				}
			}
		})
	}
}

func TestFindClosestMatch(t *testing.T) {
	tests := []struct {
		name            string
		projectedTest   m.Matrix
		projectedFaces  []m.Matrix
		wantMatchIndex  int
		wantMinDistance float64
	}{
		{
			name: "output is correct with valid inputs",
			projectedTest: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{1, 2, 3, 4},
			},
			projectedFaces: []m.Matrix{
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{
						0.353553, 3.889087,
						0.353553, 1.060660,
					},
				},
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{
						-0.353553, -3.889087,
						-0.353553, -1.060660,
					},
				},
			},
			wantMatchIndex:  1,
			wantMinDistance: 4.430569,
		},
		{
			name: "output is 0 when the image is already in the data",
			projectedTest: m.Matrix{
				Rows: 2,
				Cols: 2,
				Data: []float64{
					-0.353553, -3.889087,
					-0.353553, -1.060660,
				},
			},
			projectedFaces: []m.Matrix{
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{
						0.353553, 3.889087,
						0.353553, 1.060660,
					},
				},
				{
					Rows: 2,
					Cols: 2,
					Data: []float64{
						-0.353553, -3.889087,
						-0.353553, -1.060660,
					},
				},
			},
			wantMatchIndex:  2,
			wantMinDistance: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matchIndex, minDistance := findClosestMatch(tt.projectedTest, tt.projectedFaces)

			if matchIndex != tt.wantMatchIndex {
				t.Errorf("FindClosestMatch(): returned match index: %v, want %v", matchIndex, tt.wantMatchIndex)
			}

			if math.Abs(minDistance-tt.wantMinDistance) > EPSILON {
				t.Errorf("FindClosestMatch(): returned minDistance: %v, want %v", minDistance, tt.wantMinDistance)
			}
		})
	}
}

func TestGetSimilarity(t *testing.T) {
	tests := []struct {
		name           string
		minDistance    float64
		wantSimilarity float64
	}{
		{
			name:           "output is 100 when distance is 0",
			minDistance:    0,
			wantSimilarity: 100,
		},
		{
			name:           "output is correct when distance is 5",
			minDistance:    5,
			wantSimilarity: 80,
		},
		{
			name:           "output is correct when distance is 10",
			minDistance:    10,
			wantSimilarity: 60,
		},
		{
			name:           "output is correct when distance is 20",
			minDistance:    20,
			wantSimilarity: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := getSimilarity(tt.minDistance)

			if math.Abs(similarity-tt.wantSimilarity) > EPSILON {
				t.Errorf("GetSimilarity(): returned similarity: %v, want %v", similarity, tt.wantSimilarity)
			}
		})
	}
}

// integration test to ensure the main program works
func TestRun(t *testing.T) {
	tests := []struct {
		name              string
		timing            bool
		dataSets          []int
		testImage         []int
		k                 int
		imagesFromEachSet int
		rootDir           string
		wantMatchIndex    int
		wantSimilarity    float64
		wantErr           error
	}{
		{
			name:              "similarity is 100 if the image is in the training data",
			timing:            false,
			dataSets:          []int{1},
			testImage:         []int{1, 1},
			k:                 10,
			imagesFromEachSet: 10,
			rootDir:           "../",
			wantMatchIndex:    1,
			wantSimilarity:    100.0,
			wantErr:           nil,
		},
		{
			name:              "similarity is less than 100 if the image is not in the training data",
			timing:            false,
			dataSets:          []int{2, 3},
			testImage:         []int{20, 10},
			k:                 10,
			imagesFromEachSet: 10,
			rootDir:           "../",
			wantMatchIndex:    3,
			wantSimilarity:    78.705802,
			wantErr:           nil,
		},
		{
			name:              "too high k value fails",
			timing:            false,
			dataSets:          []int{1, 2},
			testImage:         []int{20, 10},
			k:                 100,
			imagesFromEachSet: 10,
			rootDir:           "../",
			wantMatchIndex:    0,
			wantSimilarity:    0.0,
			wantErr:           errInvalidKValue,
		},
		{
			name:              "works with many data sets (8)",
			timing:            false,
			dataSets:          []int{1, 2, 3, 4, 5, 6, 7, 8},
			testImage:         []int{20, 2},
			k:                 3,
			imagesFromEachSet: 10,
			rootDir:           "../",
			wantMatchIndex:    38,
			wantSimilarity:    43.1696864,
			wantErr:           nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matchIndex, similarity, err := Run(tt.timing, tt.dataSets, tt.testImage, tt.k, tt.imagesFromEachSet, tt.rootDir)
			if err != tt.wantErr {
				t.Errorf("Run(): returned wrong error: %v, want %v", err, tt.wantErr)
			}

			if matchIndex != tt.wantMatchIndex {
				t.Errorf("Run(): returned incorrect matchindex: %v, want %v", matchIndex, tt.wantMatchIndex)
			}

			if math.Abs(similarity-tt.wantSimilarity) > EPSILON {
				t.Errorf("Run(): returned incorrect similarity: %v, want %v", similarity, tt.wantSimilarity)
			}
		})
	}
}
