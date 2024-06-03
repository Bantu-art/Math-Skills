package math

import (
	"testing"
)

func TestCalculateVariance(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 2},
		{[]float64{2, 4, 4, 4, 5, 5, 7, 9}, 4},
		{[]float64{}, 0},
	}

	for _, test := range tests {
		result := CalculateVariance(test.numbers)
		if result != test.expected {
			t.Errorf("CalculateVariance(%v) = %v; expected %v", test.numbers, result, test.expected)
		}
	}
}
