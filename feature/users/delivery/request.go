package delivery

import "middleman-capstone/domain"

type InsertFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
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
