package main

import (
	"fmt"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")
	grid := parseLines(lines)
	fmt.Println(part1(grid))
	grid = parseLines(lines)
	fmt.Println(part2(grid))
}

func part1(g octoGrid) int {
	total := 0
	for i := 0; i < 100; i++ {
		total += g.simulateStep()
	}
	return total
}

func part2(g octoGrid) int {
	step := 1
	for {
		if g.simulateStep() == 100 {
			return step
		}
		step++
	}
}

type octoGrid struct {
	pos [][]int
}

func parseLines(lines []string) octoGrid {
	g := octoGrid{
		pos: make([][]int, len(lines)),
	}
	for row, l := range lines {
		g.pos[row] = make([]int, len(l))
		for col, c := range l {
			g.pos[row][col] = util.MustParseInt(string(c))
		}
	}
	return g
}

func (g *octoGrid) simulateStep() int {
	hasFlashed := make([][]bool, len(g.pos))
	for i := range hasFlashed {
		hasFlashed[i] = make([]bool, len(g.pos[0]))
	}

	// increase by 1
	for row := range g.pos {
		for col := range g.pos[row] {
			g.pos[row][col]++
		}
	}

	// flash anything > 9
	for g.flash(hasFlashed) {
	}

	flashes := 0
	for row := range hasFlashed {
		for col := range hasFlashed[row] {
			if hasFlashed[row][col] {
				flashes++
				g.pos[row][col] = 0
			}
		}
	}

	return flashes
}

func (g *octoGrid) flash(hasFlashed [][]bool) bool {
	flashed := false

	flashedThisTime := make([][]bool, len(hasFlashed))
	for i := range hasFlashed {
		flashedThisTime[i] = make([]bool, len(hasFlashed[i]))
	}

	for row := range g.pos {
		for col := range g.pos[row] {
			if hasFlashed[row][col] {
				continue
			}

			if g.pos[row][col] > 9 {
				hasFlashed[row][col] = true
				flashedThisTime[row][col] = true
				flashed = true
			}
		}
	}

	for row := range g.pos {
		for col := range g.pos[row] {
			if flashedThisTime[row][col] {
				for _, n := range g.getNeighbors(row, col) {
					g.pos[n[0]][n[1]]++
				}
			}
		}
	}

	return flashed
}

func (g *octoGrid) getNeighbors(row, col int) [][]int {
	var neighbors [][]int

	// three positions above this one
	if row > 0 {
		if col > 0 {
			neighbors = append(neighbors, []int{row - 1, col - 1})
		}
		neighbors = append(neighbors, []int{row - 1, col})
		if col < len(g.pos[row])-1 {
			neighbors = append(neighbors, []int{row - 1, col + 1})
		}
	}

	// this row
	if col > 0 {
		neighbors = append(neighbors, []int{row, col - 1})
	}
	if col < len(g.pos[row])-1 {
		neighbors = append(neighbors, []int{row, col + 1})
	}

	// final row
	if row < len(g.pos)-1 {
		if col > 0 {
			neighbors = append(neighbors, []int{row + 1, col - 1})
		}
		neighbors = append(neighbors, []int{row + 1, col})
		if col < len(g.pos[row])-1 {
			neighbors = append(neighbors, []int{row + 1, col + 1})
		}
	}

	return neighbors
}
