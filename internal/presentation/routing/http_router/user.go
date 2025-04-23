package http_router

import (
	"github.com/gin-gonic/gin"
	"site_builder_backend/internal/adapters/http/user_controller"
	"site_builder_backend/internal/application/use_cases/user_use_case"
	"site_builder_backend/internal/infrastructures/impl/db/mysql/user_repo"
	"site_builder_backend/internal/presentation/routing"
)

type UserRouter struct {
	user    *gin.RouterGroup
	address *gin.RouterGroup
	*routing.Services
}

func NewUserRouter(services *routing.Services) *UserRouter {
	return &UserRouter{
		user:     services.Router.Group("User"),
		address:  services.Router.Group("Address"),
		Services: services,
	}
}

func (r *UserRouter) UserRegister() {

	userReadRepo := user_repo.NewUserReadRepository(r.PostgresClient.DB, r.Logger)
	userWriteRepo := user_repo.NewUserWriteRepository(r.PostgresClient.DB, r.Logger)
	useCase := user_use_case.NewUserCommandUseCase(userReadRepo, userWriteRepo, r.Logger)
	controller := user_controller.NewUserController(useCase, r.Logger)

	r.user.POST("Login", controller.LoginUser)
}

func (r *UserRouter) AddressRegister() {

	addressReadRepo := user_repo.NewAddressReadRepository(r.PostgresClient.DB, r.Logger)
	addressWriteRepo := user_repo.NewAddressWriteRepository(r.PostgresClient.DB, r.Logger)
	useCase := user_use_case.NewAddressUseCase(addressReadRepo, addressWriteRepo, r.Logger)
	controller := user_controller.NewAddressController(useCase, r.Logger)

	r.address.POST("Create", controller.CreateAddress)
}
