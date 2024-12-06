package main

import (
	"os"
	"strings"
)

func main() {

}

func GetInput() []byte {
	input, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	return input
}

func SplitInput(input []byte) []string {
	return strings.Split(string(input), "\n")
}
