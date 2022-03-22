package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userRepo "github.com/pro-active/cta_gacha/internal/repository/user"
	"github.com/pro-active/cta_gacha/internal/usecase"
)

// rename controller -> facade
type UserFacade interface {
	GetUser(id string) (*userRepo.User, error)
	CreateUser(id string, name string, password string, email string, hash string) (*userRepo.User, error)
	UpdateUser(id string, name string, email string) (*userRepo.User, error)
}

type userFacade struct {
	usecase usecase.UserUsecase
}

func NewUserFacade(uc usecase.UserUsecase) *userFacade {
	return &userFacade{
		usecase: uc,
	}
}

func (u *userFacade) GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	user, err := u.usecase.GetUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// ToDo: delete
	userResponse := User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Hash:  user.Hash,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": userResponse,
	})
}
