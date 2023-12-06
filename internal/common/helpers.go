/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package common

import "strings"

func ReadNumbers(line string) []int {
	line = strings.Trim(line, " ")

	buffer := -1
	var numbers []int
	for i := 0; i < len(line); i++ {
		n := ToNumber(line[i])
		if n != -1 {
			if buffer == -1 {
				buffer = 0
			}

			buffer *= 10
			buffer += n
		}

		if (IsSpace(line[i]) || i+1 == len(line)) && buffer != -1 {
			numbers = append(numbers, buffer)
			buffer = -1
		}
	}

	return numbers
}

func IsSpace(c uint8) bool {
	return c == uint8(' ')
}

func ToNumber(c uint8) int {
	const zero = uint8('0')
	const nine = uint8('9')
	if c >= zero && c <= nine {
		return int(c - zero)
	} else {
		return -1
	}
}
