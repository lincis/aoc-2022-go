package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var weights = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var encoding = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var decoding = map[string]string{
	"R": "X",
	"P": "Y",
	"S": "Z",
}

var lose = map[string]string{
	"R": "S",
	"P": "R",
	"S": "P",
}

var win = map[string]string{
	"R": "P",
	"P": "S",
	"S": "R",
}

func RPSresult(step []string) int {
	result := weights[step[1]]
	step_0 := encoding[step[0]]
	step_1 := encoding[step[1]]
	if step_0 == step_1 {
		result += 3
	} else if (step_0 == "R" && step_1 == "P") || (step_0 == "P" && step_1 == "S") || (step_0 == "S" && step_1 == "R") {
		result += 6
	}

	// fmt.Println([]string{step_0, step_1})
	// fmt.Println(result)

	return result
}

func RPSstrategy(step []string) int {
	var myaction string
	step_0 := encoding[step[0]]
	if step[1] == "X" { // lose
		myaction = lose[step_0]
	} else if step[1] == "Y" { // draw
		myaction = step_0
	} else { // win
		myaction = win[step_0]
	}
	// fmt.Println([]string{step_0, step[1], myaction})
	// fmt.Println([]string{step_0, decoding[myaction]})
	return RPSresult([]string{step[0], decoding[myaction]})
}

func main() {
	fileBytes, err := ioutil.ReadFile("input/day-02.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawData := strings.Split(string(fileBytes), "\n")
	// var strategy = [][2]string{}
	points := 0
	pointsStrategy := 0
	for _, s := range rawData {
		step := strings.Split(s, " ")
		// fmt.Printlnstep
		// fmt.Println(lenstep)
		if len(step) == 2 {
			// strategy = append(strategy, [2]string{step[0], step[1]})
			points += RPSresult(step)
			pointsStrategy += RPSstrategy(step)
		}
	}

	// fmt.Println(strategy)
	fmt.Printf("Part 1, points in Rock Paper Scissors: %v\n", points)
	fmt.Printf("Part 2, points in Rock Paper Scissors: %v\n", pointsStrategy)
}
