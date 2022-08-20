package delivery

import "middleman-capstone/domain"

type InOutBounds struct {
	IdProduct int    `json:"product_id" form:"product_id" validate:"required"`
	Name      string `json:"product_name" form:"product_name"`
	Unit      string `json:"unit" form:"unit"`
	Qty       int    `json:"qty" form:"qty" validate:"required"`
}

func FromModel(data domain.InOutBounds) InOutBounds {
	return InOutBounds{
		IdProduct: data.IdProduct,
		Name:      data.Name,
		Unit:      data.Unit,
		Qty:       data.Qty,
	}
}

func FromModelList(data []domain.InOutBounds) []InOutBounds {
	res := []InOutBounds{}
	for key := range data {
		res = append(res, FromModel(data[key]))
	}
	return res
}
