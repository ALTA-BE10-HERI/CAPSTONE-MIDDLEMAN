package mysql

import (
	"fmt"
	"log"
	"middleman-capstone/config"

	adminData "middleman-capstone/feature/admins/data"

	cartData "middleman-capstone/feature/carts/data"

	inoutboundData "middleman-capstone/feature/inoutbounds/data"

	inventoryData "middleman-capstone/feature/inventories/data"
	productuserData "middleman-capstone/feature/productusers/data"
	userData "middleman-capstone/feature/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.Username,
		cfg.Password,
		cfg.Address,
		cfg.Port,
		cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateData(db *gorm.DB) {

	db.AutoMigrate(userData.User{})
	db.AutoMigrate(productuserData.ProductUser{})
	db.AutoMigrate(adminData.Product{})
	db.AutoMigrate(inventoryData.InventoryProduct{})
	db.AutoMigrate(inoutboundData.InOutBounds{})
	db.AutoMigrate(cartData.Cart{})
	db.AutoMigrate(inventoryData.Inventory{})

}
