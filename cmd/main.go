package main

import (
	"fmt"
	"os"

	"github.com/etrexel/advent-of-code-2021/internal/common"
	"github.com/spf13/cobra"
)

var rootCmdConfig = &common.ConfigOptions{}

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Runs Advent of Code solutions",
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&rootCmdConfig.Debug, "debug", false, "Enable debug logging")
	rootCmd.AddCommand(solveCmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
