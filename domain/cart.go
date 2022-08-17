package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Cart struct {
	ID         int
	Qty        int
	Subtotal   int
	Grandtotal int
	UserID     int
	ProductID  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    ProductCart
	User       UserCart
}

type ProductCart struct {
	ID           int
	ProductName  string
	Stock        int
	Unit         string
	Price        int
	ProductImage string
	UserID       int
	User         UserCart
}

type UserCart struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Role    string
	Address string
}

type CartUseCase interface {
	GetAllData(limit, offset, idFromToken int) (data []Cart, total int, err error)
	CreateData(data Cart) (row int, err error)
	UpdateData(qty, idProd, idFromToken int) (row int, err error)
	DeleteData(idProd, idFromToken int) (row int, err error)
}

//usecase
type ChartData interface {
	GetPriceProduct(id int) (price int, err error)
	GetStockProduct(idProduct int) (stok int, err error)
	GetQtyProductCart(idProduct int) (stok int, err error)
	InsertData(data Cart) (row int, err error)
	SelectData(limit, offset, idFromToken int) (data []Cart, err error)
	CheckCart(idProd, idFromToken int) (isExist bool, idCart, qty int, err error)
	UpdateDataDB(qty, idProd, idFromToken int) (row int, err error)
	DeleteDataDB(idProd, idFromToken int) (row int, err error)
}

type CartHandler interface {
	PostCart() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
}
