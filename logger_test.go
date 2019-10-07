package conner

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {

	var buf bytes.Buffer
	l := logrus.StandardLogger()
	l.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: true})
	l.SetOutput(&buf)

	logger := Logger{
		Logrus: l,
	}

	err := demoErrorStack()

	logger.Error(err)

	result := `{"a":"a","b":"b","level":"error","msg":"F1: F2: F3 Error"}` + "\n"

	if buf.String() != result {
		t.Errorf("%q\n\t%q", buf.String(), result)
	}

}
