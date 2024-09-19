package route

import (
	"firebase.google.com/go/v4/auth"
	"github.com/Zainal21/renco-boilerplate/internal/api/middleware"
	"github.com/Zainal21/renco-boilerplate/internal/repositories"
	"github.com/Zainal21/renco-boilerplate/internal/usecase"
	"github.com/Zainal21/renco-boilerplate/internal/utils"
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Setup(env *config.Env, loggerUtil logger.LoggerUtils, db *sqlx.DB, firebaseAuth *auth.Client, e *echo.Echo) {
	authUtil := utils.NewAuthUtil(env, firebaseAuth)
	userRepo := repositories.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(env, userRepo, authUtil)
	userUsecase := usecase.NewUserUsecase(env, userRepo)
	authMiddleware := middleware.NewAuthMiddleware(userUsecase, authUtil)
	validate := validator.New()

	// route group define
	rootPathGroup := e.Group("")
	rootApiGroup := e.Group("/api")

	NewHealthRoute(env, loggerUtil, rootPathGroup, validate)
	NewAuthRouter(env, loggerUtil, rootApiGroup, authUsecase, authMiddleware, validate)
}
