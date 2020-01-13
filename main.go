package main

import (
	"github.com/betorvs/gomicro/config"
	"github.com/betorvs/gomicro/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	e := echo.New()
	g := e.Group("/gomicro/v1")
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g.GET("/health", controller.CheckHealth)
	g.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.Logger.Fatal(e.Start(":" + config.Port))
}
