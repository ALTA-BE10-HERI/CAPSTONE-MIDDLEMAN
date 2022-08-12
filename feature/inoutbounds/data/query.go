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

func (iobd *inoutboundData) CekEntry(newProduct domain.InOutBounds) (cek bool, idcart, qty int) {
	var cart InOutBounds
	err := iobd.db.Model(&InOutBounds{}).Where("id_product = ? AND id_user = ?", newProduct.IdProduct, newProduct.IdUser).First(&cart)

	if err.Error != nil {
		return false, 0, 0
	}

	return true, int(cart.ID), int(cart.Qty)
}

func (iobd *inoutboundData) UpdateQty(idcart, qty int) (row int, err error) {
	res := iobd.db.Model(&InOutBounds{}).Where("id = ?", idcart).Update("qty", qty)

	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected != 1 {
		log.Println("failed to insert data")
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}
