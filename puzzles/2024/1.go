package puzzles_2024

import (
	"adventOfCode/puzzles"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var day1 = puzzles.NewPuzzle(2024, 1, func(input string) {
	fmt.Println("2024 - Day 1")
	// split input into by column into two lists

	//create 2 arrays
	var list1 = []int{}
	var list2 = []int{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var columns = strings.Split(line, "   ")
		value1, _ := strconv.Atoi(columns[0])
		value2, _ := strconv.Atoi(columns[1])
		list1 = append(list1, value1)
		list2 = append(list2, value2)
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
