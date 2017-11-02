// BinaryGap - Find longest sequence of zeros in binary representation of an integer
package _1_binarygap

import (
	"fmt"
	"strconv"
)

func main() {
	sol := solution(1376796946)
	fmt.Println("largest binary gap by solution 1 is : ", sol)
}

func solution(N int) int {
	n := int64(N)
	binary := strconv.FormatInt(n, 2)

	max := 0
	temp := 0
	for _, r := range binary {
		if r == '0' {
			temp++
		} else {
			if temp > max {
				max = temp
			}
			temp = 0
		}
	}

	return max
}
