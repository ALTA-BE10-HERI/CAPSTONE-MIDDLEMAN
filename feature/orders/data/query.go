package data

import (
	"errors"
	"fmt"
	"log"
	"middleman-capstone/domain"

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

func (od *orderData) InsertData(data []domain.Items, id int) []domain.Items {
	order := FromIP2(data, id)
	fmt.Println("order", order)
	result := od.db.Create(&order)

	if result.Error != nil {
		log.Println("cannot create data items", result.Error.Error())
		return []domain.Items{}
	}
	if result.RowsAffected < 1 {
		log.Println("failed to insert data items")
		return []domain.Items{}
	}
	return ParsePUToArr(order)
}

func (od *orderData) Insert(data domain.Order) (idOrder int, err error) {
	order2 := FromDomain(data)
	result := od.db.Create(&order2)

	if result.Error != nil {
		log.Println("cannot create data", result.Error.Error())
		return 0, err
	}

	if result.RowsAffected < 1 {
		log.Println("failed to insert data")
		return 0, err
	}
	return int(order2.ID), nil
}

func (od *orderData) GetUser(idUser int) (data domain.User, err error) {
	var user domain.User
	result := od.db.Where("id = ?", idUser).First(&user)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	if result.RowsAffected != 1 {
		return domain.User{}, fmt.Errorf("failed to get user")
	}
	return user, err
}

func (od *orderData) SelectDataAdminAll(limit, offset int) (data []domain.Order, err error) {
	var dataOrder []Order
	result := od.db.Limit(limit).Offset(offset).Find(&dataOrder)

	if result.Error != nil {
		return []domain.Order{}, result.Error
	}
	return ParseToArr(dataOrder), nil
}

func (od *orderData) SelectDataUserAll(limit, offset, idUser int) (data []domain.Order, err error) {
	var dataOrder []Order
	result := od.db.Where("user_id =?", idUser).Limit(limit).Offset(offset).Find(&dataOrder)

	if result.Error != nil {
		return []domain.Order{}, result.Error
	}
	return ParseToArr(dataOrder), nil
}

func (od *orderData) GetDetailData(idUser, idOrder int) (grandTotal int, err error) {
	dataOrder := Order{}
	result := od.db.Where("user_id =? AND id = ?", idUser, idOrder).Preload("Items").First(&dataOrder)

	if result.Error != nil {
		return 0, result.Error
	}
	return dataOrder.GrandTotal, nil
}

func (od *orderData) GetDetailItems(idOrder int) (data []domain.Items, err error) {
	var dataItem []Items
	result := od.db.Where("order_id = ?", idOrder).Find(&dataItem)

	if result.Error != nil {
		return []domain.Items{}, result.Error
	}

	return ParsePUToArr(dataItem), nil
}

func (od *orderData) AcceptPaymentData(data domain.PaymentWeb) (row int, err error) {
	var dataPayment PaymentWeb

	if data.TransactionStatus != "settlement" {
		return -1, err
	}

	updateOrder := od.db.Table("Order").Where("order_name = ?", dataPayment.OrderName).Update("status", "on process")

	if updateOrder.Error != nil {
		return 0, nil
	}

	if updateOrder.RowsAffected != 1 {
		return 0, errors.New("order not found")
	}

	return row, nil
}
