package observability

import "github.com/rs/zerolog"

func NewLogger() zerolog.Logger {
	l := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()
	return l
}
