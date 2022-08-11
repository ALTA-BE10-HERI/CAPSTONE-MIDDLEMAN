package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type InventoryProduct struct {
	ID        int
	IdUser    int
	Name      string
	Unit      string
	Qty       int
	CreatedAt time.Time
}

type InventoryHandler interface {
	Create() echo.HandlerFunc
}

type InventoryUseCase interface {
	CreateInventory(newRecap InventoryProduct, idUser int) int
}

type InventoryData interface {
	CreateInventoryData(newRecap InventoryProduct) InventoryProduct
	StockUpdate(newRecap InventoryProduct) bool
}
