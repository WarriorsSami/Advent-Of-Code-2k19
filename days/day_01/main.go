package day_01

import (
	"fmt"
	"math"
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
	masses := make([]int, 0)
	for _, line := range input {
		mass := 0
		_, err := fmt.Sscanf(line, "%d", &mass)
		if err != nil {
			return ""
		}
		masses = append(masses, mass)
	}

	totalFuel := 0
	for _, mass := range masses {
		fuel := int(math.Floor(float64(mass)/3)) - 2
		totalFuel += fuel
	}

	return fmt.Sprintf("%d", totalFuel)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	masses := make([]int, 0)
	for _, line := range input {
		mass := 0
		_, err := fmt.Sscanf(line, "%d", &mass)
		if err != nil {
			return ""
		}
		masses = append(masses, mass)
	}

	totalFuel := 0
	for _, mass := range masses {
		fuel := getFuel(mass)
		totalFuel += fuel
	}

	return fmt.Sprintf("%d", totalFuel)
}

func getFuel(mass int) int {
	fuel := int(math.Floor(float64(mass)/3)) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + getFuel(fuel)
}
