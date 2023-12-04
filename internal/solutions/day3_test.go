/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package solutions

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var source = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestSumPartNumbers(t *testing.T) {
	r := sumPartNumbers(source)
	assert.Equal(t, r, 4361)
}

func TestSumGearRations(t *testing.T) {
	r := sumGearRatios(source)
	assert.Equal(t, r, 467835)
}
