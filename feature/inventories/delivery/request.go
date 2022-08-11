package delivery

import "middleman-capstone/domain"

type InventoryFormat struct {
	Name string `json:"product_name" form:"product_name" validate:"required"`
	Qty  int    `json:"qty" form:"qty" validate:"required"`
	Unit string `json:"unit" form:"unit" validate:"required"`
}

func (inf *InventoryFormat) ToIP() domain.InventoryProduct {
	return domain.InventoryProduct{
		Name: inf.Name,
		Qty:  inf.Qty,
		Unit: inf.Unit,
	}
}
