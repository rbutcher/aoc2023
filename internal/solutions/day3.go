/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package solutions

import (
	_ "embed"
	"github.com/rbutcher/aoc2023/internal/matrix"
	"github.com/rs/zerolog/log"
)

//go:embed input/day3.txt
var day3Input string

func Day3P1() error {
	result := sumPartNumbers(day3Input)

	log.Logger.Info().
		Str("scope", "day3p1").
		Int("result", result).
		Msgf("sum_of_parts=%d", result)

	return nil
}

func Day3P2() error {
	result := sumGearRatios(day3Input)

	log.Logger.Info().
		Str("scope", "day3p1").
		Int("result", result).
		Msgf("sum_of_gear_ratios=%d", result)

	return nil
}

func sumPartNumbers(source string) int {
	m := matrix.FromString(source)
	result := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			c := m[y][x]
			if isSymbol(c) {
				pns := m.GetAdjacentNumbers(x, y)
				for _, pn := range pns {
					result += pn
				}
			}
		}
	}

	return result
}

func sumGearRatios(source string) int {
	m := matrix.FromString(source)
	result := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			c := m[y][x]
			if isGear(c) {
				pns := m.GetAdjacentNumbers(x, y)
				if len(pns) != 2 {
					continue
				}

				result += pns[0] * pns[1]
			}
		}
	}

	return result
}

func isGear(c uint8) bool {
	const gear = uint8('*')
	return c == gear
}

func isSymbol(c uint8) bool {
	const zero = uint8('0')
	const nine = uint8('9')
	const dot = uint8('.')
	return c != dot && !(c >= zero && c <= nine)
}
