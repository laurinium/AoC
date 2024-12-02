package cmd

import (
	"adventOfCode/puzzles"
	"fmt"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a specific puzzle",
	Long: `Run a specific puzzle by providing the year and day. For example:

run --year 2023 --day 1`,
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := cmd.Flags().GetInt("year")
		day, _ := cmd.Flags().GetInt("day")

		puzzle := puzzles.GetPuzzle(year, day)
		if puzzle == nil {
			fmt.Printf("Puzzle for year %d, day %d not found\n", year, day)
			return
		}

		fmt.Printf("Running puzzle for year %d, day %d\n", year, day)
		puzzle.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().Int("year", 0, "Year of the puzzle")
	runCmd.Flags().Int("day", 0, "Day of the puzzle")
}
