package solver

import (
	"errors"
	"fmt"

	"github.com/miketzian/adventofcode2023/internal/day1"
	"github.com/spf13/cobra"
)

var onlyDigits bool
var inputFile string
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Solve the Day 1 Puzzle",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		doc := day1.NewDoc(onlyDigits)
		count, err := doc.LoadFile(inputFile)
		cobra.CheckErr(err)

		fmt.Printf("Loaded %d records\n", count)
		fmt.Printf("Result: %d\n", doc.Sum())
	},
}

func init() {
	day1Cmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Check only numeric digits 0-9")
	day1Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day1Cmd)
}
