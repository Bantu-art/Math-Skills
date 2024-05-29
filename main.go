package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	for scanner.Scan() {
		line := scanner.Text()

		if line == ""{
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error: invalid number in file:", line)
			return
		}
		numbers = append(numbers, num)
	}
	if len(numbers) == 0 {
		fmt.Println("Error: the file is empty or contains no valid numbers")
		return
	}
	fmt.Println("Numbers parsed successfully:", numbers)
}
