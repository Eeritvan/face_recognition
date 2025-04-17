package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	cmd := exec.Command("./face_recognition", "-d", "1", "2", "3", "-s", "5", "5")
	result, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Command failed: %s", err)
	}

	resultStr := string(result)

	if !strings.Contains(resultStr, "Data used: [1 2 3]") {
		t.Errorf("Expected output to contain 'Data used: [1 2 3]', got:\n%s", resultStr)
	}

	if !strings.Contains(resultStr, "Test Image: set 5 | image 5") {
		t.Errorf("Expected output to contain 'Test Image: set 5 | image 5', got:\n%s", resultStr)
	}

	if !strings.Contains(resultStr, "closest match with: set 1 | image 10") {
		t.Errorf("Expected output to contain 'closest match with: set 1 | image 10', got:\n%s", resultStr)
	}

	if !strings.Contains(resultStr, "similarity: 34.5%") {
		t.Errorf("Expected output to contain 'similarity: 34.5%%', got:\n%s", resultStr)
	}
}

// func TestCLIhelp(t *testing.T) {
// 	cmd := exec.Command("./face_recognition", "-h")
// 	result, err := cmd.CombinedOutput()
// 	if err != nil {
// 		t.Errorf("Command failed: %s", err)
// 	}

// 	resultStr := string(result)

// 	if !strings.Contains(resultStr, "Data used: [1 2 3]") {
// 		t.Errorf("Expected output to contain 'Data used: [1 2 3]', got:\n%s", resultStr)
// 	}
// }

func TestCLItiming(t *testing.T) {
	cmd := exec.Command("./face_recognition", "-t")
	result, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Command failed: %s", err)
	}

	resultStr := string(result)

	if !strings.Contains(resultStr, "time to process training images") {
		t.Errorf("Expected output to contain 'time to process training images', got:\n%s", resultStr)
	}

	if !strings.Contains(resultStr, "time to compute eigenfaces") {
		t.Errorf("Expected output to contain 'time to compute eigenfaces', got:\n%s", resultStr)
	}

	if !strings.Contains(resultStr, "time to project eigenfaces") {
		t.Errorf("Expected output to contain 'time to project eigenfaces', got:\n%s", resultStr)
	}

	if !strings.Contains(resultStr, "time to load test image") {
		t.Errorf("Expected output to contain 'time to load test image', got:\n%s", resultStr)
	}

	if !strings.Contains(resultStr, "time to find closest match") {
		t.Errorf("Expected output to contain 'time to find closest match', got:\n%s", resultStr)
	}

	if !strings.Contains(resultStr, "Total time:") {
		t.Errorf("Expected output to contain 'Total time:', got:\n%s", resultStr)
	}
}
