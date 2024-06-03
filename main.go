package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"math-skills/math"
)

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
		lineNumber++
		line := scanner.Text()

		if line == "" {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Error: invalid number in your file at line %d: %s\n", lineNumber, line)
			return
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
	if len(numbers) == 0 {
		fmt.Println("Error: the file is empty or contains no valid numbers")
		return
	}
	// fmt.Println("Numbers parsed successfully:", numbers)

	Average := math.Mean(numbers)
	Median := math.Median(numbers)
	Variance := math.CalculateVariance(numbers)
	Standard_Deviation := math.CalculateStandardDeviation(numbers)
	fmt.Printf("Average: %.0f\n", Average)
	fmt.Printf("Median: %.0f\n", Median)
	fmt.Printf("Variance: %.0f\n", Variance)
	fmt.Printf("Standard Deviation:%.0f\n", Standard_Deviation)

}
