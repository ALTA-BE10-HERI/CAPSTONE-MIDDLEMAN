package delivery

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/common"
	"middleman-capstone/feature/inoutbounds/data"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type inoutboundHandler struct {
	inoutboundUseCase domain.InOutBoundUseCase
}

func New(iobuc domain.InOutBoundUseCase) domain.InOutBoundHandler {
	return &inoutboundHandler{
		inoutboundUseCase: iobuc,
	}
}

func (iobh *inoutboundHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newProduct CartFormat
		id, role := common.ExtractData(c)
		bind := c.Bind(&newProduct)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}

		cart, status := iobh.inoutboundUseCase.AddEntry(newProduct.ToModel(), id, role)

		data := data.ParseIOBToArr3(cart)

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

		if status == 403 {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"code":    status,
				"message": "forbidden",
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    status,
			"message": "success create product",
			"data":    data,
		})
	}
}

func (iobh *inoutboundHandler) ReadAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := common.ExtractData(c)

		cart, status := iobh.inoutboundUseCase.ReadEntry(id, role)
		data := data.ParseIOBToArr2(cart)

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
			"message": "get data success",
		})
	}
}

func (iobh *inoutboundHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp CartFormat2
		bind := c.Bind(&tmp)

		id, role := common.ExtractData(c)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}

		productid, err := strconv.Atoi(c.Param("idproduct"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		cart, status := iobh.inoutboundUseCase.UpdateEntry(tmp.ToModel2(), productid, id, role)
		data := data.ParseIOBToArr3(cart)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "data not found",
			})
		}

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "insufficient stock",
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
			"message": "get data success",
		})
	}
}

func (iobh *inoutboundHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("idproduct"))

		if err != nil {
			log.Println("cant convert to int", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "cant convert to int",
			})
		}

		id, role := common.ExtractData(c)

		status := iobh.inoutboundUseCase.DeleteEntry(cnv, id, role)

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

		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"code":    status,
			"message": "success delete product",
		})
	}
}
