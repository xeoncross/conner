package conner

import (
	"log"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
)

// TODO some kind of instance with helper functions that call log()?
type Logger struct {
	logger  *log.Logger
	logrus  *logrus.Logger
	zerolog *zerolog.Logger
}

func (l *Logger) log(level string, err error) {

	// var msg string
	// var values map[string]interface{}
	//
	// switch v := message.(type) {
	// case error:
	// 	msg = v.Error()
	// 	values = Values(v)
	// case string:
	// 	msg = v
	// default:
	// 	panic("Invalid message type")
	// }

	values := Values(err)

	if l.logger != nil {
		l.logger.Printf("%s %+v\n", err, values)
	}

	if l.logrus != nil {
		l.logrus.WithFields(logrus.Fields(values)).Error(err)
	}
}

// TODO manual functions for this?
// func Logrus(log *logrus.Logger, err error) {
// 	values := Values(err)
//
// 	if values != nil {
// 		log.WithFields(logrus.Fields(values)).Error(err)
// 	} else {
// 		log.Error(err)
// 	}
// }
