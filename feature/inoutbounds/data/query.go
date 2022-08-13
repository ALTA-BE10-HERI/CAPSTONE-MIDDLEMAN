package data

import (
	"log"
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type inoutboundData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.InOutBoundData {
	return &inoutboundData{
		db: db,
	}
}

func (iobd *inoutboundData) AddEntryData(newProduct domain.InOutBounds) domain.InOutBounds {
	var cart = FromIOB(newProduct)

	err := iobd.db.Create(&cart)

	if err.Error != nil {
		log.Println("cannot create data", err.Error.Error())
		return domain.InOutBounds{}
	}

	if err.RowsAffected == 0 {
		log.Println("failed to insert data")
		return domain.InOutBounds{}
	}
	return cart.ToIOB()
}

func (iobd *inoutboundData) CekUserEntry(newProduct domain.InOutBounds) (cek bool, idcart, qty int) {
	var cart InOutBounds
	err := iobd.db.Model(&InOutBounds{}).Where("id_product = ? AND id_user = ?", newProduct.IdProduct, newProduct.IdUser).First(&cart)

	if err.Error != nil {
		return false, 0, 0
	}

	return true, int(cart.ID), int(cart.Qty)
}

func (iobd *inoutboundData) CekOwnerEntry(newProduct domain.InOutBounds) (cek bool) {
	var cart domain.ProductUser
	err := iobd.db.Model(&domain.ProductUser{}).Where("id = ? AND id_user = ?", newProduct.IdProduct, newProduct.IdUser).First(&cart)

	if err.Error != nil {
		return false
	}

	if err.RowsAffected == 0 {
		log.Println("failed read data")
		return false
	}

	return true
}

func (iobd *inoutboundData) CekAdminEntry(newProduct domain.InOutBounds) (cek bool, idcart, qty int) {
	var cart InOutBounds
	err := iobd.db.Model(&InOutBounds{}).Where("id_product = ? AND role = ?", newProduct.IdProduct, "admin").First(&cart)

	if err.Error != nil {
		return false, 0, 0
	}

	return true, int(cart.ID), int(cart.Qty)
}

func (iobd *inoutboundData) UpdateQty(idcart, qty int) domain.InOutBounds {
	var cart InOutBounds
	res := iobd.db.Model(&InOutBounds{}).Where("id = ?", idcart).Update("qty", qty)
	res2 := iobd.db.Model(&InOutBounds{}).Where("id = ?", idcart).Find(&cart)

	if res.Error != nil {
		return domain.InOutBounds{}
	}

	if res.RowsAffected != 1 {
		log.Println("failed to insert data")
		return domain.InOutBounds{}
	}

	if res2.Error != nil {
		return domain.InOutBounds{}
	}

	if res2.RowsAffected != 1 {
		log.Println("failed to insert data")
		return domain.InOutBounds{}
	}

	return cart.ToIOB()
}

func (ioub *inoutboundData) UpdateEntryAdminData(productid int) domain.InOutBounds {
	var cart InOutBounds
	res0 := ioub.db.Model(&InOutBounds{}).Select("products.id, products.name, products.unit, in_out_bounds.qty").Joins("left join products on products.id = in_out_bounds.id_product").Where("in_out_bounds.id_product = ?", productid).First(&cart)

	if res0.Error != nil {
		log.Println("cannot read data", res0.Error.Error())
		return domain.InOutBounds{}
	}

	if res0.RowsAffected == 0 {
		log.Println("failed read data")
		return domain.InOutBounds{}
	}

	res := ioub.db.Model(&InOutBounds{}).Where("id_product = ? AND role = ?", productid, "admin").Updates(cart)

	if res.Error != nil {
		log.Println("cannot update data", res.Error.Error())
		return domain.InOutBounds{}
	}

	if res.RowsAffected == 0 {
		log.Println("data not found")
		return domain.InOutBounds{}
	}

	return cart.ToIOB()
}

func (ioub *inoutboundData) UpdateEntryUserData(productid int, id int) domain.InOutBounds {
	var cart InOutBounds
	res0 := ioub.db.Model(&InOutBounds{}).Select("product_users.id, product_users.name, product_users.unit, in_out_bounds.qty").Joins("left join product_users on product_users.id = in_out_bounds.id_product").Where("in_out_bounds.id_product = ? AND in_out_bounds.id_user = ?", productid, id).First(&cart)

	if res0.Error != nil {
		log.Println("cannot read data", res0.Error.Error())
		return domain.InOutBounds{}
	}

	if res0.RowsAffected == 0 {
		log.Println("failed read data")
		return domain.InOutBounds{}
	}

	res := ioub.db.Model(&InOutBounds{}).Where("id_product = ? AND id_user = ?", productid, id).Updates(cart)

	if res.Error != nil {
		log.Println("cannot update data", res.Error.Error())
		return domain.InOutBounds{}
	}

	if res.RowsAffected == 0 {
		log.Println("data not found")
		return domain.InOutBounds{}
	}

	return cart.ToIOB()
}
func (ioub *inoutboundData) ReadEntryAdminData(role string) []domain.InOutBounds {
	var cart []InOutBounds
	err := ioub.db.Model(&InOutBounds{}).Where("in_out_bounds.role = ?", role).Find(&cart)
	if err.Error != nil {
		log.Println("cannot read data", err.Error.Error())
		return []domain.InOutBounds{}
	}
	if err.RowsAffected == 0 {
		log.Println("data not found")
		return []domain.InOutBounds{}
	}
	return ParseIOBToArr(cart)
}

func (ioub *inoutboundData) ReadEntryUserData(id int) []domain.InOutBounds {
	var cart []InOutBounds
	err := ioub.db.Model(&InOutBounds{}).Where("in_out_bounds.id_user = ?", id).Find(&cart)
	if err.Error != nil {
		log.Println("cannot read data", err.Error.Error())
		return []domain.InOutBounds{}
	}
	if err.RowsAffected == 0 {
		log.Println("data not found")
		return []domain.InOutBounds{}
	}
	return ParseIOBToArr(cart)
}