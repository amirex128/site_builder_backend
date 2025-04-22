package user_use_case

import (
	"context"
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

// SendSms sends an SMS to the specified phone number with the given message
func (u *UserUseCase) SendSms(ctx context.Context, phone, message string) error {
	// Here you would implement the actual SMS sending logic
	// This could involve calling an external SMS service or updating a database
	// For now, this is just a placeholder
	return nil
}
