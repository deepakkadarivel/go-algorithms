// CyclicRotation
package main

import "fmt"

func Solution(A []int, K int) []int {
	arrayLen := len(A)
	if arrayLen == 0 {
		return A
	}
	rotatedArray := make([]int, arrayLen)

	K = K % arrayLen

	for i := K; i < arrayLen; i++ {
		rotatedArray[i] = A[i-K]
	}

	for i := 0; i < K; i++ {
		rotatedArray[i] = A[arrayLen-K+i]
	}

	return rotatedArray
}

func main() {
	A := []int{3, 8, 9, 7, 6}
	K := 3
	fmt.Println(Solution(A, K))
}


