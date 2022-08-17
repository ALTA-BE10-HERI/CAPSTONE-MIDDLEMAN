package delivery

import "middleman-capstone/domain"

type FormatOrder struct {
	GrandTotal int `json:"grand_total" form:"grand_total"`
	Items      []Items
}

type Items struct {
	ProductID int    `json:"product_id" form:"product_id"`
	Qty       int    `json:"qty" form:"qty"`
	Unit      string `json:"unit" form:"unit"`
	Subtotal  int    `json:"subtotal" form:"subtotal"`
}

func ParseIFToArr(arr []Items) []domain.Items {
	var res []domain.Items
	for _, val := range arr {
		res = append(res, val.ToIF())
	}
	return res
}

func (pf *Items) ToIF() domain.Items {
	return domain.Items{
		ProductID: pf.ProductID,
		Qty:       pf.Qty,
		Unit:      pf.Unit,
		Subtotal:  pf.Subtotal,
	}
}

func ToDomain(format FormatOrder) domain.Order {
	return domain.Order{
		GrandTotal: format.GrandTotal,
		Items:      ParseIFToArr(format.Items),
	}
}
