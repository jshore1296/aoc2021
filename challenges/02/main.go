package main

import (
	"fmt"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")

	parsed := util.SplitLines(lines, []util.Format{util.FormatString, util.FormatInt}, " ")
	//fmt.Println(parsed[0:5])
	transforms := convertToCoordinateTransforms(parsed)
	//fmt.Println(transforms[0:5])

	x, y := finalCoordinates(transforms)
	fmt.Println(x * y)

	x2, y2 := finalCoordinatesPart2(transforms)
	fmt.Println(x2 * y2)
}

type movement struct {
	X, Y int
}

func convertToCoordinateTransforms(parsed []util.ParsedLine) []movement {
	res := make([]movement, len(parsed))

	for i := range parsed {
		switch parsed[i].Strs[0] {
		case "forward":
			res[i].X = parsed[i].Ints[0]
		case "down":
			res[i].Y = parsed[i].Ints[0]
		case "up":
			res[i].Y = -parsed[i].Ints[0]
		}
	}
	return res
}

func finalCoordinates(movements []movement) (int, int) {
	var finalX, finalY int

	for _, move := range movements {
		finalX += move.X
		finalY += move.Y
	}
	return finalX, finalY
}

func finalCoordinatesPart2(movements []movement) (int, int) {
	var finalX, finalY int

	var currentAim int

	for _, move := range movements {
		finalX += move.X
		finalY += currentAim * move.X
		currentAim += move.Y
	}

	return finalX, finalY
}
