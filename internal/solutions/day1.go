/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package solutions

import (
	_ "embed"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
)

//go:embed input/day1.txt
var day1Input string

var numeralWordValues = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Day1P1() error {
	l := log.Logger.With().
		Str("context", "day1p1").
		Logger()

	lines := strings.Split(day1Input, "\n")
	result := 0
	for _, l := range lines {
		cal := getCalibrationValue(l)
		result += cal
	}

	l.Info().
		Int("result", result).
		Msgf("calibration_value=%d", result)
	return nil
}

func Day1P2() error {
	l := log.Logger.With().
		Str("context", "day1p2").
		Logger()

	lines := strings.Split(day1Input, "\n")
	result := 0
	for _, l := range lines {
		cal := getCalibrationValue(l)
		result += cal
	}

	l.Info().
		Int("result", result).
		Msgf("calibration_value=%d", result)
	return nil
}

func getCalibrationValue(source string) int {
	first := -1
	firstIndex := -1
	last := -1
	lastIndex := -1

	// Search for digits
	for i := 0; i < len(source); i++ {
		if iv, err := strconv.Atoi(string(source[i])); err == nil {
			first = iv
			firstIndex = i
			break
		}
	}

	for i := len(source) - 1; i >= 0; i-- {
		if iv, err := strconv.Atoi(string(source[i])); err == nil {
			last = iv
			lastIndex = i
			break
		}
	}

	// Search for words
	for word, wordValue := range numeralWordValues {
		firstWordIndex := strings.Index(source, word)
		if firstWordIndex != -1 && (firstWordIndex < firstIndex || firstIndex == -1) {
			first = wordValue
			firstIndex = firstWordIndex
		}

		lastWordIndex := strings.LastIndex(source, word)
		if lastWordIndex != -1 && lastWordIndex > lastIndex {
			last = wordValue
			lastIndex = lastWordIndex
		}
	}

	var result int
	if last == -1 {
		result = 10*first + first
	} else {
		result = 10*first + last
	}

	return result
}
