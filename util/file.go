package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadInts(fileName string) []int {
	lines := ReadLines(fileName)
	ints := make([]int, len(lines))
	for i, l := range lines {
		j, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ints[i] = j
	}
	return ints
}
