package main

import (
	"fmt"
	"math"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")

	template, rules := parseLines(lines)
	fmt.Println(part1(template, rules))
	template, rules = parseLines(lines)
	fmt.Println(part2(template, rules))
}

func parseLines(lines []string) (map[string]int, map[string]string) {
	parsedTemplate := make(map[string]int)
	tem := lines[0]
	for i := 1; i < len(tem); i++ {
		parsedTemplate[tem[i-1:i+1]]++
	}
	for i := 0; i < len(tem); i++ {
		parsedTemplate[string(tem[i])]++
	}
	rules := make(map[string]string)

	for _, l := range lines[2:] {
		rules[l[:2]] = l[len(l)-1:]
	}
	return parsedTemplate, rules
}

func performStep(template map[string]int, rules map[string]string) {
	newIncrs := make(map[string]int)
	for pair, newChar := range rules {
		if count, ok := template[pair]; ok {
			newIncrs[pair[:1]+newChar] += count
			newIncrs[newChar+pair[1:]] += count
			newIncrs[pair] -= count
			newIncrs[newChar] += count
		}
	}

	for pair, count := range newIncrs {
		template[pair] += count
		if template[pair] == 0 {
			delete(template, pair)
		}
	}
}

func part1(template map[string]int, rules map[string]string) int {
	return runSimulation(10, template, rules)
}

func part2(template map[string]int, rules map[string]string) int {
	return runSimulation(40, template, rules)
}

func runSimulation(steps int, template map[string]int, rules map[string]string) int {
	for i := 0; i < steps; i++ {
		performStep(template, rules)
	}

	var (
		leastCount = math.MaxInt
		mostCount  = math.MinInt
	)
	for str, count := range template {
		if len(str) > 1 {
			continue
		}
		if count < leastCount {
			leastCount = count
		}
		if count > mostCount {
			mostCount = count
		}
	}
	return mostCount - leastCount
}
