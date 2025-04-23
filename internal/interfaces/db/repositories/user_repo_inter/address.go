package user_repo_inter

import "site_builder_backend/internal/domain/user_entity"

type AddressReadRepository interface {
	FindById(id int64) (*user_entity.AddressEntity, error)
}
type AddressWriteRepository interface {
	Create(user_entity.AddressEntity) error
}
