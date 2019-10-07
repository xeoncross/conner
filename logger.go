package conner

import (
	"log"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
)

// Logger instance for logging errors with context
type Logger struct {
	logger  *log.Logger
	logrus  *logrus.Logger
	zerolog *zerolog.Logger
}

// Error logging with structured values from error chain
func (l *Logger) Error(err error) {

	values := Values(err)

	if l.logger != nil {
		l.logger.Printf("%s %+v\n", err, values)
	}

	if l.logrus != nil {
		l.logrus.WithFields(logrus.Fields(values)).Error(err)
	}

	if l.zerolog != nil {
		// https://github.com/rs/zerolog/blob/master/event.go#L148
		l.zerolog.Err(err).Fields(values).Msg("")
	}
}
