package main

import (
	"fmt"

	"github.com/Lyalyashechka/actresses/config"
	"github.com/Lyalyashechka/actresses/config/middleware"
	"github.com/Lyalyashechka/actresses/config/routing"
	actresses_handler "github.com/Lyalyashechka/actresses/internal/actresses/handler/http"
	actresses_usecase "github.com/Lyalyashechka/actresses/internal/actresses/useCase"
	hard_logger "github.com/Lyalyashechka/actresses/internal/logger"
	"github.com/labstack/echo/v4"
)

const (
	CONFIG_PATH = "config"
)

func main() {
	router := echo.New()

	configApp := config.Config{}
	err := config.LoadConfig(&configApp, CONFIG_PATH)
	if err != nil {
		panic(err)
	}

	logger, err := hard_logger.NewLogrusLogger(configApp.Logger)
	if err != nil {
		panic(err)
	}

	actressesUseCase := actresses_usecase.New(logger)
	actressesHandler := actresses_handler.New(logger, actressesUseCase)

	configRouting := routing.RoutingConfig{
		ActressesHandler: actressesHandler,
	}

	middleware.SetMiddlewares(router)
	configRouting.SetRoutes(router)

	router.Logger.Fatal(router.Start(configApp.Server.Host + ":" + configApp.Server.Port))

	fmt.Println("xyi")
}
