package data

import (
	"middleman-capstone/domain"
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	IdOutBound       string
	InventoryProduct []InventoryProduct `gorm:"foreignKey:IdOutBound;references:Idip;constraint:OnDelete:CASCADE"`
}

type InventoryProduct struct {
	ID        int
	IdUser    int
	IdProduct int `json:"product_id" form:"product_id" validate:"required"`
	Name      string
	Qty       int    `json:"qty" form:"qty" validate:"required"`
	Unit      string `json:"unit" form:"unit" validate:"required"`
	Idip      string `gorm:"primaryKey"`
	Stock     int
	CreatedAt time.Time
}

func (ip *InventoryProduct) ToPU() domain.InventoryProduct {
	return domain.InventoryProduct{
		ID:        int(ip.ID),
		IdUser:    ip.IdUser,
		IdProduct: ip.IdProduct,
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

func FromIP2(data []domain.InventoryProduct, id int, gen string) []InventoryProduct {
	var res []InventoryProduct
	for _, val := range data {
		newdata := InventoryProduct{
			IdUser:    id,
			IdProduct: val.IdProduct,
			Name:      val.Name,
			Qty:       val.Qty,
			Unit:      val.Unit,
			Idip:      gen,
			Stock:     val.Stock,
		}
		res = append(res, newdata)
	}
	return res
}

// func FromModel(data domain.Inventory, id int, gen string) Inventory {
// 	return Inventory{
// 		ID:         uint(data.ID),
// 		IdOutBound: data.IdOutBound,
// 		CreatedAt:  data.CreatedAt,
// 		// Outbound:   FromIP2(data.Outbound, id, gen),
// 	}
// }

func FromModel(data domain.Inventory, id int, gen string) Inventory {
	var res Inventory
	res.ID = uint(data.ID)
	res.IdOutBound = gen
	res.CreatedAt = data.CreatedAt
	return res
}

func (i *Inventory) ToI() domain.Inventory {
	return domain.Inventory{
		ID:         int(i.ID),
		IdOutBound: i.IdOutBound,
		CreatedAt:  i.CreatedAt,
	}
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
