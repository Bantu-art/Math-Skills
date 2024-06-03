package math

import (
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{1, 1, 1, 1}, 1},
		{[]float64{}, 0},
	}

	for _, test := range tests {
		result := Mean(test.numbers)
		if result != test.expected {
			t.Errorf("Mean(%v) = %v; expected %v", test.numbers, result, test.expected)
		}
	}
}
