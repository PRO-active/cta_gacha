package user

import userRepo "github.com/pro-active/cta_gacha/internal/repository/user"

func convertUser(input *userRepo.User) *UserOutput {
	return &UserOutput{
		Name:  input.Name,
		Email: input.Email,
	}
}

func convertUserCookie(input *userRepo.User) *UserCookie {
	return &UserCookie{
		ID:   input.ID,
		Name: input.Name,
	}
}

