package main

import (
	"aoc-2022-go/internal/utils"
	"fmt"
)

func parseForest(in []string) [][]int {
	forest := make([][]int, len(in)-1)
	for i, row := range in {
		if len(row) < 1 {
			continue
		}
		for _, tree := range row {
			forest[i] = append(forest[i], int(tree-'0'))
		}
	}
	return forest
}

func transpose(x [][]int) [][]int {
	transposed := make([][]int, len(x[0]))
	for i := 0; i < len(x[0]); i++ {
		for j := 0; j < len(x); j++ {
			transposed[i] = append(transposed[i], x[j][i])
		}
	}
	return transposed
}

func visible(x int, line []int) bool {
	isVisible := true
	for _, c := range line {
		if c >= x {
			isVisible = false
			break
		}
	}
	return isVisible
}

func treeVisible(forest, forestT [][]int, i, j int) bool {
	x := forest[i][j]
	isVisibleTop := visible(x, forestT[j][:i])
	isVisibleRight := visible(x, forest[i][(j+1):])
	isVisibleBottom := visible(x, forestT[j][(i+1):])
	isVisibleLeft := visible(x, forest[i][:j])
	return isVisibleTop || isVisibleRight || isVisibleBottom || isVisibleLeft
}

func scenicScore(forest, forestT [][]int, i, j int) int {
	score := 1
	x := forest[i][j]
	lines := [][]int{
		forestT[j][:i],       //top
		forest[i][(j + 1):],  // right
		forestT[j][(i + 1):], // bottom
		forest[i][:j],        //left
	}
	visibleTrees := 0
	for k := len(lines[0]) - 1; k >= 0; k-- {
		visibleTrees++
		if x <= lines[0][k] {
			break
		}
	}
	// fmt.Println("top", visibleTrees)
	score *= visibleTrees
	visibleTrees = 0
	for k := 0; k < len(lines[1]); k++ {
		visibleTrees++
		if x <= lines[1][k] {
			break
		}
	}
	// fmt.Println("right", visibleTrees)
	score *= visibleTrees
	visibleTrees = 0
	for k := 0; k < len(lines[2]); k++ {
		visibleTrees++
		if x <= lines[2][k] {
			break
		}
	}
	// fmt.Println("bottom", visibleTrees, lines)
	score *= visibleTrees
	visibleTrees = 0
	for k := len(lines[3]) - 1; k >= 0; k-- {
		visibleTrees++
		if x <= lines[3][k] {
			break
		}
	}
	// fmt.Println("right", visibleTrees)
	score *= visibleTrees
	return score
}

func One(in []string) int {
	visibleTrees := 0
	forest := parseForest(in)
	forestT := transpose(forest)
	for i, line := range forest {
		for j := range line {
			if treeVisible(forest, forestT, i, j) {
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

func Two(in []string) int {
	score := 0
	forest := parseForest(in)
	forestT := transpose(forest)
	for i, line := range forest {
		for j := range line {
			ijScore := scenicScore(forest, forestT, i, j)
			if ijScore > score {
				score = ijScore
			}
		}
	}
	return score
}

func main() {
	rawData := utils.ReadInputFile("input/day-08.txt")
	fmt.Println("Part One = ", One(rawData))
	fmt.Println("Part Two = ", Two(rawData))
}
