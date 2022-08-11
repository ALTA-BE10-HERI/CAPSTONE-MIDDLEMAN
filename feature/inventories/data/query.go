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

func (ind *inventoryData) CreateInventoryData(newRecap domain.InventoryProduct) domain.InventoryProduct {
	var product = FromIP(newRecap)
	err := ind.db.Create(&product)

	if err.Error != nil {
		log.Println("cannot create data", err.Error.Error())
		return domain.InventoryProduct{}
	}

	if err.RowsAffected == 0 {
		log.Println("failed to insert data")
		return domain.InventoryProduct{}
	}
	return product.ToIP()
}

func (ind *inventoryData) StockUpdate(newRecap domain.InventoryProduct) bool {
	var proder domain.ProductUser
	res2 := ind.db.Where("name = ?", newRecap.Name).First(&proder)

	if res2.Error != nil {
		log.Println("Cannot retrieve object", res2.Error.Error())
		return false
	}

	updatestock := proder.Stock - newRecap.Qty

	if updatestock > 0 {
		res3 := ind.db.Model(domain.ProductUser{}).Where("name = ?", newRecap.Name).Updates(domain.ProductUser{Stock: updatestock})
		if res3.Error != nil {
			log.Println("Cannot retrieve object", res3.Error.Error())
			return false
		}
	} else {
		log.Println("not enough stock")
		return false
	}

	return true
}
