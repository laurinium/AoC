package puzzles

import (
	"fmt"
	"os"
)

type Puzzle struct {
	Year     int
	Day      int
	Function func(input string)
}

func (p *Puzzle) Run() {
	filePath := fmt.Sprintf("inputs/%d/%d.txt", p.Year, p.Day)
	text, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	p.Function(string(text))
}

func NewPuzzle(year, day int, f func(input string)) *Puzzle {
	p := &Puzzle{
		Year:     year,
		Day:      day,
		Function: f,
	}
	return p
}
