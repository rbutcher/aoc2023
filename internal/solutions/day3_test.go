/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package solutions

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var source = `
467..114..
...*......
..35..633.
......#...
617*......
....4+.58.
..592.....
......755.
...$.*....
.664.598..
`

var array = [][]rune{
	[]rune("467..114.."),
	[]rune("...*......"),
	[]rune("..35..633."),
	[]rune("......#..."),
	[]rune("617*......"),
	[]rune("....4+.58."),
	[]rune("..592....."),
	[]rune("......755."),
	[]rune("...$.*...."),
	[]rune(".664.598.."),
}

func TestConvertInput(t *testing.T) {
	expected := array
	actual := convertInput(source)

	require.Equal(t, len(actual), len(expected))
	for i := range actual {
		assert.Equal(t, actual[i], expected[i])
	}
}

func TestIsNumber(t *testing.T) {
	var testCases = map[rune]bool{
		'A': false,
		'a': false,
		'5': true,
	}

	for r, expected := range testCases {
		t.Run(fmt.Sprintf("%v", r), func(t *testing.T) {
			assert.Equal(t, isNumber(r), expected)
		})
	}
}

func TestSumPartNumbers(t *testing.T) {
	expected := 4365
	actual := sumPartNumbers(array)
	assert.Equal(t, expected, actual)
}

func TestCalcGearRatio(t *testing.T) {
	t.Run("single digit numbers", func(t *testing.T) {
	})

}
