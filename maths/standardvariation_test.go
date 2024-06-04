package maths

import (
	"testing"
)

func TestCalculateStandardDeviation(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 1.4142135623730951},
		{[]float64{2, 4, 4, 4, 5, 5, 7, 9}, 2},
		{[]float64{}, 0},
	}

	for _, test := range tests {
		result := CalculateStandardDeviation(test.numbers)
		if result != test.expected {
			t.Errorf("CalculateStandardDeviation(%v) = %v; expected %v", test.numbers, result, test.expected)
		}
	}
}
