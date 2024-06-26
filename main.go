package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"math-skills/maths"
)

const MaxAllowedValue = 8e9

func main() {
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("Usage: go run main.go <input.txt>")
		return
	}

	input := arg[1]

	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error: unable to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []float64
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		// Skip empty lines.
		if line == "" {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Error: invalid number in your file at line %d: %s\n", lineNumber, line)
			return
		}
		if num > MaxAllowedValue {
			fmt.Printf("Error: number too big in your file at line %d: %s\n", lineNumber, line)
			return
		}
		numbers = append(numbers, num)
	}
	// Check for errors during scanning.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(numbers) == 0 {
		fmt.Println("Error: the file is empty or contains no valid numbers")
		return
	}
	// fmt.Println("Numbers parsed successfully:", numbers)

	Average := maths.Mean(numbers)
	Median := maths.Median(numbers)
	Variance := maths.CalculateVariance(numbers)
	Standard_Deviation := maths.CalculateStandardDeviation(numbers)
	fmt.Printf("Average: %.0f\n", Average)
	fmt.Printf("Median: %.0f\n", Median)
	fmt.Printf("Variance: %.0f\n", Variance)
	fmt.Printf("Standard Deviation: %.0f\n", Standard_Deviation)
}
