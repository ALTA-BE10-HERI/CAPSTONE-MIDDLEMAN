package delivery

import (
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, du domain.UserHandler) {
	e.POST("/login", du.LoginAuth())
	e.POST("/register", du.InsertUser())
	e.GET("/users", du.GetProfile(), _middleware.JWTMiddleware())
	e.DELETE("/users", du.DeleteById(), _middleware.JWTMiddleware())
	e.PUT("/users", du.UpdateUser(), _middleware.JWTMiddleware())

	productuser := e.Group("/users/products")
	productuser.POST("", du.Create(), _middleware.JWTMiddleware())
	productuser.GET("", du.ReadAll(), _middleware.JWTMiddleware())
	productuser.PUT("/:idproduct", du.Update(), _middleware.JWTMiddleware())
	productuser.DELETE("/:idproduct", du.Delete(), _middleware.JWTMiddleware())
}
