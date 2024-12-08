package main

import (
	"fmt"
	"image"

	"github.com/giuseppepaul/adventofcode/utils"
)

type Grid = map[image.Point]bool
type Frequencies = map[rune][]image.Point

func main() {

	input := utils.SplitInputLines(utils.GetInput())
	grid := Grid{}
	frequencies := Frequencies{}

	for y, row := range input {
		for x, v := range row {
			grid[image.Point{x, y}] = true
			if v != '.' {
				frequencies[v] = append(frequencies[v], image.Point{x, y})
			}
		}
	}
	fmt.Println("Part A:", partA(grid, frequencies))
	fmt.Println("Part B:", partB(grid, frequencies))
}

func partA(grid Grid, frequencies Frequencies) int {
	count := map[image.Point]bool{}

	process(frequencies, func(a, b image.Point) {
		if p := b.Add(b.Sub(a)); grid[p] {
			count[p] = true
		}
	})
	return len(count)
}

func partB(grid Grid, frequencies Frequencies) int {
	count := map[image.Point]bool{}
	var add func(a, b, c image.Point)

	add = func(start, next, vector image.Point) {
		if !grid[next] {
			return
		}
		count[next] = true
		next = next.Add(vector)
		add(start, next, vector)
	}

	process(frequencies, func(a, b image.Point) {
		add(a, b, b.Sub(a))
	})
	return len(count)
}

func process(frequencies Frequencies, cb func(a, b image.Point)) {
	for _, antennas := range frequencies {
		for _, a := range antennas {
			for _, b := range antennas {
				if a == b {
					continue
				}
				cb(a, b)
			}
		}
	}
}
