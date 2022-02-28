package extensions

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fredi12345/kuefa-karben/src/rest"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func AutoBind[T any](l *zap.Logger, fn func(c echo.Context, req T) error) echo.HandlerFunc {
	v := validator.New()
	l = l.Named("autobind")

	return func(c echo.Context) error {
		var req T
		err := c.Bind(&req)
		if err != nil {
			l.Error(fmt.Sprintf("failed to bind request for %T", req), zap.Error(err))
			return c.JSON(http.StatusBadRequest, rest.ErrorResponse{
				ErrorCode:   rest.RequestBindFailed,
				Description: "Could not bind the request to the internal request structure. ",
			})
		}

		err = v.StructCtx(c.Request().Context(), req)
		if err != nil {
			if vErr, ok := err.(validator.ValidationErrors); ok {
				attributes := make(map[string]interface{})
				for _, fieldError := range vErr {
					attributes[fieldError.StructField()] = strings.TrimSpace(fmt.Sprintf("%s %s", fieldError.Tag(), fieldError.Param()))
				}

				l.Warn(fmt.Sprintf("request validation failed for %T", req), zap.Error(err))
				return c.JSON(http.StatusBadRequest, rest.ErrorResponse{
					ErrorCode:            rest.RequestValidationFailed,
					Description:          "One or more validations failed for the incoming request.",
					AdditionalAttributes: attributes,
				})
			}
			return echo.ErrBadRequest
		}

		return fn(c, req)
	}
}
