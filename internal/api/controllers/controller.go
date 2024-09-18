package controllers

import "github.com/labstack/echo/v4"

type AuthController interface {
	GetAccessToken(c echo.Context) error
	SignUp(c echo.Context) error
}

type HealthCheckController interface {
	Ping(c echo.Context) error
}
