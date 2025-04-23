package user_repo_inter

type UserReadRepository interface {
	FindById() error
}
type UserWriteRepository interface {
	Create() error
}
