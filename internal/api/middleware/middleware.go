package middleware

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	ValidateUser() echo.MiddlewareFunc
}
