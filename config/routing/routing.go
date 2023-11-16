package routing

import (
	actresses_handler "github.com/Lyalyashechka/actresses/internal/actresses/handler/http"
	"github.com/labstack/echo/v4"
)

type RoutingConfig struct {
	ActressesHandler *actresses_handler.Handler
}

func (cr *RoutingConfig) SetRoutes(router *echo.Echo) {
	router.GET("actresess", cr.ActressesHandler.GetAllActresses)
	router.GET("actresess/top", cr.ActressesHandler.GetTopActresses)
	router.POST("actresess/vote", cr.ActressesHandler.VoteActress)
	router.Static("staticback", "static")
}
