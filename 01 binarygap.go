package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sol := solution(1376796946)
	sol2 := solution2(1376796946)
	fmt.Println("largest binary gap by solution 1 is : ", sol)
	fmt.Println("largest binary gap by solution 2 is : ", sol2)
}

func solution(N int) int {
	n := int64(N)
	binary := strconv.FormatInt(n, 2)

	s := strings.Split(binary, "1")

	max := 0

	for i := range s {
		if len(s[i]) > max {
			max = len(s[i])
		}
	}
	return max
}

func solution2(N int) int {
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
