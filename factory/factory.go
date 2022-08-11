package factory

import (
	ud "middleman-capstone/feature/users/data"
	userDelivery "middleman-capstone/feature/users/delivery"
	us "middleman-capstone/feature/users/usecase"

	pd "middleman-capstone/feature/admins/data"
	productDelivery "middleman-capstone/feature/admins/delivery"
	ps "middleman-capstone/feature/admins/usecase"

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

	productData := pd.New(db)
	productCase := ps.New(productData, validator)
	productHandler := productDelivery.New(productCase)
	productDelivery.RouteAdmin(e, productHandler)

}
