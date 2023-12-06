package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day6"
	"github.com/spf13/cobra"
)

var combineNumbers bool
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Day 6: Wait For It",
	Long: `Solve Day 6: Wait For It
=========================

For Part 2, choose -c `,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		ranges, err := day6.LoadFile(inputFile, combineNumbers)
		cobra.CheckErr(err)

		result := day6.RecordProduct(ranges)

		log.Printf("%d ranges with result: %d", len(ranges), result)
	},
}

func init() {
	day6Cmd.Flags().BoolVarP(&combineNumbers, "combine-numbers", "c", false, "Play using the rules on the reverse side.")
	day6Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day6Cmd)
}
