package user_repo_inter

type UserRepository interface {
	CreateUser() error
}
