package factory

import (
	ud "middleman-capstone/feature/users/data"
	userDelivery "middleman-capstone/feature/users/delivery"
	us "middleman-capstone/feature/users/usecase"

	pd "middleman-capstone/feature/admins/data"
	productDelivery "middleman-capstone/feature/admins/delivery"
	ps "middleman-capstone/feature/admins/usecase"

	pud "middleman-capstone/feature/productusers/data"
	productUserDelivery "middleman-capstone/feature/productusers/delivery"
	pus "middleman-capstone/feature/productusers/usecase"

	id "middleman-capstone/feature/inventories/data"
	inventoryDelivery "middleman-capstone/feature/inventories/delivery"
	is "middleman-capstone/feature/inventories/usecase"

	cd "middleman-capstone/feature/carts/data"
	cartDelivery "middleman-capstone/feature/carts/delivery"
	cs "middleman-capstone/feature/carts/usecase"

	iobd "middleman-capstone/feature/inoutbounds/data"
	inoutboundDelivery "middleman-capstone/feature/inoutbounds/delivery"
	iobs "middleman-capstone/feature/inoutbounds/usecase"

	od "middleman-capstone/feature/orders/data"
	orderDelivery "middleman-capstone/feature/orders/delivery"
	os "middleman-capstone/feature/orders/usecase"

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

	productUserData := pud.New(db)
	productUserCase := pus.New(productUserData, validator)
	productUserHandler := productUserDelivery.New(productUserCase)
	productUserDelivery.RouteProductUser(e, productUserHandler)

	inventoryData := id.New(db)
	inventoryCase := is.New(inventoryData, validator)
	inventoryHandler := inventoryDelivery.New(inventoryCase)
	inventoryDelivery.RouteInventory(e, inventoryHandler)

	cartData := cd.New(db)
	cartCase := cs.New(cartData)
	cartHandler := cartDelivery.New(cartCase)
	cartDelivery.RouteCart(e, cartHandler)

	inoutboundData := iobd.New(db)
	inoutboundCase := iobs.New(inoutboundData, validator)
	inoutboundHandler := inoutboundDelivery.New(inoutboundCase)
	inoutboundDelivery.RouteInOutBound(e, inoutboundHandler)

	orderData := od.New(db)
	orderCase := os.New(orderData, validator)
	orderHandler := orderDelivery.New(orderCase)
	orderDelivery.RouteOrder(e, orderHandler)

}
