// Given an array, X, of N integers, calculate and print the respective Mean,
// Median, and Mode on separate lines. If your array contains more than one
// modal value, choose the numerically smallest one.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Median(n int, arr []float64) float64 {
	sorted := make([]float64, len(arr))
	copy(sorted, arr)
	sort.Float64s(sorted)

	a := n/2 - 1
	b := n / 2

	xa := sorted[a]
	xb := sorted[b]

	return (xa + xb) / 2
}

func Mode(n int, arr []float64) float64 {
	m := make(map[float64]int)

	for _, x := range arr {
		m[x]++
	}

	max := arr[0]

	for x, _ := range m {
		if m[x] > m[max] {
			max = x
		}

		if m[x] == m[max] && x < max {
			max = x
		}
	}

	return max
}

func Mean(n int, arr []float64) float64 {
	var sum float64 = 0

	for _, x := range arr {
		sum += x
	}

	return sum / float64(n)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n int
	fmt.Scan(&n)

	arr := make([]float64, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.ParseFloat(scanner.Text(), 64)
		arr[i] = x
	}

	fmt.Println(Mean(n, arr))
	fmt.Println(Median(n, arr))
	fmt.Println(Mode(n, arr))
}
