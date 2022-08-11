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
