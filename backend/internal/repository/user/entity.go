package user

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string
	Hash     string
}