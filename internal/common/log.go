/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package common

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func ConfigureLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func LoggerWithScope(scope string) zerolog.Logger {
	return log.Logger.With().Str("scope", scope).Logger()
}
