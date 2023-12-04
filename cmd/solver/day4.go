package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day4"
	"github.com/spf13/cobra"
)

var bothSideRules bool
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Day 4: Scratchcards",
	Long: `Solve Day 4: Scratchcards
=========================

For Part 2, choose -r `,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		game, err := day4.LoadGame(inputFile)
		cobra.CheckErr(err)
		if bothSideRules {
			game.ApplyReverseRules()
		}

		score := game.TotalScore()
		card_count := game.TotalCards()

		log.Printf("%d cards with score: %d", card_count, score)
	},
}

func init() {
	day4Cmd.Flags().BoolVarP(&bothSideRules, "reverse", "r", false, "Play using the rules on the reverse side.")
	day4Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day4Cmd)
}
