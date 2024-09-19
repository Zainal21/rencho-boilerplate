package response

import (
	"net/http"
	"strings"
	"time"

	"github.com/Zainal21/renco-boilerplate/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidationError struct {
	Field string `json:"field"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Response struct {
	RefId            string            `json:"ref_id"`
	Code             int               `json:"code"`
	Status           string            `json:"status"`
	Data             interface{}       `json:"data,omitempty"`
	Error            string            `json:"error,omitempty"`
	BindingError     string            `json:"binding_error,omitempty"`
	ValidationErrors []ValidationError `json:"validation_errors,omitempty"`
	Timestamp        time.Time         `json:"timestamp,omitempty"`
}

func FromOK() *Response {
	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(http.StatusOK),
		Code:      http.StatusOK,
		Timestamp: time.Now().Local(),
	}
}

func FromCreated() *Response {
	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(http.StatusCreated),
		Timestamp: time.Now().Local(),
		Code:      http.StatusCreated,
	}
}

func FromData(data interface{}) *Response {
	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(http.StatusOK),
		Code:      http.StatusOK,
		Timestamp: time.Now().Local(),
		Data:      data,
	}
}

func FromError(err error) *Response {
	if strings.Contains(strings.ToLower(err.Error()), "not found") {
		return &Response{
			RefId:     util.GenerateRefID(10),
			Status:    http.StatusText(http.StatusNotFound),
			Code:      http.StatusNotFound,
			Timestamp: time.Now().Local(),
			Error:     err.Error(),
		}
	}

	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(http.StatusInternalServerError),
		Code:      http.StatusInternalServerError,
		Timestamp: time.Now().Local(),
		Error:     err.Error(),
	}
}

func FromBindingError(err error) *Response {
	return &Response{
		RefId:        util.GenerateRefID(10),
		Status:       http.StatusText(http.StatusBadRequest),
		Code:         http.StatusBadRequest,
		Timestamp:    time.Now().Local(),
		BindingError: err.Error(),
	}
}

func FromValidationError(err ValidationError) *Response {
	return &Response{
		Status: http.StatusText(http.StatusBadRequest),
		Code:   http.StatusBadRequest,
		ValidationErrors: []ValidationError{
			err,
		},
		Timestamp: time.Now().Local(),
		RefId:     util.GenerateRefID(10),
	}
}

func FromValidationErrors(_errs validator.ValidationErrors) *Response {
	var errs []ValidationError

	for _, err := range _errs {
		errs = append(errs, ValidationError{
			Field: strings.ToLower(err.StructField()),
			Name:  err.Tag(),
			Value: err.Param(),
		})
	}

	return &Response{
		RefId:            util.GenerateRefID(10),
		Status:           http.StatusText(http.StatusBadRequest),
		Code:             http.StatusBadRequest,
		ValidationErrors: errs,
		Timestamp:        time.Now().Local(),
	}
}

func FromBadRequestError(err error) *Response {
	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(http.StatusBadRequest),
		Code:      http.StatusBadRequest,
		Error:     err.Error(),
		Timestamp: time.Now().Local(),
	}
}

func FromForbiddenError(err error) *Response {
	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(http.StatusForbidden),
		Code:      http.StatusForbidden,
		Error:     err.Error(),
		Timestamp: time.Now().Local(),
	}
}

func FromNotFoundError(err error) *Response {
	if err != nil {
		return &Response{
			RefId:     util.GenerateRefID(10),
			Status:    http.StatusText(http.StatusNotFound),
			Code:      http.StatusNotFound,
			Error:     err.Error(),
			Timestamp: time.Now().Local(),
		}
	}

	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(http.StatusNotFound),
		Code:      http.StatusNotFound,
		Timestamp: time.Now().Local(),
	}
}

func FromInternalServerError() *Response {
	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(http.StatusInternalServerError),
		Code:      http.StatusNotFound,
		Timestamp: time.Now().Local(),
	}
}

func (r *Response) WithCode(code int) *Response {
	return &Response{
		RefId:     util.GenerateRefID(10),
		Status:    http.StatusText(code),
		Code:      code,
		Timestamp: time.Now().Local(),
	}
}

func (r *Response) WithEcho(c echo.Context) error {
	return c.JSON(r.Code, r)
}
