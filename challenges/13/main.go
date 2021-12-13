package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")

	paper, instr := parseLines(lines)

	fmt.Println(part1(paper, instr))
	part2(paper, instr)
}

type instruction struct {
	axis  string
	value int
}

func parseLines(lines []string) ([][]int, []instruction) {
	var paper [][]int
	for i, l := range lines {
		if l == "" {
			lines = lines[i+1:]
			break
		}
		coords := strings.Split(l, ",")
		paper = append(paper, []int{util.MustParseInt(coords[0]), util.MustParseInt(coords[1])})
	}

	var instr []instruction
	for _, l := range lines {
		parts := strings.Split(strings.Split(l, " ")[2], "=")
		instr = append(instr, instruction{
			axis:  parts[0],
			value: util.MustParseInt(parts[1]),
		})
	}

	return paper, instr
}

func part1(paper [][]int, instr []instruction) int {
	paper = applyInstruction(paper, instr[0])
	return len(paper)
}

func applyInstruction(paper [][]int, instr instruction) [][]int {
	switch instr.axis {
	case "x":
		return applyX(instr.value, paper)
	case "y":
		return applyY(instr.value, paper)
	default:
		panic("wat")
	}
}

// say we have x = 5
// fold = 4
// we want to get x = 3
// res = fold - (coord - fold) = 2*fold - coord
func applyX(fold int, paper [][]int) [][]int {
	var final [][]int
	for _, coord := range paper {
		if coord[0] > fold {
			newX := 2*fold - coord[0]
			final = append(final, []int{newX, coord[1]})
		} else {
			final = append(final, coord)
		}
	}
	return deduplicate(final)
}

func applyY(fold int, paper [][]int) [][]int {
	var final [][]int
	for _, coord := range paper {
		if coord[1] > fold {
			newY := 2*fold - coord[1]
			final = append(final, []int{coord[0], newY})
		} else {
			final = append(final, coord)
		}
	}
	return deduplicate(final)
}

func deduplicate(paper [][]int) [][]int {
	seen := make(map[string]bool)
	for _, coord := range paper {
		seen[fmt.Sprintf("%d-%d", coord[0], coord[1])] = true
	}

	var newPaper [][]int
	for key := range seen {
		parts := strings.Split(key, "-")
		newPaper = append(newPaper, []int{util.MustParseInt(parts[0]), util.MustParseInt(parts[1])})
	}
	return newPaper
}

func part2(paper [][]int, instr []instruction) {
	for _, i := range instr {
		paper = applyInstruction(paper, i)
	}
	largestX, largestY := getLargestCoords(paper)
	fmt.Println(largestX, largestY)

	output := make([]string, largestY+1)
	for i := range output {
		output[i] = strings.Repeat(" ", largestX+1)
	}

	for _, coord := range paper {
		row := output[coord[1]]
		newRow := row[0:coord[0]] + "X" + row[coord[0]+1:]
		output[coord[1]] = newRow
	}

	for _, line := range output {
		fmt.Println(line)
	}
}

func getLargestCoords(paper [][]int) (int, int) {
	largestX, largestY := math.MinInt, math.MinInt
	for _, coord := range paper {
		if coord[0] > largestX {
			largestX = coord[0]
		}
		if coord[1] > largestY {
			largestY = coord[1]
		}
	}
	return largestX, largestY
}
