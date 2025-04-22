package user_use_case

import (
	"site_builder_backend/internal/interfaces/db/repositories/user_repo_inter"
)

type AddressUseCase struct {
	addressRepo user_repo_inter.AddressRepository
}

func NewAddressUseCase(addressRepo user_repo_inter.AddressRepository) *AddressUseCase {
	return &AddressUseCase{addressRepo: addressRepo}
}
func (u *AddressUseCase) CreateAddress() {
}
