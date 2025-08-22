package observability

import "github.com/rs/zerolog"

func NewLogger() zerolog.Logger {
	return zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()
}
