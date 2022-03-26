package userhaveitem

type UserHaveItemInput struct {
	UserID string `validation:"required, max=255"`
	ItemID string `validation:"required, max=255"`
}

