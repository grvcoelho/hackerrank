// Given an array, X, of Nintegers and an array, W, representing the
// respective weights of 's elements, calculate and print the weighted mean of
// X's elements. Your answer should be rounded to a scale of decimal place
// (i.e., 12.3 format).

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WeightedMean(n int, xs, ws []int) float64 {
	wSum := 0.0
	xSum := 0.0

	for _, w := range ws {
		wSum += float64(w)
	}

	for i, x := range xs {
		xSum += float64(x) * float64(ws[i])
	}

	mean := xSum / wSum

	return mean
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n int
	fmt.Scan(&n)

	xs := make([]int, n)
	ws := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		xs[i] = x
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		ws[i] = x
	}

	fmt.Printf("%.1f\n", WeightedMean(n, xs, ws))
}
