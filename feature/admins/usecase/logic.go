package usecase

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/admins/data"

	"github.com/go-playground/validator"
)

type productUseCase struct {
	productData domain.ProductData
	validate    *validator.Validate
}

func New(pd domain.ProductData, v *validator.Validate) domain.ProductUseCase {
	return &productUseCase{
		productData: pd,
		validate:    v,
	}
}

func (pc *productUseCase) CreateProduct(newProduct domain.Product, idAdmin int) int {
	var product = data.FromModel(newProduct)
	validError := pc.validate.Struct(product)

	if validError != nil {
		log.Println("Validation Error : ", validError)
		return 400
	}
	product.ToModel()
	product.IdAdmin = idAdmin
	insert := pc.productData.CreateProductData(product.ToModel())

	if insert.ID == 0 {
		log.Println("error after creating data")
		return 500

	}
	return 201
}

func (pc *productUseCase) GetAllProduct(limit, offset int) (data []domain.Product, err error) {
	res, err := pc.productData.GetAllProductData(limit, offset)
	return res, err
}

func (pc *productUseCase) UpdateProduct(updatedData domain.Product, idProduct int) (row int, err error) {
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

	row, err = pc.productData.UpdateProductData(qry, idProduct)

	return row, err
}

func (pc *productUseCase) DeleteProduct(idProduct int) int {
	row, err := pc.productData.DeleteProductData(idProduct)

	if err != nil {
		log.Println("data not found")
		return 404
	}

	if row < 1 {
		log.Println("internal server error")
		return 500
	}
	return 204
}
