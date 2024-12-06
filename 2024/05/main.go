package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	rules, updates := handleInput(input)
	partATotal := 0
	partBTotal := 0

	sortFn := func(a, b string) int {
		for _, rule := range rules {
			if rule[0] == a && rule[1] == b {
				return -1
			}
		}
		return 0
	}

	for _, update := range updates {
		if slices.IsSortedFunc(update, sortFn) {
			partATotal += getMiddleValueAsInt(update)
		} else {
			slices.SortFunc(update, sortFn)
			partBTotal += getMiddleValueAsInt(update)
		}
	}

	fmt.Println("Part A: ", partATotal)
	fmt.Println("Part B: ", partBTotal)
}

func getMiddleValueAsInt(update []string) int {
	v, _ := strconv.Atoi(update[len(update)/2])
	return v
}

func handleInput(input []byte) (rules [][]string, updates [][]string) {
	split := strings.Split(string(input), "\n\n")

	for _, v := range strings.Split(split[0], "\n") {
		r := strings.Split(v, "|")
		rules = append(rules, r)
	}

	for _, v := range strings.Split(split[1], "\n") {
		u := strings.Split(v, ",")
		updates = append(updates, u)
	}

	return rules, updates
}
