package user_controller

import (
	"github.com/gin-gonic/gin"
	"site_builder_backend/internal/application/dto/user/address_dto"
	"site_builder_backend/internal/application/use_cases/user_use_case"
	"site_builder_backend/pkg/logger"
)

type AddressController struct {
	useCase *user_use_case.AddressUseCase
	l       *logger.ZapLogger
}

func NewAddressController(useCase *user_use_case.AddressUseCase, l *logger.ZapLogger) *AddressController {
	return &AddressController{
		useCase: useCase,
		l:       l,
	}
}

func (u *AddressController) CreateAddress(c *gin.Context) {
	var login address_dto.CreateAddressDto
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{})
	}
	u.useCase.Create()
}
