package data

import (
	"errors"
	"fmt"
	"log"
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductData {
	return &productData{
		db: db,
	}
}

func (pd *productData) CreateProductData(newProduct domain.Product) domain.Product {
	var product = FromModel(newProduct)
	err := pd.db.Create(&product)

	if err.Error != nil {
		log.Println("cannot create data", err.Error.Error())
		return domain.Product{}
	}

	if err.RowsAffected == 0 {
		log.Println("failed to insert data")
		return domain.Product{}
	}
	return product.ToModel()
}

func (pd *productData) GetAllProductData(limit, offset int) (data []domain.Product, err error) {
	dataProduct := []Product{}
	res := pd.db.Model(&Product{}).Find(&dataProduct)
	if res.Error != nil {
		return []domain.Product{}, nil
	}
	fmt.Println(ParseProductToArr(dataProduct))
	return ParseProductToArr(dataProduct), nil
}

func (pd *productData) UpdateProductData(updatedData domain.Product) domain.Product {
	var product = FromModel(updatedData)
	result := pd.db.Model(&Product{}).Where("id = ?", product.ID).Updates(product)
	if result.Error != nil {
		log.Println("cannot update data", result.Error.Error())
		return domain.Product{}
	}

	if result.RowsAffected == 0 {
		log.Println("data not found")
		return domain.Product{}
	}

	return product.ToModel()
}

func (pd *productData) DeleteProductData(idProduct int) (row int, err error) {
	result := pd.db.Where("id = ?", idProduct).Delete(&Product{})

	if result.Error != nil {
		log.Println("cannot delete data", result.Error.Error())
		return 0, result.Error
	}
	if result.RowsAffected < 1 {
		log.Println("no data deleted", result.Error.Error())
		return 0, errors.New("failed to delete data")
	}
	return int(result.RowsAffected), nil
}
