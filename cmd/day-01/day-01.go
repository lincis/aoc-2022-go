package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sumArr(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	fileBytes, err := ioutil.ReadFile("input/day-01.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawData := strings.Split(string(fileBytes), "\n")
	var calories = []int{}
	var sum int = 0
	for _, s := range rawData {
		i, err := strconv.Atoi(s)
		if err != nil {
			calories = append(calories, sum)
			sum = 0
		} else {
			sum += i
		}
	}

	sort.Ints(calories)

	fmt.Println("Part 1")
	fmt.Println(calories[len(calories)-1])

	fmt.Println("Part 2")
	fmt.Println(sumArr(calories[len(calories)-3:]))
}
