package user

import (
	"errors"

	userRepo "github.com/pro-active/cta_gacha/internal/repository/user"
	"github.com/pro-active/cta_gacha/util"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetUser(id string) (*UserOutput, error)
	CreateUser(name string, password string, email string) (*UserOutput, error)
	UpdateUser(id string, name string, email string) (*UserOutput, error)
	Login(name, password string) (*UserCookie, error)
}

type userUsecase struct {
	repository userRepo.UserRepository
}

func NewUserUsecase(u userRepo.UserRepository) *userUsecase {
	return &userUsecase{
		repository: u,
	}
}

func (u *userUsecase) GetUser(id string) (*UserOutput, error) {
	result, err := u.repository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return convertUser(result), err
}

func (u *userUsecase) CreateUser(name, password, email string) (*UserOutput, error) {
	id, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	hash := string(hashed)
	result, err := u.repository.CreateUser(id, name, password, email, hash)
	if err != nil {
		return nil, err
	}
	return convertUser(result), err
}

func (u *userUsecase) UpdateUser(id, name, email string) (*UserOutput, error) {
	result, err := u.repository.UpdateUser(id, name, email)
	if err != nil {
		return nil, err
	}
	return convertUser(result), err
}

func (u *userUsecase) Login(name, password string) (*UserCookie, error) {
	result, err := u.repository.GetUserByNameAndPassword(name, password)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("ユーザ名とパスワードが一致しません")
	}
	return convertUserCookie(result), err
}

