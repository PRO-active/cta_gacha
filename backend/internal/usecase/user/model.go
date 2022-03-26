package user

type UserInput struct {
	Name     string `validation:"required, max=255"`
	Password string `validation:"required, max=255"`
	Email    string `validation:"required, max=255"`
}

type UserOutput struct {
	Name  string `validation:"required, max=255"`
	Email string `validation:"required, max=255"`
}

type UserCookie struct {
	ID   string `validation:"required, max=255"`
	Name string `validation:"required, max=255"`
}

