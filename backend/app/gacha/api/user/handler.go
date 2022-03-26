package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	usecase "github.com/pro-active/cta_gacha/internal/usecase/user"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: uc,
	}
}

func (u *UserHandler) GetUser(ctx echo.Context) error {
	userID := ctx.Param("id")

	user, err := u.usecase.GetUser(userID)
	if err != nil {
		return err
	}

	userResponse := User{
		Name:  user.Name,
		Email: user.Email,
	}

	return ctx.JSON(http.StatusOK, userResponse)
}

func (u *UserHandler) CreateUser(ctx echo.Context) error {
	name := ctx.FormValue("name")
	password := ctx.FormValue("password")
	email := ctx.FormValue("email")

	user, err := u.usecase.CreateUser(name, password, email)
	if err != nil {
		return err
	}

	userResponse := User{
		Name:  user.Name,
		Email: user.Email,
	}

	return ctx.JSON(http.StatusOK, userResponse)
}

func (u *UserHandler) GetUserItems(ctx echo.Context) error {
	return nil
}
func (u *UserHandler) UpdateUser(ctx echo.Context) error {
	userID := ctx.Param("id")
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	user, err := u.usecase.UpdateUser(userID, name, email)
	if err != nil {
		return err
	}

	userResponse := User{
		Name:  user.Name,
		Email: user.Email,
	}
	return ctx.JSON(http.StatusOK, userResponse)
}

func (u *UserHandler) Login(ctx echo.Context) error {
	name := ctx.FormValue("name")
	password := ctx.FormValue("password")

	user, err := u.usecase.Login(name, password)
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "userID"
	cookie.Value = user.ID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	ctx.SetCookie(cookie)
	return ctx.String(http.StatusOK, "Login success")
}

