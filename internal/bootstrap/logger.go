package bootstrap

import (
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	"github.com/Zainal21/renco-boilerplate/pkg/util"
	"github.com/sirupsen/logrus"
)

func RegistryLogger(cfg *config.Env) {
	logger.Setup(logger.Config{
		Environment: util.EnvironmentTransform(cfg.AppEnv),
		Debug:       cfg.AppLoggerDebug,
		Level:       cfg.AppLoggerLevel,
		ServiceName: cfg.AppName,
		Hooks:       []logrus.Hook{}, // Add Hook list here
	})
}
