package puzzles

type PuzzleRegistry struct {
	Puzzles []*Puzzle
}

var registry = &PuzzleRegistry{}

func RegisterPuzzle(p *Puzzle) {
	registry.Puzzles = append(registry.Puzzles, p)
}

func GetPuzzles() []*Puzzle {
	return registry.Puzzles
}

func GetPuzzle(year int, day int) *Puzzle {
	for _, p := range registry.Puzzles {
		if p.Year == year && p.Day == day {
			return p
		}
	}
	return nil
}
