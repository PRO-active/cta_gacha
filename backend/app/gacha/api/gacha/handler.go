package gacha

import (
	"net/http"

	"github.com/labstack/echo/v4"
	itemModel "github.com/pro-active/cta_gacha/app/gacha/api/item"
	usecase "github.com/pro-active/cta_gacha/internal/usecase/gacha"
	itemUsecase "github.com/pro-active/cta_gacha/internal/usecase/item"
)

type GachaHandler struct {
	usecase     usecase.GachaUsecase
	itemUsecase itemUsecase.ItemUsecase
}

func NewGachaHandler(uc usecase.GachaUsecase, iu itemUsecase.ItemUsecase) *GachaHandler {
	return &GachaHandler{
		usecase:     uc,
		itemUsecase: iu,
	}
}

func (g *GachaHandler) GetGacha(ctx echo.Context) error {
	gachaID := ctx.Param("id")

	gacha, err := g.usecase.GetGacha(gachaID)
	if err != nil {
		return err
	}

	gachaResponse := convertGacha(gacha)

	return ctx.JSON(http.StatusOK, gachaResponse)
}

func (g *GachaHandler) GetGachas(ctx echo.Context) error {
	gacha, err := g.usecase.GetGachas()
	if err != nil {
		return err
	}
	gachaResponse := make([]*Gacha, 0, len(gacha))
	for i := range gacha {
		gachaResponse = append(gachaResponse, convertGacha(gacha[i]))
	}
	return ctx.JSON(http.StatusOK, gachaResponse)
}

func (g *GachaHandler) DrawGacha(ctx echo.Context) error {
	cookie, err := ctx.Cookie("userID")
	if err != nil {
		return err
	}
	userID := cookie.Value
	gachaID := ctx.Param("id")

	item, err := g.itemUsecase.DrawGacha(userID, gachaID)
	if err != nil {
		return err
	}
	itemResponse := itemModel.ConvertItem(item)
	return ctx.JSON(http.StatusOK, itemResponse)
}

func (g *GachaHandler) CreateGacha(ctx echo.Context) error {
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")

	gacha, err := g.usecase.CreateGacha(name, description)
	if err != nil {
		return err
	}

	gachaResponse := convertGacha(gacha)

	return ctx.JSON(http.StatusOK, gachaResponse)
}

func (g *GachaHandler) UpdateGacha(ctx echo.Context) error {
	gachaID := ctx.Param("id")
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	gacha, err := g.usecase.UpdateGacha(gachaID, name, description)
	if err != nil {
		return err
	}

	gachaResponse := convertGacha(gacha)
	return ctx.JSON(http.StatusOK, gachaResponse)
}

func (g *GachaHandler) DeleteGacha(ctx echo.Context) error {
	gachaID := ctx.Param("id")
	gacha, err := g.usecase.DeleteGacha(gachaID)
	if err != nil {
		return err
	}

	gachaResponse := convertGacha(gacha)
	return ctx.JSON(http.StatusOK, gachaResponse)
}

