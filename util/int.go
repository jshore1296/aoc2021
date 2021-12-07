package util

import (
	"strconv"
	"strings"
)

func MustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func GetIntSlice(s string) []int {
	res := make([]int, 0)
	for _, si := range strings.Split(s, ",") {
		res = append(res, MustParseInt(si))
	}
	return res
}
