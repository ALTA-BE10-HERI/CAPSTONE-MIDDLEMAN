package data

import (
	"middleman-capstone/domain"
	"time"

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

type InventoryProduct struct {
	ID        int `gorm:"autoIncrement"`
	IdUser    int
	Name      string `json:"product_name" form:"product_name" validate:"required"`
	Unit      string `json:"unit" form:"unit" validate:"required"`
	Qty       int    `json:"qty" form:"qty" validate:"required"`
	CreatedAt time.Time
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

func (ip *InventoryProduct) ToIP() domain.InventoryProduct {
	return domain.InventoryProduct{
		ID:        ip.ID,
		IdUser:    ip.IdUser,
		Name:      ip.Name,
		Unit:      ip.Unit,
		Qty:       ip.Qty,
		CreatedAt: ip.CreatedAt,
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

func FromIP(data domain.InventoryProduct) InventoryProduct {
	var res InventoryProduct
	res.ID = data.ID
	res.IdUser = data.IdUser
	res.Name = data.Name
	res.Unit = data.Unit
	res.Qty = data.Qty
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
