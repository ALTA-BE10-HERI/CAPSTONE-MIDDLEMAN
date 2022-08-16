package delivery

import "middleman-capstone/domain"

type ProductUser struct {
	ID    int    `json:"id"`
	Name  string `json:"product_name" form:"product_name" validate:"required"`
	Unit  string `json:"unit" form:"unit" validate:"required"`
	Stock int    `json:"stock" form:"stock" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Image string `json:"product_image" form:"product_image" validate:"required"`
}

func FromModel(data domain.ProductUser) ProductUser {
	return ProductUser{
		ID:    data.ID,
		Name:  data.Name,
		Unit:  data.Unit,
		Stock: data.Stock,
		Price: data.Price,
		Image: data.Image,
	}
}

func FromModelList(data []domain.ProductUser) []ProductUser {
	res := []ProductUser{}
	for key := range data {
		res = append(res, FromModel(data[key]))
	}
	return res
}
