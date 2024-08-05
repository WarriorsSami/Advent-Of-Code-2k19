package day_02

import (
	"fmt"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	newInput := strings.Split(input[0], ",")
	program := make([]int, len(newInput))

	for i, v := range newInput {
		if code, err := strconv.Atoi(v); err == nil {
			program[i] = code
		}
	}

	program[1] = 12
	program[2] = 2

	for i := 0; i < len(program); i += 4 {
		opcode := program[i]

		switch opcode {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		case 99:
			break
		}
	}

	return strconv.Itoa(program[0])
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			newInput := strings.Split(input[0], ",")
			memory := make([]int, len(newInput))

			for i, v := range newInput {
				if code, err := strconv.Atoi(v); err == nil {
					memory[i] = code
				}
			}

			memory[1] = noun
			memory[2] = verb

			for i := 0; i < len(memory); i += 4 {
				opcode := memory[i]

				switch opcode {
				case 1:
					memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
				case 2:
					memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
				case 99:
					break
				}
			}

			if memory[0] == 19690720 {
				return strconv.Itoa(100*noun + verb)
			}
		}
	}

	return "0"
}
