package math

import (
	"sort"
)

func Median(nbr []float64) float64 {
	if len(nbr) == 0 {
		return 0
	}

	// Sort the slice in ascending order.
	sort.Float64s(nbr)

	// Sort the slice in ascending order.
	n := len(nbr)
	if n%2 == 1 {
		// If odd, return the middle element.
		return nbr[n/2]
	}
	// If even, get the two middle elements.
	num1 := nbr[n/2]
	num2 := nbr[(n/2)-1]
	return (num1 + num2) / 2
}
