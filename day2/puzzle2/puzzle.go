package puzzle1

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

		for i := start; i <= end; i++ {

			// Solution Idea 1 - Factor Approach
			// 1. Compute all factors for a number
			// 2. Split number into sections using each factor
			// 3. Iterate over sections and compare until failure
			//
			// Solution Idea 2 - Backtracking (Claude Description of Regex Backtrace)
			// 1. Start at full number
			// 2. Recursively try and match following substring until no values left
		}

	}

	fmt.Println(invalidSum)
}
