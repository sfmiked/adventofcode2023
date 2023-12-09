package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day9"
	"github.com/spf13/cobra"
)

var addToLeft bool
var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "Day 9: Mirage Maintenance",
	Long: `Solve Day 9: Mirage Maintenance
=========================

For Part 2, choose -j `,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		var result int
		var err error

		result, err = day9.PlayFile(inputFile, addToLeft)

		cobra.CheckErr(err)

		log.Printf("Result: %d", result)
	},
}

func init() {
	day9Cmd.Flags().BoolVarP(&addToLeft, "add-to-left", "l", false, "Add on the left instead of the right")
	day9Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day9Cmd)
}
