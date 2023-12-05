/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day1

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetCalibrationValueNumeralsOnly(t *testing.T) {
	testCases := map[string]int{
		"1abc2":       12,
		"pqr3stu8vwx": 38,
		"a1b2c3d4e5f": 15,
		"treb7uchet":  77,
	}

	for tc, expected := range testCases {
		t.Run(fmt.Sprintf("%s=>%d", tc, expected), func(t *testing.T) {
			actual := getCalibrationValue(tc)
			assert.Equal(t, actual, expected)
		})
	}
}

func TestGetCalibrationValueMixedNumeralsAndWords(t *testing.T) {
	testCases := map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,
	}

	for tc, expected := range testCases {
		t.Run(fmt.Sprintf("%s=>%d", tc, expected), func(t *testing.T) {
			actual := getCalibrationValue(tc)
			assert.Equal(t, actual, expected)
		})
	}
}
