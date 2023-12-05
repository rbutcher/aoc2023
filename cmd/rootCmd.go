/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package cmd

import (
	"fmt"
	"github.com/rbutcher/aoc2023/internal/solutions"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type solution struct {
	P1 func() error
	P2 func() error
}

var rootCmd = &cobra.Command{
	Use:   "aoc2023",
	Short: "aoc2023 is a runner for solutions to the Advent of Code 2023",
	RunE: func(cmd *cobra.Command, args []string) error {
		solutions := []solution{
			{P1: solutions.Day0P1, P2: solutions.Day0P2},
			{P1: solutions.Day1P1, P2: solutions.Day1P2},
			{P1: solutions.Day2p1, P2: solutions.Day2p2},
			{P1: solutions.Day3P1, P2: solutions.Day3P2},
			{P1: solutions.Day4P1, P2: solutions.Day4P2},
		}

		day := viper.GetInt("day")
		s := solutions[day]
		if err := s.P1(); err != nil {
			return err
		}

		if err := s.P2(); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(configureLogging)

	rootCmd.PersistentFlags().IntP("day", "d", 0, "The problem/day to be run.")
	rootCmd.PersistentFlags().Bool("debug", false, "A flag denoting if debug printing is enabled.")
	viper.BindPFlag("day", rootCmd.PersistentFlags().Lookup("day"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

func configureLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
