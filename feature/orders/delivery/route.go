package delivery

import (
	"middleman-capstone/domain"

	_middleware "middleman-capstone/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteOrder(e *echo.Echo, do domain.OrderHandler) {

	order := e.Group("/orders")
	// order.GET("/admins", do.GetAllAdmin(), _middleware.JWTMiddleware())
	order.POST("/users", do.Create(), _middleware.JWTMiddleware())
	// order.POST("/user", do.CreateItems(), _middleware.JWTMiddleware())
}
