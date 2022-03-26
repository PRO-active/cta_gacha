package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, handler *UserHandler) {
	e.GET("/users/:id", handler.GetUser)
	e.GET("/users/:id/items", handler.GetUserItems)
	e.POST("/users/login", handler.Login)
	e.POST("/users", handler.CreateUser)
	e.PATCH("/users/:id", handler.UpdateUser)
	// TODO: 消す
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
}

