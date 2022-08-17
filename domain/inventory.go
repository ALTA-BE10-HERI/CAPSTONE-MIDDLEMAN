package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Inventory struct {
	ID               int
	OutBound         string
	UserID           int
	InventoryProduct []InventoryProduct
	CreatedAt        time.Time
}

type InventoryProduct struct {
	ID          int
	InventoryID int
	UserID      int
	ProductID   int
	Name        string
	Qty         int
	Unit        string
	Stock       int
	Idip        string
	CreatedAt   time.Time
}

type InventoryHandler interface {
	Create() echo.HandlerFunc
	ReadUserDetail() echo.HandlerFunc
	ReadUserHistory() echo.HandlerFunc
}

type InventoryUseCase interface {
	CreateUserDetailInventory(newRecap Inventory, idUser int) (Inventory, int)
	ReadUserOutBoundDetail(id int, outboundIDGenerate string) ([]InventoryProduct, int, string)
	ReadUserOutBoundHistory(id int) ([]Inventory, int)
}

type InventoryData interface {
	CekStok(newRecap []InventoryProduct, id int) bool
	CreateUserDetailInventoryData(newRecap []InventoryProduct, id int, gen string, invenid int) []InventoryProduct
	CreateUserInventoryData(newRecap Inventory, id int, gen string) Inventory
	RekapStock(newRecap []InventoryProduct, id int, gen string) bool
	DeleteInOutBound(id int) (err string)
	ReadUserOutBoundDetailData(id int, outboundIDGenerate string) []InventoryProduct
	ReadUserOutBoundHistoryData(id int) []Inventory
}
