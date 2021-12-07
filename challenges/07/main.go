package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")
	input := util.GetIntSlice(lines[0])
	sort.Sort(sort.IntSlice(input))
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []int) int {
	median := input[len(input)/2]

	totalFuel := 0
	for _, crabShip := range input {
		totalFuel += int(math.Abs(float64(crabShip - median)))
	}
	return totalFuel
}

func part2(input []int) int {
	total := 0
	for _, i := range input {
		total += i
	}
	mean := total / len(input)

	valuesToTry := []int{mean - 1, mean, mean + 1}

	lowestTotalFuel := math.MaxInt
	for _, v := range valuesToTry {
		totalFuel := 0
		for _, crabShip := range input {
			totalFuel += part2Fuel(int(math.Abs(float64(crabShip - v))))
		}
		if totalFuel < lowestTotalFuel {
			lowestTotalFuel = totalFuel
		}
	}
	return lowestTotalFuel
}

func part2Fuel(distance int) int {
	return distance * (distance + 1) / 2
}
