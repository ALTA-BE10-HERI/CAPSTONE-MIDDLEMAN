package delivery

import (
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteInventory(e *echo.Echo, ih domain.InventoryHandler) {
	inventoryuser := e.Group("/users/inventory")
	inventoryuser.POST("", ih.Create(), _middleware.JWTMiddleware())
	inventoryuser.GET("", ih.ReadUserHistory(), _middleware.JWTMiddleware())
	inventoryuser.GET("/:idoutbound", ih.ReadUserDetail(), _middleware.JWTMiddleware())
}
