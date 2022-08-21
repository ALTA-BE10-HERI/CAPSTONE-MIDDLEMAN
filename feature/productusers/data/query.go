package data

import (
	"log"
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type productUserData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductUserData {
	return &productUserData{
		db: db,
	}
}

func (pud *productUserData) CreateProductData(newProduct domain.ProductUser) domain.ProductUser {
	var product = FromPU(newProduct)
	err := pud.db.Create(&product)

	if err.Error != nil {
		log.Println("cannot create data", err.Error.Error())
		return domain.ProductUser{}
	}

	if err.RowsAffected == 0 {
		log.Println("failed to insert data")
		return domain.ProductUser{}
	}
	return product.ToPU()
}

func (pud *productUserData) ReadAllProductData(id int) []domain.ProductUser {
	var product []ProductUser
	err := pud.db.Where("id_user = ?", id).Find(&product)
	if err.Error != nil {
		log.Println("cannot read data", err.Error.Error())
		return []domain.ProductUser{}
	}
	if err.RowsAffected == 0 {
		log.Println("data not found")
		return []domain.ProductUser{}
	}
	return ParsePUToArr(product)
}

func (pud *productUserData) UpdateProductData(data map[string]interface{}, productid, id int) domain.ProductUser {
	var product ProductUser
	res := pud.db.Model(&ProductUser{}).Where("id = ? AND id_user = ?", productid, id).Updates(data)

	if res.Error != nil {
		log.Println("cannot update data", res.Error.Error())
		return domain.ProductUser{}
	}

	if res.RowsAffected == 0 {
		log.Println("data not found")
		return domain.ProductUser{}
	}

	res0 := pud.db.Model(&ProductUser{}).Where("id = ? AND id_user = ?", productid, id).First(&product)

	if res0.Error != nil {
		log.Println("cannot read data", res0.Error.Error())
		return domain.ProductUser{}
	}

	if res0.RowsAffected == 0 {
		log.Println("failed read data")
		return domain.ProductUser{}
	}

	return product.ToPU()
}

func (pud *productUserData) DeleteProductData(productid, id int) (err string) {
	res := pud.db.Where("id = ? AND id_user = ?", productid, id).Delete(&ProductUser{})

	if res.Error != nil {
		log.Println("cannot delete data")
		return "cannot delete data"
	}
	if res.RowsAffected < 1 {
		log.Println("no data deleted")
		return "no data deleted"
	}
	return ""
}

func (pud *productUserData) SearchRestoData(search string, idUser int) (result []domain.ProductUser, err error) {
	var dataProductUser []ProductUser
	res := pud.db.Where("name like ? AND id_user = ?", "%"+search+"%", idUser).Find(&dataProductUser)
	log.Println("cek id user :", idUser)
	if res.Error != nil {
		return []domain.ProductUser{}, res.Error
	}
	return toModelList(dataProductUser), nil
}
