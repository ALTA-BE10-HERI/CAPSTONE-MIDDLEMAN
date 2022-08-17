package delivery

import "middleman-capstone/domain"

type FormatOrder struct {
	GrandTotal int `json:"grand_total" form:"grand_total"`
	Items      []Items
}

type Items struct {
	ProductID int    `json:"product_id" form:"product_id"`
	Subtotal  int    `json:"subtotal" form:"subtotal"`
	Unit      string `json:"unit" form:"unit"`
	Qty       int    `json:"qty" form:"qty"`
}

func (i *Items) ToDomainItems() domain.Items {
	return domain.Items{
		ProductID: i.ProductID,
		Subtotal:  i.Subtotal,
		Unit:      i.Unit,
		Qty:       i.Qty,
	}
}

func ParseToArrItems(arr []Items) []domain.Items {
	var res []domain.Items
	for _, val := range arr {
		res = append(res, val.ToDomainItems())
	}
	return res
}

func FromDomainItems(data domain.Items) Items {
	var res Items
	res.ProductID = data.ProductID
	res.Subtotal = data.Subtotal
	res.Unit = data.Unit
	res.Qty = data.Qty
	return res
}
