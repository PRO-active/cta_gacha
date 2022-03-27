package user

import (
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, handler *UserHandler) {
	e.GET("/users/:id", handler.GetUser)
	e.GET("/users/:id/items", handler.GetUserItems)
	e.POST("/users/login", handler.Login)
	e.POST("/users", handler.CreateUser)
	e.PATCH("/users/:id", handler.UpdateUser)
}
