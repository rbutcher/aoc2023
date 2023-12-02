/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package solutions

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func logWithScope(scope string) zerolog.Logger {
	return log.Logger.With().Str("scope", scope).Logger()
}
