package data

import (
	"middleman-capstone/domain"
	"time"
)

type InventoryProduct struct {
	ID        int `gorm:"autoIncrement"`
	IdUser    int
	Name      string `json:"product_name" form:"product_name" validate:"required"`
	Unit      string `json:"unit" form:"unit" validate:"required"`
	Qty       int    `json:"qty" form:"qty" validate:"required"`
	CreatedAt time.Time
}

func (ip *InventoryProduct) ToIP() domain.InventoryProduct {
	return domain.InventoryProduct{
		ID:        ip.ID,
		IdUser:    ip.IdUser,
		Name:      ip.Name,
		Unit:      ip.Unit,
		Qty:       ip.Qty,
		CreatedAt: ip.CreatedAt,
	}
}

func FromIP(data domain.InventoryProduct) InventoryProduct {
	var res InventoryProduct
	res.ID = data.ID
	res.IdUser = data.IdUser
	res.Name = data.Name
	res.Unit = data.Unit
	res.Qty = data.Qty
	res.CreatedAt = data.CreatedAt
	return res
}
