package data

import (
	"middleman-capstone/domain"
	"time"
)

type InventoryProduct struct {
	ID        int `gorm:"autoIncrement"`
	IdUser    int
	IdProduct int `json:"product_id" form:"product_id" validate:"required"`
	Name      string
	Qty       int    `json:"qty" form:"qty" validate:"required"`
	Unit      string `json:"unit" form:"unit" validate:"required"`
	Idip      string ``
	Stock     int
	CreatedAt time.Time
}

type Inventory struct {
	ID         int
	IdOutBound string
	CreatedAt  time.Time
}

func (ip *InventoryProduct) ToIP() domain.InventoryProduct {
	return domain.InventoryProduct{
		ID:        ip.ID,
		IdUser:    ip.IdUser,
		Unit:      ip.Unit,
		Qty:       ip.Qty,
		CreatedAt: ip.CreatedAt,
	}
}

func ParseIPToArr(arr []InventoryProduct) []domain.InventoryProduct {
	var res []domain.InventoryProduct
	for _, val := range arr {
		res = append(res, val.ToIP())
	}
	return res
}

func FromIP2(data []domain.InventoryProduct, id int, gen string) []InventoryProduct {
	var res []InventoryProduct
	for _, val := range data {
		newdata := InventoryProduct{
			IdUser:    id,
			IdProduct: val.IdProduct,
			Unit:      val.Unit,
			Qty:       val.Qty,
			Idip:      gen,
		}
		res = append(res, newdata)
	}
	return res
}

func FromIP3(data []domain.InventoryProduct) []InventoryProduct {
	var res []InventoryProduct
	for _, val := range data {
		newdata := InventoryProduct{
			IdUser:    val.IdUser,
			IdProduct: val.IdProduct,
			Name:      val.Name,
			Qty:       val.Qty,
			Unit:      val.Unit,
			Stock:     val.Stock,
		}
		res = append(res, newdata)
	}
	return res
}

func (i *Inventory) ToI() domain.Inventory {
	return domain.Inventory{
		ID:         int(i.ID),
		IdOutBound: i.IdOutBound,
		CreatedAt:  i.CreatedAt,
	}
}

func FromModel(data domain.Inventory) Inventory {
	return Inventory{
		ID:         data.ID,
		IdOutBound: data.IdOutBound,
		CreatedAt:  data.CreatedAt,
	}
}

func (ip *InventoryProduct) ToPU() domain.InventoryProduct {
	return domain.InventoryProduct{
		Idip:      ip.Idip,
		IdProduct: ip.IdProduct,
		Name:      ip.Name,
		Qty:       ip.Qty,
		Unit:      ip.Unit,
	}
}

func ParsePUToArr(arr []InventoryProduct) []domain.InventoryProduct {
	var res []domain.InventoryProduct
	for _, val := range arr {
		res = append(res, val.ToPU())
	}
	return res
}

func ParsePUToArr2(arr []domain.InventoryProduct) []map[string]interface{} {
	var arrmap []map[string]interface{}
	for i := 0; i < len(arr); i++ {
		var res = map[string]interface{}{}
		res["product_id"] = arr[i].IdProduct
		res["product_name"] = arr[i].Name
		res["qty"] = arr[i].Qty
		res["unit"] = arr[i].Unit

		arrmap = append(arrmap, res)
	}
	return arrmap
}
