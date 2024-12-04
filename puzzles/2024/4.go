package puzzles_2024

import (
	"adventOfCode/puzzles"
	"fmt"
	"strings"
)

var day4 = puzzles.NewPuzzle(2024, 4, func(input string) {
	fmt.Println("2024 - Day 4")
	var lines = strings.Split(input, "\n")
	// read the input into two dimensional array called grid (2d array)
	var grid = [][]string{}
	for _, line := range lines {
		var row = []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}
	//find letter 'A' in the grid
	count := 0
	for i, row := range grid {
		for j, char := range row {
			if char == "A" {
				rowA := i
				colA := j
				positions := [][]int{
					{rowA - 1, colA - 1},
					{rowA - 1, colA + 1},
					{rowA + 1, colA - 1},
					{rowA + 1, colA + 1},
				}
				var corners = []string{}
				for _, pos := range positions {
					if pos[0] >= 0 && pos[0] < len(grid) && pos[1] >= 0 && pos[1] < len(grid[0]) {
						corners = append(corners, grid[pos[0]][pos[1]])
					}
				}
				if len(corners) == 4 {
					if corners[0] == "S" && corners[3] == "M" && corners[1] == "M" && corners[2] == "S" || corners[0] == "M" && corners[3] == "S" && corners[1] == "S" && corners[2] == "M" {
						count++
					}

				}
			}
		}
	}
	fmt.Println(count)
})

func init() {
	puzzles.RegisterPuzzle(day4)
}
