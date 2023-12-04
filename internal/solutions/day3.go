/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package solutions

import (
	_ "embed"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"strings"
)

const (
	stateFindNumber = iota
	stateReadNumber
	statePartNumber
	stateGearRatio
	stateFindGear
)

//go:embed input/day3.txt
var day3Input string

func isDebug() bool {
	return viper.GetBool("debug")
}

func Day3P1() error {
	l := logWithScope("day3p1")

	input := convertInput(day3Input)
	result := sumPartNumbers(input)

	l.Info().Int("result", result).Msgf("sum_of_part_numbers:%d", result)
	return nil
}

func Day3P2() error {
	return nil
}

func convertInput(input string) [][]rune {
	var result [][]rune
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	for _, line := range lines {
		result = append(result, []rune(line))
	}
	return result
}

func sumPartNumbers(input [][]rune) int {
	state := stateFindNumber

	x := 0
	y := 0
	done := false

	iNumStart := -1
	iNumEnd := -1
	result := 0

	for !done {
		switch state {
		case stateFindNumber:
			// Find a number within the current line
			if x == len(input[y]) {
				y++
				x = 0
			}

			if y == len(input) {
				done = true
				continue
			}

			if isNumber(input[y][x]) {
				iNumStart = x
				state = stateReadNumber
			} else {
				x++
			}

		case stateReadNumber:
			// Read in the whole number (between 1 and 3 digits)
			remainingInLine := len(input[y]) - x - 1
			if remainingInLine >= 2 && isNumber(input[y][x+1]) && isNumber(input[y][x+2]) {
				// Three digits
				iNumStart = x
				iNumEnd = x + 2
			} else if remainingInLine >= 1 && isNumber(input[y][x+1]) {
				// Two digits
				iNumStart = x
				iNumEnd = x + 1
			} else {
				// One digit
				iNumStart = x
				iNumEnd = x
			}
			state = statePartNumber

		case statePartNumber:
			// Check if this number is a part number
			if iNumStart > 0 { // Check left
				if isSymbol(input[y][iNumStart-1]) {
					pn := toPartNum(input, y, iNumStart, iNumEnd)
					result += pn
					x = iNumEnd + 1
					state = stateFindNumber
					continue
				}
			}

			if iNumEnd < len(input[y])-1 { // Check right
				if isSymbol(input[y][iNumEnd+1]) {
					pn := toPartNum(input, y, iNumStart, iNumEnd)
					result += pn
					x = iNumEnd + 1
					state = stateFindNumber
					continue
				}
			}

			if y > 0 { // Check upper
				uStart := iNumStart
				if iNumStart-1 >= 0 { // check diag up left
					uStart = iNumStart - 1
				}

				uEnd := iNumEnd
				if iNumEnd+1 <= len(input[y])-1 { // check diag up right
					uEnd = iNumEnd + 1
				}

				found := false
				for i := uStart; i <= uEnd; i++ {
					if isSymbol(input[y-1][i]) {
						pn := toPartNum(input, y, iNumStart, iNumEnd)
						result += pn
						x = iNumEnd + 2
						state = stateFindNumber
						found = true
						break
					}
				}

				if found {
					continue
				}
			}

			if y < len(input)-1 { // check lower
				lStart := iNumStart
				if iNumStart-1 >= 0 { // check diag up left
					lStart = iNumStart - 1
				}

				lEnd := iNumEnd
				if iNumEnd+1 <= len(input[y])-1 { // check diag up right
					lEnd = iNumEnd + 1
				}

				found := false
				for i := lStart; i <= lEnd; i++ {
					if isSymbol(input[y+1][i]) {
						pn := toPartNum(input, y, iNumStart, iNumEnd)
						result += pn
						x = iNumEnd + 1
						state = stateFindNumber
						found = true
						break
					}
				}

				if found {
					continue
				}
			}

			// Not a part number, skip it
			if isDebug() {
				surroundPrint(input, y, iNumStart, iNumEnd)
			}

			x = iNumEnd + 1
			state = stateFindNumber
			continue

		default:
			log.Logger.Fatal().Msgf("Unknown parser state: %d", state)
		}
	}

	return result
}

func calcGearRatio(input [][]rune) {
	state := stateFindGear
	x := 0
	y := 0
	done := false

	for !done {
		switch state {
		case stateFindGear:
			// Find a gear by walking the arrays
			if x == len(input[y]) {
				y++
				x = 0
			}

			if y >= len(input) {
				done = true
				continue
			}

			if isGear(input[y][x]) {
				state = stateFindNumber
			} else {
				x++
			}

		case stateFindNumber:

		case stateGearRatio:

		default:
			log.Fatal().Msgf("Unknown parser state :%d", state)
		}
	}
}

func isNumber(r rune) bool {
	ir := int(r)
	return ir >= int('0') && ir <= int('9')
}

func isSymbol(r rune) bool {
	return !isNumber(r) && int(r) != int('.')
}

func isGear(r rune) bool {
	return r == '*'
}

func toInt(r rune) int {
	return int(r) - int('0')
}

func toPartNum(input [][]rune, y, start, end int) int {
	result := 0
	for i := start; i <= end; i++ {
		result += toInt(input[y][i])
		if i != end {
			result *= 10
		}
	}

	return result
}

func surroundPrint(input [][]rune, y, start, end int) {
	b := &strings.Builder{}
	b.WriteString(fmt.Sprintf("line:\"%d\" start:\"%d\" end:\"%d\"\n", y+1, start+1, end+1))
	if y-1 >= 0 { // above
		startX := start
		if start-1 >= 0 {
			startX = start - 1
		}

		endX := end
		if end+1 <= len(input[y])-1 {
			endX = end + 1
		}

		for i := startX; i <= endX; i++ {
			b.WriteRune(input[y-1][i])
		}
	}
	b.WriteRune('\n')

	// current line
	s := start
	if start-1 >= 0 {
		s = start - 1
	}

	e := end
	if end+1 <= len(input[y])-1 {
		e = end + 1
	}

	for i := s; i <= e; i++ {
		b.WriteRune(input[y][i])
	}

	b.WriteRune('\n')
	if y+1 <= len(input[y])-1 { // below
		startX := start
		if start-1 >= 0 {
			startX = start - 1
		}

		endX := end
		if end+1 <= len(input[y])-1 {
			endX = end + 1
		}

		for i := startX; i <= endX; i++ {
			b.WriteRune(input[y+1][i])
		}
	}
	b.WriteRune('\n')
	fmt.Println(b.String())
}
