package usecase

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/inventories/data"

	"github.com/go-playground/validator"
)

type inventoryUseCase struct {
	inventoryData domain.InventoryData
	validate      *validator.Validate
}

func New(ind domain.InventoryData, v *validator.Validate) domain.InventoryUseCase {
	return &inventoryUseCase{
		inventoryData: ind,
		validate:      v,
	}
}

func (iuc *inventoryUseCase) CreateInventory(newRecap domain.InventoryProduct, id int) int {
	var product = data.FromIP(newRecap)
	validError := iuc.validate.Struct(product)

	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	product.IdUser = id
	create := iuc.inventoryData.CreateInventoryData(product.ToIP())

	if create.ID == 0 {
		log.Println("error after creating data")
		return 500
	}
	return 201
}
