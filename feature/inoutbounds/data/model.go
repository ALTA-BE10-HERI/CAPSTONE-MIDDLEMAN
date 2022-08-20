package data

import (
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type InOutBounds struct {
	gorm.Model
	IdUser    int
	IdProduct int    `json:"product_id" form:"product_id" validate:"required"`
	Name      string `json:"product_name" form:"product_name"`
	Unit      string `json:"unit" form:"unit"`
	Qty       int    `json:"qty" form:"qty" validate:"required"`
	Role      string
	Note      string
}

func (iob *InOutBounds) ToIOB() domain.InOutBounds {
	return domain.InOutBounds{
		ID:        int(iob.ID),
		IdUser:    iob.IdUser,
		IdProduct: iob.IdProduct,
		Name:      iob.Name,
		Unit:      iob.Unit,
		Qty:       iob.Qty,
		Role:      iob.Role,
		Note:      iob.Note,
	}
}

func ParseIOBToArr(arr []InOutBounds) []domain.InOutBounds {
	var res []domain.InOutBounds
	for _, val := range arr {
		res = append(res, val.ToIOB())
	}
	return res
}

func FromIOB(data domain.InOutBounds) InOutBounds {
	var res InOutBounds
	res.ID = uint(data.ID)
	res.IdUser = data.IdUser
	res.IdProduct = data.IdProduct
	res.Name = data.Name
	res.Unit = data.Unit
	res.Qty = data.Qty
	res.Role = data.Role
	return res
}

func ParseIOBToArr3(arr domain.InOutBounds) map[string]interface{} {
	var res = map[string]interface{}{}
	res["product_id"] = arr.ID
	res["product_name"] = arr.Name
	res["unit"] = arr.Unit
	res["qty"] = arr.Qty

	return res
}
