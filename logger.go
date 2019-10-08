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

// Debug message logging with structured values
func (l *Logger) Debug(msg string, fields map[string]interface{}) {

	if l.Logger != nil {
		l.Logger.Printf("DEBUG: %s %+v\n", msg, fields)
	}

	if l.Logrus != nil {
		l.Logrus.WithFields(logrus.Fields(fields)).Debug(msg)
	}

	if l.Zerolog != nil {
		l.Zerolog.Debug().Fields(fields).Msg(msg)
	}
}

// Info message logging with structured values
func (l *Logger) Info(msg string, fields map[string]interface{}) {

	if l.Logger != nil {
		l.Logger.Printf("INFO : %s %+v\n", msg, fields)
	}

	if l.Logrus != nil {
		l.Logrus.WithFields(logrus.Fields(fields)).Info(msg)
	}

	if l.Zerolog != nil {
		l.Zerolog.Info().Fields(fields).Msg(msg)
	}
}

// Error logging with structured values from error chain
func (l *Logger) Error(err error) {

	fields := Values(err)

	if l.Logger != nil {
		l.Logger.Printf("ERROR: %s %+v\n", err, fields)
	}

	if l.Logrus != nil {
		l.Logrus.WithFields(logrus.Fields(fields)).Error(err)
	}

	if l.Zerolog != nil {
		l.Zerolog.Err(err).Fields(fields).Msg("")
	}
}
