package routes

import (
	"github.com/labstack/echo/v4"
	"golang-echo-kafka-example/controller"
)

func GetUserApiRoutes(e *echo.Echo, userController *controller.UserController) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("/signup", userController.SaveUser)
	}
}
