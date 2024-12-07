package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/giuseppepaul/adventofcode/utils"
)

var directions = []string{"^", ">", "v", "<"}

var dirMap = map[string]Point{
	"^": {X: 0, Y: -1},
	">": {X: 1, Y: 0},
	"v": {X: 0, Y: 1},
	"<": {X: -1, Y: 0},
}

type Route = map[string]Point
type GridObjects = map[string]Point

type Point struct {
	X int
	Y int
}

type Guard struct {
	Location  Point
	Direction string
}

// var example = `....#.....
// .........#
// ..........
// ..#.......
// .......#..
// ..........
// .#..^.....
// ........#.
// #.........
// ......#...`

type Grid = [][]string

func main() {
	input := utils.SplitInputLines(utils.GetInput())
	// input := strings.Split(example, "\n")
	guard := Guard{}

	for y, r := range input {
		for x, v := range r {
			if isGuard(string(v)) {
				guard.Location.X = x
				guard.Location.Y = y
				guard.Direction = string(v)
			}
		}
	}

	grid := makeGrid(input)
	res, route, _ := move(grid, guard, map[string]Point{})
	count := 0

	for _, y := range res {
		for _, x := range y {
			if x == "X" {
				count += 1
			}
		}
	}

	fmt.Println("Part A: ", count)

	fmt.Println("Part B: ", partB(grid, route, guard))
}

func move(g Grid, curr Guard, route Route) (Grid, Route, bool) {
	routeKey := genRouteKeyWithDir(curr)

	if visited(route, routeKey) {
		return g, route, false
	}

	route[routeKey] = curr.Location
	g[curr.Location.Y][curr.Location.X] = "X"
	next, ok := getNextPoint(g, curr)
	if ok {
		return move(g, next, route)
	}
	return g, route, true
}

func visited(route Route, key string) bool {
	_, ok := route[key]
	return ok
}

func genRouteKeyWithDir(g Guard) string {
	return genRouteKey(g.Location) + "-" + g.Direction
}

func genRouteKey(p Point) string {
	return strconv.Itoa(p.Y) + "-" + strconv.Itoa(p.X)
}

func isEmpty(g Grid, p Point) bool {
	return g[p.Y][p.X] != "#"
}

func getNextPoint(g Grid, curr Guard) (next Guard, ok bool) {
	dirIndex := slices.Index(directions, curr.Direction)

	a := []string{}
	a = append(a, directions[dirIndex:]...)
	a = append(a, directions[:dirIndex]...)

	for _, v := range a {
		x := curr.Location.X + dirMap[v].X
		y := curr.Location.Y + dirMap[v].Y
		t := Point{X: x, Y: y}

		if !inGrid(t, g) {
			return curr, false
		}

		if isEmpty(g, t) {
			return Guard{Location: t, Direction: v}, true
		}
	}
	return curr, false
}

func inGrid(p Point, g Grid) bool {
	return !(p.X < 0 || p.X >= len(g[0]) || p.Y < 0 || p.Y >= len(g))
}

func isGuard(g string) bool {
	return g == "^" || g == ">" || g == "v" || g == "<"
}

func makeGrid(input []string) Grid {
	grid := Grid{}
	for _, row := range input {
		cols := []string{}
		for _, v := range row {
			cols = append(cols, string(v))
		}
		grid = append(grid, cols)
	}

	return grid
}

func partB(grid Grid, route Route, startPoint Guard) int {
	// y-x-dir
	matches := map[string]bool{}

	// Add obstacles to each point in the route one by one and test if it loops
	for _, p := range route {
		// if the start point exists in the route - skip it regardless of direction
		if p.X == startPoint.Location.X && p.Y == startPoint.Location.Y {
			continue
		}

		g := copyGrid(grid)
		g[p.Y][p.X] = "#"
		_, _, ok := move(g, startPoint, map[string]Point{})

		if !ok {
			matches[genRouteKey(p)] = true
		}
	}
	return len(matches)
}

func copyGrid(g Grid) Grid {
	newGrid := Grid{}

	for _, y := range g {
		cols := []string{}
		cols = append(cols, y...)
		newGrid = append(newGrid, cols)
	}

	return newGrid
}
