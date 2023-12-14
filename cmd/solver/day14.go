package solver

import (
	"errors"
	"log"

	"github.com/miketzian/adventofcode2023/internal/day14"
	"github.com/spf13/cobra"
)

var predictSpinLoad int
var day14Cmd = &cobra.Command{
	Use:   "day14",
	Short: "Day 14: Parabolic Reflector Dish",
	Long: `Solve Day 14: Parabolic Reflector Dish
=========================

For Part 2, choose -p`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" {
			cobra.CheckErr(errors.New("--input is required"))
		}

		dish, err := day14.LoadPlatform(inputFile)
		cobra.CheckErr(err)

		if predictSpinLoad > 0 {
			predictor := dish.Predictor()
			log.Printf("Result: %d", predictor.PredictLoad(predictSpinLoad))
		} else {
			dish.Tilt(day14.North)
			log.Printf("Result: %d", dish.Load())
		}
	},
}

func init() {
	day14Cmd.Flags().IntVarP(&predictSpinLoad, "predict-spin-load", "p", 0, "Predict the load after x spins")
	day14Cmd.Flags().StringVar(&inputFile, "input", "", "Input file")

	rootCmd.AddCommand(day14Cmd)
}
