package data

import (
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      int
	GrandTotal  int
	Status      string
	PaymentLink string
	OrderName   string
	Items       []Items
}

type Items struct {
	ID          int `gorm:"autoIncrement"`
	OrderID     int
	ProductID   int `json:"product_id" form:"product_id"`
	ProductName string
	Subtotal    int    `json:"subtotal" form:"subtotal"`
	Unit        string `json:"unit" form:"unit"`
	Qty         int    `json:"qty" form:"qty"`
}

type OrderPayment struct {
	Name       string
	Email      string
	Phone      int
	GrandTotal int
}

type PaymentWeb struct {
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	OrderName         string `json:"order_name" form:"order_name"`
}

func FromIP2(data []domain.Items, id int) []Items {
	var res []Items
	for _, val := range data {
		newdata := Items{
			ID:          val.ID,
			OrderID:     id,
			ProductID:   val.ProductID,
			ProductName: val.ProductName,
			Subtotal:    val.Subtotal,
			Qty:         val.Qty,
			Unit:        val.Unit,
		}
		res = append(res, newdata)
	}
	return res
}

func ParsePUToArr(arr []Items) []domain.Items {
	var res []domain.Items
	for _, val := range arr {
		res = append(res, val.ToPU())
	}
	return res
}

func (ip *Items) ToPU() domain.Items {
	return domain.Items{
		ID:          int(ip.ID),
		OrderID:     ip.OrderID,
		ProductID:   ip.ProductID,
		ProductName: ip.ProductName,
		Subtotal:    ip.Subtotal,
		Qty:         ip.Qty,
		Unit:        ip.Unit,
	}
}

func (o *Order) ToDomain() domain.Order {
	return domain.Order{
		OrderName:  o.OrderName,
		GrandTotal: o.GrandTotal,
		Status:     o.Status,
		CreatedAt:  o.CreatedAt,
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
	res.OrderName = data.OrderName
	res.PaymentLink = data.PaymentLink
	return res
}

func FromDomainItems(data domain.Items) Items {
	var res Items
	res.ID = data.ID
	res.OrderID = data.OrderID
	res.ProductID = data.ProductID
	res.ProductName = data.ProductName
	res.Subtotal = data.Subtotal
	res.Unit = data.Unit
	res.Qty = data.Qty
	return res
}

func (od *Order) ToDomainDetail() domain.Order {
	return domain.Order{
		OrderName:  od.OrderName,
		Status:     od.Status,
		GrandTotal: od.GrandTotal,
		CreatedAt:  od.CreatedAt,
	}
}

func ParseToArrDetail(arr []domain.Items, grandTotal int, orderName string) map[string]interface{} {
	var arrmap []map[string]interface{}
	var res2 = map[string]interface{}{}
	for i := 0; i < len(arr); i++ {
		var res = map[string]interface{}{}
		res["product_id"] = arr[i].ProductID
		res["product_name"] = arr[i].ProductName
		res["qty"] = arr[i].Qty
		res["subtotal"] = arr[i].Subtotal

		arrmap = append(arrmap, res)
	}
	res2["id_order"] = orderName
	res2["grand_total"] = grandTotal
	res2["items"] = arrmap
	return res2
}

func (od *Order) ToOD() domain.Order {
	return domain.Order{
		ID:         int(od.ID),
		OrderName:  od.OrderName,
		GrandTotal: od.GrandTotal,
		Status:     od.Status,
		UserID:     od.UserID,
	}
}

func ParseToArrConfirm(arr domain.Order) map[string]interface{} {
	var res = map[string]interface{}{}
	res["order_id"] = arr.OrderName
	res["grand_total"] = arr.GrandTotal
	res["status"] = arr.Status
	return res
}

func FromIP3(data []domain.Items) []Items {
	var res []Items
	for _, val := range data {
		newdata := Items{
			ID:          val.ID,
			OrderID:     val.OrderID,
			ProductID:   val.ProductID,
			ProductName: val.ProductName,
			Subtotal:    val.Subtotal,
			Qty:         val.Qty,
			Unit:        val.Unit,
		}
		res = append(res, newdata)
	}
	return res
}
