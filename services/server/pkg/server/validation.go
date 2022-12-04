package server

import (
	"fmt"
	"net/http"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct{}

func (cv *CustomValidator) Validate(i interface{}) error {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		opt.SkipOnEmpty = false
	})

	valid := validate.Struct(i)
	if !valid.Validate() {
		fmt.Println(valid.Errors)
		return echo.NewHTTPError(http.StatusBadRequest, valid.Errors)
	}
	return nil
}
