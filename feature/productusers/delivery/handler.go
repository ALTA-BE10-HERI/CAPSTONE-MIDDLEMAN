package delivery

import (
	"fmt"
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/common"
	"middleman-capstone/feature/productusers/data"
	_helper "middleman-capstone/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type productUserHandler struct {
	productUserUseCase domain.ProductUserUseCase
}

func New(puuc domain.ProductUserUseCase) domain.ProductUserHandler {
	return &productUserHandler{
		productUserUseCase: puuc,
	}
}

func (puh *productUserHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newProduct ProductFormat
		id, role := common.ExtractData(c)
		bind := c.Bind(&newProduct)

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

		fileData, fileInfo, fileErr := c.Request().FormFile("product_image")

		// return err jika missing file
		if fileErr == http.ErrMissingFile || fileErr != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
		}

		// cek ekstension file upload
		extension, err_check_extension := _helper.CheckFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
		}

		// check file size
		err_check_size := _helper.CheckFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
		}

		// memberikan nama file
		fileName := time.Now().Format("2006-01-0215:04:05") + "-s3" + "." + extension
		url, errUploadImg := _helper.UploadImageToS3(fileName, fileData)
		if errUploadImg != nil {
			fmt.Println(errUploadImg)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file "))
		}

		newProduct.Image = url
		product, status := puh.productUserUseCase.CreateProduct(newProduct.ToPU(), id)
		data := data.ParsePUToArr3(product)

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
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    status,
			"message": "success create product",
			"data":    data,
		})
	}
}

func (puh *productUserHandler) ReadAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := common.ExtractData(c)

		if role != "user" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		product, status := puh.productUserUseCase.ReadAllProduct(id)
		data := data.ParsePUToArr2(product)

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

func (puh *productUserHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp ProductFormat
		bind := c.Bind(&tmp)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		id, role := common.ExtractData(c)

		if role != "user" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		fileData, fileInfo, fileErr := c.Request().FormFile("product_image")

		if fileData != nil {
			// return err jika missing file
			if fileErr == http.ErrMissingFile || fileErr != nil {
				return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
			}

			// cek ekstension file upload
			extension, err_check_extension := _helper.CheckFileExtension(fileInfo.Filename)
			if err_check_extension != nil {
				return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
			}

			// check file size
			err_check_size := _helper.CheckFileSize(fileInfo.Size)
			if err_check_size != nil {
				return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
			}

			// memberikan nama file
			fileName := time.Now().Format("2006-01-0215:04:05") + "-s3" + "." + extension
			url, errUploadImg := _helper.UploadImageToS3(fileName, fileData)
			if errUploadImg != nil {
				fmt.Println(errUploadImg)
				return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file "))
			}

			tmp.Image = url
		} else {
			tmp.Image = ""
		}

		productid, err := strconv.Atoi(c.Param("idproduct"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		products, status := puh.productUserUseCase.UpdateProduct(tmp.ToPU(), productid, id)
		data := data.ParsePUToArr3(products)

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
			"message": "success update data",
		})
	}
}

func (puh *productUserHandler) Delete() echo.HandlerFunc {
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
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		status := puh.productUserUseCase.DeleteProduct(cnv, id)

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

func (puh *productUserHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.QueryParam("productname")

		res, err := puh.productUserUseCase.SearchRestoBusiness(search)
		dares := data.ParsePUToArr2(res)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("failed to search data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", dares))
	}

}
