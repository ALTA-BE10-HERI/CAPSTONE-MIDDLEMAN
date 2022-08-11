package delivery

import "middleman-capstone/domain"

type ProductFormat struct {
	Name  string `json:"product_name" form:"product_name" validate:"required"`
	Unit  string `json:"unit" form:"unit" validate:"required"`
	Stock int    `json:"stock" form:"stock" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Image string `json:"product_image" form:"product_image" validate:"required"`
}

func (pf *ProductFormat) ToPU() domain.ProductUser {
	return domain.ProductUser{
		Name:  pf.Name,
		Unit:  pf.Unit,
		Stock: pf.Stock,
		Price: pf.Price,
		Image: pf.Image,
	}
}
