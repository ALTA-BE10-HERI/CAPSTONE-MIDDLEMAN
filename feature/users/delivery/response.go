package delivery

import (
	"middleman-capstone/domain"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func FromModel(data domain.User) User {
	return User{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Phone:    data.Phone,
		Password: data.Password,
		Address:  data.Address,
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
		ID:    int(pu.ID),
		Name:  pu.Name,
		Unit:  pu.Unit,
		Stock: pu.Stock,
		Price: pu.Price,
		Image: pu.Image,
	}
}

func FromPU(data domain.ProductUser) ProductUser {
	var res ProductUser
	res.Name = data.Name
	res.Unit = data.Unit
	res.Stock = data.Stock
	res.Price = data.Price
	res.Image = data.Image
	res.CreatedAt = data.CreatedAt
	return res
}

func ParseGETProfile(arr domain.User) []map[string]interface{} {
	var arrmap []map[string]interface{}
	var res = map[string]interface{}{}
	res["id"] = arr.ID
	res["name"] = arr.Name
	res["email"] = arr.Email
	res["phone"] = arr.Phone
	res["address"] = arr.Address

	arrmap = append(arrmap, res)
	return arrmap
}
