package math

func CalculateVariance(nbr []float64) float64 {
	if len(nbr) == 0 {
		return 0
	}
	mean := Mean(nbr)
	var squaredDiffs float64
	for _, number := range nbr {
		diff := number - mean
		squaredDiffs += diff * diff
	}
	variance := squaredDiffs / float64(len(nbr))

	return variance
}
