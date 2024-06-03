# Statistical Analysis Tool

This Go program reads a dataset of numbers from a file and calculates various statistical metrics: average, median, variance, and standard deviation. It handles errors gracefully and skips invalid lines.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [License](#license)

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://learn.zone01kisumu.ke/git/bbantu/math-skills
    cd math-skills
    ```

2. **Initialize the Go module:**
    ```bash
    go mod init math-skills
    go mod tidy
    ```

## Usage

1. **Prepare your input file:**
   Create a text file (e.g., `input.txt`) with one number per line. Empty lines will be skipped.

    Example `input.txt`:
    ```
    40.8
    50.9
    60

    70.1
    ```

2. **Run the program:**
    ```bash
    go run main.go input.txt
    ```

    **Output:**
    ```
    Average: 55
    Median: 55
    Variance: 138
    Standard Deviation:12
    ```
    3. **Handling large numbers:**
    ```
   The program limits the input data to avoid processing large numbers that may lead to inaccuracies in the variance and standard deviation calculations. Numbers exceeding `1e6` (1,000,000) will be considered too large, and the program will prompt the user to consider scaling down the data for accurate results.
   ```

## Testing

Unit tests are written using the Go testing package. To run the tests, use the following command:

```bash
cd math
```
```bash
go test -v
```

