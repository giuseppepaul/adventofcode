package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Part A: ", partA(file))
	file.Seek(0, io.SeekStart)
	fmt.Println("Part B: ", partB(file))
}

func partA(file *os.File) int {
	listA, listB := createLists(file)

	slices.Sort(listA)
	slices.Sort(listB)

	diff := 0

	for i := range listA {
		a := listA[i]
		b := listB[i]

		if a > b {
			diff = diff + (a - b)
		} else if b > a {
			diff = diff + (b - a)
		}
	}

	return diff
}

func partB(file *os.File) int {
	listA, listB := createLists(file)

	diff := 0

	for _, v := range listA {
		count := 0
		for _, val := range listB {
			if val == v {
				count += 1
			}
		}
		diff += v * count
	}

	return diff
}

func createLists(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)

	listA := []int{}
	listB := []int{}

	for scanner.Scan() {
		locations := strings.Fields(scanner.Text())
		locA, _ := strconv.Atoi(locations[0])
		locB, _ := strconv.Atoi(locations[1])
		listA = append(listA, locA)
		listB = append(listB, locB)
	}

	return listA, listB
}
