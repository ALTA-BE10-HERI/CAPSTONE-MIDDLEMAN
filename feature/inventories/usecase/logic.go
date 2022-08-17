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

func (iuc *inventoryUseCase) CreateUserDetailInventory(newRecap domain.Inventory, id int) (domain.Inventory, int) {

	outboundIDGenerate := strconv.FormatInt(time.Now().Unix(), 10)

	validError := iuc.validate.Var(newRecap, "gt=0")
	if validError != nil {
		log.Println("Validation error : ", validError)
		return domain.Inventory{}, 400
	}

	cekstok := iuc.inventoryData.CekStok(newRecap.InventoryProduct, id)
	if !cekstok {
		log.Println("insufficient amount")
		return domain.Inventory{}, 404
	}

	outbound := iuc.inventoryData.CreateUserInventoryData(newRecap, id, outboundIDGenerate)
	if outbound.ID == 0 {
		log.Println("error after creating data")
		return domain.Inventory{}, 500
	}

	create := iuc.inventoryData.CreateUserDetailInventoryData(newRecap.InventoryProduct, id, outboundIDGenerate, outbound.ID)
	if len(create) == 0 {
		log.Println("error after creating data")
		return domain.Inventory{}, 500
	}

	updatestok := iuc.inventoryData.RekapStock(newRecap.InventoryProduct, id, outboundIDGenerate)
	if !updatestok {
		log.Println("insufficient amount")
		return domain.Inventory{}, 404
	}

	delete := iuc.inventoryData.DeleteInOutBound(id)

	if delete == "cannot delete data" {
		log.Println("internal server error")
		return domain.Inventory{}, 500
	}
	if delete == "no data deleted" {
		log.Println("data not found")
		return domain.Inventory{}, 404
	}

	return outbound, 201
}

func (iuc *inventoryUseCase) ReadUserOutBoundDetail(id int, outboundIDGenerate string) ([]domain.InventoryProduct, int, string) {
	product := iuc.inventoryData.ReadUserOutBoundDetailData(id, outboundIDGenerate)
	if len(product) == 0 {
		log.Println("data not found")
		return []domain.InventoryProduct{}, 200, ""
	}

	return product, 200, outboundIDGenerate
}

func (iuc *inventoryUseCase) ReadUserOutBoundHistory(id int) ([]domain.Inventory, int) {
	product := iuc.inventoryData.ReadUserOutBoundHistoryData(id)
	if len(product) == 0 {
		log.Println("data not found")
		return []domain.Inventory{}, 200
	}

	return product, 200
}
