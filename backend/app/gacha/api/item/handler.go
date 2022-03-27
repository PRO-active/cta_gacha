package item

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pro-active/cta_gacha/external/aws/s3"
	usecase "github.com/pro-active/cta_gacha/internal/usecase/item"
	"github.com/pro-active/cta_gacha/util"
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
	rarity := ctx.FormValue("rarity")

	file, err := ctx.FormFile("image")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// S3に送る
	uuid, err := util.GenerateUUID()
	filename := fmt.Sprintf("%s_%s", uuid, file.Filename)
	req := s3.NewUploadRequest(os.Getenv("BUCKET_NAME"), filename, src)
	imgPath, err := s3.Upload(context.Background(), req)
	if err != nil {
		return err
	}

	item, err := g.usecase.CreateItem(name, gachaID, userID, imgPath, rarity)
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
