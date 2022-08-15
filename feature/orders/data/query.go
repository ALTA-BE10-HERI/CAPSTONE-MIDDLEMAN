package data

import (
	"errors"
	"middleman-capstone/domain"
	dataC "middleman-capstone/feature/carts/data"

	"gorm.io/gorm"
)

type orderData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OrderData {
	return &orderData{
		db: db,
	}
}

func (od *orderData) InsertData(data domain.Order) (row int, err error) {
	order := FromDomain(data)
	result := od.db.Create(&order)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to create data order")
	}
	return int(result.RowsAffected), nil
}

func (od *orderData) GrandTotal(idCart int) (grandTotal int, err error) {
	grandTotalCart := []dataC.Cart{}
	result := od.db.Preload("Product").Find(&grandTotalCart, idCart)

	if result.Error != nil {
		return -1, result.Error
	}
	for _, v := range grandTotalCart {
		grandTotal += (v.Qty * v.Product.Price)
	}

	return grandTotal, nil
}

func (od *orderData) SelectDataAdminAll(limit, offset int) (data []domain.Order, err error) {
	dataOrder := []Order{}
	result := od.db.Find(&dataOrder)

	if result.Error != nil {
		return []domain.Order{}, result.Error
	}
	return ParseToArr(dataOrder), nil
}
