package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Part A:", partA(file))
	file.Seek(0, io.SeekStart)
	fmt.Println("Part B:", partB(file))
}

func partA(file *os.File) int {
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		lvls := strings.Fields(scanner.Text())
		levels := []int{}

		for _, v := range lvls {
			level, _ := strconv.Atoi(v)
			levels = append(levels, level)
		}

		if reportIsSafe(levels) {
			count += 1
		}
	}

	return count
}

func partB(file *os.File) int {
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		lvls := strings.Fields(scanner.Text())
		levels := []int{}

		for _, v := range lvls {
			level, _ := strconv.Atoi(v)
			levels = append(levels, level)
		}

		if reportIsSafeWithDampener(levels) {
			count += 1
		}
	}

	return count
}

func reportIsSafe(report []int) bool {
	dir := "asc"

	for i := 0; i < len(report)-1; i++ {
		a := report[i]
		b := report[i+1]

		// setup
		if i == 0 {
			if a == b {
				return false
			}
			if a > b {
				dir = "desc"
			}
		}

		if dir == "asc" {
			if !diffLevels(b, a) {
				return false
			}
		} else {
			if !diffLevels(a, b) {
				return false
			}
		}
	}

	return true
}

func reportIsSafeWithDampener(report []int) bool {
	dir := "asc"

	for i := 0; i < len(report)-1; i++ {
		a := report[i]
		b := report[i+1]
		isSafe := true

		// setup
		if i == 0 {
			if a > b {
				dir = "desc"
			} else {
				dir = "asc"
			}
		}

		if dir == "asc" {
			if !diffLevels(b, a) {
				isSafe = false
			}
		} else {
			if !diffLevels(a, b) {
				isSafe = false
			}
		}

		if !isSafe {
			for x := 0; x < len(report); x++ {
				removeItem := []int{}
				removeItem = append(removeItem, report[:x]...)
				removeItem = append(removeItem, report[x+1:]...)
				if reportIsSafe(removeItem) {
					return true
				}
			}
			return false
		}
	}

	return true
}

func diffLevels(bigLvl int, smallLvl int) bool {
	if bigLvl < smallLvl || bigLvl == smallLvl || bigLvl-smallLvl > 3 {
		return false
	}
	return true
}
