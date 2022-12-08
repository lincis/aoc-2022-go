package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("input/day-03.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var weights = make(map[rune]int)
	weight := 1

	for l := 'a'; l <= 'z'; l++ {
		weights[l] = weight
		weight++
	}
	for l := 'A'; l <= 'Z'; l++ {
		weights[l] = weight
		weight++
	}

	misplacedWeight := 0
	badgeWeight := 0
	rawData := strings.Split(string(fileBytes), "\n")
	var groupRucksacks [3]string
	for i, rucksack := range rawData {
		vol := len(rucksack)
		if vol < 2 {
			continue
		}
		// fmt.Println(rucksack)
		left := []rune(rucksack)[0 : vol/2]
		right := []rune(rucksack)[vol/2:]
		// fmt.Println(left)
		// fmt.Println(right)
		misplaced := '0'
		for _, l := range left {
			for _, r := range right {
				if l == r {
					misplaced = l
					break
				}
			}
			if misplaced != '0' {
				break
			}
		}
		// fmt.Printf("Misplaced = %v, weight = %v \n", strconv.QuoteRune(misplaced), weights[misplaced])
		misplacedWeight += weights[misplaced]

		remainder := i % 3
		groupRucksacks[remainder] = rucksack
		if remainder == 2 {
			fmt.Println(groupRucksacks)
			common := '0'
			for _, f := range groupRucksacks[0] {
				for _, s := range groupRucksacks[1] {
					if f == s {
						for _, t := range groupRucksacks[2] {
							if f == t {
								common = f
								break
							}
						}
					}
					if common != '0' {
						break
					}
				}
				if common != '0' {
					break
				}
			}
			badgeWeight += weights[common]
			fmt.Printf("Badge = %v, weight = %v \n", strconv.QuoteRune(common), weights[common])
		}
	}

	fmt.Printf("Part 1 weight of misplaced items = %v \n", misplacedWeight)
	fmt.Printf("Part 2 weight of badge items = %v \n", badgeWeight)
}
