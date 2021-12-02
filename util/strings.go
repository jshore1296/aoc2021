package util

import (
	"strconv"
	"strings"
)

type ParsedLine struct {
	Strs []string
	Ints []int
}

type Format int

const (
	FormatInt Format = iota
	FormatString
)

func SplitLines(lines []string, format []Format, delimiter string) []ParsedLine {
	res := make([]ParsedLine, len(lines))

	for i, line := range lines {
		elems := strings.Split(line, delimiter)
		if len(elems) != len(format) {
			panic("too many elemsn in " + line)
		}

		for j := range elems {
			switch format[j] {
			case FormatInt:
				parsed, err := strconv.ParseInt(elems[j], 10, 64)
				if err != nil {
					panic(err)
				}
				res[i].Ints = append(res[i].Ints, int(parsed))
			case FormatString:
				res[i].Strs = append(res[i].Strs, elems[j])
			}
		}
	}

	return res
}
