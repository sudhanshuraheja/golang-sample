package logger

import (
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sudhanshuraheja/golang-sample/pkg/config"
)

var logger *logrus.Logger

func Init() {
	level, err := logrus.ParseLevel(config.LogLevel())
	if err != nil {
		log.Fatalf(err.Error())
	}

	logger = &logrus.Logger{
		Out:       os.Stdout,
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
		Formatter: &logrus.TextFormatter{},
	}
}

func AddHook(hook logrus.Hook) {
	logger.Hooks.Add(hook)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

func Errorrf(r *http.Request, format string, args ...interface{}) {
	httpRequestLogEntry(r).Errorf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

func Inforf(r *http.Request, format string, args ...interface{}) {
	httpRequestLogEntry(r).Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func httpRequestLogEntry(r *http.Request) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"RequestMethod": r.Method,
		"Host":          r.Host,
		"Path":          r.URL.Path,
	})
}
