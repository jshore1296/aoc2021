package main

import (
	"fmt"
	"strings"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")

	state := getInitialState(lines[0])
	fmt.Println(runSimulation(state, 80)) // part 1

	state = getInitialState(lines[0])
	fmt.Println(runSimulation(state, 256)) // part 2
}

func getInitialState(line string) []int64 {
	state := make([]int64, 9)
	for _, s := range strings.Split(line, ",") {
		j := util.MustParseInt(s)
		state[j]++
	}
	return state
}

func runSimulation(state []int64, iterations int) int64 {
	for i := 0; i < iterations; i++ {
		start := state[0]
		state = append(state[1:], start)
		state[6] += start
	}

	total := int64(0)

	for _, fish := range state {
		total += fish
	}

	return total
}
