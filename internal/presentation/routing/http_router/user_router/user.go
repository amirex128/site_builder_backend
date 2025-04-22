package user_router

import (
	"site_builder_backend/internal/adapters/controllers/http/user_controller"
	"site_builder_backend/internal/application/use_cases/user_use_case"
	"site_builder_backend/internal/infrastructures/impl/db/mysql/user_repo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRegister(g *gin.RouterGroup, db *gorm.DB) {

	repo := user_repo.NewUserRepository(db)
	useCase := user_use_case.NewUserUseCase(repo)
	controller := user_controller.NewUserController(useCase)

	g.POST("Login", controller.LoginUser)
}
