package delivery

import "middleman-capstone/domain"

type InsertFormat struct {
	Name     string `json:"name" form:"name" validate:"min=2"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone" gorm:"unique" validate:"min=10"`
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
