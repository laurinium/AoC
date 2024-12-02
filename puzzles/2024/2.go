package puzzles_2024

import (
	"adventOfCode/puzzles"
	"database/sql"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var day2 = puzzles.NewPuzzle(2024, 2, func(input string) {
	fmt.Println("2024 - Day 2")
	// split input into by column into two lists

	//create 2 arrays

	lines := strings.Split(input, "\n")
	var count = 0
	for _, line := range lines {

		var columns = strings.Split(line, " ")
		var prevValue *int
		prevValue = nil
		for _, column := range columns {
			value1, _ := strconv.Atoi(column)
			if prevValue != nil {
				distance := value1 - *prevValue
				if distance < 0 {
					distance = distance * -1
					if distance > 1 {
						fmt.Println("Invalid")
						return

					}
					count++
				}

			}
			prevValue = value1
		}
	}

	//sort the arrays
	sort.Ints(list1)
	sort.Ints(list2)

	var distances = []int{}
	// loop through the arrays
	for i := 0; i < len(list1); i++ {
		distances = append(distances, list2[i]-list1[i])
	}

	//sum the distances
	var sum = 0
	for _, distance := range distances {
		if distance < 0 {
			distance = distance * -1
		}
		sum += distance
	}
	fmt.Println(sum)

	var simScores = []int{}
	//part 2
	for i := 0; i < len(list1); i++ {
		var num1 = list1[i]
		// count the number of times num1 appears in list2
		var count = 0
		for j := 0; j < len(list2); j++ {
			if list2[j] == num1 {
				count++
			}
		}
		simScores = append(simScores, num1*count)
	}

	//sum the simScores
	var simSum = 0
	for _, simScore := range simScores {
		simSum += simScore
	}
	fmt.Println(simSum)
})

func init() {
	puzzles.RegisterPuzzle(day1)
}
