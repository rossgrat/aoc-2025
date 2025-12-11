package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 5949 is wrong
// 5941 is right

func puzzle2() {
	input, err := os.ReadFile("input.txt")
	inputStr := string(input)
	if err != nil {
		fmt.Println("Failed to read input file.")
		os.Exit(1)
	}

	lines := strings.Split(inputStr, "\n")
	start := 50
	newStart := 0
	zeroes := 0
	for _, line := range lines {
		dir := string(line[0])
		valStr := line[1:]

		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Println("Failed to convert str to int.")
			os.Exit(1)
		}

		fmt.Println(start, "->", dir, val)

		switch dir {
		case "L":
			newStart = universalModulo(start, -val, 100)
		case "R":
			newStart = universalModulo(start, val, 100)
		default:
			fmt.Println("Unrecognized character.")
			os.Exit(1)
		}

		rotations := val / 100
		for range rotations {
			fmt.Println("Click (rotation)")
		}
		zeroes = zeroes + rotations

		if newStart == 0 {
			zeroes = zeroes + 1
			fmt.Println("Click (start == 0)")
		} else if dir == "L" && newStart > start && start != 0 {
			zeroes = zeroes + 1
			fmt.Println("Click (Rotate L: newStart > start)")
		} else if dir == "R" && newStart < start && start != 0 {
			zeroes = zeroes + 1
			fmt.Println("Click (Rotate R: newStart < start)")
		}

		start = newStart

		fmt.Println()

	}
	fmt.Println("\n\nZeroes: ", zeroes)
}

// Could divide val / total = clicks
// If result is a whole number, ending on 0 is accounted for
// If there is a remainder, no worries, doesn't matter

// Would be val /

func rotateLeft(start int, val int, zeroes *int) int {
	for range val {
		start = start - 1
		if start == -1 {
			start = 99
		}

		if start == 0 {
			*zeroes = *zeroes + 1
			fmt.Println("Click")
		}
	}
	return start
}

func rotateRight(start int, val int, zeroes *int) int {
	for range val {
		start = start + 1
		if start == 100 {
			start = 0
		}
		if start == 0 {
			*zeroes = *zeroes + 1
			fmt.Println("Click")
		}
	}
	return start
}
