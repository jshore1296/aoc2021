package main

import (
	"fmt"
	"sort"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	bits := util.ReadBinary("input")
	mostCommon := getMostCommonBits(bits)
	gamma, epsilon := getPart1Rates(bits, mostCommon)
	fmt.Println(gamma * epsilon) // part 1

	//fmt.Println(bits)

	oxygen, co2 := getPart2Rates(bits)
	fmt.Println(oxygen * co2) // part2
}

func getMostCommonBits(bits [][]bool) []bool {
	mostCommon := make([]bool, len(bits[0]))
	for i := 0; i < len(bits[0]); i++ {
		mostCommon[i] = getNumOnes(bits, i) >= len(bits)/2
	}
	return mostCommon
}

func getNumOnes(bits [][]bool, idx int) int {
	ones := 0
	for j := 0; j < len(bits); j++ {
		if bits[j][idx] {
			ones++
		}
	}

	return ones
}

func getPart1Rates(bits [][]bool, mostCommon []bool) (int, int) {
	gamma, epsilon := 0, 0

	for _, b := range mostCommon {
		if b {
			gamma++
		} else {
			epsilon++
		}
		gamma = gamma << 1
		epsilon = epsilon << 1
	}
	gamma = gamma >> 1
	epsilon = epsilon >> 1
	return gamma, epsilon
}

func getPart2Rates(bits [][]bool) (int, int) {
	sort.Slice(bits, func(i, j int) bool {
		for idx := 0; idx < len(bits[i]); idx++ {
			if bits[i][idx] != bits[j][idx] {
				return bits[i][idx]
			}
		}
		return false
	})

	oxygenBits, co2Bits := bits, bits
	//fmt.Println(oxygenBits)
	for i := 0; i < len(oxygenBits[0]); i++ {
		if len(oxygenBits) == 1 {
			break
		}
		ones := getNumOnes(oxygenBits, i)
		if ones >= len(oxygenBits)-ones {
			oxygenBits = oxygenBits[:ones]
		} else {
			oxygenBits = oxygenBits[ones:]
		}
		//fmt.Println(i, oxygenBits)
	}

	for i := 0; i < len(co2Bits[0]); i++ {
		if len(co2Bits) == 1 {
			break
		}
		ones := getNumOnes(co2Bits, i)
		if ones >= len(co2Bits)-ones {
			co2Bits = co2Bits[ones:]
		} else {
			co2Bits = co2Bits[:ones]
		}
		//fmt.Println(i, co2Bits)
	}

	oxygenRate, co2Rate := 0, 0

	for i := 0; i < len(oxygenBits[0]); i++ {
		if oxygenBits[0][i] {
			oxygenRate++
		}

		if co2Bits[0][i] {
			co2Rate++
		}

		if i == len(oxygenBits[0])-1 {
			break
		}

		oxygenRate = oxygenRate << 1
		co2Rate = co2Rate << 1
	}

	return oxygenRate, co2Rate
}
