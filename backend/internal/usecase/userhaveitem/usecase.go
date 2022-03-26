package userhaveitem

import (
	userhaveitemRepo "github.com/pro-active/cta_gacha/internal/repository/userhaveitem"
	"github.com/pro-active/cta_gacha/util"
)

type UserHaveItemUsecase interface {
	GetUserHaveItems(userID string) ([]*userhaveitemRepo.UserHaveItem, error)
	CreateUserHaveItem(id string, userID string, itemID string) (*userhaveitemRepo.UserHaveItem, error)
}

type userhaveitemUsecase struct {
	repository userhaveitemRepo.UserHaveItemRepository
}

func NewUserHaveItemUsecase(i userhaveitemRepo.UserHaveItemRepository) *userhaveitemUsecase {
	return &userhaveitemUsecase{
		repository: i,
	}
}

func (i *userhaveitemUsecase) GetUserHaveItems(userID string) ([]*userhaveitemRepo.UserHaveItem, error) {
	return i.repository.GetUserHaveItemsByUserID(userID)
}

func (i *userhaveitemUsecase) CreateUserHaveItem(userID, itemID string) (*userhaveitemRepo.UserHaveItem, error) {
	id, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	return i.repository.CreateUserHaveItem(id, userID, itemID)
}

