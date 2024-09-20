package http

import (
	"errors"
	"fmt"
	"time"

	docs "github.com/Zainal21/renco-boilerplate/docs"
	"github.com/Zainal21/renco-boilerplate/internal/api/route"
	"github.com/Zainal21/renco-boilerplate/internal/bootstrap"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	"github.com/Zainal21/renco-boilerplate/pkg/response"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/time/rate"
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

	// config cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		MaxAge: 300,
		AllowOrigins: []string{
			"http://*",
			"https://*",
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}))

	// config rate limiter
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(1), Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return response.FromForbiddenError(errors.New("error while extracting identifier")).WithEcho(context)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return response.FromForbiddenError(errors.New("to Many Request")).WithEcho(context)
		},
	}

	e.Use(middleware.RateLimiterWithConfig(config))

	route.Setup(env, loggerUtil, db, firebaseAuth, e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", env.Host, env.Port)))
}
