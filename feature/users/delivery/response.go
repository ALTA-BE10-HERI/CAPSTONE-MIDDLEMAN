package delivery

import (
	"middleman-capstone/domain"
	"time"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Role     string
}

func FromModel(data domain.User) User {
	return User{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
	}
}

func FromModelList(data []domain.User) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromModel(data[key]))
	}
	return result
}

type ProductUser struct {
	ID        int
	IdUser    int
	Name      string
	Unit      string
	Stock     int
	Price     int
	Image     string
	CreatedAt time.Time
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

// func ParsePUToArr(arr []domain.ProductUser) []domain.ProductUser {
// 	var res []domain.ProductUser
// 	for _, val := range arr {
// 		res = append(res, val.FromPU())
// 	}
// 	return res
// }
