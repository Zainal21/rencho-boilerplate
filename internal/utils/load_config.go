package utils

import (
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func LoadConfig(path string) *config.Env {
	env := config.Env{}
	viper.SetConfigFile(path)

	e := echo.New()

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		e.Logger.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		e.Logger.Infof("The App is running in development env")
	}

	return &env
}
