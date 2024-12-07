package utils

import (
	"os"
	"strings"
)

func GetInput() []byte {
	input, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	return input
}

func SplitInputLines(input []byte) []string {
	return strings.Split(string(input), "\n")
}
