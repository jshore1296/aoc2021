package main

import (
	"strconv"
	"strings"
)

func parseBoards(input []string) []bingoBoard {
	boards := make([]bingoBoard, 0)

	currentBoard := newBoard()
	i := 0
	for _, line := range input {
		if line == "" {
			boards = append(boards, currentBoard)
			currentBoard = newBoard()
			i = 0
			continue
		}
		nums := strings.Split(line, " ")
		j := 0
		for _, numStr := range nums {
			if numStr == "" {
				continue
			}
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				panic(err)
			}
			currentBoard.coordinates[int(num)] = []int{i, j}
			j++
		}
		i++
	}
	boards = append(boards, currentBoard)
	return boards
}

type bingoBoard struct {
	markings        [][]bool // 0,0 = top left, 4,4 = bottom right
	inRow, inColumn []int
	coordinates     map[int][]int // bingo number -> coordinate on board
}

func newBoard() bingoBoard {
	b := bingoBoard{
		markings:    make([][]bool, 5),
		coordinates: make(map[int][]int),
		inRow:       make([]int, 5),
		inColumn:    make([]int, 5),
	}
	for i := range b.markings {
		b.markings[i] = make([]bool, 5)
	}

	return b
}

func (b *bingoBoard) markNumber(num int) {
	coords, ok := b.coordinates[num]
	if !ok {
		return
	}
	b.markings[coords[0]][coords[1]] = true
	b.inRow[coords[0]]++
	b.inColumn[coords[1]]++
}

func (b *bingoBoard) hasWon() bool {
	for _, numMarked := range b.inRow {
		if numMarked == 5 {
			return true
		}
	}

	for _, numMarked := range b.inColumn {
		if numMarked == 5 {
			return true
		}
	}
	return false
}

func (b *bingoBoard) isMarked(num int) bool {
	coords, ok := b.coordinates[num]
	if !ok {
		return false
	}
	return b.markings[coords[0]][coords[1]]
}
