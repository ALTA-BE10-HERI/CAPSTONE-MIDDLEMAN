package delivery

import (
	"middleman-capstone/domain"
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
