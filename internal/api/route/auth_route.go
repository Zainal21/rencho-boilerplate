package route

import (
	controller "github.com/Zainal21/renco-boilerplate/internal/api/controllers"
	"github.com/Zainal21/renco-boilerplate/internal/api/middleware"
	"github.com/Zainal21/renco-boilerplate/internal/domain"
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewAuthRouter(env *config.Env, loggerUtil logger.LoggerUtils, rootGroup *echo.Group, authUsecase domain.AuthUsecase, authMiddleware middleware.AuthMiddleware, validate *validator.Validate) {
	ct := controller.NewAuthController(env, loggerUtil, authUsecase, validate)

	publicGroup := rootGroup.Group("/v1/auth")
	privateGroup := rootGroup.Group("/v1/auth")
	privateGroup.Use(authMiddleware.ValidateUser())

	publicGroup.POST("/signup", ct.SignUp)
	publicGroup.POST("/token", ct.GetAccessToken)
}
