package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

func init() {
	logrus.SetFormatter(&Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	//logrus.SetOutput(os.Stdout) // TODO: write to files or es.
}

func SetLevel(level string) {
	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func SetLogFile(path string) {
	if path != "" {
		path = strings.TrimSuffix(path, "/") + "/log"
	} else {
		path = "./log"
	}

	writer, err := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),
	)
	if err != nil {
		fmt.Print("init log file error:", err)
		os.Exit(1)
	}

	mutileWriter := io.MultiWriter(writer, os.Stdout)

	logrus.SetOutput(mutileWriter)

}
