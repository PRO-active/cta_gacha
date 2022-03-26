package gacha

import (
	repository "github.com/pro-active/cta_gacha/internal/repository/gacha"
	itemRepo "github.com/pro-active/cta_gacha/internal/repository/item"
	userRepo "github.com/pro-active/cta_gacha/internal/repository/user"
	uhiRepo "github.com/pro-active/cta_gacha/internal/repository/userhaveitem"
	usecase "github.com/pro-active/cta_gacha/internal/usecase/gacha"
	itemUsecase "github.com/pro-active/cta_gacha/internal/usecase/item"
	"gorm.io/gorm"
)

func InitializeGachaHandler(db *gorm.DB) *GachaHandler {
	gachaRepository := repository.NewGachaRepository(db)
	itemRepository := itemRepo.NewItemRepository(db)
	userRepository := userRepo.NewUserRepository(db)
	uhiRepository := uhiRepo.NewUserHaveItemRepository(db)
	gachaUsecase := usecase.NewGachaUsecase(gachaRepository)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepository, userRepository, gachaRepository, uhiRepository)
	gachaHandler := NewGachaHandler(gachaUsecase, itemUsecase)
	return gachaHandler
}

