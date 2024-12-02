package puzzles_2024

import (
	"adventOfCode/puzzles"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

var day2 = puzzles.NewPuzzle(2024, 2, func(input string) {
	fmt.Println("2024 - Day 2")
	// split input into by column into two lists

	//create 2 arrays

	lines := strings.Split(input, "\n")
	var count = 0
	var lineIsValid = true
	var badLevels = 0
	for _, line := range lines {

		var columns = strings.Split(line, " ")
		var prevValue sql.NullInt64
		prevValue = sql.NullInt64{}
		var prevDir sql.NullBool
		prevDir = sql.NullBool{}
		for _, column := range columns {
			value1, _ := strconv.ParseInt(column, 10, 64)
			var dir = sql.NullBool{}
			if prevValue.Valid {
				// check if the direction is the same as the previous
				if value1 > prevValue.Int64 {
					dir.Bool = true
				} else {
					dir.Bool = false
				}
				dir.Valid = true
				if prevDir.Valid {
					if dir.Bool != prevDir.Bool {
						// break both loops
						fmt.Println("Direction Error: ", columns)
						lineIsValid = false
						badLevels++
					}
				}
				// check if the distance is less than 3 and greater than 0
				distance := prevValue.Int64 - value1
				if distance < 0 {
					distance = distance * -1
				}
				if distance <= 0 || distance > 3 {
					lineIsValid = false
					badLevels++
					fmt.Println("Distance Error: ", columns)
				}
				prevDir.Bool = dir.Bool
				prevDir.Valid = true
			}
			// set the previous value
			prevValue.Int64 = value1
			prevValue.Valid = true
			// last column
		}
		if !lineIsValid {
			if badLevels == 1 {
				count++
			} else {
				fmt.Println("Too many errors", badLevels)
			}
		}

		if lineIsValid {
			count++
		} else {
			fmt.Println("Invalid line")
		}
		lineIsValid = true
		badLevels = 0
		//os.Exit(0)
	}
	fmt.Println(count)
})

func init() {
	puzzles.RegisterPuzzle(day2)
}
