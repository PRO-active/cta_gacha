package userhaveitem

import (
	"gorm.io/gorm"
)

type UserHaveItemRepository interface {
	GetUserHaveItemsByUserID(userID string) ([]*UserHaveItem, error)
	CreateUserHaveItem(id string, userID string, itemID string) (*UserHaveItem, error)
}

type userhaveitemRepository struct {
	db *gorm.DB
}

func NewUserHaveItemRepository(db *gorm.DB) *userhaveitemRepository {
	return &userhaveitemRepository{
		db: db,
	}
}

func (u *userhaveitemRepository) GetUserHaveItemsByUserID(userID string) ([]*UserHaveItem, error) {
	var userhaveitems []*UserHaveItem
	if err := u.db.Find(userhaveitems, "userid=?", userID).Error; err != nil {
		return nil, err
	}
	return userhaveitems, nil
}

func (u *userhaveitemRepository) CreateUserHaveItem(id, userID, itemID string) (*UserHaveItem, error) {
	userhaveitem := &UserHaveItem{ID: id, UserID: userID, ItemID: itemID}
	if err := u.db.Create(userhaveitem).Error; err != nil {
		return nil, err
	}
	return userhaveitem, nil
}
