package delivery

import (
	"middleman-capstone/domain"

	_middleware "middleman-capstone/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteOrder(e *echo.Echo, do domain.OrderHandler) {

	order := e.Group("/orders")
	order.GET("/admins", do.GetAllAdmin(), _middleware.JWTMiddleware())
	order.POST("/users", do.Create(), _middleware.JWTMiddleware())
	order.GET("/users", do.GetAllUser(), _middleware.JWTMiddleware())
	order.GET("/users/:idorder", do.GetDetail(), _middleware.JWTMiddleware())
	order.POST("/payment", do.Payment())
	order.PUT("/confirm/:idorder", do.Confirm(), _middleware.JWTMiddleware())
	order.PUT("/done/:idorder", do.Done(), _middleware.JWTMiddleware())
}
