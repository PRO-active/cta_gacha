package item

import (
	gachaRepo "github.com/pro-active/cta_gacha/internal/repository/gacha"
	repository "github.com/pro-active/cta_gacha/internal/repository/item"
	userRepo "github.com/pro-active/cta_gacha/internal/repository/user"
	uhiRepo "github.com/pro-active/cta_gacha/internal/repository/userhaveitem"
	usecase "github.com/pro-active/cta_gacha/internal/usecase/item"
	"gorm.io/gorm"
)

func InitializeItemHandler(db *gorm.DB) *ItemHandler {
	gachaRepository := gachaRepo.NewGachaRepository(db)
	itemRepository := repository.NewItemRepository(db)
	userRepository := userRepo.NewUserRepository(db)
	uhiRepository := uhiRepo.NewUserHaveItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepository, userRepository, gachaRepository, uhiRepository)
	itemHandler := NewItemHandler(itemUsecase)
	return itemHandler
}

