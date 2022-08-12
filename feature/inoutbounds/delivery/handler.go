package delivery

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/common"
	"net/http"

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
		id, _ := common.ExtractData(c)
		bind := c.Bind(&newProduct)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}

		status := iobh.inoutboundUseCase.AddEntry(newProduct.ToModel(), id)

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
