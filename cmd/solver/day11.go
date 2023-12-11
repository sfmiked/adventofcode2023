package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day11"
	"github.com/spf13/cobra"
)

var expandingValue int
var day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: "Day 11: Cosmic Expansion",
	Long: `Solve Day 11: Cosmic Expansion
=========================

For Part 2, choose -e 1000000 `,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		var result int
		var err error

		obs, err := day11.NewObservations(inputFile, expandingValue)
		cobra.CheckErr(err)

		galaxies := obs.Galaxies()

		for i, g := range galaxies {
			for i2, g2 := range galaxies {
				if i2 <= i {
					continue
				}
				dist, err := obs.GalaxyDistance(g, g2)
				cobra.CheckErr(err)
				result += dist
			}
		}

		// part 1 9684228
		// part 2 483844716556
		log.Printf("Result: %d", result)
	},
}

func init() {
	day11Cmd.Flags().IntVarP(&expandingValue, "expanding-value", "e", 2, "The width of expanding space")
	day11Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day11Cmd)
}
