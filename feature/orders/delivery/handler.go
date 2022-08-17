package delivery

import (
	"log"
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"
	_helper "middleman-capstone/helper"
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
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("sucess"))
	}
}

func (oh *OrderHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newOrder FormatOrder
		id, _ := _middleware.ExtractData(c)
		bind := c.Bind(&newOrder)

		if bind != nil {
			log.Println("cant bind data")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})

		}

		status := oh.orderUseCase.CreateOrder(ToDomain(newOrder), id)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "wrong input",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "success create product",
		})
	}
}

// func (oh *OrderHandler) CreateItems() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var newOrder FormatOrder
// 		bind := c.Bind(&newOrder)

// 		if bind != nil {
// 			return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 				"message": "error to bind",
// 				"code":    400,
// 			})
// 		}
// 		data := ParseToArrItems(newOrder.Items)
// 		row, err := oh.orderUseCase.CreateItems(data)
// 		fmt.Println("row", row)
// 		fmt.Println("err", err)
// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"msg": "berhasil",
// 		})
// 	}
// }
