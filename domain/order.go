package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Order struct {
	ID          int
	UserID      int
	GrandTotal  int
	Status      string
	PaymentLink string
	OrderName   string
	CreatedAt   time.Time
	Items       []Items
}

type Items struct {
	ID          int `gorm:"autoIncrement"`
	OrderID     int
	ProductID   int
	ProductName string
	Subtotal    int
	Unit        string
	Qty         int
}

type PaymentWeb struct {
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	OrderName         string `json:"order_name" form:"order_name"`
}

//logic
type OrderUseCase interface {
	GetAllAdmin(limit, offset int) (data []Order, err error)
	CreateOrder(dataOrder Order, idUser int) int
	// CreateItems(data []Items) (row int, err error)
	GetAllUser(limit, offset, idUser int) (data []Order, err error)
	GetDetail(idUser, idOrder int) (grandTotal int, err error)
	GetItems(idOrder int) (data []Items, err error)
	Payment(grandTotal, idUser int) (orderName, url, token string, dataUser User)
	AcceptPayment(data PaymentWeb) (row int, err error)
}

//query
type OrderData interface {
	SelectDataAdminAll(limit, offset int) (data []Order, err error)
	SelectDataUserAll(limit, offset, idUser int) (data []Order, err error)
	// CreateItems(data []Items, orderID int) (row int, err error)
	InsertData(data []Items, id int) []Items
	Insert(data Order) (idOrder int, err error)
	GetUser(idUser int) (data User, err error)
	GetDetailData(idUser, idOrder int) (grandTotal int, err error)
	GetDetailItems(idOrder int) (data []Items, err error)
	AcceptPaymentData(data PaymentWeb) (row int, err error)
}

//handler
type OrderHandler interface {
	GetAllAdmin() echo.HandlerFunc
	Create() echo.HandlerFunc
	GetAllUser() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Payment() echo.HandlerFunc
	// CreateItems() echo.HandlerFunc
}
