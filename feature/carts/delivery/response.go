package delivery

import "middleman-capstone/domain"

type Cart struct {
	ID       int     `json:"id"`
	Qty      int     `json:"qty"`
	Subtotal int     `json:"subtotal"`
	UserID   int     `json:"user_id"`
	Product  Product `json:"product"`
}

type Product struct {
	ID           int    `json:"id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	Price        int    `json:"price"`
	Unit         string `json:"unit"`
	Qty          int    `json:"qty"`
}

func FromModel(data domain.Cart) Cart {
	return Cart{
		ID:       data.ID,
		Qty:      data.Qty,
		Subtotal: data.Subtotal,
		UserID:   data.UserID,
		Product: Product{
			ID:           data.Product.ID,
			ProductName:  data.Product.ProductName,
			ProductImage: data.Product.ProductImage,
			Price:        data.Product.Price,
			Unit:         data.Product.Unit,
			Qty:          data.Product.Stock,
		},
	}
}
func FromModelList(data []domain.Cart) []Cart {
	result := []Cart{}
	for key := range data {
		result = append(result, FromModel(data[key]))
	}
	return result
}
