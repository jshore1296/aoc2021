package main

import (
	"container/heap"
	"math"

	"github.com/jshore1296/aoc2021/util"
)

type NodeList []*node

func (n *NodeList) Push(x interface{}) {
	x.(*node).heapIndex = len(*n)
	*n = append(*n, x.(*node))
}

func (n *NodeList) Pop() interface{} {
	elem := (*n)[len(*n)-1]
	elem.heapIndex = -1
	*n = (*n)[:len(*n)-1]
	return elem
}

func (n *NodeList) Len() int {
	return len(*n)
}

func (n *NodeList) Less(i, j int) bool {
	return (*n)[i].lowestRiskPathTotal < (*n)[j].lowestRiskPathTotal
}

func (n *NodeList) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
	(*n)[i].heapIndex = i
	(*n)[j].heapIndex = j
}

func parseInput(lines []string, numCopies int) *NodeList {
	res := make([][]int, len(lines)*numCopies)
	for row, line := range lines {
		for col, riskStr := range line {
			risk := util.MustParseInt(string(riskStr))

			for rowDiff := 0; rowDiff < numCopies; rowDiff++ {
				currentRow := row + len(lines)*rowDiff
				if res[currentRow] == nil {
					res[currentRow] = make([]int, len(line)*numCopies)
				}
				for colDiff := 0; colDiff < numCopies; colDiff++ {
					currentCol := col + len(lines[row])*colDiff
					currentRisk := (risk + rowDiff + colDiff) % 9
					if currentRisk == 0 {
						currentRisk = 9
					}
					res[currentRow][currentCol] = currentRisk
				}
			}
		}
	}

	return generateNodes(res)
}

func generateNodes(risks [][]int) *NodeList {
	finalList := make(NodeList, 0)
	res := make(map[int]map[int]*node)

	for row, line := range risks {
		res[row] = make(map[int]*node)
		for col, risk := range line {
			n := &node{
				row:                 row,
				col:                 col,
				riskLevel:           risk,
				lowestRiskPath:      make([]*node, 0),
				lowestRiskPathTotal: math.MaxInt,
			}
			if row > 0 {
				upperNeighbor := res[row-1][col]
				upperNeighbor.neighbors = append(upperNeighbor.neighbors, n)
				n.neighbors = append(n.neighbors, upperNeighbor)
			}
			if col > 0 {
				leftNeighbor := res[row][col-1]
				leftNeighbor.neighbors = append(leftNeighbor.neighbors, n)
				n.neighbors = append(n.neighbors, leftNeighbor)
			}
			res[row][col] = n
			heap.Push(&finalList, n)
		}
	}
	res[0][0].lowestRiskPath = []*node{res[0][0]}
	res[0][0].lowestRiskPathTotal = 0

	return &finalList
}
