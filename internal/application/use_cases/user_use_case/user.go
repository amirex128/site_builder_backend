package user_use_case

import (
	"site_builder_backend/internal/interfaces/db/repositories/user_repo_inter"
)

type UserUseCase struct {
	userRepo user_repo_inter.UserRepository
}

func NewUserUseCase(userRepo user_repo_inter.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (u *UserUseCase) CreateUser() {
	err := u.userRepo.CreateUser()
	if err != nil {
		return
	}
}
