package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pro-active/cta_gacha/app/gacha/api/gacha"
	"github.com/pro-active/cta_gacha/app/gacha/api/item"
	"github.com/pro-active/cta_gacha/app/gacha/api/user"
	"gorm.io/gorm"
)

func InitalizeServer(db *gorm.DB) *echo.Echo {

	e := echo.New()
	handlers := initializeHandlers(db)

	user.UserRoute(e, handlers.User)
	item.ItemRoute(e, handlers.Item)
	gacha.GachaRoute(e, handlers.Gacha)
	e.GET("/server-health", healthCheck)
	return e
}

func healthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

type handlers struct {
	User  *user.UserHandler
	Item  *item.ItemHandler
	Gacha *gacha.GachaHandler
}

func initializeHandlers(db *gorm.DB) handlers {
	return handlers{
		User:  user.InitializeUserHandler(db),
		Item:  item.InitializeItemHandler(db),
		Gacha: gacha.InitializeGachaHandler(db),
	}
}
