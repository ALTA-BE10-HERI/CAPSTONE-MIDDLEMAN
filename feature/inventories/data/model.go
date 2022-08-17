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
	ID          int `gorm:"autoIncrement"`
	InventoryID int
	UserID      int
	ProductID   int `json:"product_id" form:"product_id" validate:"required"`
	Name        string
	Qty         int    `json:"qty" form:"qty" validate:"required"`
	Unit        string `json:"unit" form:"unit" validate:"required"`
	Stock       int
	Idip        string
	CreatedAt   time.Time
}

func (ip *InventoryProduct) ToPU() domain.InventoryProduct {
	return domain.InventoryProduct{
		ID:        int(ip.ID),
		UserID:    ip.UserID,
		ProductID: ip.ProductID,
		Name:      ip.Name,
		Qty:       ip.Qty,
		Unit:      ip.Unit,
		Idip:      ip.Idip,
	}
}

func ParsePUToArr(arr []InventoryProduct) []domain.InventoryProduct {
	var res []domain.InventoryProduct
	for _, val := range arr {
		res = append(res, val.ToPU())
	}
	return res
}

func FromIP2(data []domain.InventoryProduct) []InventoryProduct {
	var res []InventoryProduct
	for _, val := range data {
		newdata := InventoryProduct{
			ID:          val.ID,
			InventoryID: val.InventoryID,
			UserID:      val.ID,
			ProductID:   val.ProductID,
			Name:        val.Name,
			Qty:         val.Qty,
			Unit:        val.Unit,
			Stock:       val.Stock,
			Idip:        val.Idip,
			CreatedAt:   val.CreatedAt,
		}
		res = append(res, newdata)
	}
	return res
}

func FromIP3(data []domain.InventoryProduct, invenid int, gen string, id int) []InventoryProduct {
	var res []InventoryProduct
	for _, val := range data {
		newdata := InventoryProduct{
			ID:          val.ID,
			InventoryID: invenid,
			UserID:      id,
			ProductID:   val.ProductID,
			Name:        val.Name,
			Qty:         val.Qty,
			Unit:        val.Unit,
			Stock:       val.Stock,
			Idip:        gen,
			CreatedAt:   val.CreatedAt,
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

func ParsePUToArr2(arr []domain.InventoryProduct, invenid string) map[string]interface{} {
	var arrmap []map[string]interface{}
	var res2 = map[string]interface{}{}
	for i := 0; i < len(arr); i++ {
		var res = map[string]interface{}{}
		res["product_id"] = arr[i].ProductID
		res["product_name"] = arr[i].Name
		res["qty"] = arr[i].Qty
		res["unit"] = arr[i].Unit

		arrmap = append(arrmap, res)
	}
	res2["inventory_id"] = invenid
	res2["items"] = arrmap
	return res2
}

func ParsePUToArr3(arr []domain.Inventory) []map[string]interface{} {
	var arrmap []map[string]interface{}
	for i := 0; i < len(arr); i++ {
		var res = map[string]interface{}{}
		res["inventory_id"] = arr[i].OutBound
		res["date"] = arr[i].CreatedAt

		arrmap = append(arrmap, res)
	}
	return arrmap
}

func ParsePUToArr4(arr []Inventory) []domain.Inventory {
	var res []domain.Inventory
	for _, val := range arr {
		res = append(res, val.ToPU2())
	}
	return res
}

func (ip *Inventory) ToPU2() domain.Inventory {
	return domain.Inventory{
		ID:       int(ip.ID),
		OutBound: ip.OutBound,
		UserID:   ip.UserID,
	}
}

func ParseToArr(arr domain.Inventory) map[string]interface{} {
	var res = map[string]interface{}{}
	res["inventory_id"] = arr.OutBound
	res["created_at"] = arr.CreatedAt
	return res
}
