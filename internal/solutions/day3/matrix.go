/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day3

import (
	"strconv"
	"strings"
)

const (
	zero = uint8('0')
	nine = uint8('9')
)

type Matrix []string

// FromString creates a new Matrix from a string by splitting on new line characters.
func FromString(source string) Matrix {
	return strings.Split(source, "\n")
}

func (m Matrix) GetAdjacentNumbers(x, y int) []int {
	numbers := make(map[int]bool)
	for _, d := range directions {
		sx := d.X + x
		if sx < 0 || sx >= len(m[y]) {
			continue
		}

		sy := d.Y + y
		if sy < 0 || sy >= len(m) {
			continue
		}

		c := m[sy][sx]
		if isNumber(c) {
			n := m.readNumber(sx, sy)
			numbers[n] = true
		}
	}

	var result []int
	for k := range numbers {
		result = append(result, k)
	}

	return result
}

func (m Matrix) readNumber(x, y int) int {
	start := x
	end := x
	for {
		cont := false
		if start-1 >= 0 && isNumber(m[y][start-1]) {
			cont = true
			start = start - 1
		}

		if end+1 < len(m[y]) && isNumber(m[y][end+1]) {
			cont = true
			end = end + 1
		}

		if !cont {
			break
		}
	}

	rn := m[y][start : end+1]
	result, err := strconv.Atoi(rn)
	if err != nil {
		return -1
	}

	return result
}

func isNumber(c uint8) bool {
	return c >= zero && c <= nine
}
