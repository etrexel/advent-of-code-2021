package main

import (
	"github.com/etrexel/advent-of-code-2021/internal/solve"
	"github.com/spf13/cobra"
)

var SolveConfig = &solve.ConfigOptions{
	ConfigOptions: rootCmdConfig,
}

func solveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "solve -d 1",
		Aliases: []string{"s"},
		Short:   "Solve a day's solution",
		Long:    "Run the solution for a particular day",
		RunE:    solveRun,
	}

	cmd.Flags().IntVarP(&SolveConfig.Day, "day", "d", 0, "Day to run solution for")
	cmd.Flags().StringVarP(&SolveConfig.InputPath, "input", "i", "./data/", "Path to input data file")

	return cmd
}

func solveRun(cmd *cobra.Command, args []string) error {
	return solve.Run(SolveConfig)
}
