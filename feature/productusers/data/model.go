package data

import (
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type ProductUser struct {
	gorm.Model
	IdUser int
	Name   string `json:"product_name" form:"product_name" validate:"required"`
	Unit   string `json:"unit" form:"unit" validate:"required"`
	Stock  int    `json:"stock" form:"stock" validate:"required"`
	Price  int    `json:"price" form:"price" validate:"required"`
	Image  string `json:"product_image" form:"product_image"`
	Reff   int
}

func (pu *ProductUser) ToPU() domain.ProductUser {
	return domain.ProductUser{
		ID:        int(pu.ID),
		IdUser:    pu.IdUser,
		Name:      pu.Name,
		Unit:      pu.Unit,
		Stock:     pu.Stock,
		Price:     pu.Price,
		Image:     pu.Image,
		CreatedAt: pu.CreatedAt,
	}
}

func ParsePUToArr(arr []ProductUser) []domain.ProductUser {
	var res []domain.ProductUser
	for _, val := range arr {
		res = append(res, val.ToPU())
	}
	return res
}

func FromPU(data domain.ProductUser) ProductUser {
	var res ProductUser
	res.ID = uint(data.ID)
	res.IdUser = data.IdUser
	res.Name = data.Name
	res.Unit = data.Unit
	res.Stock = data.Stock
	res.Price = data.Price
	res.Image = data.Image
	res.CreatedAt = data.CreatedAt
	return res
}

func ParsePUToArr2(arr []domain.ProductUser) []map[string]interface{} {
	var arrmap []map[string]interface{}
	for i := 0; i < len(arr); i++ {
		var res = map[string]interface{}{}
		res["id"] = arr[i].ID
		res["product_name"] = arr[i].Name
		res["unit"] = arr[i].Unit
		res["stock"] = arr[i].Stock
		res["price"] = arr[i].Price
		res["product_image"] = arr[i].Image

		arrmap = append(arrmap, res)
	}
	return arrmap
}

func ParsePUToArr3(arr domain.ProductUser) map[string]interface{} {
	var res = map[string]interface{}{}
	res["id"] = arr.ID
	res["product_name"] = arr.Name
	res["unit"] = arr.Unit
	res["stock"] = arr.Stock
	res["price"] = arr.Price
	res["product_image"] = arr.Image

	return res
}

func toModelList(data []ProductUser) []domain.ProductUser {
	res := []domain.ProductUser{}

	for key := range data {
		res = append(res, data[key].ToPU())
	}
	return res
}
