package gacha

import (
	"github.com/labstack/echo/v4"
)

func GachaRoute(e *echo.Echo, handler *GachaHandler) {
	e.GET("/gachas", handler.GetGachas)
	e.GET("/gachas/:id", handler.GetGacha)
	e.POST("/gachas", handler.CreateGacha)
	e.POST("/gachas/:id", handler.DrawGacha)
	e.PATCH("/gachas/:id", handler.UpdateGacha)
	e.DELETE("gachas/:id", handler.DeleteGacha)
}

