package puzzles_2024

import (
	"adventOfCode/puzzles"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var day3 = puzzles.NewPuzzle(2024, 3, func(input string) {
	fmt.Println("2024 - Day 3")

	donts := strings.Split(input, "don't()")
	dos := []string{}
	for _, dont := range donts {
		fmt.Println("Dont: ", dont, "\n")
		res := strings.Split(dont, "do()")
		dos = append(dos, res[1:]...)
	}
	sum := 0
	for _, do := range dos {
		sum += getSum(do)
	}
	sum += getSum(donts[0])

	fmt.Println("Sum: ", sum)
})

func getSum(input string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	// Compile the regex with proper flags
	re := regexp.MustCompile(pattern)

	sum := 0
	// Find matches
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		// convert the strings to integers
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		res := num1 * num2
		sum += res
	}
	fmt.Println("Sum: ", sum)
	return sum
}

func init() {
	puzzles.RegisterPuzzle(day3)
}
