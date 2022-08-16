package delivery

import "middleman-capstone/domain"

type ProductFormat struct {
	Name  string `json:"product_name" form:"product_name" validate:"required"`
	Unit  string `json:"unit" form:"unit" validate:"required"`
	Stock int    `json:"stock" form:"stock" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Image string `json:"product_image" form:"product_image" validate:"required"`
}

func (pf *ProductFormat) ToModel() domain.Product {
	return domain.Product{
		Name:  pf.Name,
		Unit:  pf.Unit,
		Stock: pf.Stock,
		Price: pf.Price,
		Image: pf.Image,
	}
}

func ParsePUToArr2(arr []domain.Product) []map[string]interface{} {
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
