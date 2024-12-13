package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/giuseppepaul/adventofcode/utils"
)

var deltaPoints = []image.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}

func main() {
	input := utils.GetInput()
	// input, _ := os.ReadFile("example.txt")
	grid := map[image.Point]string{}

	for y, row := range strings.Split(string(input), "\n") {
		for x := 0; x < len(row); x++ {
			grid[image.Point{x, y}] = string(row[x])
		}
	}

	visited := map[image.Point]bool{}
	var run func(p image.Point, region int, perimeter int, sides int) (int, int, int)

	run = func(p image.Point, area int, perimeter int, sides int) (int, int, int) {
		if _, ok := visited[p]; ok {
			return area, perimeter, sides
		}

		area += 1
		visited[p] = true

		for _, op := range deltaPoints {
			d := p.Add(op)

			// Not in the grid
			if _, ok := grid[d]; !ok || grid[d] != grid[p] {
				// Add a fence
				perimeter += 1

				// reverse the delta
				reverseDelta := p.Add(image.Point{-op.Y, op.X})
				oppositePoint := reverseDelta.Add(op)

				if grid[reverseDelta] != grid[p] || grid[oppositePoint] == grid[p] {
					// found a corner, add a side
					sides += 1
				}
				continue
			}
			// recurse if its in the region
			area, perimeter, sides = run(d, area, perimeter, sides)
		}

		return area, perimeter, sides
	}

	partA := 0
	partB := 0

	for p := range grid {
		if _, ok := visited[p]; !ok {
			area, perimeter, sides := run(p, 0, 0, 0)
			partA += area * perimeter
			partB += area * sides
		}
	}

	fmt.Println("Part A: ", partA)
	fmt.Println("Part B: ", partB)
}
