package main

import "fmt"

func slicing() {
	x := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	y := []int{1, 2, 3, 4, 5}
	i := 1
	fmt.Println("[i][:]", x[i][:])
	fmt.Println("[:][i]", x[:][i])
	fmt.Println("[i][i:]", x[i][i:])
	fmt.Println("[:2][2:]", x[:2][0:])
	fmt.Println(y[3:1])
}

func main() {
	slicing()
}
