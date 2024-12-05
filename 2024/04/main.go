package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type XY = struct {
	X int
	Y int
}

type Grid = map[int]map[int]string
type SearchDeltas = []XY

var partA_deltas = []XY{
	// North
	{X: 0, Y: -1},
	// North East
	{X: 1, Y: -1},
	// East
	{X: 1, Y: 0},
	// South East
	{X: 1, Y: 1},
	// South
	{X: 0, Y: 1},
	// South West
	{X: -1, Y: 1},
	// West
	{X: -1, Y: 0},
	// North West
	{X: -1, Y: -1},
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	grid := makeGrid(file)

	fmt.Println("Part A: ", partA(grid, "XMAS", partA_deltas))
	fmt.Println("Part B: ", partB(grid))
}

func getWord(grid Grid, cur XY, dir XY, wordLen int) (string, error) {
	word, _ := getLetter(grid, cur)
	curLoc := cur

	for i := 0; i < wordLen-1; i++ {
		curLoc = XY{X: curLoc.X + dir.X, Y: curLoc.Y + dir.Y}
		letter, ok := getLetter(grid, curLoc)
		if !ok {
			return "", errors.New("not enough letters")
		}
		word += letter
	}

	return word, nil
}

func getLetter(grid Grid, point XY) (string, bool) {
	value, ok := grid[point.Y][point.X]
	return value, ok
}

func partA(grid Grid, needle string, search_deltas SearchDeltas) int {

	words := []string{}
	// Loop rows
	for y := 0; y < len(grid); y++ {
		// Loop cols
		for x := 0; x < len(grid[y]); x++ {
			for _, d := range search_deltas {
				word, err := getWord(grid, XY{X: x, Y: y}, d, len(needle))

				if err == nil {
					words = append(words, word)
				}
			}
		}
	}

	return matches(words, needle)
}

func partB(grid Grid) int {
	count := 0
	points := []XY{}

	// Loop rows
	for y := 0; y < len(grid); y++ {
		// Loop cols
		for x := 0; x < len(grid[y]); x++ {
			point := XY{X: x, Y: y}
			char, _ := getLetter(grid, point)
			if char == "A" {
				points = append(points, point)
			}
		}
	}

	for _, p := range points {
		ne, _ := getLetter(grid, XY{X: p.X + 1, Y: p.Y - 1})
		se, _ := getLetter(grid, XY{X: p.X + 1, Y: p.Y + 1})
		sw, _ := getLetter(grid, XY{X: p.X - 1, Y: p.Y + 1})
		nw, _ := getLetter(grid, XY{X: p.X - 1, Y: p.Y - 1})

		a := nw + "A" + se
		b := ne + "A" + sw

		if (a == "MAS" || a == "SAM") && (b == "MAS" || b == "SAM") {
			count += 1
		}
	}

	return count
}

func matches(haystack []string, needle string) int {
	matches := 0
	for _, v := range haystack {
		if v == needle {
			matches += 1
		}
	}
	return matches
}

func makeGrid(file *os.File) Grid {
	scanner := bufio.NewScanner(file)
	grid := Grid{}
	count := 0

	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		grid[count] = make(map[int]string)

		for i, v := range chars {
			grid[count][i] = v
		}

		count += 1
	}

	return grid
}
