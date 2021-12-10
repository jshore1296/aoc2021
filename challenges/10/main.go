package main

import (
	"fmt"
	"sort"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	total := 0
	for _, l := range lines {
		if corrupt, score := isLineCorrupt(l); corrupt {
			total += score
		}
	}
	return total
}

func part2(lines []string) int {

	incompleteLines := make([]string, 0)
	for _, l := range lines {
		if corrupt, _ := isLineCorrupt(l); !corrupt {
			incompleteLines = append(incompleteLines, l)
		}
	}

	scores := make([]int, len(incompleteLines))
	for i, l := range incompleteLines {
		scores[i] = completeLine(l)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func isLineCorrupt(line string) (bool, int) {
	stack := stack{
		r: make([]rune, 0),
	}

	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack.Push(r)
		case ')':
			if stack.Pop() != '(' {
				return true, 3
			}
		case ']':
			if stack.Pop() != '[' {
				return true, 57
			}
		case '}':
			if stack.Pop() != '{' {
				return true, 1197
			}
		case '>':
			if stack.Pop() != '<' {
				return true, 25137
			}
		}
	}
	return false, 0
}

type stack struct {
	r []rune
}

func (s *stack) Push(r rune) {
	s.r = append(s.r, r)
}

func (s *stack) Pop() rune {
	if len(s.r) == 0 {
		return rune('X')
	}
	r := s.r[len(s.r)-1]
	s.r = s.r[:len(s.r)-1]
	return r
}

func completeLine(l string) int {
	stack := stack{
		r: make([]rune, 0),
	}

	for _, r := range l {
		switch r {
		case '(', '[', '{', '<':
			stack.Push(r)
		case ')', ']', '}', '>':
			stack.Pop()
		default:
			panic("wat?")
		}
	}

	score := 0
	for r := stack.Pop(); r != 'X'; r = stack.Pop() {
		score *= 5
		switch r {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}
	return score
}
