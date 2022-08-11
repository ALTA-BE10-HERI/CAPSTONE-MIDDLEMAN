package delivery

import (
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteAdmin(e *echo.Echo, dp domain.ProductHandler) {

	admin := e.Group("/admins")
	admin.GET("", dp.GetAll())
	admin.POST("", dp.Create(), _middleware.JWTMiddleware())
	admin.PUT("/:idproduct", dp.Update(), _middleware.JWTMiddleware())
	admin.DELETE("/:idproduct", dp.Delete(), _middleware.JWTMiddleware())
}
