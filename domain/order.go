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
	OrderID    string
	User       User
	CreatedAt  time.Time
}

//logic
type OrderUseCase interface {
	GetAllAdmin(limit, offset int) (data []Order, err error)
}

//query
type OrderData interface {
	SelectDataAdminAll(limit, offset int) (data []Order, err error)
}

//handler
type OrderHandler interface {
	GetAllAdmin() echo.HandlerFunc
}
