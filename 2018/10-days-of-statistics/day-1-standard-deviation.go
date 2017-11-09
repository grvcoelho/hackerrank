// Given an array, X , of N integers, calculate and print the standard deviation.
// Your answer should be in decimal form, rounded to a scale of 1 decimal place
// (i.e., 12.3 format). An error margin of +-1 will be tolerated for the standard
// deviation.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Mean(n int, xs []int) float64 {
	sum := 0.0

	for _, x := range xs {
		sum += float64(x)
	}

	return sum / float64(n)
}

func Variance(n int, xs []int) float64 {
	mean := Mean(n, xs)
	sum := 0.0

	for _, x := range xs {
		diff := float64(x) - mean
		sum += math.Pow(diff, 2)
	}

	return sum / float64(n)
}

func StandardDeviation(n int, xs []int) float64 {
	variance := Variance(n, xs)

	return math.Sqrt(variance)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n int
	fmt.Scan(&n)

	xs := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		xs[i] = x
	}

	fmt.Printf("%.1f\n", StandardDeviation(n, xs))
}
