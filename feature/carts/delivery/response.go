package delivery

import "middleman-capstone/domain"

type ResponseCart struct {
	ID           int    `json:"id"`
	ProductName  string `json:"product_name"`
	Price        int    `json:"price"`
	Unit         string `json:"unit"`
	ProductImage string `json:"product_image"`
	Qty          int    `json:"qty"`
	Subtotal     int    `json:"subtotal"`
}

func FromModel(data domain.Cart) ResponseCart {
	return ResponseCart{
		ID:           data.Product.ID,
		ProductName:  data.Product.ProductName,
		ProductImage: data.Product.ProductImage,
		Price:        data.Product.Price,
		Unit:         data.Product.Unit,
		Qty:          data.Qty,
		Subtotal:     data.Subtotal,
	}
}

func FromModelList(data []domain.Cart) []ResponseCart {
	result := []ResponseCart{}
	for key := range data {
		result = append(result, FromModel(data[key]))
	}
	return result
}
