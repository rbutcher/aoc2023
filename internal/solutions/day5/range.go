/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day5

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

type mapRange struct {
	SourceStart      int
	SourceEnd        int
	DestinationStart int
	DestinationEnd   int
	RangeLength      int
}

func MapRangeFromLine(line string) mapRange {
	numbers := readNumbers(line)
	if len(numbers) < 3 {
		log.Fatal().Str("numbers", fmt.Sprintf("%v", numbers))
	}

	return mapRange{
		DestinationStart: numbers[0],
		DestinationEnd:   numbers[0] + numbers[2],
		SourceStart:      numbers[1],
		SourceEnd:        numbers[1] + numbers[2],
		RangeLength:      numbers[2],
	}
}

type seedMap struct {
	SourceType      string
	DestinationType string
	Mappings        []mapRange
}

func (s *seedMap) GetNextMapping(input int) int {
	for _, m := range s.Mappings {
		if input >= m.SourceStart && input <= m.SourceEnd {
			return input - m.SourceStart + m.DestinationStart
		}
	}

	return input
}

func parseInput(source string) ([]int, []*seedMap) {
	lines := strings.Split(source, "\n")

	start := strings.Index(lines[0], ": ") + 2
	seeds := readNumbers(lines[0][start:])

	maps := make(map[string]*seedMap)
	currentSource := ""
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}

		if strings.Contains(line, ":") {
			s := strings.Split(strings.TrimSuffix(line, " map:"), "-")
			m := &seedMap{
				SourceType:      s[0],
				DestinationType: s[2],
				Mappings:        make([]mapRange, 0),
			}

			maps[m.SourceType] = m
			currentSource = m.SourceType
			continue
		}

		m := maps[currentSource]
		m.Mappings = append(m.Mappings, MapRangeFromLine(line))
	}

	var seedMaps []*seedMap
	for _, v := range maps {
		seedMaps = append(seedMaps, v)
	}

	return seeds, seedMaps
}

func readNumbers(line string) []int {
	line = strings.Trim(line, " ")

	buffer := 0
	var numbers []int
	for i := 0; i < len(line); i++ {
		n := toNumber(line[i])
		if n != -1 {
			buffer *= 10
			buffer += n
		}

		if isSpace(line[i]) || i+1 == len(line) {
			numbers = append(numbers, buffer)
			buffer = 0
		}
	}

	return numbers
}

func isSpace(c uint8) bool {
	return c == uint8(' ')
}

func toNumber(c uint8) int {
	const zero = uint8('0')
	const nine = uint8('9')
	if c >= zero && c <= nine {
		return int(c - zero)
	} else {
		return -1
	}
}
