package gacha

import (
	gachaRepo "github.com/pro-active/cta_gacha/internal/repository/gacha"
	"github.com/pro-active/cta_gacha/util"
)

type GachaUsecase interface {
	GetGacha(id string) (*GachaOutput, error)
	GetGachas() ([]*GachaOutput, error)
	CreateGacha(name string, description string) (*GachaOutput, error)
	UpdateGacha(id string, name string, description string) (*GachaOutput, error)
	DeleteGacha(id string) (*GachaOutput, error)
}

type gachaUsecase struct {
	repository gachaRepo.GachaRepository
}

func NewGachaUsecase(g gachaRepo.GachaRepository) *gachaUsecase {
	return &gachaUsecase{
		repository: g,
	}
}

func (g *gachaUsecase) GetGacha(id string) (*GachaOutput, error) {
	result, err := g.repository.GetGacha(id)
	if err != nil {
		return nil, err
	}
	return convertGacha(result), err
}

func (g *gachaUsecase) GetGachas() ([]*GachaOutput, error) {
	results, err := g.repository.GetGachas()
	if err != nil {
		return nil, err
	}
	gachas := make([]*GachaOutput, 0, len(results))
	for i := range results {
		gachas = append(gachas, convertGacha(results[i]))
	}
	return gachas, err
}
func (g *gachaUsecase) CreateGacha(name, description string) (*GachaOutput, error) {
	id, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	result, err := g.repository.CreateGacha(id, name, description)
	if err != nil {
		return nil, err
	}
	return convertGacha(result), err
}

func (g *gachaUsecase) UpdateGacha(id, name, description string) (*GachaOutput, error) {
	result, err := g.repository.UpdateGacha(id, name, description)
	if err != nil {
		return nil, err
	}
	return convertGacha(result), err
}

func (g *gachaUsecase) DeleteGacha(id string) (*GachaOutput, error) {
	result, err := g.repository.DeleteGacha(id)
	if err != nil {
		return nil, err
	}
	return convertGacha(result), err
}

