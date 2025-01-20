package routes

import (
	"leaky-bucket/api/controller"
	"leaky-bucket/api/routes/middlewares"

	"github.com/labstack/echo/v4"
)

func MessageRoute(e *echo.Echo) {
	messageController := controller.NewController()
	e.GET("/message", messageController.GetMessages, middlewares.LeakyBucket)
}
