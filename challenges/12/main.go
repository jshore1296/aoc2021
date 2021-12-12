package main

import (
	"fmt"
	"strings"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")

	caveMap := parseInput(lines)
	fmt.Println(part1(caveMap))
	fmt.Println(part2(caveMap))
}

func parseInput(lines []string) map[string]*cave {
	res := make(map[string]*cave)
	for _, l := range lines {
		caves := strings.Split(l, "-")
		cave0, ok := res[caves[0]]
		if !ok {
			cave0 = &cave{
				name:      caves[0],
				connected: make([]*cave, 0),
			}
		}
		cave1, ok := res[caves[1]]
		if !ok {
			cave1 = &cave{
				name:      caves[1],
				connected: make([]*cave, 0),
			}
		}
		cave0.connected = append(cave0.connected, cave1)
		cave1.connected = append(cave1.connected, cave0)
		res[caves[0]] = cave0
		res[caves[1]] = cave1
	}

	return res
}

type cave struct {
	name      string
	connected []*cave
}

func (c cave) isBig() bool {
	return strings.ToUpper(c.name) == c.name
}

func (c *cave) String() string {
	b := strings.Builder{}
	b.WriteString(c.name + ": [")
	for _, n := range c.connected {
		b.WriteString(n.name + ", ")
	}
	b.WriteString("]")
	return b.String()
}

func part1(caveMap map[string]*cave) int {
	return pathsFrom([]string{"start"}, "end", caveMap)
}

func part2(caveMap map[string]*cave) int {
	return pathsFromVisitTwice([]string{"start"}, "end", caveMap)
}

func pathsFrom(currentPath []string, end string, caveMap map[string]*cave) int {
	start := currentPath[len(currentPath)-1]
	currentCave := caveMap[start]
	if currentCave.name == end {
		return 1
	}

	total := 0
	for _, neighbor := range currentCave.connected {
		if !neighbor.isBig() {
			if alreadyVisited(currentPath, neighbor.name) {
				continue
			}
		}
		newPath := append(currentPath, neighbor.name)
		total += pathsFrom(newPath, end, caveMap)
	}
	return total
}

func alreadyVisited(path []string, cave string) bool {
	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}

func pathsFromVisitTwice(currentPath []string, end string, caveMap map[string]*cave) int {
	start := currentPath[len(currentPath)-1]
	currentCave := caveMap[start]
	if currentCave.name == end {
		return 1
	}

	total := 0
	for _, neighbor := range currentCave.connected {
		if !canVisitAgain(currentPath, neighbor) {
			continue
		}
		newPath := append(currentPath, neighbor.name)
		total += pathsFromVisitTwice(newPath, end, caveMap)
	}
	return total
}

func canVisitAgain(path []string, c *cave) bool {
	if c.name == path[0] {
		// cannot visit "start" again
		return false
	}
	if c.isBig() {
		return true
	}
	visitCountsSmallCaves := make(map[string]int)
	hasCaveWithTwo := false
	for _, p := range path {
		check := cave{name: p}
		if check.isBig() {
			continue
		}
		visitCountsSmallCaves[p]++
		if visitCountsSmallCaves[p] == 2 {
			hasCaveWithTwo = true
		}
	}
	return !(hasCaveWithTwo && visitCountsSmallCaves[c.name] > 0)
}
