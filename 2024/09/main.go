package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	// input := "2333133121414131402"

	fmt.Println("Part A: ", partA(createBlocks(string(input))))
	fmt.Println("Part B: ", partB(createBlocks(string(input))))
}

func partA(blocks []string) int {
	for x := 0; x < len(blocks); x++ {
		if string(blocks[x]) != "." {
			continue
		}

		val, idx, size := lastNonEmpty(blocks, x, len(blocks)-1)
		if size == 0 {
			break
		}

		blocks[x] = val
		blocks[idx] = "."
	}

	return checksum(blocks)
}

func partB(blocks []string) int {

	var findEmpty = func(minSize int) (int, bool) {
		for i := 0; i < len(blocks)-1; i++ {
			if blocks[i] != "." {
				continue
			}
			size := 1
			for x := i + 1; x < len(blocks)-1; x++ {
				if blocks[x] != blocks[i] {
					break
				}
				size += 1
			}

			if size >= minSize {
				return i, true
			}
		}
		return 0, false
	}

	i := len(blocks) - 1

	for i > 0 {
		val, idx, size := lastNonEmpty(blocks, 0, i)
		if size == 0 {
			break
		}

		emptyIndex, ok := findEmpty(size)

		if ok && emptyIndex < (idx-size) {
			for x := 0; x < size; x++ {
				blocks[emptyIndex+x] = val
				blocks[idx-x] = "."
			}
		}

		i = idx - size
	}
	return checksum(blocks)
}

func checksum(input []string) int {
	checksum := 0
	for i, v := range input {
		ival, _ := strconv.Atoi(string(v))
		checksum += i * ival
	}
	return checksum
}

func createBlocks(input string) []string {
	blocks := []string{}
	id := 0

	for i, num := range input {
		char := "."
		count, _ := strconv.Atoi(string(num))

		if i%2 == 0 {
			char = strconv.Itoa(id)
			id += 1
		}

		for x := 0; x < count; x++ {
			blocks = append(blocks, char)
		}
	}
	return blocks
}

func lastNonEmpty(blocks []string, offset int, limit int) (value string, index int, size int) {
	for x := limit; x >= offset; x-- {
		if blocks[x] == "." {
			continue
		}
		size := 1

		for i := x - 1; i >= offset; i-- {
			if blocks[i] != blocks[x] {
				break
			}
			size += 1
		}

		return blocks[x], x, size
	}
	return "", 0, 0
}
