package routing

import (
	"site_builder_backend/internal/adapters/http/user_controller"
	"site_builder_backend/internal/application/use_cases/user_use_case"
)

type ControllerServices struct {
	UserController    *user_controller.UserController
	AddressController *user_controller.AddressController
}

func NewControllerServices(services *Services) *ControllerServices {

	userUseCase := user_use_case.NewUserUseCase(services.UserReadRepo, services.UserWriteRepo, services.Logger)
	userController := user_controller.NewUserController(userUseCase, services.Logger)

	addressUseCase := user_use_case.NewAddressUseCase(services.AddressReadRepo, services.AddressWriteRepo, services.Logger)
	addressController := user_controller.NewAddressController(addressUseCase, services.Logger)

	return &ControllerServices{
		UserController:    userController,
		AddressController: addressController,
	}
}
