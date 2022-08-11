package delivery

import (
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteProductUser(e *echo.Echo, puh domain.ProductUserHandler) {
	productuser := e.Group("/users/products")
	productuser.POST("", puh.Create(), _middleware.JWTMiddleware())
	productuser.GET("", puh.ReadAll(), _middleware.JWTMiddleware())
	productuser.PUT("/:idproduct", puh.Update(), _middleware.JWTMiddleware())
	productuser.DELETE("/:idproduct", puh.Delete(), _middleware.JWTMiddleware())
}
