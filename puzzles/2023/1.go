package puzzles_2023

import (
	"adventOfCode/puzzles"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

var day1 = puzzles.NewPuzzle(2023, 1, func(input string) {
	fmt.Println("2023 - Day 1")
	calibrationSum := 0
	for _, line := range strings.Split(input, "\n") {
		numbers := []rune{}
		for _, char := range line {
			// if char is a number
			if char >= 48 && char <= 57 {
				numbers = append(numbers, char)
			}
		}
		twoDigits := string(numbers[0]) + string(numbers[len(numbers)-1])
		//string to int
		lineSum, _ := strconv.ParseInt(twoDigits, 10, 64)
		calibrationSum += int(lineSum)
	}
	fmt.Println("Sum of all calibration values:", calibrationSum)

	// Part 2
	replaceArray := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	fmt.Println(md5.Sum([]byte(input)))
	//replace input left to right
	for i, replace := range replaceArray {
		input = strings.ReplaceAll(input, strconv.Itoa(i+1), replace)
	}

	step2Sum := 0
	for _, line := range strings.Split(input, "\n") {
		numbers := []string{}
		for _, char := range line {
			// if char is a number
			if char >= 48 && char <= 57 {
				numbers = append(numbers, string(char))
			}
		}
		//string to int
		twoDigits := numbers[0] + numbers[len(numbers)-1]
		twoDigitsInt, _ := strconv.ParseInt(twoDigits, 10, 64)
		step2Sum += int(twoDigitsInt)
	}
	fmt.Println("Sum of all calibration values:", step2Sum)
})

func init() {
	puzzles.RegisterPuzzle(day1)
}
