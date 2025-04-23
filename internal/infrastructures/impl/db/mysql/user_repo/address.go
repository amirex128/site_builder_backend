package user_repo

import (
	"gorm.io/gorm"
	"site_builder_backend/pkg/logger"
)

type AddressReadRepository struct {
	db *gorm.DB
	l  *logger.ZapLogger
}

type AddressWriteRepository struct {
	db *gorm.DB
	l  *logger.ZapLogger
}

func NewAddressReadRepository(db *gorm.DB, l *logger.ZapLogger) *AddressReadRepository {
	return &AddressReadRepository{
		db: db,
		l:  l,
	}
}
func NewAddressWriteRepository(db *gorm.DB, l *logger.ZapLogger) *AddressWriteRepository {
	return &AddressWriteRepository{
		db: db,
		l:  l,
	}
}

func (r *AddressWriteRepository) Create() error {
	//TODO implement me
	panic("implement me")
}

func (r *AddressReadRepository) FindById() error {
	//TODO implement me
	panic("implement me")
}
