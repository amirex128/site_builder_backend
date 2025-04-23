package user_repo

import (
	"gorm.io/gorm"
	"site_builder_backend/internal/domain/user_entity"
	"site_builder_backend/pkg/logger"
)

type UserReadRepository struct {
	db *gorm.DB
	l  *logger.ZapLogger
}
type UserWriteRepository struct {
	db *gorm.DB
	l  *logger.ZapLogger
}

func NewUserReadRepository(db *gorm.DB, l *logger.ZapLogger) *UserReadRepository {
	return &UserReadRepository{
		db: db,
		l:  l,
	}
}
func NewUserWriteRepository(db *gorm.DB, l *logger.ZapLogger) *UserWriteRepository {
	return &UserWriteRepository{
		db: db,
		l:  l,
	}
}

func (r *UserWriteRepository) Create(entity user_entity.UserEntity) error {
	return r.db.Create(&entity).Error
}

func (r *UserReadRepository) FindById(id int64) (*user_entity.UserEntity, error) {
	var entity user_entity.UserEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		r.l.Error("user_repo - UserReadRepository - FindById: %v", err)
		return nil, err
	}
	return &entity, nil
}
