package main

import (
	echotemplate "github.com/foolin/echo-template"
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
	app.Renderer = echotemplate.Default()

	// Routes
	// Home
	app.GET("/", service.Home)
	app.GET("/home", service.Home)
	// Register
	// app.GET("/register", service.RegisterForm)
	// app.POST("/register", service.Register)

	app.Logger.Fatal(app.Start(":1323"))
}
