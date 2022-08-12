package delivery

import (
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteInOutBound(e *echo.Echo, iobh domain.InOutBoundHandler) {
	inoutbound := e.Group("/inoutbounds")
	inoutbound.POST("", iobh.Add(), _middleware.JWTMiddleware())
}
