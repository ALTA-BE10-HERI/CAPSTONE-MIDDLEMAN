package delivery

import (
	"fmt"
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/common"
	_helper "middleman-capstone/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUseCase domain.ProductUseCase
}

func New(ps domain.ProductUseCase) domain.ProductHandler {
	return &productHandler{
		productUseCase: ps,
	}
}

func (ph *productHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newProduct ProductFormat
		idAdmin, role := common.ExtractData(c)
		bind := c.Bind(&newProduct)

		if bind != nil {
			log.Println("cant bind data")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		if role != "admin" {
			log.Println("admin")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}
		// file, err := c.FormFile("product_image")

		// if err != nil {
		// 	log.Println(err)
		// }

		// link := awss3.DoUpload(ah.conn, *file, file.Filename)
		// newproduct.Image = link
		// =================================================================================
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

		// =================================================================================

		newProduct.Image = url
		status := ph.productUseCase.CreateProduct(newProduct.ToModel(), idAdmin)

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
			"message": "success adding a product",
		})
	}
}

func (ph *productHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitcnv, _ := strconv.Atoi(limit)
		offsetcnv, _ := strconv.Atoi(offset)
		result, err := ph.productUseCase.GetAllProduct(limitcnv, offsetcnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"code":    200,
			"message": "get data success",
		})
	}
}

func (ph *productHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp ProductFormat
		bind := c.Bind(&tmp)
		_, role := common.ExtractData(c)

		if bind != nil {
			log.Println("cant bind data")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}
		if role != "admin" {
			log.Println("not admin")
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
		idProduct, err := strconv.Atoi(c.Param("idproduct"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}

		row, err := ph.productUseCase.UpdateProduct(tmp.ToModel(), idProduct)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("there is an error in internal server"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("wrong input"))
		}

		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success update data"))
	}
}

func (ph *productHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("idproduct"))

		if err != nil {
			log.Println("cant convert id product", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "cant convert id product",
			})
		}
		_, role := common.ExtractData(c)

		if role != "admin" {
			log.Println("not admin")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		status := ph.productUseCase.DeleteProduct(cnv)

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

func (puh *productHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.QueryParam("productname")
		res, err := puh.productUseCase.SearchRestoBusiness(search)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseBadRequest("failed to search data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", res))
	}

}
