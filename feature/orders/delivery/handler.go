package delivery

import (
	"fmt"
	"log"
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"
	_data "middleman-capstone/feature/orders/data"
	_helper "middleman-capstone/helper"
	"net/http"
	"strconv"

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
		limitcnv := c.QueryParam("limit")
		offsetcnv := c.QueryParam("offset")
		limit, _ := strconv.Atoi(limitcnv)
		offset, _ := strconv.Atoi(offsetcnv)
		result, err := oh.orderUseCase.GetAllAdmin(limit, offset)
		data := ParsePUToArr2(result)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("failed to get all data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success get data", data))
	}
}

func (oh *OrderHandler) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		limitcnv := c.QueryParam("limit")
		offsetcnv := c.QueryParam("offset")
		limit, _ := strconv.Atoi(limitcnv)
		offset, _ := strconv.Atoi(offsetcnv)
		id, _ := _middleware.ExtractData(c)

		result, err := oh.orderUseCase.GetAllUser(limit, offset, id)
		data := ParsePUToArr2(result)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("failed to get data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success get data", data))

	}
}

func (oh *OrderHandler) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("idorder")
		idOrder, _ := strconv.Atoi(idParam)
		idUser, _ := _middleware.ExtractData(c)

		grandTotal, _ := oh.orderUseCase.GetDetail(idUser, idOrder)
		if grandTotal == -1 {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseInternalServerError("there is an internal server error"))
		}
		result, err := oh.orderUseCase.GetItems(idOrder)
		fmt.Println("hasil :", result)
		data := _data.ParseToArrDetail(result, grandTotal, idOrder)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("failed to get data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success get data", data))
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
		orderName, url, token, user := oh.orderUseCase.Payment(newOrder.GrandTotal, id)
		dataOrder := ToDomain(newOrder)
		dataOrder.UserID = id
		dataOrder.GrandTotal = newOrder.GrandTotal
		dataOrder.PaymentLink = url
		dataOrder.OrderName = orderName
		dataOrder.Status = "pending"
		fmt.Println("data order:", dataOrder)
		status := oh.orderUseCase.CreateOrder(dataOrder, id)
		fmt.Println("to Domain:", ToDomain(newOrder))
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
		data := map[string]interface{}{
			"order_id":    dataOrder.ID,
			"grand_total": newOrder.GrandTotal,
			"nama":        user.Name,
			"email":       user.Email,
			"link":        url,
			"token":       token,
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    status,
			"message": "success create product",
		})
	}
}

func (oh *OrderHandler) Payment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data PaymentWeb

		err := c.Bind(&data)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("wrong input data"))
		}

		dataWeb := FromWeb(data)
		fmt.Println("isi dataWeb :", dataWeb)
		fmt.Println("isi data :", data)
		response, err := oh.orderUseCase.AcceptPayment(dataWeb)

		if response == -1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseInternalServerError("there is an internal server error"))
		}

		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseInternalServerError("there is an internal server error"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
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

func (oh *OrderHandler) Confirm() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := _middleware.ExtractData(c)
		orderid := c.Param("idorder")

		if role != "admin" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		order, status := oh.orderUseCase.ConfirmOrder(orderid, id, role)
		data := _data.ParseToArrConfirm(order)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    status,
			"message": "success confirm order",
		})
	}
}

func (oh *OrderHandler) Done() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := _middleware.ExtractData(c)
		orderid := c.Param("idorder")

		if role != "admin" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		order, status := oh.orderUseCase.DoneOrder(orderid, id, role)
		data := _data.ParseToArrConfirm(order)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    status,
			"message": "success confirm order",
		})
	}
}
