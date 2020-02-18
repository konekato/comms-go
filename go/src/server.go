package main

import (
	"github.com/foolin/goview/supports/echoview"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"service"
)

func main() {
	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Set Renderer
	app.Renderer = echoview.Default()

	// Routes
	// Home
	app.GET("/", service.HomeView)
	app.GET("/home", service.HomeView)
	// Register
	app.GET("/register", service.RegisterView)
	app.POST("/register", service.Register)

	app.Logger.Fatal(app.Start(":1323"))
}
