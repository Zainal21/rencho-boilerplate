package controllers

import (
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	"github.com/Zainal21/renco-boilerplate/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BaseHealthController struct {
	env        *config.Env
	loggerUtil logger.LoggerUtils
	validate   *validator.Validate
}

// Ping implements HealthCheckController.
func (b *BaseHealthController) Ping(c echo.Context) error {
	return response.FromData(struct {
		Message string `json:"message"`
	}{
		Message: "Waras!",
	}).WithEcho(c)
}

func NewHealthController(env *config.Env, loggerUtil logger.LoggerUtils, validate *validator.Validate) HealthCheckController {
	return &BaseHealthController{
		env:        env,
		loggerUtil: loggerUtil,
		validate:   validate,
	}
}
