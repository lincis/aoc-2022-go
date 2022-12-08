package main

import (
	"aoc-2022-go/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

type input struct {
	stacks       [][]rune
	instructions [][3]int
}

func ParseInput(raw_input []string) input {
	stack_cnt := (len(raw_input[0]) + 1) / 4
	// fmt.Println("Number f stacks = ", stack_cnt, " from ", len(raw_input[0]), " / 4")
	parsed_input := input{
		stacks:       make([][]rune, stack_cnt),
		instructions: [][3]int{},
	}
	state := "stacks"
	for _, line := range raw_input {
		if len(line) < 2 {
			state = "instructions"
			continue
		}
		// fmt.Println(line)
		if state == "stacks" {
			have_stack := false
			for i, char := range line {
				if i%4 == 0 && char == '[' {
					have_stack = true
				} else if i%4 == 0 {
					have_stack = false
				}
				// fmt.Println("Have stack = ", have_stack, ", pos = ", i, " char = ", strconv.QuoteRune(char))
				if !have_stack || i%4 != 1 {
					// fmt.Println("Skipping")
					continue
				}
				stack_num := i / 4
				// fmt.Println("Append to ", stack_num, ": ", char)
				parsed_input.stacks[stack_num] = append(parsed_input.stacks[stack_num], char)
			}
		} else {
			if line[0] == ' ' {
				continue
			}
			parts := strings.Split(line, " ")
			if len(parts) == 6 {
				cnt, _ := strconv.Atoi(parts[1])
				from, _ := strconv.Atoi(parts[3])
				to, _ := strconv.Atoi(parts[5])
				parsed_input.instructions = append(parsed_input.instructions, [3]int{cnt, from, to})
			}
		}
	}
	return parsed_input
}

func Operate(in input) input {
	// fmt.Println("Input ", in)
	move := in.instructions[0]
	// fmt.Println("Instruction ", move)
	// fmt.Println(in.stacks)
	in.instructions = in.instructions[1:]
	for i := 0; i < move[0]; i++ {
		// in.stacks[move[2]-1] = append(in.stacks[move[1]-1][:1], in.stacks[move[2]-1]...)
		in.stacks[move[2]-1] = append([]rune{in.stacks[move[1]-1][0]}, in.stacks[move[2]-1]...)
		in.stacks[move[1]-1] = in.stacks[move[1]-1][1:]
	}
	return in
}

func OperateMultiple(in input) input {
	move := in.instructions[0]
	// fmt.Println("Instruction ", move)
	// fmt.Println(in.stacks)
	in.instructions = in.instructions[1:]
	in.stacks[move[2]-1] = append(append([]rune(nil), in.stacks[move[1]-1][:move[0]]...), in.stacks[move[2]-1]...)
	in.stacks[move[1]-1] = in.stacks[move[1]-1][move[0]:]
	return in
}

func One(in input) string {
	steps := len(in.instructions)
	for i := 0; i < steps; i++ {
		// fmt.Println("step = ", i)
		in = Operate(in)
	}
	// fmt.Println(in)
	message := ""
	for _, stack := range in.stacks {
		if len(stack) > 0 {
			message += string(stack[0])
		}
	}
	return message
}

func Two(in input) string {
	steps := len(in.instructions)
	for i := 0; i < steps; i++ {
		// fmt.Println("step = ", i)
		in = OperateMultiple(in)
	}
	// fmt.Println(in)
	message := ""
	for _, stack := range in.stacks {
		if len(stack) > 0 {
			message += string(stack[0])
		}
	}
	return message
}

func main() {
	raw_input := utils.ReadInputFile("input/day-05.txt")
	fmt.Println("Part 1, message = ", One(ParseInput(raw_input)))
	fmt.Println("Part 2, message = ", Two(ParseInput(raw_input)))
}
