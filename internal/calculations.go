package internal

import "math"

// CalculateLinearRegression calculates the linear regression parameters: intercept and slope
func CalculateLinearRegression(xValues, yValues []int) (float64, float64) {
	meanX := CalculateAverage(xValues)
	meanY := CalculateAverage(yValues)

	var sumXY, sumXSquared float64

	for i := 0; i < len(xValues); i++ {
		xValue := float64(xValues[i]) - meanX
		yValue := float64(yValues[i]) - meanY

		sumXY += xValue * yValue
		sumXSquared += xValue * xValue
	}

	slope := sumXY / sumXSquared
	intercept := meanY - (slope * meanX)
	return intercept, slope
}

// CalculatePCC calculates the Pearson Correlation Coefficient
func CalculatePCC(xValues, yValues []int) float64 {
	meanX := CalculateAverage(xValues)
	meanY := CalculateAverage(yValues)

	var sumXY, sumXSquared, sumYSquared float64

	for i := 0; i < len(xValues); i++ {
		xValue := float64(xValues[i]) - meanX
		yValue := float64(yValues[i]) - meanY

		sumXY += xValue * yValue
		sumXSquared += xValue * xValue
		sumYSquared += yValue * yValue
	}

	return sumXY / (math.Sqrt(sumXSquared * sumYSquared))
}

// CalculateAverage calculates the average of a list of integers
func CalculateAverage(values []int) float64 {
	total := 0
	for _, value := range values {
		total += value
	}
	return float64(total) / float64(len(values))
}
