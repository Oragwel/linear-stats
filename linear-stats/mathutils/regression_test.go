package mathutils

import (
	"math"
	"testing"
)

func TestComputeRegression(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		yValues  []float64
		expectedSlope     float64
		expectedIntercept float64
		expectedPearson   float64
	}{
		{"Simple Linear", []float64{10, 20, 30, 40, 50}, 10.0, 10.0, 1.0},
		{"Constant Values", []float64{5, 5, 5, 5, 5}, 0.0, 5.0, 0.0},
		{"Negative Slope", []float64{50, 40, 30, 20, 10}, -10.0, 50.0, -1.0},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slope, intercept, pearson := ComputeRegression(tt.yValues)

			// Allow small floating point errors
			if math.Abs(slope-tt.expectedSlope) > 1e-6 {
				t.Errorf("Expected slope %.6f, got %.6f", tt.expectedSlope, slope)
			}
			if math.Abs(intercept-tt.expectedIntercept) > 1e-6 {
				t.Errorf("Expected intercept %.6f, got %.6f", tt.expectedIntercept, intercept)
			}
			if math.Abs(pearson-tt.expectedPearson) > 1e-6 {
				t.Errorf("Expected Pearson %.10f, got %.10f", tt.expectedPearson, pearson)
			}
		})
	}
}
