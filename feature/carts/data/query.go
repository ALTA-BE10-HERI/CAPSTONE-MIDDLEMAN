package data

import (
	"errors"
	"log"
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ChartData {
	return &cartData{
		db: db,
	}
}

func (cd *cartData) SelectData(limit, offset, idFromToken int) (data []domain.Cart, err error) {
	dataCart := []Cart{}
	result := cd.db.Preload("Product").Where("user_id=? AND status = ?", idFromToken, "Pending").Find(&dataCart)
	if result.Error != nil {
		return []domain.Cart{}, result.Error
	}
	return ParseToArr(dataCart), nil
}

func (cd *cartData) InsertData(data domain.Cart) (row int, err error) {
	cart := FromDomain(data)
	result := cd.db.Create(&cart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to create data cart")
	}
	return int(result.RowsAffected), nil
}

func (cd *cartData) CheckCart(idProd, idFromToken int) (isExist bool, idCart, qty int, err error) {
	dataCart := Cart{}
	resultCheck := cd.db.Model(&Cart{}).Where("product_id = ? AND user_id = ? AND status = ?", idProd, idFromToken, "Pending").First(&dataCart)
	if resultCheck.Error != nil {
		return false, 0, 0, resultCheck.Error
	}
	return true, int(dataCart.ID), int(dataCart.Qty), nil
}

func (cd *cartData) UpdateDataDB(stock, idCart, idFromToken int) (row int, err error) {
	dataCart := Cart{}
	idCheck := cd.db.First(&dataCart, idCart)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if dataCart.UserID != idFromToken {
		log.Println("cek ", dataCart.UserID)
		return -1, errors.New("you don't have access")
	}
	result := cd.db.Model(&Cart{}).Where("id = ?", idCart).Update("stock", stock)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update data")
	}
	return int(result.RowsAffected), nil
}
