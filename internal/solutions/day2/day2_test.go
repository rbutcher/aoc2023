/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day2

import (
	"github.com/magiconair/properties/assert"
	"github.com/rs/zerolog/log"
	"testing"
)

func TestParseGame_ProperlyParsesLines(t *testing.T) {
	testCases := map[string]Game{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green": {Id: 1, Shown: []CubeSet{
			{Red: 4, Green: 0, Blue: 3},
			{Red: 1, Green: 2, Blue: 6},
			{Red: 0, Green: 2, Blue: 0},
		}},
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue": {Id: 2, Shown: []CubeSet{
			{Red: 0, Green: 2, Blue: 1},
			{Red: 1, Green: 3, Blue: 4},
			{Red: 0, Green: 1, Blue: 1},
		}},
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": {Id: 3, Shown: []CubeSet{
			{Red: 20, Green: 8, Blue: 6},
			{Red: 4, Green: 13, Blue: 5},
			{Red: 1, Green: 5, Blue: 0},
		}},
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": {Id: 4, Shown: []CubeSet{
			{Red: 3, Green: 1, Blue: 6},
			{Red: 6, Green: 3, Blue: 0},
			{Red: 14, Green: 3, Blue: 15},
		}},
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green": {Id: 5, Shown: []CubeSet{
			{Red: 6, Green: 3, Blue: 1},
			{Red: 1, Green: 2, Blue: 2},
		}},
	}

	for line, expected := range testCases {
		t.Run(line, func(t *testing.T) {
			actual := parseGame(line, log.Logger)
			assert.Equal(t, *actual, expected)
		})
	}
}

func TestIsGamePossible_ReportsProperly(t *testing.T) {
	testCases := map[string]bool{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   true,
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         true,
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": false,
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": false,
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   true,
	}

	for line, expected := range testCases {
		t.Run(line, func(t *testing.T) {
			g := parseGame(line, log.Logger)
			numCubes := &CubeSet{Red: 12, Green: 13, Blue: 14}
			actual := isGamePossible(numCubes, g)
			assert.Equal(t, actual, expected)
		})
	}
}

func TestCalculatePower_ReportsProperly(t *testing.T) {
	testCases := map[string]int{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   48,
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         12,
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": 1560,
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": 630,
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   36,
	}

	for line, expected := range testCases {
		t.Run(line, func(t *testing.T) {
			g := parseGame(line, log.Logger)
			assert.Equal(t, g.calculatePower(), expected)
		})
	}
}
