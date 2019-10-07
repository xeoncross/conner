package conner

import (
	"log"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
)

// Logger instance for logging errors with context
type Logger struct {
	Logger  *log.Logger
	Logrus  *logrus.Logger
	Zerolog *zerolog.Logger
}

// Do we need this?
// func NewLogger(log interface{}) *Logger {
// 	var logger = &Logger{}
//
// 	switch l := log.(type) {
// 	case *logrus.Logger:
// 		logger.Logrus = l
// 	case *zerolog.Logger:
// 		logger.Zerolog = l
// 	case log.Logger:
// 		logger.Logger = l
// 	}
//
// 	return logger
// }

// Error logging with structured values from error chain
func (l *Logger) Error(err error) {

	values := Values(err)

	if l.Logger != nil {
		l.Logger.Printf("%s %+v\n", err, values)
	}

	if l.Logrus != nil {
		l.Logrus.WithFields(logrus.Fields(values)).Error(err)
	}

	if l.Zerolog != nil {
		// https://github.com/rs/zerolog/blob/master/event.go#L148
		l.Zerolog.Err(err).Fields(values).Msg("")
	}
}
