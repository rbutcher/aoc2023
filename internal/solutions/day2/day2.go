/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day2

import (
	_ "embed"
	"github.com/rbutcher/aoc2023/internal/common"
	"github.com/rs/zerolog"
	"strconv"
	"strings"
)

//go:embed input.txt
var d2input string

func Day2p1() error {
	l := common.LoggerWithScope("day2p1")

	sum := 0
	numCubes := &CubeSet{Red: 12, Green: 13, Blue: 14}
	lines := strings.Split(d2input, "\n")
	for _, line := range lines {
		g := parseGame(line, l)
		if isGamePossible(numCubes, g) {
			sum += g.Id
		}
	}

	l.Info().Int("result", sum).Msgf("sum_of_game_ids: %d", sum)
	return nil
}

func Day2p2() error {
	l := common.LoggerWithScope("day2p2")

	sum := 0
	lines := strings.Split(d2input, "\n")
	for _, line := range lines {
		g := parseGame(line, l)
		sum += g.calculatePower()
	}

	l.Info().Int("result", sum).Msgf("sum_of_game_powers=%d", sum)
	return nil
}

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id    int
	Shown []CubeSet
}

func (g *Game) calculatePower() int {
	red := 0
	green := 0
	blue := 0
	for _, shown := range g.Shown {
		red = max(shown.Red, red)
		green = max(shown.Green, green)
		blue = max(shown.Blue, blue)
	}

	return red * green * blue
}

// TODO: zero-allocation version of this parsing function.
// I know there is a better way to do this but this is the first thing that came to mind.
func parseGame(line string, l zerolog.Logger) *Game {
	gameSplit := strings.Split(line, ":")
	rawId := strings.TrimPrefix(gameSplit[0], "Game ")

	id, err := strconv.Atoi(rawId)
	if err != nil {
		l.Fatal().Str("line", line).Err(err).Msg("failed to parse line into a game")
	}

	var shown []CubeSet
	shownSplit := strings.Split(gameSplit[1], ";")
	for _, rawSet := range shownSplit {
		setSplit := strings.Split(rawSet, ",")
		set := CubeSet{}
		for _, rawValue := range setSplit {
			if strings.HasSuffix(rawValue, "red") {
				rawRed := strings.TrimSuffix(rawValue, " red")
				set.Red, err = strconv.Atoi(strings.Trim(rawRed, " "))
				if err != nil {
					l.Fatal().Str("line", line).Err(err).Msg("failed to parse line into a game")
				}
				continue
			}

			if strings.HasSuffix(rawValue, "green") {
				rawGreen := strings.TrimSuffix(rawValue, " green")
				set.Green, err = strconv.Atoi(strings.Trim(rawGreen, " "))
				if err != nil {
					l.Fatal().Str("line", line).Err(err).Msg("failed to parse line into a game")
				}
				continue
			}

			if strings.HasSuffix(rawValue, "blue") {
				rawBlue := strings.TrimSuffix(rawValue, " blue")
				set.Blue, err = strconv.Atoi(strings.Trim(rawBlue, " "))
				if err != nil {
					l.Fatal().Str("line", line).Err(err).Msg("failed to parse line into a game")
				}
				continue
			}
		}

		shown = append(shown, set)
	}

	result := Game{Id: id, Shown: shown}
	return &result
}

func isGamePossible(numCubes *CubeSet, game *Game) bool {
	for _, shown := range game.Shown {
		if shown.Red > numCubes.Red || shown.Green > numCubes.Green || shown.Blue > numCubes.Blue {
			return false
		}
	}

	return true
}
