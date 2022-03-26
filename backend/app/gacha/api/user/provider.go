package user

import (
	repository "github.com/pro-active/cta_gacha/internal/repository/user"
	usecase "github.com/pro-active/cta_gacha/internal/usecase/user"
	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB) *UserHandler {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := NewUserHandler(userUsecase)
	return userHandler
}

