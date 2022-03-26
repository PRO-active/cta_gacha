package item

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/pro-active/cta_gacha/internal/usecase/item"
)

type ItemHandler struct {
	usecase usecase.ItemUsecase
}

func NewItemHandler(uc usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{
		usecase: uc,
	}
}

func (g *ItemHandler) CreateItem(ctx echo.Context) error {
	name := ctx.FormValue("name")
	gachaID := ctx.FormValue("gachaid")
	userID := ctx.FormValue("userid")
	// TODO: 画像が送られてくる？
	imgpath := ctx.FormValue("imgpath")
	rarity := ctx.FormValue("rarity")
	item, err := g.usecase.CreateItem(name, gachaID, userID, imgpath, rarity)
	if err != nil {
		return err
	}
	itemResponse := Item{
		Name:    item.Name,
		ImgPath: item.ImgPath,
		Rarity:  item.Rarity,
	}

	return ctx.JSON(http.StatusOK, itemResponse)
}
