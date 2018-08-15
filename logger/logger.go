package logger

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

type Logger struct {
	*logrus.Logger
}

func New() *Logger {
	formatter := &logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	}

	log := logrus.New()
	log.Formatter = formatter
	log.Out = colorable.NewColorableStdout()
	return &Logger{log}
}

func (log Logger) WithRequest(r *http.Request) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"method":   r.Method,
		"endpoint": r.URL.Path,
	})
}
