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
	Idip      string
	Stock     int
	CreatedAt time.Time
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

func FromIP2(data []domain.InventoryProduct, id int) []InventoryProduct {
	var res []InventoryProduct
	for _, val := range data {
		newdata := InventoryProduct{
			IdUser:    id,
			IdProduct: val.IdProduct,
			Unit:      val.Unit,
			Qty:       val.Qty,
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
