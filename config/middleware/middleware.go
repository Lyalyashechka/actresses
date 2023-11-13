package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMiddlewares(router *echo.Echo) {
	router.Use(
		Recover,
		GetCORSMiddleware(),
		RequestID(),
		middleware.Secure(),
	)
}
