package delivery

import "middleman-capstone/domain"

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"product_name" form:"product_name" validate:"required"`
	Unit  string `json:"unit" form:"unit" validate:"required"`
	Stock int    `json:"stock" form:"stock" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Image string `json:"product_image" form:"product_image" validate:"required"`
}

func FromModel(data domain.Product) Product {
	return Product{
		ID:    data.ID,
		Price: data.Price,
		Name:  data.Name,
		Image: data.Image,
		Stock: data.Stock,
		Unit:  data.Unit,
	}
}

func FromModelList(data []domain.Product) []Product {
	res := []Product{}
	for key := range data {
		res = append(res, FromModel(data[key]))
	}
	return res
}
