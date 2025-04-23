package user_use_case

import (
	"site_builder_backend/internal/interfaces/db/repositories/user_repo_inter"
	"site_builder_backend/pkg/logger"
)

type AddressUseCase struct {
	addressReadRepo  user_repo_inter.AddressReadRepository
	addressWriteRepo user_repo_inter.AddressWriteRepository
	l                *logger.ZapLogger
}

func NewAddressUseCase(addressReadRepo user_repo_inter.AddressReadRepository, addressWriteRepo user_repo_inter.AddressWriteRepository, l *logger.ZapLogger) *AddressUseCase {
	return &AddressUseCase{
		addressReadRepo:  addressReadRepo,
		addressWriteRepo: addressWriteRepo,
		l:                l,
	}
}
func (u *AddressUseCase) Create() {
}
