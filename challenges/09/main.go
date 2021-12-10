package main

import (
	"fmt"
	"sort"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")
	cf := parseLines(lines)
	fmt.Println(part1(cf))
	fmt.Println(part2(cf))
}

type caveFloor struct {
	positions   [][]int
	basinNumber [][]int
}

func parseLines(lines []string) caveFloor {
	cf := caveFloor{
		positions:   make([][]int, len(lines)),
		basinNumber: make([][]int, len(lines)),
	}
	for row, l := range lines {
		cf.positions[row] = make([]int, len(l))
		cf.basinNumber[row] = make([]int, len(l))
		for col, c := range l {
			h := util.MustParseInt(string(c))
			cf.positions[row][col] = h
			if h == 9 {
				cf.basinNumber[row][col] = -1
			}
		}
	}
	return cf
}

func (cf caveFloor) getNeighbors(row, col int) [][]int {
	var neighbors [][]int
	if row > 0 {
		neighbors = append(neighbors, []int{row - 1, col})
	}
	if row < len(cf.positions)-1 {
		neighbors = append(neighbors, []int{row + 1, col})
	}
	if col > 0 {
		neighbors = append(neighbors, []int{row, col - 1})
	}
	if col < len(cf.positions[0])-1 {
		neighbors = append(neighbors, []int{row, col + 1})
	}
	return neighbors
}

func (cf caveFloor) posIsLowerThanNeighbors(row, col int) bool {
	height := cf.positions[row][col]
	if height == 9 {
		return false
	}
	neighbors := cf.getNeighbors(row, col)
	//	fmt.Println(row, col, neighbors)
	for _, n := range neighbors {
		if cf.positions[n[0]][n[1]] < height {
			return false
		}
	}
	return true
}

func part1(cf caveFloor) int {
	sumRisk := 0

	for row := range cf.positions {
		for col := range cf.positions[row] {
			if cf.posIsLowerThanNeighbors(row, col) {
				sumRisk += cf.positions[row][col] + 1
			}
		}
	}

	return sumRisk
}

func part2(cf caveFloor) int {
	basinNumber := 1
	basinTotals := make([]int, 0)
	for row := range cf.positions {
		for col := range cf.positions[row] {
			if cf.basinNumber[row][col] == 0 {
				size := cf.mapBasin(basinNumber, row, col)
				basinNumber++
				basinTotals = append(basinTotals, size)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinTotals)))

	return basinTotals[0] * basinTotals[1] * basinTotals[2]
}

func (cf caveFloor) mapBasin(basinNumber, row, col int) int {
	if cf.basinNumber[row][col] != 0 {
		return 0
	}
	cf.basinNumber[row][col] = basinNumber

	neighbors := cf.getNeighbors(row, col)
	total := 1
	for _, n := range neighbors {
		total += cf.mapBasin(basinNumber, n[0], n[1])
	}
	return total
}
