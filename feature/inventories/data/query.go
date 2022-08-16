package data

import (
	"log"
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type inventoryData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.InventoryData {
	return &inventoryData{
		db: db,
	}
}

// func (ind *inventoryData) CreateUserInventoryData(newRecap domain.InventoryProduct) domain.InventoryProduct {
func (ind *inventoryData) CreateUserInventoryData(newRecap []domain.InventoryProduct, id int, gen string) []domain.InventoryProduct {
	var product = FromIP2(newRecap, id)
	err := ind.db.Create(product)

	if err.Error != nil {
		log.Println("cannot create data", err.Error.Error())
		return []domain.InventoryProduct{}
	}

	if err.RowsAffected == 0 {
		log.Println("failed to insert data")
		return []domain.InventoryProduct{}
	}
	return ParseIPToArr(product)
}

func (ind *inventoryData) CekStock(newRecap []domain.InventoryProduct, id int) bool {
	var product = FromIP3(newRecap)
	something := []InventoryProduct{}
	for _, val := range product {
		res2 := ind.db.Model(&InventoryProduct{}).Select("inventory_products.id_user, inventory_products.id_product, product_users.name, inventory_products.qty, inventory_products.unit, product_users.stock").Joins("left join product_users on product_users.id = inventory_products.id_product").Where("id_product = ?", val.IdProduct).First(&something)
		if res2.Error != nil {
			log.Println("Cannot retrieve object", res2.Error.Error())
			return false
		}
		for _, value := range something {
			res3 := ind.db.Model(&domain.ProductUser{}).Where("id = ?", value.IdProduct)
			if res3.Error != nil {
				log.Println("Cannot retrieve object", res2.Error.Error())
				return false
			}
		}
	}

	return true
}
