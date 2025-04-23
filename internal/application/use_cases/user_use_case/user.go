package user_use_case

import (
	"context"
	"site_builder_backend/internal/interfaces/db/repositories/user_repo_inter"
	"site_builder_backend/pkg/logger"
)

type UserUseCase struct {
	userReadRepo  user_repo_inter.UserReadRepository
	userWriteRepo user_repo_inter.UserWriteRepository
	l             *logger.ZapLogger
}

func NewUserCommandUseCase(userReadRepo user_repo_inter.UserReadRepository, userWriteRepo user_repo_inter.UserWriteRepository, l *logger.ZapLogger) *UserUseCase {
	return &UserUseCase{
		userReadRepo:  userReadRepo,
		userWriteRepo: userWriteRepo,
		l:             l,
	}
}

func (u *UserUseCase) CreateUser() {
	err := u.userWriteRepo.Create()
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
