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

	updateOrder := od.db.Table("orders").Where("order_name = ?", dataPayment.OrderName).Update("status", "waiting confirmation")

	if updateOrder.Error != nil {
		return 0, nil
	}

	if updateOrder.RowsAffected != 1 {
		return 0, errors.New("order not found")
	}

	return row, nil
}

func (od *orderData) ConfirmOrderData(ordername string, userid int) domain.Order {

	order := Order{}
	updatestatus := od.db.Table("orders").Where("order_name = ?", ordername).Update("status", "on process")

	if updatestatus.Error != nil {
		return domain.Order{}
	}

	if updatestatus.RowsAffected < 1 {
		return domain.Order{}
	}

	statusorder := od.db.Table("orders").Where("order_name = ?", ordername).First(&order)

	if statusorder.Error != nil {
		return domain.Order{}
	}

	if statusorder.RowsAffected < 1 {
		return domain.Order{}
	}

	return order.ToOD()
}

func (od *orderData) DoneOrderData(ordername string, userid int) domain.Order {

	order := Order{}
	updatestatus := od.db.Table("orders").Where("order_name = ?", ordername).Update("status", "delivered")
	if updatestatus.Error != nil {
		return domain.Order{}
	}
	if updatestatus.RowsAffected < 1 {
		return domain.Order{}
	}

	statusorder := od.db.Table("orders").Where("order_name = ?", ordername).First(&order)
	if statusorder.Error != nil {
		return domain.Order{}
	}
	if statusorder.RowsAffected < 1 {
		return domain.Order{}
	}

	return order.ToOD()
}

func (od *orderData) UpdateStokAdmin(ordername string) bool {
	var order = Order{}
	var item = []Items{}

	res0 := od.db.Model(&Order{}).Where("order_name = ?", ordername).Find(&order)
	if res0.Error != nil {
		log.Println("Cannot retrieve object", res0.Error.Error())
		return false
	}

	res := od.db.Model(&Items{}).Where("order_id = ?", order.ID).Find(&item)
	if res.Error != nil {
		log.Println("Cannot retrieve object", res.Error.Error())
		return false
	}

	for _, val := range item {
		var productadmin = domain.Product{}
		res2 := od.db.Table("products").Where("products.id = ?", val.ProductID).First(&productadmin)
		if res2.Error != nil {
			log.Println("Cannot retrieve object", res2.Error.Error())
			return false
		}

		res3 := od.db.Table("products").Where("products.id = ?", val.ProductID).Update("stock", productadmin.Stock-val.Qty)
		if res3.Error != nil {
			log.Println("Cannot retrieve object", res2.Error.Error())
			return false
		}
	}
	return true
}

func (od *orderData) CekOwnedUser(ordername string, userid int) bool {
	var order = Order{}
	var item = []Items{}

	res0 := od.db.Model(&Order{}).Where("order_name = ?", ordername).Find(&order)
	if res0.Error != nil {
		log.Println("Cannot retrieve object", res0.Error.Error())
		return false
	}

	res := od.db.Model(&Items{}).Where("order_id = ?", order.ID).Find(&item)
	if res.Error != nil {
		log.Println("Cannot retrieve object", res.Error.Error())
		return false
	}

	for _, val := range item {
		var productadmin = domain.ProductUser{}
		var productitems = Items{}
		res7 := od.db.Model(&Items{}).Where("product_id = ?", val.ProductID).First(&productitems)
		if res7.Error != nil {
			log.Println("Cannot retrieve object", res7.Error.Error())
			return false
		}
		fmt.Println("productitem", productitems)

		res2 := od.db.Table("product_users").Where("reff = ? AND id_user = ?", val.ProductID, userid).First(&productadmin)
		fmt.Println("productadmin", productadmin)
		if res2.Error != nil {
			var itemuser = domain.Items{}
			var productadmin2 = domain.ProductUser{}
			res3 := od.db.Where("product_id", val.ProductID).First(&itemuser)
			if res3.Error != nil {
				log.Println("Cannot retrieve object", res3.Error.Error())
				return false
			}
			fmt.Println("productadmin2", productadmin2)
			res4 := od.db.Create(&productadmin2)
			if res4.Error != nil {
				log.Println("Cannot retrieve object", res4.Error.Error())
				return false
			}
			res5 := od.db.Model(&domain.ProductUser{}).Where("name = ?", val.ProductName).Updates(domain.ProductUser{Reff: val.ProductID, IdUser: userid})
			if res5.Error != nil {
				log.Println("Cannot retrieve object", res5.Error.Error())
				return false
			}
		} else {
			res6 := od.db.Model(&domain.ProductUser{}).Where("reff = ? AND id_user = ?", val.ProductID, userid).Update("stock", productitems.Qty+productadmin.Stock)
			if res6.Error != nil {
				log.Println("Cannot retrieve object", res6.Error.Error())
				return false
			}
		}
	}
	return true
}
