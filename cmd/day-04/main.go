package main

import (
	"aoc-2022-go/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func One(input []string) int {
	assignements := SplitAssignements(input)
	overlaps := 0
	for _, assigned := range assignements {
		if Overlap(assigned[0], assigned[1]) {
			overlaps++
		}
	}
	return overlaps
}

func Two(input []string) int {
	assignements := SplitAssignements(input)
	overlaps := 0
	for _, assigned := range assignements {
		overlaps += OverlapAny(assigned[0], assigned[1])
	}
	return overlaps
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func SplitAssignements(input []string) [][][]int {
	var assignements [][][]int
	for _, row := range input {
		if len(row) < 1 {
			continue
		}
		pairs := strings.Split(row, ",")
		var assignement [][]int
		for _, pair := range pairs {
			limits := strings.Split(pair, "-")
			l0, _ := strconv.Atoi(limits[0])
			l1, _ := strconv.Atoi(limits[1])
			assignement = append(assignement, MakeRange(l0, l1))
		}
		assignements = append(assignements, assignement)
	}
	return assignements
}

func Overlap(a1, a2 []int) bool {
	o1 := true
	o2 := true
	for _, i1 := range a1 {
		i := 0
		for _, i2 := range a2 {
			if i1 == i2 {
				break
			}
			i++
		}
		if i == len(a2) {
			o1 = false
		}
	}
	for _, i2 := range a2 {
		i := 0
		for _, i1 := range a1 {
			if i2 == i1 {
				break
			}
			i++
		}
		if i == len(a1) {
			o2 = false
		}
	}
	return o1 || o2
}

func OverlapAny(a1, a2 []int) int {
	overlap := false
	for _, i1 := range a1 {
		for _, i2 := range a2 {
			overlap = i1 == i2
			if overlap {
				break
			}
		}
		if overlap {
			break
		}
	}
	if overlap {
		return 1
	} else {
		return 0
	}
}

func main() {
	input := utils.ReadInputFile("./input/day-04.txt")
	fmt.Println("Part 1: ", One(input))
	fmt.Println("Part 2: ", Two(input))
}
