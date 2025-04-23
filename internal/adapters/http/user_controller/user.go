package user_controller

import (
	"github.com/gin-gonic/gin"
	"site_builder_backend/internal/application/dto/user/user_dto"
	"site_builder_backend/internal/application/use_cases/user_use_case"
	"site_builder_backend/pkg/logger"
)

type UserController struct {
	useCase *user_use_case.UserUseCase
	l       *logger.ZapLogger
}

func NewUserController(useCase *user_use_case.UserUseCase, l *logger.ZapLogger) *UserController {
	return &UserController{
		useCase: useCase,
		l:       l,
	}
}

func (u *UserController) LoginUser(c *gin.Context) {
	var login user_dto.LoginDto
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{})
	}
}
