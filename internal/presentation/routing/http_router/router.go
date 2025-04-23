package http_router

import (
	"github.com/gin-gonic/gin"
	"site_builder_backend/internal/presentation/routing"
)

type Router struct {
	routerEngine *gin.Engine
	*routing.Services
	ControllerServices *routing.ControllerServices
	user               *gin.RouterGroup
	address            *gin.RouterGroup
}

func NewRouter(g *gin.Engine, services *routing.Services, controllerServices *routing.ControllerServices) *Router {
	return &Router{
		Services:           services,
		ControllerServices: controllerServices,
		user:               g.Group("User", services.AuthMiddleware.Authenticate()),
		address:            g.Group("Address", services.AuthMiddleware.Authenticate()),
	}
}

func Register(g *gin.Engine, controllerServices *routing.ControllerServices, services *routing.Services) {
	router := NewRouter(g, services, controllerServices)

	router.UserRegister()
	router.AddressRegister()

}
