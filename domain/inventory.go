package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Inventory struct {
	ID               int
	OutBound         string
	UserID           int
	Role             string
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
	Role        string
	CreatedAt   time.Time
}

type InventoryHandler interface {
	CreateUser() echo.HandlerFunc
	ReadUserDetail() echo.HandlerFunc
	ReadUserHistory() echo.HandlerFunc
	CreateAdmin() echo.HandlerFunc
	ReadAdminHistory() echo.HandlerFunc
	ReadAdminDetail() echo.HandlerFunc
}

type InventoryUseCase interface {
	CreateUserInventory(newRecap Inventory, idUser int) (Inventory, int)
	ReadUserOutBoundDetail(id int, outboundIDGenerate string) ([]InventoryProduct, int, string)
	ReadUserOutBoundHistory(id int) ([]Inventory, int)
	CreateAdminInventory(newRecap Inventory, id int, role string) (Inventory, int)
	ReadAdminOutBoundHistory() ([]Inventory, int)
	ReadAdminOutBoundDetail(outboundIDGenerate string) ([]InventoryProduct, int, string)
}

type InventoryData interface {
	CekStok(newRecap []InventoryProduct, id int) bool
	CreateUserDetailInventoryData(newRecap []InventoryProduct, id int, gen string, invenid int) []InventoryProduct
	CreateUserInventoryData(newRecap Inventory, id int, gen string) Inventory
	RekapStock(newRecap []InventoryProduct, id int, gen string) bool
	DeleteInOutBound(id int) (err string)
	ReadUserOutBoundDetailData(id int, outboundIDGenerate string) []InventoryProduct
	ReadUserOutBoundHistoryData(id int) []Inventory
	CreateAdminInventoryData(newRecap Inventory, id int, gen string) Inventory
	CreateAdminDetailInventoryData(newRecap []InventoryProduct, id int, gen string, invenid int, role string) []InventoryProduct
	RekapAdminStock(newRecap []InventoryProduct, id int, gen string) bool
	DeleteAdminInOutBound() (err string)
	ReadAdminOutBoundHistoryData() []Inventory
	ReadAdminOutBoundDetailData(outboundIDGenerate string) []InventoryProduct
}
