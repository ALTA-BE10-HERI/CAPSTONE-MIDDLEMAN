package data

import (
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	IdAdmin int
	Name    string `json:"product_name" form:"product_name" validate:"required" gorm:"unique"`
	Unit    string `json:"unit" form:"unit" validate:"required"`
	Stock   int    `json:"stock" form:"stock" validate:"required"`
	Price   int    `json:"price" form:"price" validate:"required"`
	Satuan  int    `json:"satuan" form:"satuan"`
	Image   string `json:"product_image" form:"product_image" validate:"required"`
}

func (p *Product) ToModel() domain.Product {
	return domain.Product{
		ID:        int(p.ID),
		IdAdmin:   p.IdAdmin,
		Name:      p.Name,
		Unit:      p.Unit,
		Stock:     p.Stock,
		Price:     p.Price,
		Satuan:    p.Satuan,
		Image:     p.Image,
		CreatedAt: p.CreatedAt,
	}
}

func ParseProductToArr(arr []Product) []domain.Product {
	var res []domain.Product
	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Product) Product {
	var res Product
	res.ID = uint(data.ID)
	res.IdAdmin = data.IdAdmin
	res.Name = data.Name
	res.Unit = data.Unit
	res.Stock = data.Stock
	res.Satuan = data.Satuan
	res.Price = data.Price
	res.Image = data.Image
	res.CreatedAt = data.CreatedAt
	return res
}
func toModelList(data []Product) []domain.Product {
	res := []domain.Product{}

	for key := range data {
		res = append(res, data[key].ToModel())
	}
	return res
}
