package main

import (
	"fmt"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadInts("input")
	fmt.Println(numLarger(lines, 1)) // part 1
	fmt.Println(numLarger(lines, 3)) // part 2
}

func numLarger(input []int, windowSize int) int {
	count := 0
	for i := windowSize; i < len(input); i++ {
		if input[i] > input[i-windowSize] {
			count++
		}
	}
	return count
}
