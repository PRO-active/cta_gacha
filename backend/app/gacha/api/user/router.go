package user

import (
	"log"

	"github.com/gin-gonic/gin"
	userRepo "github.com/pro-active/cta_gacha/internal/repository/user"
	"github.com/pro-active/cta_gacha/internal/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RegisterRouter(
	router chi.Router,
	adminAccountUseCase adminaccountusecase.UseCase,
) {
	private := router.With()
	private.Route("/admin_accounts", func(p chi.Router) {
		p.Post("/", userFacade.CreateAdminAccount)
	})
}

func UserRoute(user *gin.RouterGroup) {
	uc := userInjector()
	user.GET("", uc.GetUser)
	// Createとかを足す
}

func userInjector() *userFacade {
	dsn := "root:password@tcp(127.0.0.1:16002)/tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	userRepository := userRepo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userFacade := NewUserFacade(userUsecase)
	return userFacade
}
