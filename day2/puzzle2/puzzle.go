package puzzle2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fileB, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Failed to read file.")
		os.Exit(1)
	}

	file := string(fileB)

	ranges := strings.Split(file, ",")

	invalidSum := 0
	for _, r := range ranges {
		startEnd := strings.Split(r, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])

		fmt.Println("Start: ", start, " - End: ", end)

		for val := start; val <= end; val++ {
			fmt.Println("VAL", val)
			valStr := strconv.Itoa(val)

			factors := []int{}
			for length := 1; length < len(valStr); length++ {
				if len(valStr)%length == 0 {
					factors = append(factors, length)
				}
			}

			fmt.Println("Factors:", factors)

			match := matchedPattern(valStr, factors)

			if match {
				invalidSum = invalidSum + val
			}

			// Solution Idea 1 - Factor Approach
			// 1. Compute all factors for a number
			// 2. Split number into sections using each factor
			// 3. Iterate over sections and compare until failure
			//
			// Solution Idea 2 - Backtracking (Claude Description of Regex Backtrace)
			// 1. Start at full number
			// 2. Recursively try and match following substring until no values left
			fmt.Println()
			fmt.Println()
		}

	}

	fmt.Println(invalidSum)
}

func matchedPattern(valStr string, factors []int) bool {
	for _, factor := range factors {
		factorFailed := false
		fmt.Println("FACTOR", factor)
		chunk1 := valStr[:factor]
		fmt.Println("CHUNK1", chunk1)
		for iter := factor * 2; iter <= len(valStr); iter = iter + factor {
			chunk2 := valStr[iter-factor : iter]
			fmt.Println("CHUNK-N", chunk2)
			if chunk1 != chunk2 {
				factorFailed = true
				break
			}
		}
		if !factorFailed {
			return true
		}
	}
	return false
}
