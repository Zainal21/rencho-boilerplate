package route

import (
	controller "github.com/Zainal21/renco-boilerplate/internal/api/controllers"
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewHealthRoute(env *config.Env, loggerUtil logger.LoggerUtils, rootGroup *echo.Group, validate *validator.Validate) {
	ct := controller.NewHealthController(env, loggerUtil, validate)

	publicGroup := rootGroup.Group("/v1")

	publicGroup.GET("/ping", ct.Ping)
}