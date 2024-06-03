package math

func Mean(nbr []float64) float64 {
	if len(nbr) == 0 {
		return 0
	}
	var sum float64
	for _, num := range nbr {
		sum += num
	}
	average := sum / float64(len(nbr))
	return average
}
