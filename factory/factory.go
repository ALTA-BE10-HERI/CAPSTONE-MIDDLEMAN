package factory

import (
	ud "middleman-capstone/feature/users/data"
	userDelivery "middleman-capstone/feature/users/delivery"
	us "middleman-capstone/feature/users/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	validator := validator.New()

	userData := ud.New(db)
	useCase := us.New(userData, validator)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)

}
