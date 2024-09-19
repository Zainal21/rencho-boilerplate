package logger

import (
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type baseLoggerUtils struct {
	env    *config.Env
	logger *logrus.Logger
}

func SetJSONFormatter() {
	logrus.SetFormatter(&Formatter{
		ChildFormatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: MessageKey,
			},
		},
		Line:         true,
		Package:      false,
		File:         true,
		BaseNameOnly: false,
	})
}

func NewLoggerSetup(env *config.Env) LoggerUtils {
	logger := logrus.New()
	if env.AppEnv != "prod" {
		logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetLevel(logrus.InfoLevel)

	}

	return &baseLoggerUtils{env: env, logger: logger}
}

func (b *baseLoggerUtils) EchoMiddlewareFunc() func(c echo.Context, values middleware.RequestLoggerValues) error {
	return func(c echo.Context, values middleware.RequestLoggerValues) error {
		logData := b.logger.WithFields(logrus.Fields{
			"URI":    values.URI,
			"method": values.Method,
			"status": values.Status,
		})
		if values.Status >= 300 {
			logData.Errorf("failed request with status %d", values.Status)
		} else {
			logData.Infof("success request with status %d", values.Status)
		}

		return nil
	}
}

func (b *baseLoggerUtils) Fatalf(format string, args ...interface{}) {
	b.logger.Fatalf(format, args...)
}

func (b *baseLoggerUtils) Debugf(format string, args ...interface{}) {
	b.logger.Debugf(format, args...)
}

func (b *baseLoggerUtils) Infoln(args ...interface{}) {
	b.logger.Infoln(args...)
}

func (b *baseLoggerUtils) Infof(format string, args ...interface{}) {
	b.logger.Infof(format, args...)
}

func (b *baseLoggerUtils) Warnf(format string, args ...interface{}) {
	b.logger.Warnf(format, args...)
}

func (b *baseLoggerUtils) Errorf(format string, args ...interface{}) {
	b.logger.Errorf(format, args...)
}