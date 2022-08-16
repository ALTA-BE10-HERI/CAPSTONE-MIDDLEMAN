package usecase

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/productusers/data"

	"github.com/go-playground/validator"
)

type productUserUseCase struct {
	productUserData domain.ProductUserData
	validate        *validator.Validate
}

func New(puc domain.ProductUserData, v *validator.Validate) domain.ProductUserUseCase {
	return &productUserUseCase{
		productUserData: puc,
		validate:        v,
	}
}

func (puuc *productUserUseCase) CreateProduct(newProduct domain.ProductUser, id int) (domain.ProductUser, int) {
	var product = data.FromPU(newProduct)
	validError := puuc.validate.Struct(product)

	if validError != nil {
		log.Println("Validation error : ", validError)
		return domain.ProductUser{}, 400
	}

	product.IdUser = id
	create := puuc.productUserData.CreateProductData(product.ToPU())

	if create.ID == 0 {
		log.Println("error after creating data")
		return domain.ProductUser{}, 500
	}
	return create, 201
}

func (puuc *productUserUseCase) ReadAllProduct(id int) ([]domain.ProductUser, int) {
	product := puuc.productUserData.ReadAllProductData(id)
	if len(product) == 0 {
		log.Println("data not found")
		return nil, 404
	}

	return product, 200
}

func (puuc *productUserUseCase) UpdateProduct(updatedData domain.ProductUser, productid, id int) (domain.ProductUser, int) {

	if updatedData.Name == "" && updatedData.Unit == "" && updatedData.Stock == 0 && updatedData.Price == 0 && updatedData.Image == "" {
		return domain.ProductUser{}, 400
	}
	qry := map[string]interface{}{}

	if updatedData.Name != "" {
		qry["name"] = updatedData.Name
	}

	if updatedData.Unit != "" {
		qry["unit"] = updatedData.Unit
	}

	if updatedData.Stock != 0 {
		qry["stock"] = updatedData.Stock
	}

	if updatedData.Price != 0 {
		qry["price"] = updatedData.Price
	}

	if updatedData.Image != "" {
		qry["image"] = updatedData.Image
	}

	products := puuc.productUserData.UpdateProductData(qry, productid, id)

	if products.ID == 0 {
		log.Println("Empty Data")
		return domain.ProductUser{}, 404
	}
	return products, 200
}

func (puuc *productUserUseCase) DeleteProduct(productid, id int) int {
	err := puuc.productUserData.DeleteProductData(productid, id)
	if err == "cannot delete data" {
		log.Println("internal server error")
		return 500
	}
	if err == "no data deleted" {
		log.Println("data not found")
		return 404
	}
	return 204
}

func (puuc *productUserUseCase) SearchRestoBusiness(search string) (result []domain.ProductUser, err error) {
	response, err := puuc.productUserData.SearchRestoData(search)

	return response, err
}
