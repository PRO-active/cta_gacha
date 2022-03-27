package item

import usecase "github.com/pro-active/cta_gacha/internal/usecase/item"

func ConvertItem(input *usecase.ItemOutput) *Item {
	return &Item{
		Name:    input.Name,
		Rarity:  input.Rarity,
		ImgPath: input.ImgPath,
	}
}

