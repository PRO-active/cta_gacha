package gacha

import (
	"gorm.io/gorm"
)

type GachaRepository interface {
	GetGacha(id string) (*Gacha, error)
	GetGachas() ([]*Gacha, error)
	CreateGacha(id string, name string, description string) (*Gacha, error)
	UpdateGacha(id string, name string, description string) (*Gacha, error)
	DeleteGacha(id string) (*Gacha, error)
}

type gachaRepository struct {
	db *gorm.DB
}

func NewGachaRepository(db *gorm.DB) *gachaRepository {
	return &gachaRepository{
		db: db,
	}
}

func (g *gachaRepository) GetGacha(id string) (*Gacha, error) {
	gacha := &Gacha{}
	gacha.ID = id
	if err := g.db.Where("hidden = 0").First(gacha).Error; err != nil {
		return nil, err
	}
	return gacha, nil
}

func (g *gachaRepository) GetGachas() ([]*Gacha, error) {
	var gachas []*Gacha
	if err := g.db.Where("hidden = 0").Find(gachas).Error; err != nil {
		return nil, err
	}
	return gachas, nil
}

func (g *gachaRepository) CreateGacha(id string, name string, description string) (*Gacha, error) {
	gacha := &Gacha{ID: id, Name: name, Description: description}
	if err := g.db.Create(gacha).Error; err != nil {
		return nil, err
	}
	return gacha, nil
}

func (g *gachaRepository) UpdateGacha(id string, name string, description string) (*Gacha, error) {
	gacha := &Gacha{ID: id}
	if err := g.db.First(gacha).Error; err != nil {
		return nil, err
	}
	gacha.Name = name
	gacha.Description = description
	if err := g.db.Save(gacha).Error; err != nil {
		return nil, err
	}
	return gacha, nil
}

func (g *gachaRepository) DeleteGacha(id string) (*Gacha, error) {
	gacha := &Gacha{}
	gacha.ID = id
	if err := g.db.First(gacha).Error; err != nil {
		return nil, err
	}
	gacha.Hidden = true
	if err := g.db.Save(gacha).Error; err != nil {
		return nil, err
	}
	return gacha, nil
}
