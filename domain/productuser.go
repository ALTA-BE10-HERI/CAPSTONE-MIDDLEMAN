package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type ProductUser struct {
	ID        int
	IdUser    int
	Name      string
	Unit      string
	Stock     int
	Price     int
	Image     string
	Reff      int
	CreatedAt time.Time
}

type ProductUserHandler interface {
	Create() echo.HandlerFunc
	ReadAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Search() echo.HandlerFunc
}

type ProductUserUseCase interface {
	CreateProduct(newProduct ProductUser, id int) (ProductUser, int)
	ReadAllProduct(id int) ([]ProductUser, int)
	UpdateProduct(updatedData ProductUser, productid, id int) (ProductUser, int)
	DeleteProduct(productid, id int) int
	SearchRestoBusiness(search string, idUser int) (result []ProductUser, err error)
}

type ProductUserData interface {
	CreateProductData(newProduct ProductUser) ProductUser
	ReadAllProductData(id int) []ProductUser
	UpdateProductData(data map[string]interface{}, productid, id int) ProductUser
	DeleteProductData(productid, id int) (err string)
	SearchRestoData(search string, idUser int) (result []ProductUser, err error)
}
