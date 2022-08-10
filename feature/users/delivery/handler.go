package delivery

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/common"
	_middleware "middleman-capstone/feature/common"
	user "middleman-capstone/feature/users"
	_helper "middleman-capstone/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(us domain.UserUseCase) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
	}
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		err := c.Bind(&tmp)
		if err != nil {
			log.Println("cannot parse data", err)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
		}

		dataUser := tmp.ToModel()
		row, err := uh.userUsecase.AddUser(dataUser)
		if row == -1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your email is already registered"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("register success"))
	}
}

func (uh *userHandler) LoginAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		authData := user.LoginModel{}
		c.Bind(&authData)
		fromToken, e := uh.userUsecase.Login(authData)
		if e != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
		}

		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("login success", fromToken))
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := _middleware.ExtractData(c)
		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModel(data)))
	}
}

func (uh *userHandler) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idFromToken, _ := _middleware.ExtractData(c)
		if idFromToken == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("you dont have access"))
		}
		row, errDel := uh.userUsecase.DeleteCase(idFromToken)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data user"))
		}
		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success delete data"))
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		idFromToken, _ := _middleware.ExtractData(c)
		err := c.Bind(&tmp)
		if err != nil {

			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
		}
		row, err := uh.userUsecase.UpdateCase(tmp.ToModel(), idFromToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed update data users, cek your input email"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed update data users, no data"))
		}

		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success update data"))
	}
}

func (uh *userHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newProduct ProductFormat
		id, role := common.ExtractData(c)
		bind := c.Bind(&newProduct)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}

		if role != "user" {
			log.Println("not user")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "not user",
			})
		}

		// file, err := c.FormFile("product_image")

		// if err != nil {
		// 	log.Println(err)
		// }

		// link := awss3.DoUpload(ah.conn, *file, file.Filename)
		// newproduct.Image = link
		status := uh.userUsecase.CreateProduct(newProduct.ToPU(), id)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "wrong input",
			})
		}
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
			"code":    status,
			"message": "success create product",
		})
	}
}

func (uh *userHandler) ReadAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := common.ExtractData(c)

		if role != "user" {
			log.Println("not user")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "not user",
			})
		}

		product, status := uh.userUsecase.ReadAllProduct(id)
		data := ParsePUToArr(product)

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

func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp ProductFormat
		bind := c.Bind(&tmp)

		qry := map[string]interface{}{}
		id, role := common.ExtractData(c)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}

		if role != "user" {
			log.Println("not user")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "not user",
			})
		}

		productid, err := strconv.Atoi(c.Param("idproduct"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		if tmp.Name != "" {
			qry["product_name"] = tmp.Name
		}

		if tmp.Unit != "" {
			qry["unit"] = tmp.Unit
		}

		if tmp.Stock != 0 {
			qry["stock"] = tmp.Stock
		}

		if tmp.Price != 0 {
			qry["price"] = tmp.Price
		}

		// file, err := c.FormFile("product_image")

		// if err != nil {
		// 	log.Println(err)
		// }

		// link := awss3.DoUpload(ah.conn, *file, file.Filename)
		// tmp.Image = link

		status := uh.userUsecase.UpdateProduct(tmp.ToPU(), productid, id)

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
			"message": "update success",
		})
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
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

		if role != "user" {
			log.Println("not user")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "not user",
			})
		}

		status := uh.userUsecase.DeleteProduct(cnv, id)

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
			"code":    status,
			"message": "success delete product",
		})
	}
}
