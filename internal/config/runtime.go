package config

import "github.com/sirupsen/logrus"

var (
	LogLevel     string
	Output       string
	Bootstrapper string
)

func init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)
}
