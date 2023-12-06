/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day5

import (
	_ "embed"
	"github.com/rs/zerolog/log"
)

//go:embed input.txt
var day5input string

func Day5P1() error {
	result := findLowestLocation(day5input)
	log.Info().Str("scope", "day5p1").Int("result", result).Send()
	return nil
}

func Day5P2() error {
	result := expandedFindLowestLocation(day5input)
	log.Info().Str("scope", "day5p2").Int("result", result).Send()
	return nil
}

func findLowestLocation(source string) int {
	seeds, maps := parseInput(source)

	result := -1
	for _, seed := range seeds {
		currentSource := "seed"
		currentValue := seed

		for currentSource != "location" {
			for _, m := range maps {
				if m.SourceType == currentSource {
					log.Debug().Int("source_value", currentValue).Str("source_type", currentSource).Send()
					currentValue = m.GetNextMapping(currentValue)
					currentSource = m.DestinationType
				}
			}
		}

		log.Debug().Int("source_value", currentValue).Str("source_type", currentSource).Send()
		if result == -1 || result > currentValue {
			result = currentValue
		}
	}

	return result
}

func expandedFindLowestLocation(source string) int {
	seeds, maps := parseInput(source)

	results := make(chan int)

	numCalcs := len(seeds) / 2
	for i := 0; i < len(seeds); i += 2 {
		index := i
		go func() {
			for j := seeds[index]; j < seeds[index]+seeds[index+1]; j++ {
				currentSource := "seed"
				currentValue := j

				for currentSource != "location" {
					for _, m := range maps {
						if m.SourceType == currentSource {
							currentValue = m.GetNextMapping(currentValue)
							currentSource = m.DestinationType
						}
					}
				}
				results <- currentValue
			}
			log.Info().
				Int("start", seeds[index]).
				Int("end", seeds[index]+seeds[index+1]).
				Msg("done")

			numCalcs--
			if numCalcs == 0 {
				close(results)
			}
		}()
	}

	result := -1
	for v := range results {
		if result == -1 || result > v {
			result = v
			log.Info().Int("new_smallest", result).Send()
		}
	}

	return result
}
