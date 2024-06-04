package maths

import (
	"testing"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{1, 2, 3, 4}, 2.5},
		{[]float64{}, 0},
	}

	for _, test := range tests {
		result := Median(test.numbers)
		if result != test.expected {
			t.Errorf("Median(%v) = %v; expected %v", test.numbers, result, test.expected)
		}
	}
}
