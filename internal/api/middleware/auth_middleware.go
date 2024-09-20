package middleware

import (
	"errors"
	"strings"

	"github.com/Zainal21/renco-boilerplate/internal/domain"
	"github.com/Zainal21/renco-boilerplate/pkg/response"
	"github.com/labstack/echo/v4"
)

type baseAuthMiddleware struct {
	userUsecase    domain.UserUsecase
	jwtAuthUseCase domain.JwtAuthUseCase
}

func NewAuthMiddleware(userUsecase domain.UserUsecase, jwtAuthUseCase domain.JwtAuthUseCase) AuthMiddleware {
	return &baseAuthMiddleware{
		userUsecase:    userUsecase,
		jwtAuthUseCase: jwtAuthUseCase,
	}
}

func (b *baseAuthMiddleware) ValidateUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerToken := c.Request().Header.Get("authorization")
			if !strings.HasPrefix(bearerToken, "Bearer ") {
				return response.FromForbiddenError(errors.New("invalid access token")).WithEcho(c)
			}

			token := strings.Split(bearerToken, " ")[1]
			firebaseUID, err := b.jwtAuthUseCase.VerifyToken(token)
			if err != nil {
				return response.FromForbiddenError(err).WithEcho(c)
			}
			user, err := b.userUsecase.GetUserByFirebaseUID(c.Request().Context(), firebaseUID)
			if user == nil {
				return response.FromForbiddenError(errors.New("access denied")).WithEcho(c)
			}
			if err != nil {
				return response.FromInternalServerError().WithEcho(c)
			}
			c.Set("user", user)

			return next(c)
		}
	}
}
