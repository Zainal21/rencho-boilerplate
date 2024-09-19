// Package logger
package logger

import (
	"github.com/sirupsen/logrus"
)

func Setup(cfg Config) {
	if cfg.Debug {
		logrus.SetLevel(logrus.TraceLevel)
		return
	}

	lvl, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		lvl = logrus.InfoLevel
	}

	logrus.SetLevel(lvl)

	for _, hook := range cfg.Hooks {
		logrus.AddHook(hook)
	}
}
