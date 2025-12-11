package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// input.txt Answer = 989

func puzzle1() {
	input, err := os.ReadFile("input.txt")
	inputStr := string(input)
	if err != nil {
		fmt.Println("Failed to read input file.")
		os.Exit(1)
	}

	lines := strings.Split(inputStr, "\n")
	start := 50
	zeroes := 0
	for _, line := range lines {
		dir := string(line[0])
		valStr := line[1:]

		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Println("Failed to convert str to int.")
			os.Exit(1)
		}

		fmt.Println(dir, val)

		switch dir {
		case "L":
			val = -val
			fmt.Println(start)
		case "R":
			// Noop
		default:
			fmt.Println("Unrecognized character.")
			os.Exit(1)
		}

		start = universalModulo(start, val, 100)

		if start == 0 {
			zeroes = zeroes + 1
		}
	}
	fmt.Println("\n\nZeroes: ", zeroes)
}

func universalModulo(start int, val int, total int) int {
	return ((((start + val) % total) + total) % total)
}

// func rotateLeft(start int, val int) int {
// 	return ((start + val + 100) % 100)
// 	// for range val {
// 	// 	start = start - 1
// 	// 	if start == -1 {
// 	// 		start = 99
// 	// 	}
// 	// }
// 	return start
// }

// func rotateRight(start int, val int) int {
// 	return ((start + val) % 100)
// 	// for range val {
// 	// 	start = start + 1
// 	// 	if start == 100 {
// 	// 		start = 0
// 	// 	}
// 	// }
// 	// return start
// }
