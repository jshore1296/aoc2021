package main

import (
	"fmt"
	"strings"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")

	segments := parseLineSegments(lines)

	floor := getOceanFloor(segments)

	fmt.Println(part1(floor, segments))
	floor = getOceanFloor(segments)
	fmt.Println(part2(floor, segments))
}

func parseLineSegments(input []string) []lineSegment {
	res := make([]lineSegment, len(input))
	for i, l := range input {
		pairs := strings.Split(l, " -> ")
		start := strings.Split(pairs[0], ",")
		end := strings.Split(pairs[1], ",")
		res[i].x1, res[i].y1 = util.MustParseInt(start[0]), util.MustParseInt(start[1])
		res[i].x2, res[i].y2 = util.MustParseInt(end[0]), util.MustParseInt(end[1])
	}
	return res
}

func getOceanFloor(input []lineSegment) [][]int {
	largestX, largestY := 0, 0

	for _, l := range input {
		if l.x1 > largestX {
			largestX = l.x1
		}
		if l.x2 > largestX {
			largestX = l.x2
		}
		if l.y1 > largestY {
			largestY = l.y1
		}
		if l.y2 > largestY {
			largestY = l.y2
		}
	}

	oceanFloor := make([][]int, largestX+1)
	for i := range oceanFloor {
		oceanFloor[i] = make([]int, largestY+1)
	}
	return oceanFloor
}

func part1(oceanFloor [][]int, segments []lineSegment) int {
	for _, seg := range segments {
		coords := seg.getCoordinates(false)
		for _, c := range coords {
			oceanFloor[c[0]][c[1]]++
		}
	}

	return getDangerZones(oceanFloor)
}

func getDangerZones(oceanFloor [][]int) int {
	dangerZones := 0
	for x := range oceanFloor {
		for y := range oceanFloor[x] {
			if oceanFloor[x][y] > 1 {
				dangerZones++
			}
		}
	}
	return dangerZones
}

func part2(oceanFloor [][]int, segments []lineSegment) int {
	for _, seg := range segments {
		coords := seg.getCoordinates(true)
		for _, c := range coords {
			oceanFloor[c[0]][c[1]]++
		}
	}

	return getDangerZones(oceanFloor)
}
