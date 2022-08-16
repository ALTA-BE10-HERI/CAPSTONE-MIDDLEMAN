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

type Inventory struct {
	ID         int
	IdOutBound string
	CreatedAt  time.Time
}

type InventoryHandler interface {
	Create() echo.HandlerFunc
	ReadUser() echo.HandlerFunc
}

type InventoryUseCase interface {
	CreateUserDetailInventory(newRecap []InventoryProduct, idUser int) int
	ReadUserOutBoundDetail(id int, outboundIDGenerate string) ([]InventoryProduct, int)
}

type InventoryData interface {
	CekStok(newRecap []InventoryProduct, id int) bool
	CreateUserDetailInventoryData(newRecap []InventoryProduct, id int, gen string) []InventoryProduct
	CreateUserInventoryData(newRecap Inventory) Inventory
	RekapStock(newRecap []InventoryProduct, id int) bool
	DeleteInOutBound(id int) (err string)
	ReadUserOutBoundDetailData(id int, outboundIDGenerate string) []InventoryProduct
}
