/*
Find the unpaired element in the slice.

The array contains an odd number of elements, and each element of the array can be paired with
another element that has same value,
except for the one element that is left unpaired

Assume that:
	-N is an odd integer with the range [1..1, 000, 000];
	-each element of array A is an integer with a range [1..1, 000, 000, 000];
	-all but except one value in A occur even number of times.
*/

package main

import "fmt"

/*
Solution:
	Takes a slice of int and returns an array
 */

func Solution(A []int) int {
	odd := 0
	for _, value := range A {
		odd ^= value
	}

	return odd
}

func main() {
	t := []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println(Solution(t))
}
