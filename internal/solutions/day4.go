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

const (
	zero = uint8('0')
	nine = uint8('9')
)

//go:embed input/day4.txt
var day4Input string

func Day4P1() error {
	l := log.Logger.With().
		Str("scope", "day4p1").
		Logger()

	result := scoreCards(day4Input)
	l.Info().Int("result", result).Send()

	return nil
}

func Day4P2() error {
	l := log.Logger.With().
		Str("scope", "day4p2").
		Logger()

	result := countCards(day4Input)
	l.Info().Int("result", result).Send()

	return nil
}

func scoreCards(source string) int {
	lines := strings.Split(source, "\n")
	score := 0
	for _, line := range lines {
		start := strings.Index(line, ": ") + 2
		end := strings.Index(line, "|")
		var winningNumbers []int
		buffer := 0
		for i := start; i < end; i++ {
			n := toNumber(line[i])
			if n != -1 {
				buffer *= 10
				buffer += n
			} else if isSpace(line[i]) && buffer != 0 {
				winningNumbers = append(winningNumbers, buffer)
				buffer = 0
			}
		}

		buffer = 0
		cardScore := 0
		for i := end; i < len(line); i++ {
			n := toNumber(line[i])
			if n != -1 {
				buffer *= 10
				buffer += n
			}

			if buffer != 0 && (isSpace(line[i]) || i+1 == len(line)) {
				for j := 0; j < len(winningNumbers); j++ {
					if winningNumbers[j] == buffer {
						if cardScore == 0 {
							cardScore = 1
						} else {
							cardScore *= 2
						}
					}
				}
				buffer = 0
			}
		}

		score += cardScore
	}
	return score
}

func countCards(source string) int {
	lines := strings.Split(source, "\n")
	var cards []card
	for _, line := range lines {
		cards = append(cards, cardFromLine(line))
	}

	for i, c := range cards {
		if c.Matches != 0 {
			for j := i + 1; j <= i+c.Matches; j++ {
				cards[j].NumCopies += c.NumCopies
			}
		}
	}

	count := 0
	for _, c := range cards {
		count += c.NumCopies
	}

	return count
}

type card struct {
	Id             int
	WinningNumbers []int
	MatchNumbers   []int
	Matches        int
	NumCopies      int
}

func cardFromLine(source string) card {
	// Parse card id
	start := strings.Index(source, ":") - 1
	for isNumber(source[start-1]) {
		start--
	}

	end := strings.Index(source, ":")
	cardId, err := strconv.Atoi(source[start:end])
	if err != nil {
		log.Fatal().Str("og", source[start:end]).Msgf("failed to parse cardId")
	}

	start = end
	end = strings.Index(source, "|")
	winningNums := parseNumbers(source[start:end])

	start = end + 1
	matchNums := parseNumbers(source[start:])

	matches := 0
	for _, w := range winningNums {
		for _, m := range matchNums {
			if m == w {
				matches++
			}
		}
	}

	return card{
		Id:             cardId,
		WinningNumbers: winningNums,
		MatchNumbers:   matchNums,
		Matches:        matches,
		NumCopies:      1,
	}
}

func parseNumbers(source string) []int {
	var result []int
	buffer := 0
	for i := 0; i < len(source); i++ {
		n := toNumber(source[i])
		if n != -1 {
			buffer *= 10
			buffer += n
		}

		if buffer != 0 && (isSpace(source[i]) || i+1 == len(source)) {
			result = append(result, buffer)
			buffer = 0
		}
	}

	return result
}

func isNumber(c uint8) bool {
	return c >= zero && c <= nine
}

func isSpace(c uint8) bool {
	return c == uint8(' ')
}

func toNumber(c uint8) int {
	if !isNumber(c) {
		return -1
	}

	return int(c - zero)
}
