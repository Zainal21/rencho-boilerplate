package utils

import (
	"fmt"

	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	"github.com/spf13/viper"
)

func LoadConfig(path string) *config.Env {
	env := config.Env{}
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Can't find the file .env %+v)", err))
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Environment can't be loaded %+v)", err))
	}

	if env.AppEnv == "development" {
		logger.Info(fmt.Sprintf("The App is running in development mode)"))
	}

	return &env
}
