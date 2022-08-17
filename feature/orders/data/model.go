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
	Items      []Items
}

type Items struct {
	ID          int `gorm:"autoIncrement"`
	OrderID     int
	ProductID   int
	ProductName string
	Subtotal    int
	Qty         int
}

func (o *Order) ToDomain() domain.Order {
	return domain.Order{
		ID:         int(o.ID),
		Status:     o.Status,
		GrandTotal: o.GrandTotal,
		UserID:     o.UserID,
	}
}

func (i *Items) ToDomainItems() domain.Items {
	return domain.Items{
		ID:          int(i.ID),
		OrderID:     int(i.OrderID),
		ProductID:   i.ProductID,
		ProductName: i.ProductName,
		Subtotal:    i.Subtotal,
		Qty:         i.Qty,
	}
}

func ParseToArr(arr []Order) []domain.Order {
	var res []domain.Order
	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ParseToArrItems(arr []Items) []domain.Items {
	var res []domain.Items
	for _, val := range arr {
		res = append(res, val.ToDomainItems())
	}
	return res
}

func FromDomItems(data []domain.Items, orderID int) []Items {
	var res []Items
	for _, val := range data {
		newdata := Items{
			OrderID:     orderID,
			ProductName: val.ProductName,
			ProductID:   val.ProductID,
			Subtotal:    val.Subtotal,
			Qty:         val.Qty,
		}
		res = append(res, newdata)
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

func FromDomainItems(data domain.Items) Items {
	var res Items
	res.OrderID = data.OrderID
	res.ProductID = data.ProductID
	res.ProductName = data.ProductName
	res.Subtotal = data.Subtotal
	res.Qty = data.Qty
	return res
}
