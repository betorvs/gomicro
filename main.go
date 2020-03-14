package main

import (
	"github.com/betorvs/gomicro/config"
	"github.com/betorvs/gomicro/controller"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Enable /metrics for prometheus
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	g := e.Group("/gomicro/v1")
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g.GET("/health", controller.CheckHealth)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
