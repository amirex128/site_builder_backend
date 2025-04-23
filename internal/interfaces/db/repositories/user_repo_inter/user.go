package user_repo_inter

import "site_builder_backend/internal/domain/user_entity"

type UserReadRepository interface {
	FindById(id int64) (*user_entity.UserEntity, error)
}
type UserWriteRepository interface {
	Create(user_entity.UserEntity) error
}
