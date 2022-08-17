package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Order struct {
	ID         int
	UserID     int
	GrandTotal int
	Payment    string
	Status     string
	CreatedAt  time.Time
	Items      []Items
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

//logic
type OrderUseCase interface {
	// GetAllAdmin(limit, offset int) (data []Order, err error)
	CreateOrder(dataOrder Order, id int) int
	// CreateItems(data []Items) (row int, err error)
}

//query
type OrderData interface {
	// SelectDataAdminAll(limit, offset int) (data []Order, err error)
	// CreateItems(data []Items, orderID int) (row int, err error)
	InsertData(data []Items, id int) []Items
}

//handler
type OrderHandler interface {
	// GetAllAdmin() echo.HandlerFunc
	Create() echo.HandlerFunc
	// CreateItems() echo.HandlerFunc
}
