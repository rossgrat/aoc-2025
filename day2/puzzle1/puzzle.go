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
			iStr := strconv.Itoa(i)
			if len(iStr)%2 != 0 {
				continue
			}
			left := iStr[:len(iStr)/2]
			right := iStr[len(iStr)/2:]

			fmt.Println(left, right)
			if left == right {
				invalidSum = invalidSum + i
			}
		}

	}

	fmt.Println(invalidSum)
}
