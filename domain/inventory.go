package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type InventoryProduct struct {
	ID        int
	IdUser    int
	IdProduct int
	Name      string
	Qty       int
	Unit      string
	Idip      string
	Stock     int
	CreatedAt time.Time
}

type InventoryHandler interface {
	Create() echo.HandlerFunc
}

type InventoryUseCase interface {
	CreateUserInventory(newRecap []InventoryProduct, idUser int) int
}

type InventoryData interface {
	CreateUserInventoryData(newRecap []InventoryProduct, id int, gen string) []InventoryProduct
	CekStock(newRecap []InventoryProduct, id int) bool
}
