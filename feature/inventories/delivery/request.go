package delivery

import "middleman-capstone/domain"

type InputFormat struct {
	Items []InventoryFormat `json:"items" form:"items" validate:"required"`
}
type InventoryFormat struct {
	ProductID int    `json:"product_id" form:"product_id" validate:"required"`
	Qty       int    `json:"qty" form:"qty" validate:"required"`
	Unit      string `json:"unit" form:"unit" validate:"required"`
}

func ParseIFToArr(arr []InventoryFormat) []domain.InventoryProduct {
	var res []domain.InventoryProduct
	for _, val := range arr {
		res = append(res, val.ToIF())
	}
	return res
}

func (pf *InventoryFormat) ToIF() domain.InventoryProduct {
	return domain.InventoryProduct{
		ProductID: pf.ProductID,
		Qty:       pf.Qty,
		Unit:      pf.Unit,
	}
}

func ToDomain(format InputFormat) domain.Inventory {
	return domain.Inventory{
		InventoryProduct: ParseIFToArr(format.Items),
	}
}
