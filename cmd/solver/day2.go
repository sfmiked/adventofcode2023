package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day2"
	"github.com/spf13/cobra"
)

var filterGames bool
var powerSum bool
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Day 2: Cube Conundrum",
	Long: `Solve Day 2: Cube Conundrum
===========================

For Part 1, filter the results with -f.
For Part 2, choose the power function with -p.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		games, err := day2.LoadGameFile(inputFile)
		cobra.CheckErr(err)

		sum, count := 0, 0

		for _, game := range games {
			if filterGames {
				red, green, blue := game.DicesRequired()
				if red > 12 || green > 13 || blue > 14 {
					continue
				}
			}
			if powerSum {
				sum += game.Power()
			} else {
				sum += game.Id()
			}
			count += 1
		}

		log.Printf("%d games with result: %d", count, sum)
	},
}

func init() {
	day2Cmd.Flags().BoolVarP(&filterGames, "filter", "f", false, "Calculate the power of the matching games")
	day2Cmd.Flags().BoolVarP(&powerSum, "power", "p", false, "Calculate the power of the matching games")
	day2Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day2Cmd)
}
