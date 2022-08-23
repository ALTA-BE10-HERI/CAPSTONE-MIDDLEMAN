package delivery

import (
	"middleman-capstone/domain"
	"time"
)

type ResponOrder struct {
	Status     string        `json:"status" form:"status" validate:"required"`
	OrderName  string        `json:"order_id" form:"order_id" validate:"required"`
	GrandTotal int           `json:"grand_total" form:"grand_total" validate:"required"`
	Items      []ResponItems `json:"items" form:"items" validate:"required"`
}

type ResponOrderUser struct {
	OrderName  string    `json:"order_id" form:"order_id" validate:"required"`
	Status     string    `json:"status" form:"status" validate:"required"`
	GrandTotal int       `json:"grand_total" form:"grand_total" validate:"required"`
	CreatedAt  time.Time `json:"date" form:"date" validate:"required"`
}

type ResponItems struct {
	ProductID   int    `json:"product_id" form:"product_id" validate:"required"`
	ProductName string `json:"product_name" form:"product_name"`
	Qty         int    `json:"qty" form:"qty" validate:"required"`
	Subtotal    int    `json:"subtotal" form:"subtotal" validate:"required"`
}

func FromModel(data domain.Items) ResponItems {
	return ResponItems{
		ProductID:   data.ProductID,
		ProductName: data.ProductName,
		Qty:         data.Qty,
		Subtotal:    data.Subtotal,
	}
}

func FromModel2(data []domain.Items, grand int, status string, ordername string) ResponOrder {
	return ResponOrder{
		Status:     status,
		OrderName:  ordername,
		GrandTotal: grand,
		Items:      FromModelList(data),
	}
}

func FromModelList(data []domain.Items) []ResponItems {
	res := []ResponItems{}
	for key := range data {
		res = append(res, FromModel(data[key]))
	}
	return res
}

func FromModelUser(data domain.Order) ResponOrderUser {
	return ResponOrderUser{
		OrderName:  data.OrderName,
		Status:     data.Status,
		GrandTotal: data.GrandTotal,
		CreatedAt:  data.CreatedAt,
	}
}

func FromModelListUser(data []domain.Order) []ResponOrderUser {
	res := []ResponOrderUser{}
	for key := range data {
		res = append(res, FromModelUser(data[key]))
	}
	return res
}
