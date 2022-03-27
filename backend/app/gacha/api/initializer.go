package api

import (
	userrepository "github.com/pro-active/cta_gacha/internal/repository/user"
	userusecase "github.com/pro-active/cta_gacha/internal/usecase/user"
)

func NewApi(
	db database.Database,
) *api {
	api := &api{
		db: db,
	}

	initRepository(api)
	initUseCase(api)

	return api
}

func initRepository(a *api) {
	a.userRepository = userrepository.NewUserRepository(a.db)
}

func initUseCase(a *api) {
	a.userUseCase = userusecase.NewUserUsecase(
		a.db,
	)
}
