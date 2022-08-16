package usecase

import (
	"log"
	"middleman-capstone/domain"
	"strconv"
	"time"

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

func (iuc *inventoryUseCase) CreateUserInventory(newRecap []domain.InventoryProduct, id int) int {
	orderIDGenerate := strconv.FormatInt(time.Now().Unix(), 10)
	validError := iuc.validate.Var(newRecap, "gt=0")
	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	create := iuc.inventoryData.CreateUserInventoryData(newRecap, id, orderIDGenerate)

	if len(create) == 0 {
		log.Println("data not found")
		return 404
	}

	stok := iuc.inventoryData.CekStock(newRecap, id)
	if !stok {
		log.Println("insufficient amount")
		return 404
	}

	return 201
}
