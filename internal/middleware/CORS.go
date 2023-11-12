package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	allowOrigins  = []string{"*"}
	allowMethods  = []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions}
	allowHeaders  = []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}
	exposeHeaders = []string{"Authorization"}
)

func GetCORSMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowMethods:     allowMethods,
		AllowHeaders:     allowHeaders,
		ExposeHeaders:    exposeHeaders,
		AllowCredentials: true,
	})
}
