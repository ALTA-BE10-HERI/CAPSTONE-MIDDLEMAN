package delivery

import (
	"middleman-capstone/domain"
	"time"
)

type FormatOrder struct {
	GrandTotal int `json:"grand_total" form:"grand_total"`
	Items      []Items
}

type PaymentWeb struct {
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	OrderName         string `json:"order_id" form:"order_id"`
}

type Items struct {
	ProductID int    `json:"product_id" form:"product_id"`
	Qty       int    `json:"qty" form:"qty"`
	Unit      string `json:"unit" form:"unit"`
	Subtotal  int    `json:"subtotal" form:"subtotal"`
}

type Order struct {
	ID         int
	GrandTotal int
	Status     string
	CreatedAt  time.Time
}

func FromWeb(web PaymentWeb) domain.PaymentWeb {
	return domain.PaymentWeb{
		TransactionStatus: web.TransactionStatus,
		OrderName:         web.OrderName,
	}
}

func ParseIFToArr(arr []Items) []domain.Items {
	var res []domain.Items
	for _, val := range arr {
		res = append(res, val.ToIF())
	}
	return res
}

func (pf *Items) ToIF() domain.Items {
	return domain.Items{
		ProductID: pf.ProductID,
		Qty:       pf.Qty,
		Unit:      pf.Unit,
		Subtotal:  pf.Subtotal,
	}
}

func ToDomain(format FormatOrder) domain.Order {
	return domain.Order{
		GrandTotal: format.GrandTotal,
		Items:      ParseIFToArr(format.Items),
	}
}

func (o *Order) ToDomain() domain.Order {
	return domain.Order{
		ID:         int(o.ID),
		Status:     o.Status,
		GrandTotal: o.GrandTotal,
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

func ParsePUToArr2(arr []domain.Order) []map[string]interface{} {
	var arrmap []map[string]interface{}
	for i := 0; i < len(arr); i++ {
		var res = map[string]interface{}{}
		res["id"] = arr[i].ID
		res["status"] = arr[i].Status
		res["date"] = arr[i].CreatedAt
		res["grand_total"] = arr[i].GrandTotal

		arrmap = append(arrmap, res)
	}
	return arrmap
}
