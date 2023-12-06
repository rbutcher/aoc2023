/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day6

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var day6sample = `Time:      7  15   30
Distance:  9  40  200`

func TestMultiplyWinningHoldTimeInstances(t *testing.T) {
	actual := multiplyWinningHoldTimeInstances(day6sample)
	assert.Equal(t, actual, 288)
}

func TestSingleRaceHoldTimeInstances(t *testing.T) {
	actual := singleRaceHoldTimeInstances(day6sample)
	assert.Equal(t, actual, 71503)
}
