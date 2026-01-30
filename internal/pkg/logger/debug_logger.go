package mylog

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func Debug() {
	Infof("haloo")
}

func Infof(format string, args ...any) {
	logrus.SetFormatter(config())
	colortext := ColorizeInfo(format)
	logrus.Infof(colortext, args...)
}

func Errorf(format string, args ...any) {
	logrus.SetFormatter(config())
	colortext := ColorizeError(format)
	logrus.Errorf(colortext, args...)
}

func Panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func Infoln(v ...any) {
	logrus.SetFormatter(config())

	if len(v) == 1 {
		b, err := json.MarshalIndent(v[0], "", "  ")
		if err == nil {
			logrus.Info("\n" + ColorizeINFOJSON(string(b)))
			return
		}
	}

	logrus.Info(strings.TrimSuffix(fmt.Sprintln(v...), "\n"))
}

func Errorln(v ...any) {
	logrus.SetFormatter(config())

	if len(v) == 1 {
		b, err := json.MarshalIndent(v[0], "", "  ")
		if err == nil {
			logrus.Error("\n" + ColorizeErrorJSON(string(b)))
			return
		}
	}

	logrus.Error(strings.TrimSuffix(fmt.Sprintln(v...), "\n"))
}

func config() logrus.Formatter {
	return &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
}
