package user_router

import (
	"site_builder_backend/internal/adapters/http/user_controller"
	"site_builder_backend/internal/application/use_cases/user_use_case"
	"site_builder_backend/internal/infrastructures/impl/db/mysql/user_repo"
	"site_builder_backend/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func AddressRegister(g *gin.RouterGroup, db *postgres.Postgres) {

	repo := user_repo.NewAddressRepository(db.DB)
	useCase := user_use_case.NewAddressUseCase(repo)
	controller := user_controller.NewAddressController(useCase)

	g.POST("Create", controller.CreateAddress)
}
