package delivery

import "middleman-capstone/domain"

type CartFormat struct {
	IdProduct int `json:"product_id" form:"product_id" validate:"required"`
	Qty       int `json:"qty" form:"qty" validate:"required"`
}

func (cf *CartFormat) ToModel() domain.InOutBounds {
	return domain.InOutBounds{
		IdProduct: cf.IdProduct,
		Qty:       cf.Qty,
	}
}
