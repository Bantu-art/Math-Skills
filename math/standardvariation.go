package math

import "math"

// CalculateStandardDeviation returns the standard deviation of a slice of float64 numbers
func CalculateStandardDeviation(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	// Calculate the variance of the numbers.
	variance := CalculateVariance(numbers)
	return math.Sqrt(variance)
}
