package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")

	numStrs := strings.Split(lines[0], ",")

	numbers := make([]int, len(numStrs))
	for i, s := range numStrs {
		j, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		numbers[i] = int(j)
	}

	boardInput := lines[2:]

	boards := parseBoards(boardInput)

	fmt.Println(part1(numbers, boards))

	boards = parseBoards(boardInput)
	fmt.Println(part2(numbers, boards))
}

func part1(numbers []int, boards []bingoBoard) int {
	winningBoard, winningNumber := getWinningBoard(numbers, boards)

	return getScore(winningBoard, winningNumber)
}

func part2(numbers []int, boards []bingoBoard) int {
	lastWinningBoard, winningNumber := getLastWinningBoard(numbers, boards)

	return getScore(lastWinningBoard, winningNumber)
}

func getScore(board bingoBoard, num int) int {
	sumUnmarked := 0
	for num := range board.coordinates {
		if board.isMarked(num) {
			continue
		}
		sumUnmarked += num
	}

	return sumUnmarked * num
}

func getWinningBoard(numbers []int, boards []bingoBoard) (bingoBoard, int) {
	for _, num := range numbers {
		for _, board := range boards {
			board.markNumber(num)
			if board.hasWon() {
				return board, num
			}
		}
	}
	panic("no winning board")
}

func getLastWinningBoard(numbers []int, boards []bingoBoard) (bingoBoard, int) {
	for _, num := range numbers {
		nextBoards := make([]bingoBoard, 0, len(boards))
		for _, board := range boards {
			board.markNumber(num)
			if board.hasWon() {
				if len(boards) == 1 {
					return board, num
				}
			} else {
				nextBoards = append(nextBoards, board)
			}
		}
		boards = nextBoards
	}
	panic("no last winning board?????")
}
