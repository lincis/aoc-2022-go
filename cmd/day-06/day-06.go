package main

import (
	"aoc-2022-go/internal/utils"
	"fmt"
)

func allDifferent[I comparable](in []I) bool {
	result := true
	for _, r1 := range in {
		matches := 0
		for _, r2 := range in {
			if r1 == r2 {
				matches++
			}
		}
		if matches > 1 {
			result = false
			break
		}
	}
	return result
}

func StartMarker(in string, length int) int {
	marker := -1
	in_runes := []rune(in)
	for i := length; i < len(in_runes); i++ {
		if allDifferent(in_runes[(i - length):i]) {
			marker = i
			break
		}
	}
	return marker
}

func One(in string) int {
	return StartMarker(in, 4)
}

func Two(in string) int {
	return StartMarker(in, 14)
}

func main() {
	raw_input := utils.ReadInputFile("./input/day-06.txt")
	fmt.Println("Part One = ", One(raw_input[0]))
	fmt.Println("Part Two = ", Two(raw_input[0]))
}
