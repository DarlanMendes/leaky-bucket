package main

import (
	"leaky-bucket/api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.MessageRoute(e)
	e.Logger.Fatal(e.Start(":9000"))
}
