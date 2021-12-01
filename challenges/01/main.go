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
	for i := 1; i <= len(input)-windowSize; i++ {
		if windowTotal(input, i, windowSize) > windowTotal(input, i-1, windowSize) {
			count++
		}
	}
	return count
}

func windowTotal(input []int, idx, window int) int {
	sum := 0
	for i := idx; i < idx+window; i++ {
		sum += input[i]
	}
	return sum
}
