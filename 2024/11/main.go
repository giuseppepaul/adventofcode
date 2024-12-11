package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input := "125 17"
	input, _ := os.ReadFile("input.txt")
	blinks := 75
	stones := map[int]int{}

	for _, v := range strings.Split(string(input), " ") {
		s, _ := strconv.Atoi(v)
		stones[s] = 1
	}

	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}

	output := 0
	for _, count := range stones {
		output += count
	}
	fmt.Println("Part B: ", output)

	// Part A - brute force :D
	// var output string

	// for i := 0; i < blinks; i++ {
	// 	output = ""
	// 	for _, v := range strings.Split(input, " ") {
	// 		if v == "0" {
	// 			output += " 1"
	// 			continue
	// 		}

	// 		if len(v)%2 == 0 {
	// 			index := len(v) / 2
	// 			right := strings.TrimLeft(v[index:], "0")
	// 			if len(right) == 0 {
	// 				right = "0"
	// 			}
	// 			output += " " + v[:index] + " " + right
	// 			continue
	// 		}
	// 		iVal, _ := strconv.Atoi(v)
	// 		output += " " + strconv.Itoa(iVal*2024)
	// 	}
	// 	input = strings.TrimSpace(output)
	// }

	// fmt.Println("Part A: ", len(strings.Split(input, " ")))
}

func blink(stones map[int]int) map[int]int {
	update := map[int]int{}
	for k, count := range stones {
		if k == 0 {
			update[1] += count
			continue
		} else if skey := strconv.Itoa(k); len(skey)%2 == 0 {
			i := len(skey) / 2
			right, _ := strconv.Atoi(skey[i:])
			left, _ := strconv.Atoi(skey[:i])
			update[right] += count
			update[left] += count
			continue
		}
		update[k*2024] += count
	}
	return update
}
