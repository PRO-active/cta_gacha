package item

type ItemInput struct {
	Name    string `validation:"required, max=255"`
	GachaID string `validation:"required, max=255"`
	UserID  string `validation:"required, max=255"`
	ImgPath string `validation:"required, max=255"`
	Rairty  string `validation:"required, max=255"`
}

type ItemOutput struct {
	Name    string `validation:"required, max=255"`
	Rarity  string `validation:"required, max=255"`
	ImgPath string `validation:"required, max=255"`
}

