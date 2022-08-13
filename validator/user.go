package validator

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var ErrorHandlerUser = func(err error, c echo.Context) {

	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required",
					strings.ToLower(err.Field()))
			case "min":
				report.Message = fmt.Sprintf("%s minimum format %s characters",
					strings.ToLower(err.Field()), err.Param())

			case "number":
				report.Message = fmt.Sprintf("%s value required number",
					err.Field())
			}

			break
		}
	}

	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
