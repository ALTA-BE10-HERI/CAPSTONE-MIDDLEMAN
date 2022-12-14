package delivery

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/common"
	"middleman-capstone/feature/inventories/data"
	"net/http"

	"github.com/labstack/echo/v4"
)

type inventoryHandler struct {
	inventoryUseCase domain.InventoryUseCase
}

func New(iuc domain.InventoryUseCase) domain.InventoryHandler {
	return &inventoryHandler{
		inventoryUseCase: iuc,
	}
}

func (ih *inventoryHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newInventory InputFormat
		id, role := common.ExtractData(c)
		bind := c.Bind(&newInventory)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		if role != "user" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		inventory, status := ih.inventoryUseCase.CreateUserInventory(ToDomain(newInventory), id)
		data := data.ParseToArr(inventory)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "please check your outbounds, it must have a data",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    status,
			"message": "success input data",
			"data":    data,
		})
	}
}

func (ih *inventoryHandler) ReadUserDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv := c.Param("idoutbound")
		id, role := common.ExtractData(c)

		if role != "user" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		product, status, invenid := ih.inventoryUseCase.ReadUserOutBoundDetail(id, cnv)
		data := FromModel2(product, invenid)

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    status,
			"message": "get data success",
		})
	}
}

func (ih *inventoryHandler) ReadUserHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := common.ExtractData(c)

		if role != "user" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		product, status := ih.inventoryUseCase.ReadUserOutBoundHistory(id)
		data := FromModelListHistory(product)

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

func (ih *inventoryHandler) CreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newInventory InputFormat
		id, role := common.ExtractData(c)
		bind := c.Bind(&newInventory)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		if role != "admin" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		inventory, status := ih.inventoryUseCase.CreateAdminInventory(ToDomain(newInventory), id, role)
		data := data.ParseToArr(inventory)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "please check your inbounds, it must have a data",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    status,
			"message": "success input data",
			"data":    data,
		})
	}
}

func (ih *inventoryHandler) ReadAdminHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := common.ExtractData(c)

		if role != "admin" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		product, status := ih.inventoryUseCase.ReadAdminOutBoundHistory()
		data := FromModelListHistory(product)

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

func (ih *inventoryHandler) ReadAdminDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv := c.Param("idinbound")
		_, role := common.ExtractData(c)

		if role != "admin" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		product, status, invenid := ih.inventoryUseCase.ReadAdminOutBoundDetail(cnv)
		data := FromModel2(product, invenid)

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    status,
			"message": "get data success",
		})
	}
}
