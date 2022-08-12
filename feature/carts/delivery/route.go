package delivery

import (
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteCart(e *echo.Echo, dp domain.CartHandler) {
	e.POST("/carts", dp.PostCart(), _middleware.JWTMiddleware())
	e.GET("/carts", dp.GetAll(), _middleware.JWTMiddleware())
	e.PUT("/carts/:id", dp.UpdateCart(), _middleware.JWTMiddleware())
	e.DELETE("/carts/:id", dp.DeleteCart(), _middleware.JWTMiddleware())
}
