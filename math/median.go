package math

import (
	"sort"
)

func Median(nbr []float64) float64 {
	if len(nbr) == 0 {
		return 0
	}
	sort.Float64s(nbr)

	n := len(nbr)
	if n%2 == 1 {
		return nbr[n/2]
	}
	num1 := nbr[n/2]
	num2 := nbr[(n/2)-1]
	return (num1 + num2) / 2
}
