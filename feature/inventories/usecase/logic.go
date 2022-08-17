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

func (iuc *inventoryUseCase) CreateUserDetailInventory(newRecap domain.Inventory, id int) int {

	outboundIDGenerate := strconv.FormatInt(time.Now().Unix(), 10)

	validError := iuc.validate.Var(newRecap, "gt=0")
	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	cekstok := iuc.inventoryData.CekStok(newRecap.InventoryProduct, id, outboundIDGenerate)
	if !cekstok {
		log.Println("insufficient amount")
		return 404
	}

	create := iuc.inventoryData.CreateUserDetailInventoryData(newRecap.InventoryProduct, id, outboundIDGenerate)
	if len(create) == 0 {
		log.Println("error after creating data")
		return 500
	}

	updatestok := iuc.inventoryData.RekapStock(newRecap.InventoryProduct, id, outboundIDGenerate)
	if !updatestok {
		log.Println("insufficient amount")
		return 404
	}

	delete := iuc.inventoryData.DeleteInOutBound(id)

	if delete == "cannot delete data" {
		log.Println("internal server error")
		return 500
	}
	if delete == "no data deleted" {
		log.Println("data not found")
		return 404
	}

	// outbound := iuc.inventoryData.CreateUserInventoryData(newRecap, id, outboundIDGenerate)
	// if outbound.ID == 0 {
	// 	log.Println("error after creating data")
	// 	return 500
	// }

	return 201
}

func (iuc *inventoryUseCase) ReadUserOutBoundDetail(id int, outboundIDGenerate string) ([]domain.InventoryProduct, int) {
	product := iuc.inventoryData.ReadUserOutBoundDetailData(id, outboundIDGenerate)
	if len(product) == 0 {
		log.Println("data not found")
		return []domain.InventoryProduct{}, 200
	}

	return product, 200
}
