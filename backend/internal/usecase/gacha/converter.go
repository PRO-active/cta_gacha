package gacha

import gachaRepo "github.com/pro-active/cta_gacha/internal/repository/gacha"

func convertGacha(input *gachaRepo.Gacha) *GachaOutput {
	return &GachaOutput{
		Name:        input.Name,
		Description: input.Description,
	}
}

