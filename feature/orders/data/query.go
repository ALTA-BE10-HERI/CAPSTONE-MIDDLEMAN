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

func (od *orderData) GetDetailData(orderName string) (grandTotal, idOrder int, status string, err error) {
	dataOrder := Order{}
	result := od.db.Where("order_name = ?", orderName).Preload("Items").First(&dataOrder)

	if result.Error != nil {
		return 0, 0, "", result.Error
	}
	return dataOrder.GrandTotal, int(dataOrder.ID), dataOrder.Status, nil
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

	updateOrder := od.db.Table("orders").Where("order_name = ?", data.OrderName).Update("status", "waiting confirmation")

	if updateOrder.Error != nil {
		return 0, nil
	}

	if updateOrder.RowsAffected != 1 {
		return 0, errors.New("order not found")
	}

	return row, nil
}

func (od *orderData) CancelPaymentData(data domain.PaymentWeb) (row int, err error) {

	updateOrder := od.db.Table("orders").Where("order_name = ?", data.OrderName).Update("status", "canceled")

	if updateOrder.Error != nil {
		return 0, nil
	}

	if updateOrder.RowsAffected != 1 {
		return 0, errors.New("order not found")
	}

	return row, nil
}

func (od *orderData) GetIncomingData(limit, offset int) (data []domain.Order, err error) {
	var dataOrder []Order
	result := od.db.Limit(limit).Offset(offset).Where("status = ?", "waiting confirmation").Find(&dataOrder)
	if result.Error != nil {
		log.Println("error get data")
		return []domain.Order{}, nil
	}
	return ParseToArr(dataOrder), nil
}

func (od *orderData) ConfirmOrderData(orderName string) domain.Order {
	var dataOrder domain.Order
	updatestatus := od.db.Table("orders").Where("order_name = ? AND status = ?", orderName, "waiting confirmation").Update("status", "on process")
	if updatestatus.Error != nil {
		log.Println("error update data")
		return domain.Order{}
	}
	if updatestatus.RowsAffected < 1 {
		log.Println("status not updated")
		return domain.Order{}
	}

	statusorder := od.db.Table("orders").Where("order_name = ?", orderName).First(&dataOrder)
	if statusorder.Error != nil {
		return domain.Order{}
	}
	if statusorder.RowsAffected < 1 {
		return domain.Order{}
	}

	return dataOrder
}

func (od *orderData) GetUserByOrderName(orderName string) (data domain.Order, err error) {
	var order domain.Order
	result := od.db.Where("order_name = ?", orderName).First(&order)

	if result.Error != nil {
		return domain.Order{}, result.Error
	}

	if result.RowsAffected != 1 {
		return domain.Order{}, fmt.Errorf("failed to get order")
	}
	return order, err
}

func (od *orderData) DoneOrderData(ordername string) domain.Order {

	order := Order{}
	updatestatus := od.db.Table("orders").Where("order_name = ? AND status = ? ", ordername, "on process").Update("status", "delivered")
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

func (od *orderData) CekUser(ordername string) (products []domain.Items, id int) {
	var order = Order{}
	var item = []Items{}

	res0 := od.db.Model(&Order{}).Where("order_name = ?", ordername).Find(&order)
	if res0.Error != nil {
		log.Println("Cannot retrieve object", res0.Error.Error())
		return []domain.Items{}, 0
	}

	res := od.db.Model(&Items{}).Where("order_id = ?", order.ID).Find(&item)
	if res.Error != nil {
		log.Println("Cannot retrieve object", res.Error.Error())
		return []domain.Items{}, 0
	}
	return ParsePUToArr(item), order.UserID
}

func (od *orderData) CekOwned(product domain.Items, userid int) bool {
	item := FromDomainItems(product)
	productuser := domain.ProductUser{}
	res := od.db.Table("product_users").Where("reff = ? AND id_user = ?", item.ProductID, userid).First(&productuser)
	if res.Error != nil {
		log.Println("Cannot retrieve object", res.Error.Error())
		return false
	}
	return true
}

func (od *orderData) CreateNewProduct(product domain.Items, userid int) bool {
	item := FromDomainItems(product)
	productuser := domain.ProductUser{}
	productuser2 := domain.ProductUser{}
	productadmin := domain.Product{}
	res0 := od.db.Model(&domain.Product{}).Where("id = ?", item.ProductID).First(&productadmin)
	if res0.Error != nil {
		log.Println("Cannot retrieve object", res0.Error.Error())
		return false
	}
	res := od.db.Create(&productuser)
	if res.Error != nil {
		log.Println("Cannot retrieve object", res.Error.Error())
		return false
	}
	res2 := od.db.Model(&domain.ProductUser{}).Where("id = ?", productuser.ID).Updates(domain.ProductUser{IdUser: userid, Name: item.ProductName, Stock: item.Qty * productadmin.Satuan, Reff: item.ProductID})
	if res2.Error != nil {
		log.Println("Cannot retrieve object", res2.Error.Error())
		return false
	}
	res3 := od.db.Model(&domain.ProductUser{}).Select("products.name, products.image").Joins("left join products on product_users.reff = products.id").Where("product_users.reff = ?", item.ProductID).First(&productuser2)
	if res3.Error != nil {
		log.Println("Cannot retrieve object", res3.Error.Error())
		return false
	}
	res4 := od.db.Model(&domain.ProductUser{}).Where("product_users.reff = ?", item.ProductID).Updates(&productuser2)
	if res4.Error != nil {
		log.Println("Cannot retrieve object", res4.Error.Error())
		return false
	}
	return true
}

func (od *orderData) UpdateNewProduct(product domain.Items, userid int) bool {
	item := FromDomainItems(product)
	productuser := domain.ProductUser{}
	productadmin := domain.Product{}
	res0 := od.db.Model(&domain.Product{}).Where("id = ?", item.ProductID).First(&productadmin)
	if res0.Error != nil {
		log.Println("Cannot retrieve object", res0.Error.Error())
		return false
	}
	res := od.db.Model(&domain.ProductUser{}).Where("reff = ? AND id_user = ?", item.ProductID, userid).First(&productuser)
	if res.Error != nil {
		log.Println("Cannot retrieve object", res.Error.Error())
		return false
	}
	res2 := od.db.Model(&domain.ProductUser{}).Where("reff = ? AND id_user = ?", item.ProductID, userid).Update("stock", item.Qty*productadmin.Satuan+productuser.Stock)
	if res.Error != nil {
		log.Println("Cannot retrieve object", res2.Error.Error())
		return false
	}
	return true

}

func (od *orderData) DeleteCart(userid int) bool {
	res := od.db.Where("user_id = ?", userid).Delete(&domain.Cart{})

	if res.Error != nil {
		log.Println("cannot delete data")
		return false
	}
	if res.RowsAffected < 1 {
		log.Println("no data deleted")
		return false
	}
	return true
}
