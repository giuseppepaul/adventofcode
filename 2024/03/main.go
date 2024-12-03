package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Part A: ", partA(string(file)))
	fmt.Println("Part B: ", partB(string(file)))
}

func partA(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	total := 0

	for _, v := range matches {
		left, _ := strconv.Atoi(v[1])
		right, _ := strconv.Atoi(v[2])
		total += left * right
	}
	return total
}

func partB(input string) int {
	re := regexp.MustCompile(`(?sm)(don't\(\).*?do\(\))`)
	removeDonts := re.ReplaceAllString(input, "")
	return partA(removeDonts)
}
