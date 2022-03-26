package item

type Item struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	GachaID string
	UserID  string
	Rarity  string
	ImgPath string
}

