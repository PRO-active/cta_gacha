package userhaveitem

type UserHaveItem struct {
	ID     string `gorm:"primaryKey"`
	ItemID string
	UserID string
}

