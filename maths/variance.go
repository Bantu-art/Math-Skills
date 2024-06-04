package maths

func CalculateVariance(nbr []float64) float64 {
	if len(nbr) == 0 {
		return 0
	}
	// Calculate the mean (average) of the numbers in the slice.
	mean := Mean(nbr)
	var squaredDiffs float64
	for _, number := range nbr {

		// Calculate the difference between the number and the mean.
		diff := number - mean
		// Square the difference and add it to the sum of squared differences.
		squaredDiffs += diff * diff
	}
	// Calculate the variance by dividing the sum of squared differences by the number of elements.
	variance := squaredDiffs / float64(len(nbr))

	return variance
}
