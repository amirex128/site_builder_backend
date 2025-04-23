package http_router

import ()

func (r *Router) UserRegister() {

	r.user.POST("Login", r.ControllerServices.AddressController.CreateAddress)
}

func (r *Router) AddressRegister() {

	r.address.POST("Create", r.ControllerServices.AddressController.CreateAddress, r.Services.AuthMiddleware.CheckPolicy())
}
