package solver

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inputFile string
var rootCmd = &cobra.Command{
	Use:   "aoc2023",
	Short: "Advent of Code 2023 Solver",
	Long:  `A tool to solve Advent of Code 2023`,
	// Run: func(cmd *cobra.Command, args []string) {
	// we don't need this run command as we assume the sub-commands
	// will provide all the functionality
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
