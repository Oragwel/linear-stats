package mathutils

import "math"

// ComputeRegression calculates the linear regression line and Pearson correlation coefficient
func ComputeRegression(yValues []float64) (float64, float64, float64) {
	n := float64(len(yValues))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i, y := range yValues {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}

	// Compute linear regression
	denominatorSlope := n*sumX2 - sumX*sumX
	denominatorPearson := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))

	var slope, intercept, pearson float64
	if denominatorSlope != 0 {
		slope = (n*sumXY - sumX*sumY) / denominatorSlope
		intercept = (sumY - slope*sumX) / n
	}

	// Compute Pearson correlation coefficient
	if denominatorPearson != 0 {
		pearson = (n*sumXY - sumX*sumY) / denominatorPearson
	}

	return slope, intercept, pearson
}
