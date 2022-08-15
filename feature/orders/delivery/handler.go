package delivery

import (
	"middleman-capstone/domain"
	"middleman-capstone/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUseCase domain.OrderUseCase
}

func New(os domain.OrderUseCase) domain.OrderHandler {
	return &OrderHandler{
		orderUseCase: os,
	}
}

func (oh *OrderHandler) GetAllAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 		limitcnv := c.QueryParam("limit")
		// 		offsetcnv := c.QueryParam("offset")
		// 		limit, _ := strconv.Atoi(limitcnv)
		// 		offset, _ := strconv.Atoi(offsetcnv)
		// 		result, err := oh.orderUseCase.GetAllAdmin(limit, offset)
		// 		if err != nil {
		// 			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("failed to get all data"))
		// 		}
		// 		var data = map[string]interface{}{
		// 			"data": FromModelList(result),
		// 		}
		return c.JSON(http.StatusOK, helper.ResponseOkNoData("sucess"))
	}
}
