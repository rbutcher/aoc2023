/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package cmd

import (
	"fmt"
	"github.com/rbutcher/aoc2023/internal/common"
	"github.com/rbutcher/aoc2023/internal/solutions/day0"
	"github.com/rbutcher/aoc2023/internal/solutions/day1"
	"github.com/rbutcher/aoc2023/internal/solutions/day2"
	"github.com/rbutcher/aoc2023/internal/solutions/day3"
	"github.com/rbutcher/aoc2023/internal/solutions/day4"
	"github.com/rbutcher/aoc2023/internal/solutions/day5"
	"github.com/rbutcher/aoc2023/internal/solutions/day6"
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
			{P1: day0.Day0P1, P2: day0.Day0P2},
			{P1: day1.Day1P1, P2: day1.Day1P2},
			{P1: day2.Day2p1, P2: day2.Day2p2},
			{P1: day3.Day3P1, P2: day3.Day3P2},
			{P1: day4.Day4P1, P2: day4.Day4P2},
			{P1: day5.Day5P1, P2: day5.Day5P2},
			{P1: day6.Day6P1, P2: day6.Day6P2},
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
	cobra.OnInitialize(common.ConfigureLogger)

	rootCmd.PersistentFlags().IntP("day", "d", 0, "The problem/day to be run.")
	rootCmd.PersistentFlags().Bool("debug", false, "A flag denoting if debug printing is enabled.")
	viper.BindPFlag("day", rootCmd.PersistentFlags().Lookup("day"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}
