package item

import (
	"github.com/labstack/echo/v4"
)

func ItemRoute(e *echo.Echo, handler *ItemHandler) {
	e.POST("/items", handler.CreateItem)
}

