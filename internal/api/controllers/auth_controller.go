package controllers

import (
	"errors"

	"github.com/Zainal21/renco-boilerplate/internal/domain"
	"github.com/Zainal21/renco-boilerplate/internal/dtos"
	"github.com/Zainal21/renco-boilerplate/pkg/config"
	"github.com/Zainal21/renco-boilerplate/pkg/logger"
	"github.com/Zainal21/renco-boilerplate/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type baseAuthController struct {
	env         *config.Env
	loggerUtil  logger.LoggerUtils
	authUsecase domain.AuthUsecase
	validate    *validator.Validate
}

func NewAuthController(env *config.Env, loggerUtil logger.LoggerUtils, authUsecase domain.AuthUsecase, validate *validator.Validate) AuthController {
	return &baseAuthController{
		env:         env,
		loggerUtil:  loggerUtil,
		authUsecase: authUsecase,
		validate:    validate,
	}
}

// SignUp godoc
//
//	@Summary	Create user
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		credential body	domain.AuthControllerPayloadSignUp true	"email and password"
//	@Success	201
//	@Failure	400	"validation error | user already exist"
//	@Failure	500	"Internal Server Error"
//	@Router		/auth/signup [post]
func (b *baseAuthController) SignUp(c echo.Context) error {
	var payload dtos.AuthControllerPayloadSignUp
	err := c.Bind(&payload)
	if err != nil {
		return response.FromBindingError(err).WithEcho(c)
	}
	err = b.validate.Struct(&payload)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return response.FromValidationErrors(validationErrors).WithEcho(c)
		}
	}

	err = b.authUsecase.SignUp(c.Request().Context(), payload.Email, payload.Password, payload.IsAdmin)
	if err != nil {
		if err.Error() == "user already exist" {
			return response.FromBadRequestError(err).WithEcho(c)
		}

		return response.FromError(err).WithEcho(c)
	}

	return response.FromCreated().WithEcho(c)
}

// GetAccessToken godoc
//
//	@Summary	Get access token
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		credential body domain.AuthControllerPayloadGetAccessToken	true "email and password"
//	@Success	200
//	@Failure	400	"validation error"
//	@Failure	404	"user not found"
//	@Failure	500	"Internal Server Error"
//	@Router		/auth/access-token [post]
func (b *baseAuthController) GetAccessToken(c echo.Context) error {
	var payload dtos.AuthControllerPayloadGetAccessToken
	err := c.Bind(&payload)
	if err != nil {
		return response.FromBindingError(err).WithEcho(c)
	}
	err = b.validate.Struct(&payload)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return response.FromValidationErrors(validationErrors).WithEcho(c)
		}
	}

	accessToken, err := b.authUsecase.GetAccessToken(c.Request().Context(), payload.Email, payload.Password)
	if err != nil {
		if err.Error() == "user not found" {
			return response.FromNotFoundError(err).WithEcho(c)
		}

		return response.FromError(err).WithEcho(c)
	}

	return response.FromData(accessToken).WithEcho(c)
}
