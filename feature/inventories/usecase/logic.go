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

func (iuc *inventoryUseCase) CreateUserInventory(newRecap domain.Inventory, id int) (domain.Inventory, int) {

	outboundIDGenerate := strconv.FormatInt(time.Now().Unix(), 10)

	validError := iuc.validate.Var(newRecap, "gt=0")
	if validError != nil {
		log.Println("Validation error : ", validError)
		return domain.Inventory{}, 400
	}

	delete := iuc.inventoryData.DeleteInOutBound(id)
	if delete == "no data deleted" {
		log.Println("data not found")
		return domain.Inventory{}, 404
	}

	cekstok := iuc.inventoryData.CekStok(newRecap.InventoryProduct, id)
	if !cekstok {
		log.Println("insufficient amount")
		return domain.Inventory{}, 400
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
		return domain.Inventory{}, 400
	}

	return outbound, 201
}

func (iuc *inventoryUseCase) ReadUserOutBoundDetail(id int, outboundIDGenerate string) ([]domain.InventoryProduct, int, string) {
	product := iuc.inventoryData.ReadUserOutBoundDetailData(id, outboundIDGenerate)
	if len(product) == 0 {
		log.Println("data not found")
		return []domain.InventoryProduct{}, 404, ""
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

func (iuc *inventoryUseCase) CreateAdminInventory(newRecap domain.Inventory, id int, role string) (domain.Inventory, int) {

	inboundIDGenerate := strconv.FormatInt(time.Now().Unix(), 10)

	validError := iuc.validate.Var(newRecap, "gt=0")
	if validError != nil {
		log.Println("Validation error : ", validError)
		return domain.Inventory{}, 400
	}

	delete := iuc.inventoryData.DeleteAdminInOutBound()
	if delete == "no data deleted" {
		log.Println("data not found")
		return domain.Inventory{}, 404
	}

	inbound := iuc.inventoryData.CreateAdminInventoryData(newRecap, id, inboundIDGenerate)
	if inbound.ID == 0 {
		log.Println("error after creating data")
		return domain.Inventory{}, 500
	}

	create := iuc.inventoryData.CreateAdminDetailInventoryData(newRecap.InventoryProduct, id, inboundIDGenerate, inbound.ID, role)
	if len(create) == 0 {
		log.Println("error after creating data")
		return domain.Inventory{}, 500
	}

	updatestok := iuc.inventoryData.RekapAdminStock(newRecap.InventoryProduct, id, inboundIDGenerate)
	if !updatestok {
		log.Println("insufficient amount")
		return domain.Inventory{}, 400
	}

	return inbound, 201
}

func (iuc *inventoryUseCase) ReadAdminOutBoundHistory() ([]domain.Inventory, int) {
	product := iuc.inventoryData.ReadAdminOutBoundHistoryData()
	if len(product) == 0 {
		log.Println("data not found")
		return []domain.Inventory{}, 200
	}

	return product, 200
}

func (iuc *inventoryUseCase) ReadAdminOutBoundDetail(inboundIDGenerate string) ([]domain.InventoryProduct, int, string) {
	product := iuc.inventoryData.ReadAdminOutBoundDetailData(inboundIDGenerate)
	if len(product) == 0 {
		log.Println("data not found")
		return []domain.InventoryProduct{}, 404, ""
	}

	return product, 200, inboundIDGenerate
}
