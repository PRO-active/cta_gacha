package gacha

import usecase "github.com/pro-active/cta_gacha/internal/usecase/gacha"

func convertGacha(input *usecase.GachaOutput) *Gacha {
	return &Gacha{
		Name:        input.Name,
		Description: input.Description,
	}
}

