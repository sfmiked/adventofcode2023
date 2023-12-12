package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day12"
	"github.com/spf13/cobra"
)

var fiveTimes bool
var day12Cmd = &cobra.Command{
	Use:   "day12",
	Short: "Day 12: Hot Springs",
	Long: `Solve Day 12: Hot Springs
=========================

For Part 2, choose -e`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		result, err := day12.CheckArrangementsFromFile(inputFile, fiveTimes)
		cobra.CheckErr(err)

		log.Printf("Result: %d", result)
	},
}

func init() {
	day12Cmd.Flags().BoolVarP(&fiveTimes, "expand", "e", false, "Should the inputs be expanded")
	day12Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day12Cmd)
}
