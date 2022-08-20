package delivery

import (
	"middleman-capstone/domain"
	"time"
)

type Inventory struct {
	OutBound         string             `json:"inventory_id" form:"inventory_id" validate:"required"`
	InventoryProduct []InventoryProduct `json:"items" form:"items" validate:"required"`
}

type InventoryHistory struct {
	OutBound  string    `json:"inventory_id" form:"inventory_id" validate:"required"`
	CreatedAt time.Time `json:"date" form:"date" validate:"required"`
}

type InventoryProduct struct {
	ProductID int    `json:"product_id" form:"product_id" validate:"required"`
	Name      string `json:"product_name" form:"product_id"`
	Qty       int    `json:"qty" form:"qty" validate:"required"`
	Unit      string `json:"unit" form:"unit" validate:"required"`
}

func FromModel(data domain.InventoryProduct) InventoryProduct {
	return InventoryProduct{
		ProductID: data.ProductID,
		Name:      data.Name,
		Unit:      data.Unit,
		Qty:       data.Qty,
	}
}

func FromModel2(data []domain.InventoryProduct, inven string) Inventory {
	return Inventory{
		OutBound:         inven,
		InventoryProduct: FromModelList(data),
	}
}

func FromModelList(data []domain.InventoryProduct) []InventoryProduct {
	res := []InventoryProduct{}
	for key := range data {
		res = append(res, FromModel(data[key]))
	}
	return res
}

func FromModelHistory(data domain.Inventory) InventoryHistory {
	return InventoryHistory{
		OutBound:  data.OutBound,
		CreatedAt: data.CreatedAt,
	}
}

func FromModelListHistory(data []domain.Inventory) []InventoryHistory {
	res := []InventoryHistory{}
	for key := range data {
		res = append(res, FromModelHistory(data[key]))
	}
	return res
}
