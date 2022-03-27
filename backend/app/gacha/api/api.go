package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pro-active/cta_gacha/app/gacha/api/user"

	userrepository "github.com/pro-active/cta_gacha/internal/repository/user"
	userusecase "github.com/pro-active/cta_gacha/internal/usecase/user"
)

type (
	api struct {
		db database.Database

		// repository
		userRepository userrepository.UserRepository

		// useCase
		userUseCase userusecase.UserUsecase
	}
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	userGroup := r.Group("/users")
	user.UserRoute(userGroup)

	// gachaとかを足す

	return r
}
