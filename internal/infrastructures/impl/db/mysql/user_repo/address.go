package user_repo

import "gorm.io/gorm"

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

func (r *AddressRepository) CreateAddress() error {
	return nil
}
