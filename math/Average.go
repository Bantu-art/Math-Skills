package math

// Mean calculates the average of a slice of float64 numbers.
func Mean(nbr []float64) float64 {
	// Check if the slice is empty.
	if len(nbr) == 0 {
		return 0
	}
	var sum float64
	// Loop through each number in the slice.
	for _, num := range nbr {
		sum += num
	}
	// Calculate the average by dividing the sum by the number of elements.
	average := sum / float64(len(nbr))
	return average
}
