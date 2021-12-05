package main

import (
	"math"
)

type lineSegment struct {
	x1, y1 int
	x2, y2 int
}

func (l lineSegment) isHorizontal() bool {
	return l.y1 == l.y2
}

func (l lineSegment) isVertical() bool {
	return l.x1 == l.x2
}

func (l lineSegment) isDiagonal() bool {
	return math.Abs(float64(l.x2-l.x1)) == math.Abs(float64(l.y2-l.y1))
}

func (l lineSegment) getCoordinates(inclDiagonal bool) [][]int {
	coords := make([][]int, 0)
	if l.isHorizontal() {
		start, end := l.x1, l.x2
		if start > end {
			start, end = end, start
		}
		for x := start; x <= end; x++ {
			coords = append(coords, []int{x, l.y1})
		}
	} else if l.isVertical() {
		start, end := l.y1, l.y2
		if start > end {
			start, end = end, start
		}
		for y := start; y <= end; y++ {
			coords = append(coords, []int{l.x1, y})
		}
	} else if l.isDiagonal() && inclDiagonal {
		xDiff := 1
		totalCoords := l.x2 - l.x1 + 1
		if l.x1 > l.x2 {
			xDiff = -1
			totalCoords = l.x1 - l.x2 + 1
		}
		yDiff := 1
		if l.y1 > l.y2 {
			yDiff = -1
		}
		for i := 0; i < totalCoords; i++ {
			coords = append(coords, []int{l.x1 + xDiff*i, l.y1 + yDiff*i})
		}
	}
	return coords
}
