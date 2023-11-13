package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				log.Error(err)
				context.NoContent(http.StatusInternalServerError)
			}
		}()

		return next(context)
	}
}
