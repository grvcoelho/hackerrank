// Given an array, X, of N integers, calculate the respective first quartile
// (Q1), second quartile (Q2), and third quartile (Q3). It is guaranteed that
// Q1, Q2, and Q3 are integers.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Sort(xs []int) []int {
	ys := make([]int, len(xs))
	copy(ys, xs)
	sort.Ints(ys)
	return ys
}

func Median(xs []int) int {
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

func Quartiles(n int, xs []int) (int, int, int) {
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

	q1, q2, q3 := Quartiles(n, xs)

	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q3)
}
