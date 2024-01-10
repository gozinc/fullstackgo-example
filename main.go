package main

import (
	"context"
	"fullstackgo/view"
	"fullstackgo/view/layout"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func main() {
	echoServer := echo.New()
	echoServer.HideBanner = true

	echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://i.imgur.com"},
		AllowMethods: []string{"GET"},
	}))

	echoServer.Static("/static/", "./static")

	echoServer.GET("/", func(c echo.Context) error {
		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, "host", c.Request().Host)

		return view.Index(global.Count).Render(ctx, c.Response().Writer)
	})

	echoServer.POST("/", func(c echo.Context) error {
		ctx := c.Request().Context()

		global.Count += 1

		return layout.CountsButton(global.Count).Render(ctx, c.Response().Writer)
	})

	echoServer.Logger.Fatal(echoServer.Start("127.0.0.1:3000"))
}
