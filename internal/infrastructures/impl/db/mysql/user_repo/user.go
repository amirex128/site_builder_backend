package user_repo

import (
	"gorm.io/gorm"
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

func (r *UserWriteRepository) Create() error {
	return nil
}

func (r *UserReadRepository) FindById() error {
	return nil
}
