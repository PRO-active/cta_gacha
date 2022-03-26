package item

import itemRepo "github.com/pro-active/cta_gacha/internal/repository/item"

func convertItem(input *itemRepo.Item) *ItemOutput {
	return &ItemOutput{
		Name:    input.Name,
		Rarity:  input.Rarity,
		ImgPath: input.ImgPath,
	}
}

