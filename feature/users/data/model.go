package data

import (
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Phone    string
	Email    string `gorm:"unique" validate:"required,email"`
	Address  string
	Password string
	Role     string `gorm:"default:user"`
}

type ProductUser struct {
	gorm.Model
	IdUser int
	Name   string `json:"product_name" form:"product_name" validate:"required"`
	Unit   string `json:"unit" form:"unit" validate:"required"`
	Stock  int    `json:"stock" form:"stock" validate:"required"`
	Price  int    `json:"price" form:"price" validate:"required"`
	Image  string `json:"product_image" form:"product_image"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:        int(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Password:  u.Password,
		Address:   u.Address,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
func ParseToArr(arr []User) []domain.User {
	var res []domain.User
	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Email = data.Email
	res.Name = data.Name
	res.Password = data.Password
	res.Phone = data.Phone
	res.Address = data.Address
	res.Role = data.Role
	return res
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
	res.IdUser = data.IdUser
	res.Name = data.Name
	res.Unit = data.Unit
	res.Stock = data.Stock
	res.Price = data.Price
	res.Image = data.Image
	res.CreatedAt = data.CreatedAt
	return res
}
