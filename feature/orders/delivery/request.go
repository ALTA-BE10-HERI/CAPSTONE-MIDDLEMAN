package delivery

type FormatOrder struct {
	GrandTotal int
	Items      Items
}

type Order struct {
	ID         int
	GrandTotal int
	Items      Items
}

type Items struct {
	ID       int    `json:"id"`
	Subtotal int    `json:"subtotal"`
	Unit     string `json:"unit"`
	Qty      int    `json:"qty"`
}

// func FromModel(data domain.Order) Order {
// 	return Order{
// 		GrandTotal: data.GrandTotal,
// 		Items: Items{
// 			ID:       data.Product.ID,
// 			Qty:      data.Cart.Qty,
// 			Unit:     data.Cart.Product.Unit,
// 			Subtotal: data.Cart.Subtotal,
// 		},
// 	}
// }

// func FromModelList(data []domain.Order) []Order {
// 	result := []Order{}
// 	for key := range data {
// 		result = append(result, FromModel(data[key]))
// 	}
// 	return result
// }
