package user_repo_inter

type AddressReadRepository interface {
	FindById() error
}
type AddressWriteRepository interface {
	Create() error
}
