package data

import (
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     int
	GrandTotal int
	Status     string
	CartID     int
	Cart       Cart `gorm:"foreignKey:CartID;references:ID;constraint:OnDelete:CASCADE"`
	// OrderDetail []OrderDetail
}

// type OrderDetail struct {
// 	ID          int `gorm:"autoIncrement"`
// 	OrderID     int
// 	ProductID   int
// 	ProductName string `gorm:"column:product_name"`
// 	Price       int
// 	Qty         int
// }

type Cart struct {
	gorm.Model
	Qty       int
	Status    string
	Subtotal  int
	UserID    int
	ProductID int
	User      []User
	Product   []Product
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Phone    string `gorm:"unique"`
	Role     string
	Address  string
	Product  []Product
}

type Product struct {
	gorm.Model
	Stock  int
	Status string
	Name   string
	Image  string
	Unit   string
	Price  int
	UserID int
	Cart   []Cart
	User   User
}

func (o *Order) ToDomain() domain.Order {
	return domain.Order{
		ID:         int(o.ID),
		Status:     o.Status,
		GrandTotal: o.GrandTotal,
		UserID:     o.UserID,
	}
}

func ParseToArr(arr []Order) []domain.Order {
	var res []domain.Order
	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func FromDomain(data domain.Order) Order {
	var res Order
	res.Status = data.Status
	res.GrandTotal = data.GrandTotal
	res.UserID = data.UserID
	return res
}
