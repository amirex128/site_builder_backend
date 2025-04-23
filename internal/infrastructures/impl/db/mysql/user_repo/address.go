package user_repo

import (
	"gorm.io/gorm"
	"site_builder_backend/internal/domain/user_entity"
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

func (r *AddressWriteRepository) Create(entity user_entity.AddressEntity) error {
	return r.db.Create(&entity).Error
}

func (r *AddressReadRepository) FindById(id int64) (*user_entity.AddressEntity, error) {
	var entity user_entity.AddressEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		r.l.Error("user_repo - AddressReadRepository - FindById: %v", err)
		return nil, err
	}
	return &entity, nil
}
