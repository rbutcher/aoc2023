/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day6

import (
	_ "embed"
	"github.com/rbutcher/aoc2023/internal/common"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
)

//go:embed input.txt
var day6Input string

func Day6P1() error {
	result := multiplyWinningHoldTimeInstances(day6Input)
	log.Info().Int("result", result).Str("scope", "day6p1").Send()
	return nil
}

func Day6P2() error {
	result := singleRaceHoldTimeInstances(day6Input)
	log.Info().Int("result", result).Str("scope", "day6p2").Send()
	return nil
}

func singleRaceHoldTimeInstances(source string) int {
	r := parseRace(source)
	return len(r.CalculateWinningHoldTimes())
}

func multiplyWinningHoldTimeInstances(source string) int {
	races := parseRaces(source)

	result := -1
	for _, r := range races {
		holdTimes := r.CalculateWinningHoldTimes()
		if result == -1 {
			result = len(holdTimes)
		} else {
			result *= len(holdTimes)
		}
	}

	return result
}

type race struct {
	Time     int
	Distance int
}

func parseRaces(source string) []race {
	lines := strings.Split(source, "\n")

	start := strings.Index(lines[0], ":") + 1
	times := common.ReadNumbers(lines[0][start:])

	start = strings.Index(lines[1], ":") + 1
	distances := common.ReadNumbers(lines[1][start:])

	var result []race
	for i := 0; i < len(times); i++ {
		result = append(result, race{
			Time:     times[i],
			Distance: distances[i],
		})
	}

	return result
}

func parseRace(source string) race {
	lines := strings.Split(source, "\n")

	start := strings.Index(lines[0], ":") + 1
	l := strings.Replace(lines[0][start:], " ", "", -1)
	time, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal().Err(err).Str("value", l).Msg("failed to parse value to int")
	}

	start = strings.Index(lines[1], ":") + 1
	l = strings.Replace(lines[1][start:], " ", "", -1)
	distance, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal().Err(err).Str("value", l).Msg("failed to parse value to int")
	}

	return race{
		Time:     time,
		Distance: distance,
	}
}

func (r race) CalculateWinningHoldTimes() []int {
	var winning []int
	for hold := 0; hold <= r.Time; hold++ {
		d := hold * (r.Time - hold)
		if d > r.Distance {
			winning = append(winning, hold)
		}
	}

	return winning
}
