package http

import (
	"fmt"

	docs "github.com/Zainal21/renco-boilerplate/docs"
	"github.com/Zainal21/renco-boilerplate/internal/api/route"
	"github.com/Zainal21/renco-boilerplate/internal/bootstrap"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Renco Boilerplate
//	@version		0.1
//	@description	An API Boilerplate

//	@license.name	MIT
//	@license.url	https://opensource.org/license/MIT

//	@BasePath	/api/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Firebase auth access token, get it from POST /auth/access-token

func Start() {
	logger.SetJSONFormatter()
	app := bootstrap.App()
	env := app.Env
	db := app.DB
	defer app.CloseDBConnection()
	firebaseAuth := app.FirebaseAuth
	loggerUtil := logger.NewLoggerSetup(env)
	docs.SwaggerInfo.Host = env.Host

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:        true,
		LogError:      true,
		LogMethod:     true,
		LogStatus:     true,
		LogValuesFunc: loggerUtil.EchoMiddlewareFunc(),
	}))

	route.Setup(env, loggerUtil, db, firebaseAuth, e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", env.Host, env.Port)))
}
