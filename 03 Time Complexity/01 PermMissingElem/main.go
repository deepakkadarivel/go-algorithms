package main

import "fmt"

func Solution(A []int) int {
	n := len(A) + 1

	totalSum := n * (n +1) / 2
	actualSum := 0

	for _, value := range A {
		actualSum += value
	}

	return totalSum - actualSum
}

func main() {
	a := []int{1, 2, 3, 5}
	fmt.Println(Solution(a))
}
