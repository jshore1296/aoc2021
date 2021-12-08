package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")
	panes := parseLines(lines)
	fmt.Println(part1(panes))
	fmt.Println(part2(panes))
}

type displayPane struct {
	patterns        []string
	displayedNumber []string
}

func parseLines(lines []string) []displayPane {
	res := make([]displayPane, len(lines))
	for i, l := range lines {
		parts := strings.Split(l, " | ")
		res[i].patterns = sortStringChars(strings.Split(parts[0], " "))
		res[i].displayedNumber = sortStringChars(strings.Split(parts[1], " "))
	}
	return res
}

func sortStringChars(input []string) []string {
	res := make([]string, len(input))
	for i, s := range input {
		runes := []rune(s[:])
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		res[i] = string(runes)
	}
	return res
}

func part1(panes []displayPane) int {
	total := 0
	for _, p := range panes {
		for _, d := range p.displayedNumber {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				total++
			}
		}
	}
	return total
}

func part2(panes []displayPane) int {
	total := 0
	for _, p := range panes {
		total += solvePane(p)
	}
	return total
}

func solvePane(pane displayPane) int {
	solved := make([]string, 10)

	pane.patterns, solved[1] = getWithLen(pane.patterns, 2)
	pane.patterns, solved[4] = getWithLen(pane.patterns, 4)
	pane.patterns, solved[7] = getWithLen(pane.patterns, 3)
	pane.patterns, solved[8] = getWithLen(pane.patterns, 7)
	pane.patterns, solved[9] = getLenXStringWithAllCharsInY(pane.patterns, 6, solved[4]) // only length 6 pattern that contains all letters for 4
	pane.patterns, solved[0] = getLenXStringWithAllCharsInY(pane.patterns, 6, solved[1])
	pane.patterns, solved[6] = getWithLen(pane.patterns, 6)
	pane.patterns, solved[3] = getLenXStringWithAllCharsInY(pane.patterns, 5, solved[1])

	// "determiner" will be present in 5 and _not_ in 2 - it's the only character in "1" that is also in "6"
	var determiner string
	for _, c := range solved[1] {
		if strings.Contains(solved[6], string(c)) {
			determiner = string(c)
			break
		}
	}
	if determiner == "" {
		panic("neither of the digits in 1 were found in 6????")
	}

	pane.patterns, solved[5] = getLenXStringWithAllCharsInY(pane.patterns, 5, determiner)
	pane.patterns, solved[2] = getWithLen(pane.patterns, 5)

	//fmt.Println(pane.patterns)
	//fmt.Println(solved)

	final := 0
	for _, pattern := range pane.displayedNumber {
		final *= 10
		for i, candidate := range solved {
			if candidate == pattern {
				final += i
			}
		}
	}

	return final
}

func getWithLen(input []string, length int) ([]string, string) {
	for i, s := range input {
		if len(s) == length {
			return append(input[0:i], input[i+1:]...), s
		}
	}
	panic("no string with length " + strconv.Itoa(length))
}

func getLenXStringWithAllCharsInY(input []string, x int, neededChars string) ([]string, string) {
	for i, s := range input {
		if len(s) != x {
			continue
		}
		containsAll := true
		for _, c := range neededChars {
			if !strings.Contains(s, string(c)) {
				containsAll = false
				break
			}
		}
		if containsAll {
			return append(input[:i], input[i+1:]...), s
		}
	}
	panic("no string found")
}
