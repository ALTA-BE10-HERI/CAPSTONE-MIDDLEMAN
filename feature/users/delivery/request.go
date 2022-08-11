package delivery

import "middleman-capstone/domain"

type InsertFormat struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Role     string `gorm:"default:user"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		Name:     i.Name,
		Email:    i.Email,
		Password: i.Password,
		Phone:    i.Phone,
		Address:  i.Address,
		Role:     i.Role,
	}
}

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

type InventoryFormat struct {
	Name string `json:"product_name" form:"product_name" validate:"required"`
	Qty  int    `json:"qty" form:"qty" validate:"required"`
	Unit string `json:"unit" form:"unit" validate:"required"`
}

func (inf *InventoryFormat) ToIP() domain.InventoryProduct {
	return domain.InventoryProduct{
		Name: inf.Name,
		Qty:  inf.Qty,
		Unit: inf.Unit,
	}
}
