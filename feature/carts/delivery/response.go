package delivery

import "middleman-capstone/domain"

type Cart struct {
	ID         int `json:"id"`
	Qty        int `json:"qty"`
	TotalPrice int `json:"totalprice"`
	UserID     int `json:"user_id"`
	Product    Product
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
		TotalPrice: data.TotalPrice,
		Product: Product{
			ID:           data.Product.ID,
			ProductName:  data.Product.ProductName,
			Price:        data.Product.Price,
			ProductImage: data.Product.ProductImage,
			Qty:          data.Product.Qty,
		},
	}
}
