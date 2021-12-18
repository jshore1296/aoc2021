package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"time"

	"github.com/jshore1296/aoc2021/util"
)

func main() {
	lines := util.ReadLines("input")

	now := time.Now()
	graph := parseInput(lines, 1)
	fmt.Println(getPathWeight(graph, len(lines)-1)) // part 1
	fmt.Println(time.Since(now))
	now = time.Now()
	graph = parseInput(lines, 5)
	fmt.Println(getPathWeight(graph, 5*len(lines)-1)) // part 2
	fmt.Println(time.Since(now))
}

func getPathWeight(unvisited *NodeList, maxRow int) int {
	visited := calculateShortestPaths(unvisited)
	//fmt.Println(visited[maxRow][maxRow])
	return visited[maxRow][maxRow].lowestRiskPathTotal
}

type node struct {
	row, col  int
	riskLevel int
	heapIndex int

	neighbors           []*node
	lowestRiskPath      []*node
	lowestRiskPathTotal int
}

func (n node) Key() string {
	return strconv.Itoa(n.row) + "-" + strconv.Itoa(n.col)
}

func (n *node) String() string {
	return fmt.Sprintf(`{row: %d, col: %d, riskLevel: %d, lowestRiskPathTotal: %d, neighbors: %s, lowestRiskPath: %s}`, n.row, n.col, n.riskLevel, n.lowestRiskPathTotal, getKeys(n.neighbors), getKeys(n.lowestRiskPath))
}

func getKeys(in []*node) []string {
	out := make([]string, len(in))
	for i, n := range in {
		out[i] = n.Key()
	}
	return out
}

func calculateShortestPaths(unvisited *NodeList) map[int]map[int]*node {
	visited := make(map[int]map[int]*node)

	// for len(unvisited) > 0
	for unvisited.Len() > 0 {
		currentNode := heap.Pop(unvisited).(*node)
		for _, neighbor := range currentNode.neighbors {
			if currentNode.lowestRiskPathTotal+neighbor.riskLevel < neighbor.lowestRiskPathTotal {
				neighbor.lowestRiskPathTotal = currentNode.lowestRiskPathTotal + neighbor.riskLevel
				neighbor.lowestRiskPath = make([]*node, len(currentNode.lowestRiskPath))
				copy(neighbor.lowestRiskPath, currentNode.lowestRiskPath)
				neighbor.lowestRiskPath = append(neighbor.lowestRiskPath, neighbor)
				heap.Fix(unvisited, neighbor.heapIndex)
			}
		}

		rowMap, ok := visited[currentNode.row]
		if !ok {
			rowMap = make(map[int]*node)
		}
		rowMap[currentNode.col] = currentNode
		visited[currentNode.row] = rowMap
	}
	// get lowest total risk
	// calculate new lowest risk of all neighbors, update lowest total risk path if it's new
	// remove current node from unvisited, add to visited

	return visited
}
