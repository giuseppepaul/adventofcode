package main

import (
	"fmt"
	"image"
	"strconv"

	"github.com/giuseppepaul/adventofcode/utils"
)

var example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

type GridMap = map[image.Point]int
type VisitedMap = map[image.Point]bool

var searchPoints = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {

	input := utils.SplitInputLines(utils.GetInput())
	// input := utils.SplitInputLines([]byte(example))
	grid := GridMap{}

	for y, row := range input {
		for x, v := range row {
			vInt, _ := strconv.Atoi(string(v))
			grid[image.Point{x, y}] = vInt
		}
	}

	countA := 0
	countB := 0

	for point := range grid {
		if grid[point] != 0 {
			continue
		}
		countA += search(grid, point, VisitedMap{})
		countB += search(grid, point, nil)
	}

	fmt.Println("Part A:", countA)
	fmt.Println("Part B:", countB)
}

func search(grid GridMap, curr image.Point, visited VisitedMap) int {
	if grid[curr] == 9 {
		if visited != nil {
			if visited[curr] {
				return 0
			}
			visited[curr] = true
		}
		return 1
	}
	count := 0
	next := grid[curr] + 1

	// Check all the surrounding points
	for _, point := range searchPoints {
		if p := curr.Add(point); grid[p] == next {
			count += search(grid, p, visited)
		}
	}

	return count
}
