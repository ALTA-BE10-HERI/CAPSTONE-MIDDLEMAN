package data

import (
	"middleman-capstone/domain"
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	OutBound         string
	UserID           int
	InventoryProduct []InventoryProduct
}

type InventoryProduct struct {
	ID int `gorm:"autoIncrement"`
	// InventoryID int
	UserID      int
	ProductID   int `json:"product_id" form:"product_id" validate:"required"`
	ProductName string
	Qty         int    `json:"qty" form:"qty" validate:"required"`
	Unit        string `json:"unit" form:"unit" validate:"required"`
	Stock       int
	Idip        string
	CreatedAt   time.Time
}

func (ip *InventoryProduct) ToPU() domain.InventoryProduct {
	return domain.InventoryProduct{
		ID:          int(ip.ID),
		UserID:      ip.UserID,
		ProductID:   ip.ProductID,
		ProductName: ip.ProductName,
		Qty:         ip.Qty,
		Unit:        ip.Unit,
		Idip:        ip.Idip,
	}
}

func ParsePUToArr(arr []InventoryProduct) []domain.InventoryProduct {
	var res []domain.InventoryProduct
	for _, val := range arr {
		res = append(res, val.ToPU())
	}
	return res
}

func FromIP2(data []domain.InventoryProduct, id int, gen string) []InventoryProduct {
	var res []InventoryProduct
	for _, val := range data {
		newdata := InventoryProduct{
			UserID:      id,
			ProductID:   val.ProductID,
			ProductName: val.ProductName,
			Qty:         val.Qty,
			Unit:        val.Unit,
			Stock:       val.Stock,
			Idip:        gen,
		}
		res = append(res, newdata)
	}
	return res
}

func FromModel(data domain.Inventory, id int, gen string) Inventory {
	var res Inventory
	res.ID = uint(data.ID)
	res.OutBound = gen
	res.UserID = id
	res.CreatedAt = data.CreatedAt
	return res
}

func (i *Inventory) ToI() domain.Inventory {
	return domain.Inventory{
		ID:        int(i.ID),
		OutBound:  i.OutBound,
		CreatedAt: i.CreatedAt,
	}
}

func ParsePUToArr2(arr []domain.InventoryProduct) []map[string]interface{} {
	var arrmap []map[string]interface{}
	for i := 0; i < len(arr); i++ {
		var res = map[string]interface{}{}
		res["product_id"] = arr[i].ProductID
		res["product_name"] = arr[i].ProductName
		res["qty"] = arr[i].Qty
		res["unit"] = arr[i].Unit

		arrmap = append(arrmap, res)
	}
	return arrmap
}
