package item

import (
	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItem(id string) (*Item, error)
	GetItems(gachaID string) ([]*Item, error)
	CreateItem(id, name, gachaID, userID, imgPath, rairty string) (*Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{
		db: db,
	}
}

func (i *itemRepository) GetItem(id string) (*Item, error) {
	item := &Item{}
	item.ID = id
	if err := i.db.First(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemRepository) GetItems(gachaID string) ([]*Item, error) {
	var items []*Item
	if err := i.db.Where("gacha_id = ?", gachaID).Find(items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemRepository) CreateItem(id, name, gachaID, userID, imgPath, rairty string) (*Item, error) {
	item := &Item{ID: id, Name: name, GachaID: gachaID, UserID: userID, ImgPath: imgPath, Rarity: rairty}
	if err := i.db.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

