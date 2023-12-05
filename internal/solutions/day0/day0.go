/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package day0

import (
	_ "embed"
	"github.com/rs/zerolog/log"
)

//go:embed input.txt
var day0Input string

func Day0P1() error {
	l := log.With().Int("day", 0).Int("part", 1).Logger()
	l.Info().Msg(day0Input)
	return nil
}

func Day0P2() error {
	l := log.With().Int("day", 0).Int("part", 2).Logger()
	l.Info().Msg(day0Input)
	return nil
}
