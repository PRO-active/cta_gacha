package gacha

type GachaInput struct {
	Name        string `validation:"required, max=255"`
	Description string `validation:"required, max=255"`
}

type GachaOutput struct {
	Name        string `validation:"required, max=255"`
	Description string `validation:"required, max=255"`
}

