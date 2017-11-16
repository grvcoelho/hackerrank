// The interquartile range of an array is the difference between its first (Q1)
// and third (Q3) quartiles (i.e., Q3 - Q1). Given an array, X, of n integers
// and an array, F, representing the respective frequencies of X's elements,
// construct a data set, S, where each xi occurs at frequency fi. Then
// calculate and print S's interquartile range, rounded to a scale of 1 decimal
// place (i.e., 12.3 format).

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Sort(xs []float64) []float64 {
	ys := make([]float64, len(xs))
	copy(ys, xs)
	sort.Float64s(ys)
	return ys
}

func Median(xs []float64) float64 {
	var a, b int

	if len(xs)%2 == 0 {
		a = len(xs)/2 - 1
		b = len(xs) / 2
	} else {
		a = len(xs) / 2
		b = len(xs) / 2
	}

	return (xs[a] + xs[b]) / 2
}

func Quartiles(n int, xs []float64) (float64, float64, float64) {
	var c1, c2 int
	xs = Sort(xs)

	if n%2 == 0 {
		c1 = n / 2
		c2 = n / 2
	} else {
		c1 = n / 2
		c2 = c1 + 1
	}

	q1 := Median(xs[:c1])
	q2 := Median(xs)
	q3 := Median(xs[c2:])

	return q1, q2, q3
}

func InterQuartiles(n int, xs []float64) float64 {
	q1, _, q3 := Quartiles(n, xs)

	return q3 - q1
}

func Frequency(xs []float64, fs []int) []float64 {
	total := 0

	for _, f := range fs {
		total += f
	}

	arr := make([]float64, total)
	index := 0

	for i, x := range xs {
		for j := 0; j < fs[i]; j++ {
			arr[index] = x
			index++
		}
	}

	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n int
	fmt.Scan(&n)

	xs := make([]float64, n)
	fs := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.ParseFloat(scanner.Text(), 64)
		xs[i] = x
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		f, _ := strconv.Atoi(scanner.Text())
		fs[i] = f
	}

	arr := Frequency(xs, fs)

	result := InterQuartiles(len(arr), arr)
	fmt.Printf("%.1f\n", result)
}
