package delivery

import (
	"log"
	"middleman-capstone/domain"
	_middleware "middleman-capstone/feature/common"
	user "middleman-capstone/feature/users"
	_helper "middleman-capstone/helper"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

		validate := validator.New()
		if errValidate := validate.Struct(tmp); errValidate != nil {
			return errValidate
		}

		if err != nil {
			log.Println("cannot parse data", err)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
		}

		dataUser := tmp.ToModel()
		row, err := uh.userUsecase.AddUser(dataUser)
		if row == -1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("please make sure all fields are filled in correctly"))
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("your email or handphone number is already registered"))
		}
		return c.JSON(http.StatusCreated, _helper.ResponseCreate("register success"))
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

		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("get data success", FromModel(data)))
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
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("failed to bind data, check your input"))
		}
		validate := validator.New()
		if errValidate := validate.Struct(tmp); errValidate != nil {
			return errValidate
		}
		row, _ := uh.userUsecase.UpdateCase(tmp.ToModel(), idFromToken)
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("failed update data users, your email or phone already registerd"))
		}
		if row == 404 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("nothing to update data"))
		}
		if row == 401 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("your format phone is wrong"))
		}
		if row == 400 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("your format email is wrong"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success update data"))
	}
}
