/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package matrix

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdjacentNumbers(t *testing.T) {
	t.Run("single digit single number returns expected", func(t *testing.T) {
		testCases := map[string]int{
			".......\n...&4..\n.......": 4,
			".......\n..8&...\n.......": 8,
			"...6...\n...&...\n.......": 6,
			"..7....\n...&...\n.......": 7,
			"....3..\n...&...\n.......": 3,
			".......\n...&...\n...1...": 1,
			".......\n...&...\n..9....": 9,
			".......\n...&...\n....2..": 2,
		}

		for source, expected := range testCases {
			t.Run(fmt.Sprintf("%d", expected), func(t *testing.T) {
				m := FromString(source)
				actual := m.GetAdjacentNumbers(3, 1)
				require.Equal(t, 1, len(actual))
				assert.Equal(t, actual[0], expected)
			})
		}
	})

	t.Run("two digit single number returns expected", func(t *testing.T) {
		testCases := map[string]int{
			"...42..\n...%...\n.......": 42,
			"..34...\n...%...\n.......": 34,
			".56....\n...%...\n.......": 56,
			"....78.\n...%...\n.......": 78,
			".......\n...%...\n...48..": 48,
			".......\n...%...\n..44...": 44,
			".......\n...%...\n.35....": 35,
			".......\n...%...\n....88.": 88,
			".......\n...%76.\n.......": 76,
			".......\n.67%...\n.......": 67,
		}

		for source, expected := range testCases {
			t.Run(fmt.Sprintf("%d", expected), func(t *testing.T) {
				m := FromString(source)
				actual := m.GetAdjacentNumbers(3, 1)
				require.Equal(t, 1, len(actual))
				assert.Equal(t, actual[0], expected)
			})
		}
	})

	t.Run("multiple numbers returns expected", func(t *testing.T) {
		testCases := map[string][]int{
			"..42...\n.13$...\n464....": {42, 13, 464},
			"....456\n...$233\n.......": {456, 233},
			"...89..\n...$4..\n..333..": {89, 4, 333},
		}

		for source, expected := range testCases {
			t.Run(fmt.Sprintf("%v", expected), func(t *testing.T) {
				m := FromString(source)
				actual := m.GetAdjacentNumbers(3, 1)

				require.Equal(t, len(expected), len(actual))
				for _, v := range actual {
					require.True(t, contains(expected, v))
				}
			})
		}
	})

	t.Run("against the bounds does not index out of bounds", func(t *testing.T) {
		testCases := []struct {
			Source   string
			X        int
			Y        int
			Expected []int
		}{
			{"%48.\n364.", 0, 0, []int{48, 364}}, // top-left
			{".56%\n.482", 3, 0, []int{56, 482}}, // top-right
			{"432.\n%1.", 0, 1, []int{432, 1}},   // bottom-left
			{".456\n.78%", 3, 1, []int{456, 78}}, // bottom-right
		}

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%v", tc.Expected), func(t *testing.T) {
				m := FromString(tc.Source)
				actual := m.GetAdjacentNumbers(tc.X, tc.Y)
				require.Equal(t, len(tc.Expected), len(actual))
				for _, v := range actual {
					require.True(t, contains(tc.Expected, v))
				}
			})
		}
	})
}

func contains(slice []int, item int) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
