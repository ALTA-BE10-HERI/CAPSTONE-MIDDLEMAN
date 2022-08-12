package delivery

import "middleman-capstone/domain"

type Cart struct {
	ID       int `json:"id"`
	Qty      int `json:"qty"`
	Subtotal int `json:"subtotal"`
	UserID   int `json:"user_id"`
	Product  Product
}

type Product struct {
	ID           int    `json:"id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	Price        int    `json:"price"`
	Unit         int    `json:"unit"`
	Qty          int    `json:"qty"`
}

func FromModel(data domain.Cart) Cart {
	return Cart{
		Subtotal: data.Subtotal,
		Product: Product{
			ID:           data.Product.ID,
			ProductName:  data.Product.ProductName,
			Price:        data.Product.Price,
			ProductImage: data.Product.ProductImage,
			Qty:          data.Product.Qty,
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
