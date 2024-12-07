package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/giuseppepaul/adventofcode/utils"
)

var example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func main() {
	input := utils.SplitInputLines(utils.GetInput())
	// input := strings.Split(example, "\n")
	partATotal := 0
	partBTotal := 0

	for _, v := range input {
		s := strings.Split(v, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(s[0]))
		inputs := []int{}

		for _, strNum := range strings.Split(strings.TrimSpace(s[1]), " ") {
			n, _ := strconv.Atoi(strNum)
			inputs = append(inputs, n)
		}

		if validateCalibration(target, inputs, false) {
			partATotal += target
		}

		if validateCalibration(target, inputs, true) {
			partBTotal += target
		}
	}
	fmt.Println("Part A: ", partATotal)
	fmt.Println("Part B: ", partBTotal)
}

func validateCalibration(target int, values []int, concat bool) bool {
	var run func(index, currentValue int) bool
	run = func(index, currentValue int) bool {
		if index == len(values) {
			return currentValue == target
		}

		nextIndex := index + 1

		if run(nextIndex, currentValue*values[index]) || run(nextIndex, currentValue+values[index]) {
			return true
		}

		if concat {
			nextVal, _ := strconv.Atoi(strconv.Itoa(currentValue) + strconv.Itoa(values[index]))
			if run(nextIndex, nextVal) {
				return true
			}
		}

		return false
	}

	return run(1, values[0])
}
