package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type LoggerUtils interface {
	Debugf(format string, args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	EchoMiddlewareFunc() func(c echo.Context, values middleware.RequestLoggerValues) error
}

type Config struct {
	Debug       bool   `json:"debug"`
	Environment string `json:"environment"`
	Level       string `json:"level"`
	ServiceName string `json:"service_name"`
	Hooks       []logrus.Hook
}
