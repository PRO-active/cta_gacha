package gacha

type Gacha struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	UserID      string
	Description string
	Hidden      bool
}
