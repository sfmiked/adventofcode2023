package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day7"
	"github.com/spf13/cobra"
)

var withJokers bool
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Day 7: Camel Cards",
	Long: `Solve Day 7: Camel Cards
=========================

For Part 2, choose -j `,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		game := day7.NewGame(withJokers)

		err := game.LoadHands(inputFile)
		cobra.CheckErr(err)

		result := game.TotalWinnings()

		log.Printf("%d hands with result: %d", game.Len(), result)
	},
}

func init() {
	day7Cmd.Flags().BoolVarP(&withJokers, "with-jokers", "j", false, "Play using the rules on the reverse side.")
	day7Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day7Cmd)
}
