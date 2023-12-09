package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day8"
	"github.com/spf13/cobra"
)

var endingZ bool
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Day 8: Haunted Wasteland",
	Long: `Solve Day 8: Haunted Wasteland
=========================

For Part 2, choose -j `,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		book, err := day8.LoadFile(inputFile)
		cobra.CheckErr(err)

		var result int

		if endingZ {
			result, err = book.PlayEnding("Z")
		} else {
			result, err = book.Play("AAA", func(input string) bool {
				return input == "ZZZ"
			})
		}
		cobra.CheckErr(err)

		log.Printf("Result: %d", result)
	},
}

func init() {
	day8Cmd.Flags().BoolVarP(&endingZ, "ending-z", "z", false, "Play using the rules on the reverse side.")
	day8Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day8Cmd)
}
